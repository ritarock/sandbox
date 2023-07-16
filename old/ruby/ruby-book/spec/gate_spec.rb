require './spec/spec_helper'
require './lib/gate'
require './lib/ticket'

RSpec.describe 'Gate' do
  before do
    @umeda = Gate.new(:umeda)
    @juso = Gate.new(:juso)
    @mikuni = Gate.new(:mikuni)
  end

  describe 'Umeda to Juso' do
    it 'is OK' do
      ticket = Ticket.new(150)
      @umeda.enter(ticket)
      expect(@juso.exit(ticket)).to be_truthy
    end
  end

  describe 'Umeda to Mikuni' do
    context 'fare is not enough' do
      it 'is NG' do
        ticket = Ticket.new(150)
        @umeda.enter(ticket)
        expect(@mikuni.exit(ticket)).to be_falsey
      end
    end
    context 'fare is enough' do
      it 'is OK' do
        ticket = Ticket.new(190)
        @umeda.enter(ticket)
        expect(@mikuni.exit(ticket)).to be_truthy
      end
    end
  end

  describe 'Juso to Mikuni' do
    it 'is OK' do
      ticket = Ticket.new(150)
      @juso.enter(ticket)
      expect(@mikuni.exit(ticket)).to be_truthy
    end
  end
end

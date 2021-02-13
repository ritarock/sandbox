require './spec/spec_helper'
require './lib/bank'
require './lib/team'

RSpec.describe 'Deep freezable' do
  describe 'to array' do
    it 'freezes deeply' do
      expect(Team::COUNTRIES).to eq ['Japan', 'US', 'India']
      expect(Team::COUNTRIES).to be_frozen
      expect(Team::COUNTRIES.all? { |country| country.frozen? }).to be_truthy
    end
  end

  describe 'to hash' do
    it 'freezes deeply' do
      expect(Bank::CURRENCIES).to eq({ 'Japan' => 'yen', 'US' => 'dollar', 'India' => 'rupee' })
      expect(Bank::CURRENCIES).to be_frozen
      expect(Bank::CURRENCIES.all? { |key, value| key.frozen? && value.frozen? }).to be_truthy
    end
  end
end

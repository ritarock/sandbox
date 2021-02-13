require './spec/spec_helper'
require './lib/effects'

RSpec.describe Effects do
  describe '.reverse' do
    it 'returns valid string' do
      effect = Effects.reverse
      expect(effect.call('Ruby is fun!')).to eq 'ybuR si !nuf'
    end
  end

  describe '.echo' do
    it 'returns valid string' do
      effect = Effects.echo(2)
      expect(effect.call('Ruby is fun!')).to eq 'RRuubbyy iiss ffuunn!!'

      effect = Effects.echo(3)
      expect(effect.call('Ruby is fun!')).to eq 'RRRuuubbbyyy iiisss fffuuunnn!!!'
    end
  end

  describe '.loud' do
    it 'returns valid string' do
      effect = Effects.loud(2)
      expect(effect.call('Ruby is fun!')).to eq 'RUBY!! IS!! FUN!!!'

      effect = Effects.loud(3)
      expect(effect.call('Ruby is fun!')).to eq 'RUBY!!! IS!!! FUN!!!!'
    end
  end
end

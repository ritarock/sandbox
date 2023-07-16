class WordSynth
  def initialize
    @effects = []
  end

  def add_effect(effect)
    @effects << effect
  end

  def play(original_words)
    @effects.inject(original_words) do |word, effect|
      effect.call(word)
    end
  end
end

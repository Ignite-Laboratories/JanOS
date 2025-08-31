package english

// PartOfSpeech represents a lemma that's an adjective, adverb, noun, or verb.  As a lemma can be considered as
// any number of these, this is your first branch towards Synset definitions.
//
// See Lexeme, Prototype, PunctuationMark, Quote, Dash, Function, ControlCharacter, PartOfSpeech, and Synset
type PartOfSpeech struct {
	Lexeme

	Adjective []Synset
	Adverb    []Synset
	Noun      []Synset
	Verb      []Synset
}

func (l PartOfSpeech) String() string {
	return l.Lemma
}

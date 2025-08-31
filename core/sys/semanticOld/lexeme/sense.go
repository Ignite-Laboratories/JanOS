package lexeme

import "core/enum/partOfSpeech"

// Sense is analogous to a flattened synset meant to provide direct access to all traversal dimensions associated
// with a single lemma.  A Sense is what you reach when you have navigated deeper into a Lexicon, essentially.
//
// - Sense, Adjective, Adverb, Noun, and Verb
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Sense struct {
	Source Lexicon

	PartOfSpeech partOfSpeech.PartOfSpeech

	Synonyms []Sense
	Antonyms []Sense

	Hypernyms []Sense
	Hyponyms  []Sense

	Holonyms []Sense
	Meronyms []Sense

	Pertainyms []Sense
}

func (s Sense) String() string {
	return s.Source.String()
}

// Adjective is a kind of Sense that provides adjective-related dimensionality.
//
// - Sense, Adjective, Adverb, Noun, and Verb
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Adjective Sense

func (a Adjective) String() string {
	return a.Source.String()
}

// Adverb is a kind of Sense that provides adverb-related dimensionality.
//
// - Sense, Adjective, Adverb, Noun, and Verb
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Adverb Sense

func (r Adverb) String() string {
	return r.Source.String()
}

// Noun is a kind of Sense that provides noun-related dimensionality.
//
// - Sense, Adjective, Adverb, Noun, and Verb
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Noun Sense

func (n Noun) String() string {
	return n.Source.String()
}

// Verb is a kind of Sense that provides verb-related dimensionality.
//
// - Sense, Adjective, Adverb, Noun, and Verb
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Verb Sense

func (v Verb) String() string {
	return v.Source.String()
}

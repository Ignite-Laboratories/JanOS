package lexeme

// Lexicon represents the bridge between a Lemma and it's senses.  A Sense is the contextual relation between
// a Lemma as a particular part of speech and a synset.  A synset defines the definition (read: gloss) of a
// synonymous set of lemma.  In WordNet, this is where the pointers for traversal are attained - in JanOS,
// these are "flattened" into their appropriate Sense types.
//
// - Sense, Adjective, Adverb, Noun, and Verb
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Lexicon struct {
	Lemma string

	Adjective []Adjective
	Adverb    []Adverb
	Noun      []Noun
	Verb      []Verb
}

func (l Lexicon) String() string {
	return l.Lemma
}

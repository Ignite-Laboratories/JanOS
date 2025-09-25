package lexeme

// Function words are the semantic glue that holds language together.  They compose of many different categories,
// each of which is a type of function word.  If you wish to describe 'any function word' use this interface - otherwise
// you may describe it using the below references:
//
// - AuxiliaryVerb, Conjunction, Preposition, Determiner, Pronoun, Quantifier, and Particle.
//
// NOTE: Function words are always stored in lowercase form for disambiguation.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Function Base

// A Conjunction links together words, phrases, or clauses of equal or unequal grammatical rank - see Conjunctions.
//
// - AuxiliaryVerb, Conjunction, Preposition, Determiner, Pronoun, Quantifier, and Particle.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Conjunction Function

// Conjunctions - see Conjunction.
var Conjunctions = []Function{
	{"and"}, {"but"}, {"or"}, {"not"},
	{"so"}, {"yet"}, {"for"}, {"although"},
	{"because"}, {"since"}, {"unless"}, {"while"},
	{"whereas"}, {"if"}, {"though"},
}

func (f Conjunction) String() string {
	return f.Lexeme
}

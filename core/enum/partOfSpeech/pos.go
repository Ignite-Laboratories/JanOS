package partOfSpeech

// A PartOfSpeech is a category to which a word is assigned in accordance with its syntactic functions.
//
// See Adjective, Adverb, Noun, and Verb
type PartOfSpeech string

const (
	// An Adjective is a word or phrase naming an attribute, added to or grammatically related to a noun to modify or describe it.
	//
	// See Adjective, Adverb, Noun, and Verb
	Adjective PartOfSpeech = "a"

	// An Adverb is a word or phrase that modifies or qualifies an adjective, verb, or other adverb or a word group, expressing
	// a relation of place, time, circumstance, manner, cause, degree, etc...
	//
	// See Adjective, Adverb, Noun, and Verb
	Adverb PartOfSpeech = "r"

	// A Noun is a word (other than a pronoun) used to identify any of a class of people, places, or things (common noun),
	// or to name a particular one of these (proper noun.)
	//
	// See Adjective, Adverb, Noun, and Verb
	Noun PartOfSpeech = "n"

	// A Verb is a word used to describe an action, state, or occurrence, and forming the main part of the predicate of a sentence.
	//
	// See Adjective, Adverb, Noun, and Verb
	Verb PartOfSpeech = "v"
)

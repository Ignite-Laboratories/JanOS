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

// An AuxiliaryVerb helps a main verb to express tense, mood, or voice - see AuxiliaryVerbs.
//
// - AuxiliaryVerb, Conjunction, Preposition, Determiner, Pronoun, Quantifier, and Particle.
//
// See std.Lexeme, Base, Function
type AuxiliaryVerb Function

// AuxiliaryVerbs - see AuxiliaryVerb
var AuxiliaryVerbs = []Function{
	{"be"}, {"am"}, {"is"}, {"are"},
	{"was"}, {"were"}, {"being"}, {"been"},
	{"have"}, {"has"}, {"had"}, {"having"},
	{"do"}, {"does"}, {"did"}, {"will"},
	{"would"}, {"shall"}, {"should"},
	{"can"}, {"could"}, {"may"}, {"might"}, {"must"},
}

func (f AuxiliaryVerb) String() string {
	return f.Lexeme
}

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

// A Preposition establishes a relationship, typically spatial or temporal, between a noun or pronoun and another word - see Prepositions.
//
// - AuxiliaryVerb, Conjunction, Preposition, Determiner, Pronoun, Quantifier, and Particle.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Preposition Function

// Prepositions - See Preposition
var Prepositions = []Function{
	{"of"}, {"to"}, {"in"}, {"on"},
	{"at"}, {"by"}, {"for"}, {"with"},
	{"about"}, {"over"}, {"under"}, {"between"},
	{"through"}, {"during"}, {"against"}, {"among"},
	{"across"}, {"behind"}, {"beyond"}, {"inside"},
	{"outside"}, {"without"}, {"within"}, {"upon"},
}

func (f Preposition) String() string {
	return f.Lexeme
}

// A Determiner specifies the reference of a noun, providing context about its quantity, definiteness, or ownership - see Determiners.
//
// - AuxiliaryVerb, Conjunction, Preposition, Determiner, Pronoun, Quantifier, and Particle.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Determiner Function

// Determiners - See Determiner
var Determiners = []Function{
	{"a"}, {"an"}, {"the"}, {"this"},
	{"that"}, {"these"}, {"those"}, {"my"},
	{"your"}, {"his"}, {"her"}, {"its"},
	{"our"}, {"their"}, {"some"}, {"any"},
	{"no"}, {"each"}, {"every"}, {"either"},
	{"neither"}, {"much"}, {"many"}, {"few"},
	{"several"}, {"all"}, {"both"},
}

func (f Determiner) String() string {
	return f.Lexeme
}

// A Pronoun substitutes for a noun or noun phrase, often to prevent repetition - see Pronouns.
//
// - AuxiliaryVerb, Conjunction, Preposition, Determiner, Pronoun, Quantifier, and Particle.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Pronoun Function

// Pronouns - See Pronoun
var Pronouns = []Function{
	{"I"}, {"you"}, {"he"}, {"she"},
	{"it"}, {"we"}, {"they"}, {"me"},
	{"him"}, {"her"}, {"us"}, {"them"},
	{"myself"}, {"yourself"}, {"himself"}, {"herself"},
	{"itself"}, {"ourselves"}, {"yourselves"}, {"themselves"},
	{"who"}, {"whom"}, {"whose"}, {"which"},
	{"what"}, {"anyone"}, {"someone"}, {"everyone"},
	{"nobody"}, {"nothing"}, {"anything"}, {"everything"},
}

func (f Pronoun) String() string {
	return f.Lexeme
}

// A Quantifier indicates the quantity or amount of a noun without being a specific number - see Quantifiers.
//
// - AuxiliaryVerb, Conjunction, Preposition, Determiner, Pronoun, Quantifier, and Particle.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Quantifier Function

// Quantifiers - See Quantifier
var Quantifiers = []Function{
	{"few"}, {"little"}, {"much"}, {"many"},
	{"several"}, {"all"}, {"some"}, {"any"},
	{"enough"}, {"more"}, {"most"}, {"less"},
	{"least"},
}

func (f Quantifier) String() string {
	return f.Lexeme
}

// A Particle modifies or connects other parts of a phrase - see Particles.
//
// - AuxiliaryVerb, Conjunction, Preposition, Determiner, Pronoun, Quantifier, and Particle.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Particle Function

// Particles - See Particle
var Particles = []Function{
	{"not"}, {"no"}, {"nor"}, {"only"},
	{"just"}, {"even"}, {"still"}, {"yet"},
	{"already"}, {"also"}, {"too"},
	{"there"}, {"here"}, {"then"}, {"now"},
	{"how"}, {"when"}, {"where"}, {"why"},
	{"yes"}, {"no"}, {"okay"}, {"well"},
}

func (f Particle) String() string {
	return f.Lexeme
}

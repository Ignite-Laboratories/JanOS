package lexeme

// Base represents a concrete basic std.Lexeme.
//
// - Punctuation and Prototype.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Base struct {
	Lexeme string
}

func (b Base) String() string {
	return b.Lexeme
}

// Punctuation is a Base std.Lexeme used as a punctuation mark, such as '!' or '?' - or even the quote characters, themselves.
//
// - Punctuation and Prototype.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Punctuation Base

func (p Punctuation) String() string {
	return p.Lexeme
}

// Prototype is a Base std.Lexeme of unknown classification.  This collection should be used as a source from which to
// build a Lexicon over time.
//
// - Punctuation and Prototype.
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Prototype Base

func (p Prototype) String() string {
	return p.Lexeme
}

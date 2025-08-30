package lexeme

// Base represents a concrete std.Lexeme.
//
// The most primitive lexemes are Punctuation and Proto.
//
// See std.Lexeme, lexeme.Base, lexeme.Function, lexeme.Lexicon, and lexeme.Special
type Base struct {
	Lexeme string
}

func (b Base) String() string {
	return b.Lexeme
}

// Punctuation is any std.Lexeme used as a punctuation mark, such as '!' or '?' - or even the quote characters, themselves.
//
// The most primitive lexemes are Punctuation and Proto.
//
// See std.Lexeme, lexeme.Base
type Punctuation Base

func (p Punctuation) String() string {
	return p.Lexeme
}

// Proto represents any std.Lexeme of unknown classification.  This collection should be used as a source from which to
// build the Lexicon over time.
//
// The most primitive lexemes are Punctuation and Proto.
//
// See std.Lexeme, lexeme.Base
type Proto Base

func (p Proto) String() string {
	return p.Lexeme
}

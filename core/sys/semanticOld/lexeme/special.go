package lexeme

// Special represents any std.Lexeme of either a special character Symbol (like '#' or '@') or a Control signal.  As we
// construct the lexicon of semantically parsing data, this acts as a way to guide logic.
//
// - Control and Symbol
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Special Base

func (s Special) String() string {
	return s.Lexeme
}

// Control is any std.Lexeme used to control the flow of logic.  For instance, this might direct the cursor to "branch"
// or change the current direction of travel.
//
// - Control and Symbol
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Control Special

func (c Control) String() string {
	return c.Lexeme
}

// Symbol is any std.Lexeme of a special symbol character not considered a Punctuation mark - for instance, '@' or '#'.
//
// - Control and Symbol
//
// See std.Lexeme, Base, Function, Lexicon, and Special
type Symbol Base

func (s Symbol) String() string {
	return s.Lexeme
}

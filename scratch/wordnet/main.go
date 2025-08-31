package main

import (
	"core/std"
	"fmt"
)

func main() {
	fma := std.NewFMA[float32](0.1, 1.2, 3.3, "𝑎")
	fmt.Println(fma)
	fmt.Println(fma.ForceForceForce())

	f := Proto{"Hello"}

	Test(f)
}

func Test(a std.Lexeme) {
	fmt.Println(a)
}

type Base struct {
	Lemma string
}

// Punctuation is any std.Lexeme used as a punctuation mark, such as '!' or '?' - or even the quote characters, themselves.
//
// See std.Lexeme, Proto, Special, Punctuation, Function, and Lexicon
type Punctuation Base

func (p Punctuation) String() string {
	return p.Lemma
}

// Proto represents any std.Lexeme of unknown classification.  This collection should be used as a source from which to
// build the Lexicon over time.
//
// See std.Lexeme, Proto, Special, Punctuation, Function, and Lexicon
type Proto Base

func (p Proto) String() string {
	return p.Lemma
}

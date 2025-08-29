// Package italic provides quick access to the italic case-invariant 26 English alphabet characters.
//
// See Lower and Upper
package italic

import (
	"core/enum/italic/lower"
	"core/enum/italic/upper"
	"core/sys/support"
	"strings"
)

var Characters = append(lower.Characters, upper.Characters...)

type _letters struct {
	// A represents the case-invariant English letter '𝐴' or '𝑎'
	A string
	// B represents the case-invariant English letter '𝐵' or '𝑏'
	B string
	// C represents the case-invariant English letter '𝐶' or '𝑐'
	C string
	// D represents the case-invariant English letter '𝐷' or '𝑑'
	D string
	// E represents the case-invariant English letter '𝐸' or '𝑒'
	E string
	// F represents the case-invariant English letter '𝐹' or '𝑓'
	F string
	// G represents the case-invariant English letter '𝐺' or '𝑔'
	G string
	// H represents the case-invariant English letter '𝐻' or 'ℎ'
	H string
	// I represents the case-invariant English letter '𝐼' or '𝑖'
	I string
	// J represents the case-invariant English letter '𝐽' or '𝑗'
	J string
	// K represents the case-invariant English letter '𝐾' or '𝑘'
	K string
	// L represents the case-invariant English letter '𝐿' or '𝑙'
	L string
	// M represents the case-invariant English letter '𝑀' or '𝑚'
	M string
	// N represents the case-invariant English letter '𝑁' or '𝑛'
	N string
	// O represents the case-invariant English letter '𝑂' or '𝑜'
	O string
	// P represents the case-invariant English letter '𝑃' or '𝑝'
	P string
	// Q represents the case-invariant English letter '𝑄' or '𝑞'
	Q string
	// R represents the case-invariant English letter '𝑅' or '𝑟'
	R string
	// S represents the case-invariant English letter '𝑆' or '𝑠'
	S string
	// T represents the case-invariant English letter '𝑇' or '𝑡'
	T string
	// U represents the case-invariant English letter '𝑈' or '𝑢'
	U string
	// V represents the case-invariant English letter '𝑉' or '𝑣'
	V string
	// W represents the case-invariant English letter '𝑊' or '𝑤'
	W string
	// X represents the case-invariant English letter '𝑋' or '𝑥'
	X string
	// Y represents the case-invariant English letter '𝑌' or '𝑦'
	Y string
	// Z represents the case-invariant English letter '𝑍' or '𝑧'
	Z string
}

// Lower provides access to the lowercase English alphabet.
//
// See lower.A, lower.B, lower.C, lower.D, lower.E, lower.F, lower.G, lower.H, lower.I, lower.J, lower.K, lower.L, lower.M, lower.N, lower.O, lower.P, lower.Q, lower.R, lower.S, lower.T, lower.U, lower.V, lower.W, lower.X, lower.Y, and lower.Z
var Lower = _letters{
	A: lower.A,
	B: lower.B,
	C: lower.C,
	D: lower.D,
	E: lower.E,
	F: lower.F,
	G: lower.G,
	H: lower.H,
	I: lower.I,
	J: lower.J,
	K: lower.K,
	L: lower.L,
	M: lower.M,
	N: lower.N,
	O: lower.O,
	P: lower.P,
	Q: lower.Q,
	R: lower.R,
	S: lower.S,
	T: lower.T,
	U: lower.U,
	V: lower.V,
	W: lower.W,
	X: lower.X,
	Y: lower.Y,
	Z: lower.Z,
}

// Upper provides access to the uppercase English alphabet.
//
// See upper.A, upper.B, upper.C, upper.D, upper.E, upper.F, upper.G, upper.H, upper.I, upper.J, upper.K, upper.L, upper.M, upper.N, upper.O, upper.P, upper.Q, upper.R, upper.S, upper.T, upper.U, upper.V, upper.W, upper.X, upper.Y, and upper.Z
var Upper = _letters{
	A: upper.A,
	B: upper.B,
	C: upper.C,
	D: upper.D,
	E: upper.E,
	F: upper.F,
	G: upper.G,
	H: upper.H,
	I: upper.I,
	J: upper.J,
	K: upper.K,
	L: upper.L,
	M: upper.M,
	N: upper.N,
	O: upper.O,
	P: upper.P,
	Q: upper.Q,
	R: upper.R,
	S: upper.S,
	T: upper.T,
	U: upper.U,
	V: upper.V,
	W: upper.W,
	X: upper.X,
	Y: upper.Y,
	Z: upper.Z,
}

func Italicize(input string) string {
	input = support.SetCase(strings.Clone(input), false)
	runes := []rune(input)
	for i, r := range runes {
		switch r {
		case 'a':
			runes[i] = rune(lower.A[0])
		case 'A':
			runes[i] = rune(upper.A[0])
		case 'b':
			runes[i] = rune(lower.B[0])
		case 'B':
			runes[i] = rune(upper.B[0])
		case 'c':
			runes[i] = rune(lower.C[0])
		case 'C':
			runes[i] = rune(upper.C[0])
		case 'd':
			runes[i] = rune(lower.D[0])
		case 'D':
			runes[i] = rune(upper.D[0])
		case 'e':
			runes[i] = rune(lower.E[0])
		case 'E':
			runes[i] = rune(upper.E[0])
		case 'f':
			runes[i] = rune(lower.F[0])
		case 'F':
			runes[i] = rune(upper.F[0])
		case 'g':
			runes[i] = rune(lower.G[0])
		case 'G':
			runes[i] = rune(upper.G[0])
		case 'h':
			runes[i] = rune(lower.H[0])
		case 'H':
			runes[i] = rune(upper.H[0])
		case 'i':
			runes[i] = rune(lower.I[0])
		case 'I':
			runes[i] = rune(upper.I[0])
		case 'j':
			runes[i] = rune(lower.J[0])
		case 'J':
			runes[i] = rune(upper.J[0])
		case 'k':
			runes[i] = rune(lower.K[0])
		case 'K':
			runes[i] = rune(upper.K[0])
		case 'l':
			runes[i] = rune(lower.L[0])
		case 'L':
			runes[i] = rune(upper.L[0])
		case 'm':
			runes[i] = rune(lower.M[0])
		case 'M':
			runes[i] = rune(upper.M[0])
		case 'n':
			runes[i] = rune(lower.N[0])
		case 'N':
			runes[i] = rune(upper.N[0])
		case 'o':
			runes[i] = rune(lower.O[0])
		case 'O':
			runes[i] = rune(upper.O[0])
		case 'p':
			runes[i] = rune(lower.P[0])
		case 'P':
			runes[i] = rune(upper.P[0])
		case 'q':
			runes[i] = rune(lower.Q[0])
		case 'Q':
			runes[i] = rune(upper.Q[0])
		case 'r':
			runes[i] = rune(lower.R[0])
		case 'R':
			runes[i] = rune(upper.R[0])
		case 's':
			runes[i] = rune(lower.S[0])
		case 'S':
			runes[i] = rune(upper.S[0])
		case 't':
			runes[i] = rune(lower.T[0])
		case 'T':
			runes[i] = rune(upper.T[0])
		case 'u':
			runes[i] = rune(lower.U[0])
		case 'U':
			runes[i] = rune(upper.U[0])
		case 'v':
			runes[i] = rune(lower.V[0])
		case 'V':
			runes[i] = rune(upper.V[0])
		case 'w':
			runes[i] = rune(lower.W[0])
		case 'W':
			runes[i] = rune(upper.W[0])
		case 'x':
			runes[i] = rune(lower.X[0])
		case 'X':
			runes[i] = rune(upper.X[0])
		case 'y':
			runes[i] = rune(lower.Y[0])
		case 'Y':
			runes[i] = rune(upper.Y[0])
		case 'z':
			runes[i] = rune(lower.Z[0])
		case 'Z':
			runes[i] = rune(upper.Z[0])
		}
	}
	return string(runes)
}

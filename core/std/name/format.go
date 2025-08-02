package name

// Format represents the format this name should follow.
//
// See NameDB, SurnameDB, Tiny, and Multi.
type Format interface {
	NameDB | SurnameDB | Tiny | Multi
}

// NameDB represents a set of cultural names, typically one word in length.
//
// All credit goes to Kevin MacLeod of Incompetech for such a wonderful source database!
// https://incompetech.com
//
// Please check his stuff out, he's quite clever!
//
// See SurnameDB, Tiny, Multi, and Format.
type NameDB string

// SurnameDB represents a set of surnames.
//
// This list was gathered from The Internet Surname Database - please check out their stuff!
// https://surnamedb.com
//
// See NameDB, Tiny, Multi, and Format.
type SurnameDB string

// Tiny is a database that generates names satisfying tiny's implicit naming requirements.
// Currently, these are our explicit filters -
//
//   - Only the standard 26 letters of the English alphabet (case-insensitive)
//   - No whitespace or special characters (meaning only single word alpha-explicit names)
//   - At least three characters in length
//   - At least 2¹⁴ unique names before beginning to recycling names
//   - Names are case-sensitive in uniqueness.
//
// These filters will never be reduced - if any changes are made, they will only be augmented.
//
// See NameDB, SurnameDB, Multi, and Format.
type Tiny string

// Multi represents a name with multiple parts.
//
// See NameDB, SurnameDB, Tiny, and Format.
type Multi string

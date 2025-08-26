package format

// Format represents the format this name should follow.
//
// See Format, NameDB, SurnameDB, Tiny, Multi, and Default.
type Format interface {
	NameDB | SurnameDB | Tiny | Multi | Default
}

// Default currently generates a Multi name.
//
// NOTE: This may change in the future!  If you need a strict format, please use that instead.
//
// See Format, NameDB, SurnameDB, Tiny, Multi, and Default.
type Default string

// NameDB is a database of cultural first names.
//
// All credit goes to Kevin MacLeod of Incompetech for such a wonderful source database -
//
// https://incompetech.com
//
// Please check his stuff out, he's quite clever!
//
// See Format, NameDB, SurnameDB, Tiny, Multi, and Default.
type NameDB string

// SurnameDB is a database of surnames.
//
// This list was gathered from The Internet Surname Database - please check out their stuff!
//
// https://surnamedb.com
//
// See Format, NameDB, SurnameDB, Tiny, Multi, and Default.
type SurnameDB string

// Tiny generates names that satisfy tiny's implicit naming requirements.
// Currently, these are our explicit filters -
//
//   - Only the standard 26 letters of the English alphabet (case-insensitive)
//   - No whitespace or special characters (meaning only single word alpha-explicit names)
//   - At least three characters in length
//   - At least 2¹⁴ unique names before beginning to recycle
//   - Names are case-insensitively unique.
//
// These filters will never be reduced - if any changes are made, they will only be augmented.
//
// NOTE: Names are currently selected from the NameDB.
//
// See Format, NameDB, SurnameDB, Tiny, Multi, and Default.
type Tiny string

// Multi generates a multipart name with an entry from NameDB, a single space, and then an entry from SurnameDB.
//
// NOTE: The cultural information will persist from the used NameDB entry, as the SurnameDB does not contain cultural information.
//
// See Format, NameDB, SurnameDB, Tiny, Multi, and Default.
type Multi string

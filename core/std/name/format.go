package name

// Format represents the format this name should follow.
//
// See NameDB, SurnameDB, Tiny, and Multi.
type Format interface {
	NameDB | SurnameDB | Tiny | Multi
}

// NameDB is a database of cultural first names.
//
// All credit goes to Kevin MacLeod of Incompetech for such a wonderful source database -
//
// https://incompetech.com
//
// Please check his stuff out, he's quite clever!
//
// See SurnameDB, Tiny, Multi, and Format.
type NameDB string

// SurnameDB is a database of surnames.
//
// This list was gathered from The Internet Surname Database - please check out their stuff!
//
// https://surnamedb.com
//
// See NameDB, Tiny, Multi, and Format.
type SurnameDB string

// Tiny generates names that satisfy tiny's implicit naming requirements.
// Currently, these are our explicit filters -
//
//   - Only the standard 26 letters of the English alphabet (case-insensitive)
//   - No whitespace or special characters (meaning only single word alpha-explicit names)
//   - At least three characters in length
//   - At least 2¹⁴ unique names before beginning to recycle names
//   - Names are case-insensitive in uniqueness.
//
// These filters will never be reduced - if any changes are made, they will only be augmented.
//
// NOTE: Names are currently selected from the NameDB.
//
// See NameDB, SurnameDB, Multi, and Format.
type Tiny string

// Multi generates a multipart name with an entry from NameDB, a single space, and then an entry from SurnameDB.
//
// NOTE: The cultural information will persist from the used NameDB entry, as the SurnameDB does not contain cultural information.
//
// See NameDB, SurnameDB, Tiny, and Format.
type Multi string

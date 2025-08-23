package scheme

// Reverse indicates that the padding elements should be reversed before application to the source data.
type Reverse byte

// Tile indicates that the padding elements should be repeated as is while applying it to the source data.
type Tile byte

// Randomize indicates that the padding elements should be randomly re-ordered before application to the source data.
type Randomize byte

// ReflectInward indicates that the padding elements should be treated as walking centrally inward.
//
// This means the elements are ordered from ABCDE to ACEDB, where elements are walked pairwise inward.
//
// If statically interleaving '11111' with 'ABCDE' to 10 elements wide - A 1 C 1 E 1 D 1 B 1
//
// If padding the left of '11111' with 'ABCDE' to 10 elements wide - A C E D B 1 1 1 1 1
//
// If padding the right of '11111' with 'ABCDE' to 10 elements wide - 1 1 1 1 1 A C E D B
type ReflectInward byte

// ReflectOutward indicates that the padding elements should be treated as walking centrally outward.
//
// This means the elements are ordered from ABCDE to DBACE, where A is placed in the middle before
// elements are walked pairwise outward.
//
// If statically interleaving '11111' with 'ABCDE' to 10 elements wide - D 1 B 1 A 1 C 1 E 1
//
// If padding the left of '11111' with 'ABCDE' to 10 elements wide - D B A C E 1 1 1 1 1
//
// If padding the right of '11111' with 'ABCDE' to 10 elements wide - 1 1 1 1 1 D B A C E
type ReflectOutward byte

package std

import "fmt"

// Bit represents one binary place. [0 - 1]
type Bit byte

// BitSanityCheck checks if the provided bytes are either 0, 1, or Nil (219) - otherwise, it panics.
func BitSanityCheck(bits ...Bit) {
	for _, b := range bits {
		if b != 0 && b != 1 {
			panic(fmt.Errorf("not a bit value: %d", b))
		}
	}
}

// String converts the provided Bit to a string "1", "0", or "-" for Nil [219] and panics if the found value is anything else.
func (b Bit) String() string {
	switch b {
	case 0:
		return "0"
	case 1:
		return "1"
	case 219:
		return "-"
	default:
		panic(fmt.Errorf("not a bit value: %d", b))
	}
}

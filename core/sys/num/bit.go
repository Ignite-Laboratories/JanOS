package num

import "fmt"

// A Bit is a byte that implicitly holds the value 0 or 1.
type Bit byte

// SanityCheck panics if the underlying byte value is not 0 or 1
func (b Bit) SanityCheck() {
	if b != 0 && b != 1 {
		panic(fmt.Sprintf("invalid bit value: %d", b))
	}
}

func (b Bit) String() string {
	return fmt.Sprintf("%d", b)
}

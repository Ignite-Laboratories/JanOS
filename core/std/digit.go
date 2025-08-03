package std

import "fmt"

// Digit represents a single placeholder digit holding a value up to base 255.
//
// NOTE: Digits can be addressed with 0x values - see String
type Digit byte

// String prints this digit value in capitalized hexadecimal form.
func (d Digit) String() string {
	return fmt.Sprintf("%X", d)
}

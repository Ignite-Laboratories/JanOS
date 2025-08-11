package std

import "fmt"

// Digit represents a single placeholder digit holding a value up to base 255.  While in digital form these merely
// hold the intrinsic value - when converting to string, these are represented in hexadecimal form.  This allows
// higher-order bases to be representable in a universal fashion, while maintaining a logical order to their identifiers.
type Digit byte

// String prints this digit value in capitalized hexadecimal form.
func (d Digit) String() string {
	return fmt.Sprintf("%X", d)
}

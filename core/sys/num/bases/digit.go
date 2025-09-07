package bases

import "fmt"

type Digit byte

func (d Digit) Increment(base uint16) Digit {
	d += 1
	if d > Digit(base)-1 {
		d = 0
	}
	return d
}

func (d Digit) ShouldRound(base uint16) bool {
	mid := Digit(base / 2)
	return d > mid
}

func (d Digit) String() string {
	return fmt.Sprintf("%02x", byte(d))
}

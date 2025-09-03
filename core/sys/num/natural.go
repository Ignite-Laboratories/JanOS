package num

import (
	"strconv"
	"strings"
)

// A Natural is a slice of Placeholder values.  It is treated as any string, where the most significant
// bit is index 0 and the least is index 𝑛.
type Natural struct {
	digits []Numeric[byte]
	base   byte
}

func newNatural(input string, base byte) Natural {
	n := Natural{
		digits: make([]Numeric[byte], len(input)),
		base:   base,
	}
	for i := len(input) - 1; i >= 0; i-- {
		v, _ := strconv.ParseUint(string(input[i]), int(base), 8)
		bnd, _ := NewNumericBounded[byte](byte(v), 0, base)
		n.digits[i] = bnd
	}
	return n
}

func (n Natural) Width() uint {
	return uint(len(n.digits))
}

func (n Natural) Base() byte {
	return n.base
}

func (n Natural) String() string {
	b := new(strings.Builder)
	for _, p := range n.digits {
		b.WriteString(p.String())
	}
	return b.String()
}

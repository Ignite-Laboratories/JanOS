package std

import "strings"

type Bridge []string

func (b Bridge) String() string {
	return strings.Join(b, " ⇝ ")
}

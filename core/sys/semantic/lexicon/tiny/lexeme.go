package tiny

import (
	"fmt"
)

func (t Lexeme) String() string {
	return fmt.Sprintf("%x", t.Placeholder)
}

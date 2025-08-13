package mouse

import (
	"github.com/ignite-laboratories/core/std"
)

// State provides the current state of the mouse.
type State struct {
	Position std.XY[int]
	Buttons  struct {
		Left, Middle, Right, Extra1, Extra2 bool
	}
}

package mouse

import "github.com/ignite-laboratories/core/std/bounded"

// State provides the current state of the mouse.
type State struct {
	Position bounded.XY[int]
	Buttons  struct {
		Left, Middle, Right, Extra1, Extra2 bool
	}
}

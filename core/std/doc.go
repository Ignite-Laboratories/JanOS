package std

import (
	"git.ignitelabs.net/core"
)

func init() {
	if core.Alive() {
	} // Kick core.init off if using a std type
}

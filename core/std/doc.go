package std

import "core"

func init() {
	if core.Alive() {
	} // Kick core.init off if using a std type
}

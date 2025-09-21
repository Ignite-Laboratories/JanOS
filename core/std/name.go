package std

import (
	"git.ignitelabs.net/core/sys/given"
	"git.ignitelabs.net/core/sys/given/format"
)

// RandomName returns a random given.Name's name string formatted to format.Default.
func RandomName() string {
	return given.Random[format.Default]().Name
}

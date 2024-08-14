package main

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"time"
)

// a sin ( w t + p)
// a - amplitude
// w - angular frequency
//   - 2 * pi * frequency
// t - time
// p - phase adjustment

func main() {
	o := Logic.NewGraphOscillator(100, 1)

	Logic.Drive(time.Millisecond*10, o)
}

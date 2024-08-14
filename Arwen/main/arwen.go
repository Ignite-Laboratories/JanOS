package main

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/Ignite-Laboratories/JanOS/Logic/Math"
)

func main() {
	d1 := Logic.NewDimension("Test", Math.Omega, 0)
	u := Logic.NewUniverse(Logic.World{}, d1)
	u.BigBang()

	u.Dimensions["Test"]
}

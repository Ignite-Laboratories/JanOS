package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/greek"
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/std/phrase"
	"github.com/ignite-laboratories/core/std/rgba"
	"github.com/ignite-laboratories/core/std/xyzw"
)

func main() {
	fmt.Println(greek.Lower.SigmaFinal)
	name.New("bob")
	name.Random[name.Tiny]()
	phrase.OfBits(1).Named("asdf")

	a := xyzw.From(byte(0), 5, 4, 3)
	ff := rgba.From(0xffffffff)
	ffN := rgba.Normalize[byte, float32](ff)
	ffS := rgba.ReScale[float32, byte](ffN)
}

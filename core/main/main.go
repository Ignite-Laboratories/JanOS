package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/greek"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/std/phrase"
	"github.com/ignite-laboratories/core/std/rgba"
	"github.com/ignite-laboratories/core/std/rgbaGeneric"
	"github.com/ignite-laboratories/core/std/xy"
	"github.com/ignite-laboratories/core/std/xyzw"
	"github.com/ignite-laboratories/core/sys/atlas"
)

func main() {
	fmt.Println(greek.Lower.SigmaFinal)
	fmt.Println(name.New("bob"))
	fmt.Println(name.Random[name.Tiny]())
	fmt.Println(name.Random[name.NameDB]())
	fmt.Println(name.Random[name.SurnameDB]())
	fmt.Println(name.Random[name.Multi]())
	fmt.Println(phrase.OfBits(1, 0, 1, 1).Named("asdf"))

	xy.fr

	fmt.Println(xy.From(555, 444, 1024, 768))

	fmt.Println(xyzw.FromInfinite[byte](0, 5, 4, 3))
	ff := rgba.FromHex[num.Note](0xaabbccdd)
	fmt.Println(ff)
	fmt.Println(ff.ABGR())
	ffN := rgba.Normalize[num.Note, float32](ff)
	fmt.Println(ffN)
	fmt.Println(rgba.ReScale[float32, num.Morsel](ffN))

	std.BitSanityCheck()

	fmt.Println(atlas.Precision)

	gen := rgbaGeneric.From(num.Morsel(33), num.Flake(55), num.Crumb(22), byte(111))
	fmt.Println(gen)
}

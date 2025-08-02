package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/greek"
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/std/phrase"
)

func main() {
	fmt.Println(greek.Lower.SigmaFinal)
	name.New("bob")
	name.Random[name.Tiny]()
	phrase.OfBits(1).Named("asdf")
}

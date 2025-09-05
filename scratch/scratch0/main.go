package main

import (
	"core/sys/num"
	"fmt"
)

func main() {
	rl1 := num.NewReal(-1234.555555555555)
	rl1.ChangeBase(10)

	rl2 := num.NewReal(5)

	fmt.Println(rl1, rl2)
}

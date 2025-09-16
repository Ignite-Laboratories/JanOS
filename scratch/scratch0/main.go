package main

import (
	"core/sys/num"
	"fmt"
)

func main() {
	rl1 := num.NewRealized(-1234.555555555555)
	rl1.ChangeBase(10)

	rl2 := num.NewRealized(5)

	fmt.Println(rl1, rl2)
}

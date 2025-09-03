package main

import (
	"core/sys/num"
	"fmt"
)

func main() {
	for i := 0; i < 256; i++ {
		hex := fmt.Sprintf("%x", i)
		b := byte(hex[0])
		fmt.Println(hex + " - " + string(b))
	}

	a := num.NewReal(0.05)
	b := num.NewReal(1.0)
	c := num.NewReal(555)
	d := num.NewReal(-22)
	e := num.NewReal(-0.00042)
	f := num.NewReal(1024.04)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

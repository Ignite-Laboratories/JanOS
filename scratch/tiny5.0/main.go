package main

import "fmt"

func main() {
	x := Test{}
	y := 5
	x.Precision = &y
	fmt.Println(*x.Precision)
}

type Test struct {
	Precision *int
}

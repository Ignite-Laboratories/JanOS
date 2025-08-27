package main

import (
	"core/std"
	"fmt"
)

func main() {
	test(std.NewXYZ(1, 5, 3).SetName("Alpha"))
}

func test[T std.Vector](value T) {
	fmt.Println(value)
	fmt.Println(value.GetComponentByName("Z"))
}

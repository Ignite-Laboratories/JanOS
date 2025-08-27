package main

import (
	"core/std"
	"fmt"
)

func main() {
	xyz := std.NewXYZ(1, 5, 3).SetName("Alpha")
	v := test(xyz)
	fmt.Println(v)
}

func test[T std.Vector](value T) std.Vector {
	fmt.Println(value)

	z, _ := value.ComponentByName("Z")
	_ = z.SetUsingAny(2)
	fmt.Println(z)
	return value
}

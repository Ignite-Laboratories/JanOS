package main

import (
	"core/std"
	"core/sys/num"
	"fmt"
)

func main() {
	fma := std.NewFMA(num.MaxValue[float64](), -2.2, 3.3, "Alpha")
	fmt.Println(fma)

	xyz := std.NewXYZ(1, 5, 3, "Omega")
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

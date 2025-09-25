package mirror

import (
	"core/sys/num"
	"fmt"
	"strings"
)

type MyStruct[T1 any, T2 any] struct {
}

func (m MyStruct[T1, T2]) Foo() {
	fmt.Println("rawr")
}

func (m MyStruct[T1, T2]) RAWR(i num.Numeric[int], x string) string {
	fmt.Println("rawr")
}

func (m *MyStruct[T1, T2]) RAWR2(i int, x strings.Builder) {
	fmt.Println("rawr")
}

func (m *MyStruct[T1, T2]) Bar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func Baz() {
}

type MyOtherStruct[T any] MyStruct[T, T]

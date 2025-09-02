package mirror

import (
	"core/sys/num/bounded"
	"strings"
)

func (m MyOtherStruct[T][T1, T2]) Foo()	{ MyStruct(m).Foo() }

func (m MyOtherStruct[T][T1, T2]) RAWR(i bounded.Numeric[int], x string) string {
	return MyStruct(m).RAWR(i, x)
}

func (m *MyOtherStruct[T][T1, T2]) RAWR2(i int, x strings.Builder)	{ (*MyStruct(m)).RAWR2(i, x) }

func (m *MyOtherStruct[T][T1, T2]) Bar()	{ (*MyStruct(m)).Bar() }

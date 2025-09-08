package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("%f\n", 1e-100)
	fmt.Printf("%f\n", 1e-10)
	fmt.Printf("%f\n", 1e-1)
}

func Test(mixed ...any) {
	for _, a := range mixed {
		op := reflect.ValueOf(a)

		// Cases to handle:
		// 0. Structure types - panic
		// 1. Complex numbers - panic
		// 2. Functions that return a structure - panic
		// 3. Primitive types - Capture
		// 4. Pointers to primitive types - Capture
		// 5. Advanced types - Capture
		// 6. Pointers to advanced types - Capture
		// 7. Functions that return all Primitive/Advanced types - Capture
		// 8. Function chains that eventually return all Primitive/Advanced types - Capture

		if !op.IsValid() {
			panic("𝑡𝑖𝑛𝑦 only supports value or pointer type operands")
		}

	}
}

func main() {
	i := 10
	p := &i

	_, getVal := makeGetter(i) // non-pointer: snapshot
	_, getPtr := makeGetter(p) // pointer: live view

	fmt.Println(getVal()) // 10
	fmt.Println(getPtr()) // 10

	i = 42
	fmt.Println(getVal()) // still 10 (snapshot)
	fmt.Println(getPtr()) // 42 (reads through pointer)
}

func makeGetter(x any) (isPtr bool, get func() any) {
	v := reflect.ValueOf(x)
	if !v.IsValid() {
		return false, func() any { return nil }
	}
	if v.Kind() == reflect.Ptr {
		return true, func() any {
			if v.IsNil() {
				return nil
			}
			return v.Elem().Interface() // read current pointee
		}
	}
	// Not a pointer: return the copied value
	return false, func() any { return x }
}

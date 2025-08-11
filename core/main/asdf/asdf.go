package main

type MyStruct[T any] struct {
	MyField func() T
}

package recorded

import "github.com/ignite-laboratories/core/std"

// Any represents any recorded comparable data type, including std.Data.
//
// NOTE: 'context' is a simplified alias for Context.
type Any[T comparable] interface {
	*std.Data[T] | *Unique[T] | *context[T] | *Experience[T]
	Yield(uint) []T
	Append(T)
	Prepend(T)
	Insert(uint, T) error
	Remove(uint) error
	Select(uint, ...uint) ([]T, error)
	SelectAll() []T
}

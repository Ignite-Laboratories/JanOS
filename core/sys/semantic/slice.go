package semantic

import (
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/semantic/emit"
	"sync"
)

type Record[T any] struct{}

func (y Record[T]) Append(fn func() []T) {

}

func (y Record[T]) Prepend(fn func() []T) {

}

func (y Record[T]) Insert(index uint, fn func() []T) error {
	return nil
}

type Slice[T any] struct {
	data []T
	cap  uint64

	Record Record[T]
}

func (r Slice[T]) Yield(fn YieldFn[T]) {

}

func test() {
	a := Slice[int]{}
	a.Yield(a.Rawr)
	a.Yield(emit.Rawr[int])
}

func (r Slice[T]) Rawr(slice Slice[T]) (mutated Slice[T], artifacts Slice[T]) {

}

//
// slice.Selection(low, ...high)
// slice.RandomSelection(...count)
// slice.All()
// slice.Reverse()
// slice.First()
// slice.Last()
// slice.Append(data)
// slice.Prepend(data)
// slice.Insert(index, data)
// slice.Remove(indices)
//
// slice.Slice.Append(func() ([]T, bool))
// slice.Slice.Prepend(func() ([]T, bool))
// slice.Slice.Insert(index, func() ([]T, bool))

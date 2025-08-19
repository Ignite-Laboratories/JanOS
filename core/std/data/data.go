package data

import (
	"github.com/ignite-laboratories/core/std"
	"sync"
)

// Std is a way to pass slices around in anonymous closures without any worries.  For instance, if you pass a slice as a
// parameter to a function that generates an anonymous method which closed on the slice, appends to the originating slice
// will not persist through to the anonymous method's closure variable.  This is because the anonymous method has closed
// on a COPY of a header pointing to the data, which Go's 'append' method then reassigned outside the anonymous method.
//
// By unifying standardized data in a wrapper structure, you can colloquially guarantee asynchronous access to dimensional
// data by gating append operations through method calls and a mutex - which is exactly what this type provides for you =)
type Std[T any] struct {
	mutex  sync.Mutex
	values []T
}

// Contextual is a kind of data.Std[T] that associates every entry with at least a moment in time.
type Contextual[TData any, TContext std.Moment[T]] Std[TContext]

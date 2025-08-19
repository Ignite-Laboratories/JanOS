package recorded

import "github.com/ignite-laboratories/core/std"

type Context[T any, TContext std.Instantaneous[T]] = context[T]
type context[T any] struct{}

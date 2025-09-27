package std

import "sync"

type Thought struct {
	Realization any
	Gate        sync.Mutex
}

func NewThought(realization any) *Thought {
	return &Thought{Realization: realization}
}

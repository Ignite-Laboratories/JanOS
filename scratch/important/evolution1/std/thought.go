package std

import "sync"

type Thought struct {
	Revelation any
	Gate       *sync.Mutex
}

func NewThought(revelation any) *Thought {
	return &Thought{Revelation: revelation}
}

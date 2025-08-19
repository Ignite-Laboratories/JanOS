package recorded

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/num"
	"sync"
)

// Unique is a way to request periodically unique random data.  You may either "seed" the set with a finite set of
// comparable values or generate random num.Primitive numbers bounded within a range.  Whenever you request a 'random' value
// from the set, it will guarantee you a random entry until it exhausts the available uniqueness.  At which point, it
// will restart the cycle until it exhausts it again while retaining all emitted values in logical order for later retrieval.
//
// For example, given a closed interval of four values -
//
//		let 𝑎 = [0,3]
//
//	 new Unique[T](𝑎).Random(10):
//	 ℕ[0](0.0) - 3
//	 ℕ[1](0.1) - 3 1
//	 ℕ[2](0.2) - 3 1 0
//	 ℕ[3](0.3) - 3 1 0 2 ← Exhaustion point
//	 ℕ[4](1.0) - 3 1 0 2 | 1
//	 ℕ[5](1.1) - 3 1 0 2 | 1 3
//	 ℕ[6](1.2) - 3 1 0 2 | 1 3 2
//	 ℕ[7](1.3) - 3 1 0 2 | 1 3 2 0 ← Exhaustion point
//	 ℕ[8](2.0) - 3 1 0 2 | 1 3 2 0 | 0
//	 ℕ[9](2.1) - 3 1 0 2 | 1 3 2 0 | 0 1 ← Resulting Unique Set
//
// Abstractly -
// In a traditional mathematical environment, you must consider the presence of 'infinity' - but in a programmatic
// one, you may consider that all num.Primitive types are bounded by their inherent bit-width.  Color spaces, for instance,
// are often bounded to ranges below 255 - meaning it wouldn't be long before you began repeating values. By tracking the
// order of emission and cyclically ensuring uniqueness, this type affords you as evenly distributed of a unique set of
// values as I can imagine.
//
// NOTE: Unique is 'resettable' by default, meaning a call to Reset will clear its contents.  You may override this
// by calling SetResettable if you wish to prevent destruction of existing data.
type Unique[T any] struct {
	generator  func() T
	size       uint64
	ordered    []T
	numeric    bool
	resettable bool
	entries    []map[any]struct{}
	mutex      sync.Mutex
}

// UniqueBounded creates a new numeric unique set bounded by the provided Bounded[T].
//
// NOTE: You may optionally designate if the set is resettable on creation.
func UniqueBounded[T num.Primitive](bounds std.Bounded[T], resettable ...bool) *Unique[T] {
	r := true
	if len(resettable) > 0 {
		r = resettable[0]
	}
	return &Unique[T]{
		generator: func() T {
			return num.RandomWithinRange[T](bounds.Minimum(), bounds.Maximum())
		},
		size:       bounds.Range(),
		ordered:    make([]T, 0),
		numeric:    true,
		resettable: r,
		entries:    make([]map[any]struct{}, 0, bounds.Range()),
	}
}

func deduplicate[T any](data []T) []T {
	seen := make(map[any]struct{}, len(data))
	unique := make([]T, 0, len(data))
	for _, v := range data {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		unique = append(unique, v)
	}
	return unique
}

// UniqueSeeded creates a new unique set from the provided data.
//
// NOTE: This will automatically weed out duplicate entries.
func UniqueSeeded[T comparable](data ...T) *Unique[T] {
	if len(data) == 0 {
		panic("cannot create a unique set without any data to seed it with")
	}

	data = deduplicate(data)

	bounds, _ := std.NewBounded[uint](0, 0, uint(len(data)-1))
	return &Unique[T]{
		generator: func() T {
			return data[num.RandomWithinRange[uint](bounds.Minimum(), bounds.Maximum())]
		},
		size:    uint64(len(data) - 1),
		ordered: make([]T, 0),
		entries: make([]map[any]struct{}, 0, len(data)-1),
	}
}

// SetResettable sets whether a call to Reset will clear the set's contents.  True allows the set contents to be reset,
// while false will prevent calls to Reset from clearing any existing content.
func (s *Unique[T]) SetResettable(resettable bool) *Unique[T] {
	s.resettable = resettable
	return s
}

// Reseed calls Reset and then re-seeds the unique set with the provided data.
//
// NOTE: For a non-seeded (bounded) set, this only calls Reset.
func (s *Unique[T]) Reseed(data ...T) {
	s.Reset()
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.numeric {
		data = deduplicate(data)

		bounds, _ := std.NewBounded[uint](0, 0, uint(len(data)-1))
		s.generator = func() T {
			return data[num.RandomWithinRange[uint](bounds.Minimum(), bounds.Maximum())]
		}
		s.entries = make([]map[any]struct{}, 0, len(data)-1)
		s.ordered = make([]T, 0)
		s.size = uint64(len(data) - 1)
	}
}

// Entries returns the current ordered collection of randomly selected entries.
func (s *Unique[T]) Entries() []T {
	return s.ordered
}

// Reset clears all entries and starts anew.
//
// NOTE: This will do nothing if the set has been marked as 'not resettable' through SetResettable and sets are resettable by default.
func (s *Unique[T]) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.resettable {
		s.entries = make([]map[any]struct{}, 0, s.size)
		s.ordered = make([]T, 0)
	}
}

// Modulate either shrinks the number of entries or grows it to the requested length.
//
// NOTE: This can be a destructive operation and is -NOT- gated through Reset!  Be cautious =)
func (s *Unique[T]) Modulate(length uint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	curr := uint(len(s.ordered))

	switch {
	case length < curr:
		trim := s.ordered[length:]
		s.ordered = s.ordered[:length]

		for _, val := range trim {
			for i := len(s.entries) - 1; i >= 0; i-- {
				if _, ok := s.entries[i][val]; ok {
					delete(s.entries[i], val)
					break
				}
			}
		}

		for len(s.entries) > 0 && len(s.entries[len(s.entries)-1]) == 0 {
			s.entries = s.entries[:len(s.entries)-1]
		}
	case length > curr:
		s.Random(length - curr)
	}
}

// Random returns an evenly distributed slice of random numbers in the Unique's Bounded range.
//
// NOTE: If no count is provided, a single value is yielded.
func (s *Unique[T]) Random(count ...uint) []T {
	c := 1
	if len(count) > 0 {
		c = int(count[0])
	}

	if s.entries == nil {
		s.entries = make([]map[any]struct{}, 0)
	}
	if s.ordered == nil {
		s.ordered = make([]T, 0)
	}
	if len(s.entries) == 0 {
		s.entries = append(s.entries, make(map[any]struct{}, 0))
	}

	out := make([]T, c)
	for i := 0; i < c; i++ {
		if uint64(len(s.entries[len(s.entries)-1])) >= s.size {
			s.mutex.Lock()
			s.entries = append(s.entries, make(map[any]struct{}, 0))
			s.mutex.Unlock()
		}

		var val T
		for {
			val = s.generator()

			if _, ok := s.entries[len(s.entries)-1][val]; !ok {
				s.mutex.Lock()
				s.entries[len(s.entries)-1][val] = struct{}{}
				out[i] = val
				s.mutex.Unlock()
				break
			}
		}
	}

	s.mutex.Lock()
	s.ordered = append(s.ordered, out...)
	s.mutex.Unlock()
	return out
}

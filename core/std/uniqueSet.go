package std

import (
	"github.com/ignite-laboratories/core/std/num"
)

// UniqueSet is a kind of numeric set that guarantees uniqueness using periodicity.  This type will produce a set of
// uniquely random numbers from a bounded range whenever asked.  Once it exhausts the available unique values, it
// resets and begins a new round of guaranteed uniqueness - while retaining the order of all emitted values.
//
// For example, given a closed interval of four values -
//
//		let 𝑎 = [0,3]
//		let gen: ℕ → 𝑎^𝑛
//
//	 gen(10) output:
//	 ℕ[0](0.0) - 3
//	 ℕ[1](0.1) - 3 1
//	 ℕ[2](0.2) - 3 1 0
//	 ℕ[3](0.3) - 3 1 0 2 ← Exhaustion point
//	 ℕ[4](1.0) - 1
//	 ℕ[5](1.1) - 1 3
//	 ℕ[6](1.2) - 1 3 2
//	 ℕ[7](1.3) - 1 3 2 0 ← Exhaustion point
//	 ℕ[8](2.0) - 0
//	 ℕ[9](2.1) - 0 1
//
// Abstractly -
// In a traditional mathematical environment, you must consider the presence of 'infinity' - but in a programmatic
// one, you may consider that all num.Primitive types are bounded by their inherent bit-width.  Color spaces, for instance,
// are often bounded to ranges below 255 - meaning it wouldn't be long before you began repeating values. By tracking the
// order of emission and cyclically ensuring uniqueness, this type affords you as evenly distributed of a unique set of
// values as I can imagine.
//
// See seed.Random
type UniqueSet[T num.Primitive] struct {
	Bounded[T]
	entries []map[T]struct{}
	ordered []T
	current uint64
}

// Reset clears all entries and starts anew.
func (s *UniqueSet[T]) Reset() {
	s.current = 0
	s.entries = make([]map[T]struct{}, 0)
	s.ordered = make([]T, 0)
}

// Modulate either shrinks the number of entries or grows it to the requested length.
func (s *UniqueSet[T]) Modulate(length uint) {
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

// Random returns an evenly distributed slice of random numbers in the UniqueSet's Bounded range.
//
// NOTE: If no count is provided, a single value is yielded.
func (s *UniqueSet[T]) Random(count ...uint) []T {
	c := 1
	if len(count) > 0 {
		c = int(count[0])
	}

	if s.entries == nil {
		s.entries = make([]map[T]struct{}, 0)
	}
	if s.ordered == nil {
		s.ordered = make([]T, 0)
	}
	if len(s.entries) == 0 {
		s.entries = append(s.entries, make(map[T]struct{}, 0))
	}

	out := make([]T, c)
	for i := 0; i < c; i++ {
		if uint64(len(s.entries[len(s.entries)-1])) >= s.Range() {
			s.entries = append(s.entries, make(map[T]struct{}, 0))
		}

		var val T
		for {
			val = num.RandomWithinRange[T](s.minimum, s.maximum)

			if _, ok := s.entries[len(s.entries)-1][val]; !ok {
				s.current++
				s.entries[len(s.entries)-1][val] = struct{}{}
				out[i] = val
				break
			}
		}
	}

	s.ordered = append(s.ordered, out...)
	return out
}

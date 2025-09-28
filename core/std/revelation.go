package std

import "time"

type Revelation[T any] struct {
	*TemporalBuffer[T]

	reveal func(last time.Time) T
	last   time.Time
}

func NewRevelation[T any](reveal func(time.Time) T, window ...*time.Duration) *Revelation[T] {
	if reveal == nil {
		panic("reveal function must not be nil")
	}

	return &Revelation[T]{
		TemporalBuffer: NewTemporalBuffer[T](window...),
		reveal:         reveal,
		last:           time.Now(),
	}
}

func (r *Revelation[T]) Reveal() T {
	if r.reveal == nil {
		panic("please create a revelation through std.NewRevelation[T]()")
	}

	result := r.reveal(r.last)
	now := time.Now()
	r.Record(now, result)
	r.last = now
	return result
}

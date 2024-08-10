package Arwen

import "gonum.org/v1/gonum/mat"

// Sampler is used to record a snapshot of readings across time.
type Sampler struct {
	Dimensions []*Dimension
	Stride     int
	Position   int
	Data       *mat.Dense
	OnSample   func(dense mat.Dense)
}

func NewSampler(stride int, onSample func(dense mat.Dense), dimensions ...*Dimension) *Sampler {
	data := make([]float64, len(dimensions)*stride)
	s := &Sampler{
		Dimensions: dimensions,
		Stride:     stride,
		Position:   0,
		Data:       mat.NewDense(len(dimensions), stride, data),
		OnSample:   onSample,
	}

	return s
}

// Tick samples the current value of its referenced dimensions.
func (s *Sampler) Tick() {
	for i, d := range s.Dimensions {
		s.Data.Set(i, s.Position, d.Value)
	}

	s.Position++
	if s.Position == s.Stride {
		s.Position = 0
		s.OnSample(*s.Data)
	}
}

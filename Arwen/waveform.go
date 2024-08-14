package Arwen

import (
	"log"
	"math"
	"time"
)

type Sample struct {
	Step time.Duration
	Data []float64
}

type waveform struct {
	playing  bool
	headTime time.Time
	data     Sample
}

func NewWaveform(data []float64, step time.Duration) *waveform {
	return &waveform{
		data: Sample{
			Step: step,
			Data: data,
		},
	}
}

func (w *waveform) Play() {
	if w.playing {
		log.Panic("waveform is already playing")
	}
	w.headTime = time.Now()
	w.playing = true
}

func (w *waveform) Stop() {
	w.playing = false
}

func (w *waveform) Sample(start time.Time, duration time.Duration) Sample {
	nowOffset := start.Sub(w.headTime)
	maxLen := float64(len(w.data.Data))
	nowIndex := float64(nowOffset / w.data.Step)
	offsetIndex := float64(duration / w.data.Step)

	// get the smaller index, or 0 if negative
	smallerIndex := int(math.Max(math.Min(nowIndex, offsetIndex), 0))
	// get the larger index, or the max index if overflowing
	largerIndex := int(math.Min(math.Max(nowIndex, offsetIndex), maxLen))

	return Sample{
		Step: w.data.Step,
		Data: w.data.Data[smallerIndex:largerIndex],
	}
}

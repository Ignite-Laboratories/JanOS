package JanOS

import (
	"log"
	"sync"
	"time"
)

/**
RESOLUTION
*/

// NOTE:
// This system is meant to be performant and has been very carefully exposed.
// We intentionally fail hard - if you are trying to look too far into the future or past then
// you need to restructure your model!  Look at the Universe.BufferLength to know what the practical
// boundaries are in your simulation and modify it if you need a different stretch of time.

// Resolution represents multiple ways to represent an array of data.
// Frequency is measured in hertz (the universal default is 44,000hz)
// Nanoseconds represents the amount of nanoseconds that a single frequency period takes to pass.
// Duration represents the amount of time that a single frequency period takes to pass.
type Resolution struct {
	Frequency   int
	Nanoseconds int64
	Duration    time.Duration
}

// NewResolution creates a structure used to represent the fidelity of time.
// The provided frequency is in hertz (the universal default is 44,000hz)
func NewResolution(frequency int) Resolution {
	r := Resolution{
		Frequency: frequency,
	}
	r.Nanoseconds = int64(float64(time.Second.Nanoseconds()) / float64(r.Frequency))
	r.Duration = time.Duration(r.Nanoseconds + 1)
	return r
}

// SetResolution updates the timeline's Resolution.  Any existing data on the timeline
// will be read in the new Resolution interval - making this essentially stretch or
// shrink the existing timeline.  As such, the best practice is to set this before
// starting to put data on the timeline - but that doesn't mean you shouldn't try!
func (tl *timeline) SetResolution(frequency int) {
	tl.lock.Lock()
	defer tl.lock.Unlock()
	calculatedResolution := NewResolution(frequency)
	tl.resolution.Frequency = calculatedResolution.Frequency
	tl.resolution.Nanoseconds = calculatedResolution.Nanoseconds
	tl.resolution.Duration = calculatedResolution.Duration
}

// GetResolution returns the current Resolution information for the provided timeline.
func (tl *timeline) GetResolution() Resolution {
	return tl.resolution
}

// ToIndex returns an index representation of the duration of time provided.
func (r *Resolution) ToIndex(duration time.Duration) int {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / float64(r.Frequency)
	// Divide the provided duration's nanoseconds by the step size to get the index count
	return int(float64(duration.Nanoseconds()) / nanoStep)
}

// ToDuration returns a time.Duration representation of the amount of indices provided
// Check Universe.Resolution for the operational frequency.
func (r *Resolution) ToDuration(steps int) time.Duration {
	// Calculate the amount of nanoseconds per index
	nanoStep := float64(time.Second.Nanoseconds()) / float64(r.Frequency)
	// Multiply the amount of nanoseconds per step by the number of steps to get the duration in nanoseconds
	return time.Duration(nanoStep * float64(steps))
}

/**
POINT VALUE
*/

// PointValue represents a value and its derivative (if available) at a single point in time.
type PointValue struct {
	Value      float64
	Derivative float64
}

// CalculateDerivative takes a slice of float64 data and converts it to a PointValue array
func CalculateDerivative(data []float64) []PointValue {
	toReturn := make([]PointValue, len(data))
	if len(data) == 0 {
		return toReturn
	}

	lastVal := data[0]
	for i, val := range data {
		toReturn[i] = PointValue{val, val - lastVal}
		lastVal = val
	}
	return toReturn
}

/**
INSTANTANEOUS VALUE
*/

// InstantaneousValue represents a value at a specified point in time.
type InstantaneousValue struct {
	Instant time.Time
	Point   PointValue
}

/**
TIME SLICE
*/

// TimeSlice represents a slice of data relative to a moment in time.
type TimeSlice struct {
	StartTime  time.Time
	Data       []PointValue
	Resolution Resolution
}

// Integrate takes a TimeSlice and for every index adds its value.
// This gives us the area under the curve at the Resolution that
// the TimeSlice was recorded at - which acts as an integral.
func (ts TimeSlice) Integrate() float64 {
	area := 0.0
	for _, val := range ts.Data {
		area += val.Value
	}
	return area
}

// GetTotalDuration gets the total time.Duration of the provided TimeSlice
func (ts TimeSlice) GetTotalDuration() time.Duration {
	return time.Duration(ts.Resolution.Duration.Nanoseconds() * int64(len(ts.Data)))
}

// DownSample returns back a different resolute version of an existing TimeSlice.
// In down-sampling the data is destructively lost - you can only interpolate it
// to up-sample again.
func (ts TimeSlice) DownSample(frequency int) TimeSlice {
	stride := int(float64(time.Second.Nanoseconds()) / float64(frequency))

	toReturn := TimeSlice{
		StartTime:  ts.StartTime,
		Data:       make([]PointValue, 0),
		Resolution: NewResolution(frequency),
	}

	for i, _ := range ts.Data {
		// if we are at a sample index...
		if i == 0 || i%stride == 0 {
			tsIndex := i * stride
			if tsIndex < len(ts.Data) {
				toReturn.Data = append(toReturn.Data, ts.Data[i*stride])
			}
		}
	}
	return toReturn
}

// UpSample returns back a different resolute version of an existing TimeSlice.
// In up-sampling it performs linear interpolation on the values between known
// points in time.
func (ts TimeSlice) UpSample(frequency int) TimeSlice {
	r := NewResolution(frequency)
	indexCount := r.ToIndex(ts.GetTotalDuration())
	result := make([]float64, indexCount)

	// First, we apply the old values to their found indices in the new Resolution
	stride := indexCount / (len(ts.Data) - 1)
	for i, _ := range ts.Data {
		destinationIndex := stride * i
		result[destinationIndex] = ts.Data[i].Value
	}

	log.Println(result)

	i := 0
	for i < len(result) {
		value := result[i]
		// Act whenever the data changes
		if i == 0 || value != 0 {
			nextValue := 0.0
			stepsForward := 0
			nextSampleIndex := i
			// From i...do we have new data?
			for nextValue == 0 && stepsForward+i < len(result)-1 { // (unless we reach the end of the data)
				// ...walk forward
				stepsForward++
				nextSampleIndex++
				// ...grab its value
				nextValue = result[nextSampleIndex]
			}

			delta := nextValue - value
			for x := 0; x < stepsForward; x++ {
				scaleFactor := float64(x) / float64(stepsForward)
				newValue := value + (delta * scaleFactor)
				result[i] = newValue
				i++
			}
		} else {
			i++
		}
	}

	return TimeSlice{
		StartTime:  ts.StartTime,
		Data:       CalculateDerivative(result),
		Resolution: r,
	}
}

func (ts TimeSlice) Mux(formula Formula, otherSlices ...TimeSlice) TimeSlice {
	outputSlice := TimeSlice{
		StartTime:  ts.StartTime,
		Data:       make([]PointValue, len(ts.Data)),
		Resolution: ts.Resolution,
	}
	// Copy off the original slice data so this operation is non-destructive to it
	copy(outputSlice.Data, ts.Data)

	oldVal := 0.0
	// For each entry in the original slice...
	for x, indexValue := range outputSlice.Data {
		// Save off it's value
		newVal := indexValue.Value

		// And get the abstract duration to this index
		passedDuration := outputSlice.Resolution.ToDuration(x)

		// Then loop over the other slices...
		for _, otherSlice := range otherSlices {
			// Get the appropriate index from the other slice
			otherIndex := otherSlice.Resolution.ToIndex(passedDuration)

			// If we can read that index...
			if otherIndex < len(otherSlice.Data) {
				// ...Update the value by performing the formula on the two values
				newVal = formula.Operation(newVal, otherSlice.Data[otherIndex].Value)
			}
		}

		// Update the new value and calculate its derivative
		outputSlice.Data[x] = PointValue{newVal, newVal - oldVal}
		oldVal = newVal
	}

	return outputSlice
}

/**
TIME LINE
*/

// timeline represents rolling data in time.
// headTime is the referential point the data currently starts at
// Resolution is the details of how to relate time a single index value
// lock is used to ensure the buffer data is only manipulated in scoped contexts.
// (i.e. any operation that modifies the data buffer should request a lock on it while running)
type timeline struct {
	data       []float64
	headTime   time.Time
	resolution Resolution
	lastValue  float64
	lock       sync.Mutex
}

// GetRelativeIndex takes a moment in time and gets its index, relative to the current head time.
func (tl *timeline) GetRelativeIndex(t time.Time) int {
	return tl.resolution.ToIndex(t.Sub(tl.headTime))
}

// GetInstant returns the value on the timeline at a moment in time.
func (tl *timeline) GetInstant(instant time.Time) InstantaneousValue {
	instantIndex := tl.GetRelativeIndex(instant)
	if instantIndex <= len(tl.data) {
		value := tl.data[instantIndex]
		return InstantaneousValue{
			Instant: instant,
			Point:   PointValue{value, 0},
		}
	}
	return InstantaneousValue{
		Instant: instant,
		Point:   PointValue{0, 0},
	}
}

// SliceEntireFuture returns the remainder of the buffer from the provided instant in time.
func (tl *timeline) SliceEntireFuture(instant time.Time) TimeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	if instantIndex <= len(tl.data) {
		return TimeSlice{
			StartTime:  instant,
			Data:       CalculateDerivative(tl.data[instantIndex:]),
			Resolution: tl.resolution,
		}
	}
	return TimeSlice{
		StartTime:  instant,
		Data:       make([]PointValue, 0),
		Resolution: tl.resolution,
	}
}

// SliceEntirePast returns the entire buffer up to the provided instant in time.
func (tl *timeline) SliceEntirePast(instant time.Time) TimeSlice {
	// We capture the head time here to ensure all calculations
	// are relative to the execution of this line of code in time.
	headTime := tl.headTime
	headIndex := tl.GetRelativeIndex(headTime)
	instantIndex := tl.GetRelativeIndex(instant)
	if instantIndex <= len(tl.data) {
		return TimeSlice{
			StartTime:  headTime,
			Data:       CalculateDerivative(tl.data[headIndex:instantIndex]),
			Resolution: tl.resolution,
		}
	}
	return TimeSlice{
		StartTime:  instant,
		Data:       make([]PointValue, 0),
		Resolution: tl.resolution,
	}
}

// SliceFutureDuration returns a slice of the future from an instant in time utilizing a time.Duration.
func (tl *timeline) SliceFutureDuration(instant time.Time, duration time.Duration) TimeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	futureIndex := instantIndex + tl.resolution.ToIndex(duration)
	if futureIndex <= len(tl.data) {
		return TimeSlice{
			StartTime:  instant,
			Data:       CalculateDerivative(tl.data[instantIndex:futureIndex]),
			Resolution: tl.resolution,
		}
	}
	return TimeSlice{
		StartTime:  instant,
		Data:       make([]PointValue, 0),
		Resolution: tl.resolution,
	}
}

// SliceFutureIndices returns an indexCount length slice of the future from an instant in time.
func (tl *timeline) SliceFutureIndices(instant time.Time, indexCount int) TimeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	if instantIndex <= len(tl.data) {
		return TimeSlice{
			StartTime:  instant,
			Data:       CalculateDerivative(tl.data[instantIndex : instantIndex+indexCount]),
			Resolution: tl.resolution,
		}
	}
	return TimeSlice{
		StartTime:  instant,
		Data:       make([]PointValue, 0),
		Resolution: tl.resolution,
	}
}

// SlicePastDuration returns an indexCount length slice of the past up to an instant in time.
func (tl *timeline) SlicePastDuration(instant time.Time, duration time.Duration) TimeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	pastIndex := instantIndex - tl.resolution.ToIndex(duration)
	if instantIndex <= len(tl.data) {
		return TimeSlice{
			StartTime:  instant.Add(-duration),
			Data:       CalculateDerivative(tl.data[pastIndex:instantIndex]),
			Resolution: tl.resolution,
		}
	}
	return TimeSlice{
		StartTime:  instant,
		Data:       make([]PointValue, 0),
		Resolution: tl.resolution,
	}
}

// SlicePastIndices returns a indexCount length slice of the past up to an instant in time.
func (tl *timeline) SlicePastIndices(instant time.Time, indexCount int) TimeSlice {
	instantIndex := tl.GetRelativeIndex(instant)
	if instantIndex <= len(tl.data) {
		return TimeSlice{
			StartTime:  instant,
			Data:       CalculateDerivative(tl.data[instantIndex-indexCount : instantIndex]),
			Resolution: tl.resolution,
		}
	}
	return TimeSlice{
		StartTime:  instant,
		Data:       make([]PointValue, 0),
		Resolution: tl.resolution,
	}
}

// AddValues seeks to the appropriate position in time and additively introduces the provided data to the buffer
func (tl *timeline) AddValues(instant time.Time, data ...float64) {
	tl.lock.Lock()
	defer tl.lock.Unlock()
	startIndex := tl.GetRelativeIndex(instant)

	// If the buffer has no data between the end of its current prediction and the start of this data...
	tlDataLen := len(tl.data)
	if startIndex >= tlDataLen {
		// ...fill the gap with initialized data to the current value of the timeline
		toFill := startIndex - tlDataLen
		initialized := NewInitializedArray(tl.data[len(tl.data)-1], toFill)
		tl.data = append(tl.data, initialized...)
	}

	for x := 0; x < len(data); x++ {
		currentIndex := startIndex + x
		// If we are out of bounds of the underlying buffer...
		if currentIndex > len(tl.data) {
			// ...append the remaining data to the buffer and exit the loop
			toAppend := data[x:]
			tl.data = append(tl.data, toAppend...)
			break
		}

		// Otherwise, just add the data to the existing data
		newVal := tl.data[currentIndex] + data[x]
		tl.data[currentIndex] = newVal
	}
}

// setValue seeks to the appropriate position in time and replaces the values on the remainder of the buffer.
// It is not exposed because you should always set the value of the signal from the signal - this is a method
// that keeps the  timeline logic together.  It keeps from spreading its implementation details into the signal
// code through which you are intended to access it.
func (tl *timeline) setValue(instant time.Time, value float64) {
	tl.lock.Lock()
	defer tl.lock.Unlock()
	tl.setValueWithoutLock(instant, value)
}

func (tl *timeline) setValueWithoutLock(instant time.Time, value float64) {
	startIndex := tl.GetRelativeIndex(instant)

	// Buffer the future timeline with its last known value if it isn't long enough
	if startIndex > len(tl.data) {
		toFill := startIndex - len(tl.data)
		lastValue := tl.data[len(tl.data)-1]
		tl.data = append(tl.data, NewInitializedArray(lastValue, toFill)...)
	}

	// Otherwise, set the value on the timeline
	result := value

	// Then grab the remaining length of the timeline...
	remainingLen := len(tl.data) - startIndex
	// ...and initialize it to the new value
	tl.data = append(tl.data[:startIndex], NewInitializedArray(result, remainingLen)...)
}

// newTimeline creates a timeline buffer.
// The duration represents the total amount of time to buffer, with now being considered relative to
// the midpoint of the buffer.  The frequency tells it how often to trim/append the buffer in time.
// onStep can be nil if you do not need that functionality.
func (signal *Signal) newTimeline(defaultValue float64, onStep func(*Signal, InstantaneousValue) float64) *timeline {
	nanoStep := float64(time.Second.Nanoseconds()) / float64(Universe.StdResolution)
	durationInIndex := int(float64(Universe.StdBufferLength.Nanoseconds()) / nanoStep)
	tl := &timeline{
		// Set the head time relative to the midpoint of the buffer
		// (we initialize with the past being empty, essentially)
		headTime:   time.Now().Add(-(Universe.StdBufferLength / 2)),
		data:       NewInitializedArray(defaultValue, durationInIndex),
		resolution: NewResolution(Universe.StdResolution),
	}

	// Spin off the loop for this timeline
	go func() {
		lastUpdate := time.Now()
		for {
			if Universe.Terminate {
				break
			}

			// Save off now so the logic operates on the same instant during this routine
			now := time.Now()
			elapsedTime := now.Sub(lastUpdate)

			// If we are above the resolution time for this timeline...
			if elapsedTime > tl.resolution.Duration {
				tl.lock.Lock()
				// Calculate new head time
				newHead := now.Add(-(Universe.StdBufferLength / 2))
				delta := tl.GetRelativeIndex(newHead)
				// Update the head time
				tl.headTime = newHead
				// Slice off the past that has passed
				if delta < len(tl.data) {
					tl.data = tl.data[delta:]
				}
				// Append new data to replace the old data
				newData := NewInitializedArray(tl.data[len(tl.data)-1], delta)
				tl.data = append(tl.data, newData...)

				// ...Get the current value of this signal and calculate it's derivative
				currentInstant := tl.GetInstant(now)
				currentValue := currentInstant.Point.Value
				currentInstant.Point.Derivative = currentValue - tl.lastValue

				// If there is a function to perform on this logical step in time...
				if onStep != nil {
					// ...send off the calculated data
					currentValue = onStep(signal, currentInstant)

					// ...update the timeline with the calculated result from onStep
					tl.setValueWithoutLock(now, currentValue)

					// -------- NOTE:
					// The derivative is only relevant to the observer's perspective!
					// That means that we are ensuring that any function that operates in the
					// current -chain- will know the appropriate derived change, but any external
					// observers (such as samplers) will gain a different perspective of the
					// derivative of this signal.  This is a part of temporality and observation
					// of numbers within space and time.  The value intrinsically exists, but it
					// may have a different representation to different observers based on the
					// context with which they view the value.
				}

				// Save off the loop logic values
				tl.lastValue = currentValue
				lastUpdate = now
				tl.lock.Unlock()
			}
			time.Sleep(time.Nanosecond * 100)
		}
		Universe.Printf(signal, "%s timeline stopped", string(signal.Symbol))
	}()

	return tl
}

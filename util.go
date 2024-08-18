package JanOS

// SelectUnique returns the unique values found by the predicate
func SelectUnique[TIn any, TOut any](data []TIn, predicate func(TIn) TOut) []TOut {
	return Unique(Select(data, predicate))
}

// Select returns the values that the predicate matches
func Select[TIn any, TOut any](data []TIn, predicate func(TIn) TOut) []TOut {
	output := make([]TOut, len(data))
	for i := range data {
		output[i] = predicate(data[i])
	}
	return output
}

// Unique returns the unique values in a slice of data.
func Unique[T any](data []T) []T {
	keys := make(map[any]bool)
	var list []T
	for _, entry := range data {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// MapContains returns true if the provided map key exists.
func MapContains[T any](data map[string]T, key string) bool {
	for k, _ := range data {
		if k == key {
			return true
		}
	}
	return false
}

// DiffSigns returns true if the two values have different signs.
func DiffSigns(a int, b int) bool {
	return a*b < 0
}

// GetLargest returns the largest found value in a slice of data.
func GetLargest(s []float64) float64 {
	maxValue := s[0]
	for _, v := range s {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// NewInitializedArray returns a new array of the provided size with
// all values initialized to the provided value.
func NewInitializedArray[T any](value T, size int) []T {
	newArray := make([]T, size)
	for i := range newArray {
		newArray[i] = value
	}
	return newArray
}

func SpacedStringSet(delimeter string, values ...string) string {
	if len(values) == 0 {
		return ""
	}
	toReturn := values[0]
	i := 1
	for i < len(values) {
		toReturn = toReturn + " " + delimeter + " " + values[i]
		i++
	}
	return toReturn
}

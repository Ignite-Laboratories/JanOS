package Logic

func SelectUnique[TIn any, TOut any](data []TIn, selector func(TIn) TOut) []TOut {
	return Unique(Select(data, selector))
}

func Select[TIn any, TOut any](data []TIn, selector func(TIn) TOut) []TOut {
	output := make([]TOut, len(data))
	for i := range data {
		output[i] = selector(data[i])
	}
	return output
}

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

func MapContains[T any](data map[string]T, entry string) bool {
	for key, _ := range data {
		if key == entry {
			return true
		}
	}
	return false
}

package Util

func DiffSigns(a int, b int) bool {
	return a*b < 0
}

func GetLargest(s []int) int {
	maxValue := s[0]
	for _, v := range s {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}
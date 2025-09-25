package internal

import "fmt"

// TODO: delete this
func shouldRound(digit byte, base uint16) bool {
	mid := byte(base / 2)
	return digit > mid
}

func PrintDigit(digit byte) string {
	return fmt.Sprintf("%02x", digit)
}

package transcendental

// Number represents a transcendental number, which is a value precalculated to atlas.Precision placeholder
// width - allowing 𝑡𝑖𝑛𝑦 to efficiently store very large commonly used numbers without each Real holding a copy of
// all digits.
//
// See Non, E, and Pi
type Number string

const (
	// Number represents a transcendental number, which is a value identified as known to atlas.Precision
	// width - allowing 𝑡𝑖𝑛𝑦 to efficiently store very large commonly used numbers.
	//
	// See Non, E, and Pi
	Non Number = ""

	// Number represents a transcendental number, which is a value identified as known to atlas.Precision
	// width - allowing 𝑡𝑖𝑛𝑦 to efficiently store very large commonly used numbers.
	//
	// See Non, E, and Pi
	E Number = "ℯ"

	// Number represents a transcendental number, which is a value identified as known to atlas.Precision
	// width - allowing 𝑡𝑖𝑛𝑦 to efficiently store very large commonly used numbers.
	//
	// See Non, E, and Pi
	Pi Number = "π"
)

func IsIdentifier(char string) Number {
	switch char {
	case "ℯ":
		return E
	case "π":
		return Pi
	default:
		return Non
	}
}

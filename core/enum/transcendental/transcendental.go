package transcendental

// Transcendental represents a transcendental number, which is a value precalculated to atlas.Precision placeholder
// width - allowing 𝑡𝑖𝑛𝑦 to efficiently store very large commonly used numbers without each Real holding a copy of
// all digits.
//
// See Non, E, and Pi
type Transcendental string

const (
	// Transcendental represents a transcendental number, which is a value identified as known to atlas.Precision
	// width - allowing 𝑡𝑖𝑛𝑦 to efficiently store very large commonly used numbers.
	//
	// See Non, E, and Pi
	Non Transcendental = ""

	// Transcendental represents a transcendental number, which is a value identified as known to atlas.Precision
	// width - allowing 𝑡𝑖𝑛𝑦 to efficiently store very large commonly used numbers.
	//
	// See Non, E, and Pi
	E Transcendental = "ℯ"

	// Transcendental represents a transcendental number, which is a value identified as known to atlas.Precision
	// width - allowing 𝑡𝑖𝑛𝑦 to efficiently store very large commonly used numbers.
	//
	// See Non, E, and Pi
	Pi Transcendental = "π"
)

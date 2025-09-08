package num

import (
	"core/sys/atlas"
	"core/sys/num/internal"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type _base byte

var Base _base

/**
Public
*/

func (_base) StringToString(source string, sourceBase uint16, targetBase uint16) (string, uint) {
	if len(source) == 0 {
		return "", 0
	}

	digits, negative := Base.StringToDigits(source, sourceBase, targetBase)

	out := make([]string, len(digits))
	for i, d := range digits {
		out[i] = internal.PrintDigit(d)
	}

	if negative {
		out = append([]string{"-"}, out...)
	}

	if targetBase > 16 {
		return strings.Join(out, " "), uint(len(digits))
	}
	return strings.Join(out, ""), uint(len(digits))
}

func (_base) StringToDigits(source string, sourceBase uint16, targetBase uint16) ([]byte, bool) {
	if sourceBase < 2 {
		panic(fmt.Sprintf("invalid base: %d", sourceBase))
	}
	if sourceBase == 2 {
		return decimalStringToBaseDigits(binaryStringToDecimalString(source), targetBase)
	} else if sourceBase == 10 {
		return decimalStringToBaseDigits(source, targetBase)
	}
	return decimalStringToBaseDigits(baseStringToDecimalString(source, sourceBase), targetBase)
}

func (_base) DigitsToString(source []byte, sourceBase uint16, targetBase uint16) (string, uint) {
	digits := make([]string, len(source))

	for i, d := range source {
		digits[i] = internal.PrintDigit(d)
	}

	var sourceStr string
	if targetBase > 16 {
		sourceStr = strings.Join(digits, " ")
	} else {
		sourceStr = strings.Join(digits, "")
	}

	return Base.StringToString(sourceStr, sourceBase, targetBase)
}

func (_base) DigitsToDigits(source []byte, sourceBase uint16, targetBase uint16) ([]byte, bool) {
	digits := make([]string, len(source))

	for i, d := range source {
		digits[i] = internal.PrintDigit(d)
	}

	var sourceStr string
	if targetBase > 16 {
		sourceStr = strings.Join(digits, " ")
	} else {
		sourceStr = strings.Join(digits, "")
	}
	return Base.StringToDigits(sourceStr, sourceBase, targetBase)
}

/**
Private
*/

// binaryStringToDecimalString converts a binary input string to a base 10 string using the double-dabble method.
func binaryStringToDecimalString(s string) string {
	const base = 1_000_000_000 // 10^9, fits in uint32 with headroom for carries

	// Preprocess: trim spaces, handle optional sign and 0b/0B prefix, ignore underscores.
	s = strings.TrimSpace(s)
	if s == "" {
		panic("empty input")
	}

	neg := false
	switch s[0] {
	case '+':
		s = s[1:]
	case '-':
		neg = true
		s = s[1:]
	}
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'b' || s[1] == 'B') {
		s = s[2:]
	}
	if s == "" {
		panic("missing digits")
	}

	// Decimal big-int digits in little-endian base 1e9.
	// digits[0] is least-significant chunk.
	digits := []uint32{0}

	seenDigit := false
	for _, r := range s {
		if r == '_' {
			continue
		}
		if unicode.IsSpace(r) {
			continue
		}
		if r != '0' && r != '1' {
			panic(fmt.Errorf("invalid binary byte: %q", r))
		}
		seenDigit = true
		bit := uint64(r - '0')

		// value = value*2 + bit, done in base 1e9 with carries.
		var carry uint64 = bit
		for i := 0; i < len(digits); i++ {
			v := uint64(digits[i])*2 + carry
			digits[i] = uint32(v % base)
			carry = v / base
		}
		for carry != 0 {
			digits = append(digits, uint32(carry%base))
			carry /= base
		}
	}
	if !seenDigit {
		panic(errors.New("no binary digits found"))
	}

	// Convert decimal chunks to string.
	// Strip leading zeros from the most significant end.
	i := len(digits) - 1
	for i > 0 && digits[i] == 0 {
		i--
	}
	// Build most significant chunk without left padding.
	out := fmt.Sprintf("%d", digits[i])
	for j := i - 1; j >= 0; j-- {
		out += fmt.Sprintf("%09d", digits[j])
	}

	// Handle sign (avoid "-0").
	if neg && out != "0" {
		out = "-" + out
	}
	return out
}

// findPeriodic finds the periodic component of a num.Realized.  It deems a real is 'periodic' by
// checking if ceil(atlas.Precision/atlas.PeriodicDenominator) worth of trailing placeholders
// all contain a periodic value.
func findPeriodic(digits []byte) (pre, period []byte, repeats int) {
	// NOTE: This is just here for posterity - my original idea was that periodicity could be 'observed'
	// off of the number of repeating values in the fractional component, since we control the limit of
	// placeholders.  This naive thinking is what got me across the line on the Realized type, so I feel
	// it's REALLY important to keep here for posterity's sake - Alex

	n := len(digits)
	if n == 0 || uint(n) < atlas.Precision {
		return digits, nil, 0
	}

	depth := 4

	// NOTE: This uses ceiling division: d = (x + d - 1) / d
	threshold := (int(atlas.Precision) + (depth - 1)) / depth

	// Reverse digits -> rev
	rev := make([]byte, n)
	for i := 0; i < n; i++ {
		rev[i] = digits[n-1-i]
	}

	// KMP prefix function on rev
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && rev[i] != rev[j] {
			j = pi[j-1]
		}
		if rev[i] == rev[j] {
			j++
		}
		pi[i] = j
	}

	// Scan for the longest suffix (largest L) that is an exact repetition of its minimal period,
	// covers at least 'threshold' digits, and repeats at least twice.
	bestL, bestK, bestR := 0, 0, 0
	for L := n; L >= threshold; L-- {
		k := L - pi[L-1] // minimal period for this prefix (i.e., suffix in original)
		if k <= 0 {
			k = L
		}
		if L%k != 0 {
			continue
		}
		r := L / k
		if r >= 2 { // ensure it's actually repeating
			bestL, bestK, bestR = L, k, r
			break // largest L due to descending scan
		}
	}

	if bestL == 0 {
		return digits, nil, 0
	}

	pre = digits[:n-bestL]
	period = digits[n-bestK : n]
	repeats = bestR
	return pre, period, repeats
}

// decimalStringToBaseDigits converts the input string to a byte slice of bases [2, 256].
//
// NOTE: This expects a positive or negative integer value but allows whitespace and underscores.
func decimalStringToBaseDigits(s string, base uint16) (digits []byte, negative bool) {
	if base < 2 || base > 256 {
		panic("base must be in [2, 256]")
	}

	// Trim and handle sign.
	s = strings.TrimSpace(s)
	if s == "" {
		panic("empty input")
	}
	neg := false
	switch s[0] {
	case '+':
		s = s[1:]
	case '-':
		neg = true
		s = s[1:]
	}
	s = strings.TrimSpace(s)

	// Clean input: keep only '0'..'9', ignore underscores/spaces.
	buf := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '_' || c == ' ' || c == '\t' || c == '\n' || c == '\r':
			continue
		case c >= '0' && c <= '9':
			buf = append(buf, byte(c))
		default:
			panic("invalid character in decimal input")
		}
	}
	if len(buf) == 0 {
		panic("missing digits")
	}

	// Strip leading zeros.
	i := 0
	for i < len(buf) && buf[i] == '0' {
		i++
	}
	if i == len(buf) {
		// It's zero.
		return []byte{0}, false
	}
	dec := buf[i:]

	// Repeated long-division by 'base', collecting remainders.
	// Each pass: dec (base-10) -> quotient (base-10), remainder in [0, base-1].
	var digitsLE []byte
	for !(len(dec) == 1 && dec[0] == '0') {
		var carry uint32 = 0
		quot := make([]byte, 0, len(dec))
		for k := 0; k < len(dec); k++ {
			val := carry*10 + uint32(dec[k]-'0') // remainder*10 + next byte
			qd := val / uint32(base)
			carry = val % uint32(base)
			if len(quot) > 0 || qd > 0 {
				quot = append(quot, byte('0'+qd))
			}
		}
		if len(quot) == 0 {
			quot = []byte{'0'}
		}
		digitsLE = append(digitsLE, byte(carry)) // remainder is a byte in target base
		dec = quot
	}

	// Convert to MSB-first order.
	for l, r := 0, len(digitsLE)-1; l < r; l, r = l+1, r-1 {
		digitsLE[l], digitsLE[r] = digitsLE[r], digitsLE[l]
	}
	// If the number is zero, strip any leading zeros to canonical [0].
	if len(digitsLE) == 0 {
		digitsLE = []byte{0}
	}
	isNeg := neg && !(len(digitsLE) == 1 && digitsLE[0] == 0)
	return digitsLE, isNeg
}

// baseStringToDecimalString converts an input string in base [2,256] to a base-10 string.
//
// Expectations:
//   - Optional leading '+' or '-'.
//   - Base 2..16: compact digits using 0-9 and A-F/a-f; underscores are ignored.
//   - Base 17..256: whitespace-separated tokens; each token is exactly two hex characters (00..FF),
//     and its value must be < base.
//
// NOTE: For base₂ conversion, BinaryStringToDecimalString is far more efficient =)
//
// Returns: (decimalString, isNegative, error).
func baseStringToDecimalString(input string, base uint16) string {
	if base < 2 || base > 256 {
		panic("base must be in [2, 256]")
	}

	s := strings.TrimSpace(input)
	if s == "" {
		panic("empty input")
	}

	neg := false
	switch s[0] {
	case '+':
		s = strings.TrimSpace(s[1:])
	case '-':
		neg = true
		s = strings.TrimSpace(s[1:])
	}
	if s == "" {
		panic("missing digits")
	}

	digits := parseDigitsWithHexBytes(s, base)

	// Remove leading zeros in the byte sequence
	i := 0
	for i < len(digits) && digits[i] == 0 {
		i++
	}
	if i == len(digits) {
		return "0"
	}
	digits = digits[i:]

	// Evaluate as decimal string: res = res*base + d
	dec := "0"
	for _, d := range digits {
		dec = mulDecBySmall(dec, base)
		dec = addSmallToDec(dec, uint16(d))
	}

	if dec == "0" {
		return "0"
	}
	if neg {
		return "-" + dec
	}
	return dec
}

// parseDigitsWithHexBytes parses digits according to the stated rules.
// - Base <= 16: compact mode (0-9, A-F/a-F), underscores ignored, no internal whitespace.
// - Base >= 17: tokenized mode; each token must be exactly two hex chars (00..FF), value < base.
func parseDigitsWithHexBytes(s string, base uint16) []uint16 {
	if base <= 16 {
		// Compact mode. Reject internal whitespace (outside of leading/trailing).
		for _, r := range s {
			if unicode.IsSpace(r) {
				panic("whitespace-separated hex-byte tokens are only for bases > 16")
			}
		}
		var out []uint16
		for _, r := range s {
			if r == '_' {
				continue
			}
			v := hexCharToVal(r)
			if v < 0 {
				panic(fmt.Errorf("invalid byte '%c' for base %d", r, base))
			}
			if uint16(v) >= base {
				panic(fmt.Errorf("byte '%c' out of range for base %d", r, base))
			}
			out = append(out, uint16(v))
		}
		if len(out) == 0 {
			panic("missing digits")
		}
		return out
	}

	// Tokenized hex-byte mode for bases 17..256.
	fields := strings.Fields(s)
	if len(fields) == 0 {
		panic("missing digits")
	}
	out := make([]uint16, len(fields))
	for i, tok := range fields {
		if len(tok) == 1 {
			tok = "0" + tok
		}
		if len(tok) != 2 {
			panic(fmt.Errorf("token %q must be two hex characters", tok))
		}
		val, err := parseHexByte(tok)
		if err != nil {
			panic(fmt.Errorf("invalid hex-byte token %q: %w", tok, err))
		}
		if uint16(val) >= base {
			panic(fmt.Errorf("byte %02x out of range for base %d", val, base))
		}
		out[i] = uint16(val)
	}
	return out
}

// hexCharToVal returns 0..15 for hex chars, else -1.
func hexCharToVal(r rune) int {
	switch {
	case r >= '0' && r <= '9':
		return int(r - '0')
	case r >= 'A' && r <= 'F':
		return int(r-'A') + 10
	case r >= 'a' && r <= 'f':
		return int(r-'a') + 10
	default:
		return -1
	}
}

// parseHexByte parses exactly two hex characters into a byte (0..255).
func parseHexByte(tok string) (uint8, error) {
	// strconv.ParseUint handles mixed case and validates characters.
	v, err := strconv.ParseUint(tok, 16, 8)
	if err != nil {
		return 0, err
	}
	return uint8(v), nil
}

// mulDecBySmall multiplies a base-10 string by m (2..256).
func mulDecBySmall(dec string, m uint16) string {
	if dec == "0" {
		return "0"
	}
	out := make([]byte, 0, len(dec)+3)
	carry := uint32(0)
	for i := len(dec) - 1; i >= 0; i-- {
		d := uint32(dec[i] - '0')
		prod := d*uint32(m) + carry
		out = append(out, byte('0'+(prod%10)))
		carry = prod / 10
		if i == 0 {
			break
		}
	}
	for carry > 0 {
		out = append(out, byte('0'+(carry%10)))
		carry /= 10
	}
	// reverse
	for l, r := 0, len(out)-1; l < r; l, r = l+1, r-1 {
		out[l], out[r] = out[r], out[l]
	}
	return string(out)
}

// addSmallToDec adds add (0..255) to a base-10 string.
func addSmallToDec(dec string, add uint16) string {
	carry := uint32(add)
	out := make([]byte, 0, len(dec)+3)
	for i := len(dec) - 1; i >= 0; i-- {
		d := uint32(dec[i] - '0')
		sum := d + carry
		out = append(out, byte('0'+(sum%10)))
		carry = sum / 10
		if i == 0 {
			break
		}
	}
	for carry > 0 {
		out = append(out, byte('0'+(carry%10)))
		carry /= 10
	}
	// reverse
	for l, r := 0, len(out)-1; l < r; l, r = l+1, r-1 {
		out[l], out[r] = out[r], out[l]
	}
	// strip leading zeros (normally none)
	i := 0
	for i < len(out)-1 && out[i] == '0' {
		i++
	}
	return string(out[i:])
}

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str, _, err := BaseToDecimalString("f", 17)
	fmt.Println(str, err)
}

// BaseToDecimalString converts an input string in base [2,256] to a base-10 string.
// Rules:
//   - Optional leading '+' or '-'.
//   - Base 2..16: compact digits using 0-9 and A-F/a-f; underscores are ignored.
//   - Base 17..256: whitespace-separated tokens; each token is exactly two hex characters (00..FF),
//     and its value must be < base.
//
// Returns: (decimalString, isNegative, error).
func BaseToDecimalString(input string, base uint16) (string, bool, error) {
	if base < 2 || base > 256 {
		return "", false, errors.New("base must be in [2, 256]")
	}

	s := strings.TrimSpace(input)
	if s == "" {
		return "", false, errors.New("empty input")
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
		return "", false, errors.New("missing digits")
	}

	digits, err := parseDigitsWithHexBytes(s, base)
	if err != nil {
		return "", false, err
	}

	// Remove leading zeros in the digit sequence
	i := 0
	for i < len(digits) && digits[i] == 0 {
		i++
	}
	if i == len(digits) {
		return "0", false, nil
	}
	digits = digits[i:]

	// Evaluate as decimal string: res = res*base + d
	dec := "0"
	for _, d := range digits {
		dec = mulDecBySmall(dec, base)
		dec = addSmallToDec(dec, uint16(d))
	}

	if dec == "0" {
		neg = false
	}
	return dec, neg, nil
}

// parseDigitsWithHexBytes parses digits according to the stated rules.
// - Base <= 16: compact mode (0-9, A-F/a-F), underscores ignored, no internal whitespace.
// - Base >= 17: tokenized mode; each token must be exactly two hex chars (00..FF), value < base.
func parseDigitsWithHexBytes(s string, base uint16) ([]uint16, error) {
	if base <= 16 {
		// Compact mode. Reject internal whitespace (outside of leading/trailing).
		for _, r := range s {
			if unicode.IsSpace(r) {
				return nil, errors.New("whitespace-separated hex-byte tokens are only for bases > 16")
			}
		}
		var out []uint16
		for _, r := range s {
			if r == '_' {
				continue
			}
			v := hexCharToVal(r)
			if v < 0 {
				return nil, fmt.Errorf("invalid digit '%c' for base %d", r, base)
			}
			if uint16(v) >= base {
				return nil, fmt.Errorf("digit '%c' out of range for base %d", r, base)
			}
			out = append(out, uint16(v))
		}
		if len(out) == 0 {
			return nil, errors.New("missing digits")
		}
		return out, nil
	}

	// Tokenized hex-byte mode for bases 17..256.
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return nil, errors.New("missing digits")
	}
	out := make([]uint16, len(fields))
	for i, tok := range fields {
		if len(tok) == 1 {
			tok = "0" + tok
		}
		if len(tok) != 2 {
			return nil, fmt.Errorf("token %q must be two hex characters", tok)
		}
		val, err := parseHexByte(tok)
		if err != nil {
			return nil, fmt.Errorf("invalid hex-byte token %q: %w", tok, err)
		}
		if uint16(val) >= base {
			return nil, fmt.Errorf("digit %02x out of range for base %d", val, base)
		}
		out[i] = uint16(val)
	}
	return out, nil
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

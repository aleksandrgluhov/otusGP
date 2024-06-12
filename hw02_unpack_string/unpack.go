package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	EscapeChar       = '\\'
	ErrInvalidEscape = errors.New("invalid escape")
	ErrInvalidString = errors.New("invalid string")
)

func Unpack(s string) (string, error) {
	// Input string is empty
	if len(s) == 0 {
		return "", nil
	}

	r := []rune(s)

	// Error: invalid string, f.e.: 4ab
	if unicode.IsDigit(r[0]) {
		return "", ErrInvalidString
	}

	// Value constructor
	var sb strings.Builder

	// Loop control variables
	skipped := false
	escaped := false

	for i, cc := range r {
		var (
			nc     rune
			repeat int
		)

		// Error: attempt to escape regular character, f.e.: \a
		if escaped && !unicode.IsDigit(cc) && (cc != EscapeChar) {
			return "", ErrInvalidEscape
		}

		// If not last iteration, compute nc values, otherwise append cc if not skipped
		if len(r)-i > 1 {
			nc = r[i+1]
			repeat, _ = strconv.Atoi(string(nc))
		} else {
			switch {
			// Error: last character is not escaped escape operation
			case !escaped && (cc == EscapeChar):
				return "", ErrInvalidEscape
			case !skipped:
				sb.WriteString(string(cc))
				continue
			}
		}

		// Error: invalid string, f.e.: a69b, ab69
		if !escaped && unicode.IsDigit(cc) && unicode.IsDigit(nc) {
			return "", ErrInvalidString
		}

		// If we've done any repeat operation before, then skip cc
		if skipped {
			skipped = false
			continue
		}

		// Skip iteration on escape operation
		if !escaped && (cc == EscapeChar) {
			escaped = true
			continue
		}

		escaped = false

		// Apply repeat operation or just append cc
		if unicode.IsDigit(nc) {
			skipped = true
			if repeat > 0 {
				sb.WriteString(strings.Repeat(string(cc), repeat))
			}
		} else {
			sb.WriteString(string(cc))
		}
	}

	return sb.String(), nil
}

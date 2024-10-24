package counter

import (
	"errors"
	"unicode/utf8"
)

var counter int

func Increment() {
	counter++
}

func Sum(a, b int) int {
	return a + b
}

func Length(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "short"
	case a < 100:
		return "long"
	}
	return "very long"
}

// DeleteVowels removes all vowels (a, e, i, o, u) from the input string.
func DeleteVowels(s string) string {
	vowels := "aeiouAEIOU"
	result := ""
	for _, char := range s {
		if !contains(vowels, char) {
			result += string(char)
		}
	}
	return result
}

// contains checks if a character is in the given string.
func contains(s string, char rune) bool {
	for _, v := range s {
		if v == char {
			return true
		}
	}
	return false
}

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}
	return utf8.RuneCount(input), nil
}

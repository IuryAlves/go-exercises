package isogram

import (
	"strings"
)

// IsIsogram: Return a boolean indicating whether a string has no repeating letters
func IsIsogram(s string) bool {
	counter := make(map[rune]int)
	s = strings.ToLower(s)

	for _, char := range s {
		// only increment counter if char is a letter
		if char > 97 && char <= 122 {
			counter[char]++
		}
		if counter[char] > 1 {
			return false
		}
	}
	return true
}

package codewars

import (
	"unicode"
)

func FirstNonRepeating(str string) string {
	// chars := []rune(str)
	// finded := false
	count := make(map[rune]int)

	// First pass: count the occurrences of each character (case-insensitive)
	for _, r := range str {
		// Convert to lowercase for comparison
		count[unicode.ToLower(r)]++
	}

	// Second pass: find the first character with count 1
	for _, r := range str {
		// Check the count in lowercase
		if count[unicode.ToLower(r)] == 1 {
			return string(r) // Return the original case
		}
	}
	// Test loops
	/* for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars); j++ {
			if chars[i] == chars[j] {
				continue
			}
		}
		if finded {
			return string(chars[i])
		}
	}
	*/
	return ""
}

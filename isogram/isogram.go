// Package isogram determines if a word or phrase is an isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks whether given word isogram or not
func IsIsogram(word string) bool {
	m := make(map[rune]int)
	for _, r := range word {
		r = unicode.ToLower(r)
		if strings.ContainsRune("-", r) || strings.ContainsRune(" ", r) {
			continue
		}
		m[r] += 1
		if m[r] > 1 {
			return false
		}
	}
	return true
}

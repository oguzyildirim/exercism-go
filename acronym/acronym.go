// Package acronym provides converter that gives you acronym of given word
package acronym

import (
	"strings"
	"unicode/utf8"
)

// Abbreviate abbreviate given word to its acronym
func Abbreviate(s string) string {
	// hyphenated words are treated as separate words
	words := strings.FieldsFunc(s, split)
	var acr string
	for _, str := range words {
		r, _ := utf8.DecodeRuneInString(str)
		acr = acr + string(r)
	}
	return strings.ToUpper(acr)
}

func split(r rune) bool {
	return r == ' ' || r == '-' || r == '_'
}

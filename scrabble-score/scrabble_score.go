// Package scrabble calculates scrabble score of word
package scrabble

import (
	"strings"
)

// Score calculates scrabble score of word
func Score(word string) int {
	var result int
	word = strings.ToLower(word)
	for _, char := range word {
		result += getCharValue(char)
	}
	return result
}

func getCharValue(r rune) int {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 0
	}
}

// Package etl provides new system for scrabble score
package etl

import "strings"

// Transform  migrates shiny new scrabble system
func Transform(outdated map[int][]string) map[string]int {
	result := make(map[string]int)
	for score, chars := range outdated {
		for _, char := range chars {
			result[strings.ToLower(char)] = score
		}
	}
	return result
}

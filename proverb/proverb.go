// Package proverb provides text of proverbial rhymes
package proverb

import "fmt"

// Proverb  generates proverbial rhymes
func Proverb(rhyme []string) []string {
	var result []string
	if len(rhyme) == 0 {
		return result
	}
	for i := 1; i < len(rhyme); i++ {
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	result = append(result, "And all for the want of a "+rhyme[0]+".")
	return result
}

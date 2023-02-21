// Package strand provides dna to rna conversion
package strand

import "strings"

// ToRNA given a DNA strand, return its RNA complement
func ToRNA(input string) string {
	return strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	).Replace(input)
}

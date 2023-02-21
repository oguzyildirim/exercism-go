// Package dna provides count algorithms for dna nucleotides
package dna

import (
	"fmt"
	"strings"
)

type Histogram map[rune]uint

// DNA is a list of nucleotides
type DNA string

// Counts counts all nucleotides in a strand
// Returns an error if strand contains an invalid nucleotide
func (d DNA) Counts() (Histogram, error) {
	d = DNA(strings.ToUpper(string(d)))
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range d {
		_, ok := h[nucleotide]
		if !ok {
			return nil, fmt.Errorf("Invalid nucleotide %v", nucleotide)
		}

		h[nucleotide]++
	}

	return h, nil
}

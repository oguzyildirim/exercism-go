// Package hamming includes algortihm that calculates hamming distance
package hamming

import "errors"

// Distance calculates hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("different length of DNAs")
	}
	hammingDistance := 0
	runeB := []rune(b)
	for pos, char := range a {
		if char != runeB[pos] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}

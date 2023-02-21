// Package darts provides calculate method for dart game
package darts

import "math"

// Score calculates the score of hit
func Score(x float64, y float64) int {
	radius := math.Sqrt(x*x + y*y)
	switch {
	case radius <= 1:
		return 10
	case radius <= 5:
		return 5
	case radius <= 10:
		return 1
	default:
		return 0
	}
}

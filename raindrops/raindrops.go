// Package raindrops includes algorithms that finds raindrops in a given number
package raindrops

import "strconv"

// Convert convert given number to raindrops
func Convert(num int) string {
	var raindrop string
	if num%3 == 0 {
		raindrop = "Pling"
	}
	if num%5 == 0 {
		raindrop = raindrop + "Plang"
	}
	if num%7 == 0 {
		raindrop = raindrop + "Plong"
	}
	if len(raindrop) == 0 {
		return strconv.Itoa(num)
	}
	return raindrop
}

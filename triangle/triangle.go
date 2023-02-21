// Package triangle provides algorithm that identify types of given triangle
package triangle

import "math"

// Kind indicates whether a triangle is isosceles, scalane, etc.
type Kind int

const (
	// NaT = not a triangle
	NaT Kind = iota
	// Equ = equilateral
	Equ
	// Iso = isosceles
	Iso
	// Sca = scalene
	Sca
)

// KindFromSides finds the type of given triangle
func KindFromSides(a, b, c float64) Kind {
	if isValidTriangle(a, b, c) {
		if isEquilateral(a, b, c) {
			return Equ
		} else if isIsosceles(a, b, c) {
			return Iso
		} else if isScalene(a, b, c) {
			return Sca
		}
	}

	return NaT
}

func isValidTriangle(a, b, c float64) bool {
	return isValidSide(a) && isValidSide(b) && isValidSide(c) && isValidUnderInequality(a, b, c)
}

func isValidSide(side float64) bool {
	return side > 0 && side != math.Inf(1)
}

func isValidUnderInequality(a, b, c float64) bool {
	return (a <= b+c) && (b <= a+c) && (c <= a+b)
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isIsosceles(a, b, c float64) bool {
	return (a == b || a == c || b == c) && !isEquilateral(a, b, c)
}

func isScalene(a, b, c float64) bool {
	return a != b && a != c && b != c
}

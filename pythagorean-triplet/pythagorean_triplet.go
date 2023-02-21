// Package pythagorean finds the Pythagorean triplets
package pythagorean

// Pythagorean triplet is a set of three natural numbers, {a, b, c}, for which, "a**2 + b**2 = c**2"
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the range min to max inclusive.
func Range(min int, max int) []Triplet {
	var result []Triplet
	for a := min; a <= max; a++ {
		b := a + 1
		c := b + 1
		for c <= max {
			for c*c < a*a+b*b {
				c += 1
			}
			if c*c == a*a+b*b {
				result = append(result, Triplet{a, b, c})
			}
			b += 1
		}
	}
	return result
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c (the perimeter) is equal to p.
func Sum(sum int) []Triplet {
	var result []Triplet
	for _, triplet := range Range(1, sum) {
		if triplet[0]+triplet[1]+triplet[2] == sum {
			result = append(result, triplet)
		}
	}
	return result
}

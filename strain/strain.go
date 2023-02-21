// Package strain implement the keep and discard operation on collections
package strain

// Ints collection of integers
type Ints []int

// Lists collection of arrays of integers
type Lists [][]int

// Strings collection of strings
type Strings []string

// Keep filters collection with a given function
func (i Ints) Keep(strainer func(int) bool) Ints {
	var result Ints
	for _, num := range i {
		if strainer(num) {
			result = append(result, num)
		}
	}
	return result
}

// Discard discards values in collection with a given function
func (i Ints) Discard(strainer func(int) bool) Ints {
	var result Ints
	for _, num := range i {
		if !strainer(num) {
			result = append(result, num)
		}
	}
	return result
}

// Keep filters collection with a given function
func (l Lists) Keep(strainer func([]int) bool) Lists {
	var result Lists
	for _, num := range l {
		if strainer(num) {
			result = append(result, num)
		}
	}
	return result
}

// Keep filters collection with a given function
func (s Strings) Keep(strainer func(string) bool) Strings {
	var result Strings
	for _, num := range s {
		if strainer(num) {
			result = append(result, num)
		}
	}
	return result
}

// Package collatzconjecture  summarize 3x+1 problem
package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return n, fmt.Errorf("non positive number %v", n)
	}
	var result int
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		result++
	}
	return result, nil
}

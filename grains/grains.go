// Package grains calculates the number of grains of wheat on a chessboard
package grains

import "errors"

// Square returns the number of grains of wheat on a chessboard
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("number must be between 1 and 64")
	}

	return (1<<n - 1), nil
}

// Total returns the number of grains of rice on the entire chessboard
func Total() uint64 {
	return 1<<64 - 1
}

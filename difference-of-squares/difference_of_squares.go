//  Package diffsquares calculates the square of the sum and the sum of the squares
package diffsquares

// SquareOfSum calculates square of the sum
func SquareOfSum(n int) int {
	var result int
	sum := (n * (n + 1)) / 2
	result = sum * sum
	return result
}

// SumOfSquares calculates sum of the squares
func SumOfSquares(n int) int {
	result := (n * (n + 1) * (2*n + 1)) / 6
	return result
}

// Difference calculates difference between the square of the sum and the sum of the squares
func Difference(n int) int {
	result := SquareOfSum(n) - SumOfSquares(n)
	return result
}

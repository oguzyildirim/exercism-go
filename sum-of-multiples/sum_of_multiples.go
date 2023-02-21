package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var result int
	for i := 1; i < limit; i++ {
		for _, n := range divisors {
			if n == 0 {
				break
			}
			if i%n == 0 {
				result += i
				break
			}
		}
	}
	return result
}

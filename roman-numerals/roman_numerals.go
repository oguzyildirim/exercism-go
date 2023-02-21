// Package romannumerals provides number conversion for roman numbers
package romannumerals

import "fmt"

// ToRomanNumeral converts numbers to roman numbers
func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n >= 3001 {
		return "", fmt.Errorf("out of range for number: %v", n)
	}
	var result string
	numbers := [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romanNumbers := [13]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	for i := 12; i > -1; i-- {
		rep := n / numbers[i]
		n -= rep * numbers[i]
		for j := 0; j < rep; j++ {
			result += romanNumbers[i]
		}
	}
	return result, nil
}

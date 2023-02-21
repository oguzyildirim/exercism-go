// Package luhn test given string whether its is valid luhn number
package luhn

import (
	"strconv"
	"strings"
)

// Valid checks whether given string is vallid luhn number
func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	_, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	isOdd := len(number) % 2
	sum := 0
	for i, num := range number {
		num := int(num - '0')
		if i%2 == isOdd {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	if len(number) < 2 || sum%10 != 0 {
		return false
	}
	return true
}

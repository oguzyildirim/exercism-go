// Package pangram determines if the given string contains every letter
package pangram

import "strings"

// IsPangram determines if the given string contains every letter
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for l := 'a'; l <= 'z'; l++ {
		if !strings.ContainsRune(s, l) {
			return false
		}
	}
	return true
}

// // IsPangram determines if the given string contains every letter
// func IsPangram(s string) bool {
// 	var bs int32
// 	for i := 0; i < len(s); i++ {
// 		c := (s[i] & 0xdf) - 'A'
// 		if c > 25 || c < 0 {
// 			continue
// 		}
// 		bs |= 1 << c
// 	}
// 	return bs == 0x3ffffff
// }

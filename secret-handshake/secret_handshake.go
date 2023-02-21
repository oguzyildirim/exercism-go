// Package secret provides some kind of way to secret messaging
package secret

// Package a deal with buddy
func Handshake(num uint) []string {
	secret := []string{}
	if num&1 == 1 {
		secret = append(secret, "wink")
	}
	if num&2 == 2 {
		secret = append(secret, "double blink")
	}
	if num&4 == 4 {
		secret = append(secret, "close your eyes")
	}
	if num&8 == 8 {
		secret = append(secret, "jump")
	}
	if num&16 == 16 {
		for i, j := 0, len(secret)-1; i < j; i, j = i+1, j-1 {
			secret[i], secret[j] = secret[j], secret[i]
		}
	}
	return secret
}

// Package twofer is a logging system with the given name
package twofer

// ShareWith logs pattern with the given name
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}

// Package accumulate applies given function to list
package accumulate

// Accumulate  apply operation on list
func Accumulate(list []string, operation func(string) string) (result []string) {
	for _, element := range list {
		o := operation(element)
		result = append(result, o)
	}
	return result
}

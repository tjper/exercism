// Package accumlate provides a library of functions that perform operations on sets of values.
package accumulate

// Accumulate performs a specified operation on a set of strings, if successful,
// the new set of strings is returned.
func Accumulate(s []string, f func(string) string) []string {
	for i, v := range s {
		s[i] = f(v)
	}
	return s
}

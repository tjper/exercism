// Package twofer implements a library for generating two for one strings.
package twofer

// ShareWith inserts a name into a string, or defaults to "you", if successful,
// a string is returned.
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}

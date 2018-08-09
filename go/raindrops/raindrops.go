// Package raindrops provides a library for raindrop themed functions.
package raindrops

import "strconv"

// Convert translates a number into a raindrop sound, if successful,
// it returns the raindrop sound text.
func Convert(n int) string {
	var out string
	if n%3 == 0 {
		out += "Pling"
	}
	if n%5 == 0 {
		out += "Plang"
	}
	if n%7 == 0 {
		out += "Plong"
	}
	if out == "" {
		out = strconv.Itoa(n)
	}
	return out
}

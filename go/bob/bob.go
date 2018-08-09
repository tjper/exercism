// Package bob provides a library for communication with bob the teenager.
package bob

import (
	"strings"
	"unicode"
)

// Hey accepts a remark, if successful,
// a deterministic string response is returned.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	var res string
	q := strings.HasSuffix(remark, "?")
	e := (strings.ToUpper(remark) == remark) && alphaExists(remark)
	switch {
	case q && e:
		res = "Calm down, I know what I'm doing!"
	case q:
		res = "Sure."
	case e:
		res = "Whoa, chill out!"
	default:
		res = "Whatever."
	}
	return res
}

// alphaExists accepts a string and determines if there is at least one letter in a string
// a bool is returned indicating if a char exits.
func alphaExists(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// Package acronym provides a library for transforming words into abbreviations.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate accepts a string arguement that is converted into an acronym, if successful,
// the acronym is returned.
func Abbreviate(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}

	acronym := []byte{s[0]}
	for i, r := range s {
		if unicode.IsSpace(r) || r == '-' {
			acronym = append(acronym, s[i+1])
		}
	}
	return strings.ToUpper(string(acronym))
}

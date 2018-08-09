// Package hamming provides a library to determine the hamming distance between two DNA strands.
package hamming

import "errors"

var ErrStrLenNotEqual = errors.New("String length not equal.")

// Distance determines the hamming distance between two strings a, b, if successful,
// the distance is returned.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, ErrStrLenNotEqual
	}

	var cnt int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt, nil
}

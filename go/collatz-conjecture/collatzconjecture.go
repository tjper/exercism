// Package collatzconjecture provides a library of algorithms to solve Collatz Conjecture
package collatzconjecture

import "errors"

var (
	ErrNEqualsZero = errors.New("n equals zero.")
	ErrNIsNegative = errors.New("n is Negative.")
)

// CollatzConjecture determines the number steps required to solve Collatz Conjecture for n, if successful,
// the number of steps is returned.
func CollatzConjecture(n int) (int, error) {
	if n == 0 {
		return 0, ErrNEqualsZero
	}
	if n < 0 {
		return 0, ErrNIsNegative
	}

	var stepCnt int
	for {
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		stepCnt++
	}
	return stepCnt, nil
}

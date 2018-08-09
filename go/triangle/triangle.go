// Package triangle provides a library for interacting with "triangles".
package triangle

import "math"

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// KindFromSides determines the kind of triangle based on the length of its sides, if successful,
// the Kind of triangle is returned.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	for _, v := range sides {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return NaT
		}
	}

	sortSides(sides)
	if (sides[0]+sides[1] < sides[2]) || sides[0] <= 0 {
		return NaT
	}

	var eqCnt int
	m := make(map[float64]bool)

	for _, v := range sides {
		if _, exists := m[v]; exists {
			eqCnt++
		} else {
			m[v] = true
		}
	}

	var k Kind
	switch eqCnt {
	case 0:
		k = Sca
	case 1:
		k = Iso
	case 2:
		k = Equ
	}
	return k
}

// I'm aware of the sort package, but didn't want to add the dependency.
// My bench marking metrics also dramatically improved by writing my own function. Is this bad practice?

// sortSides orders the triangle sides from min to max, if successful,
// the resulting slice is returned.
func sortSides(sides []float64) {
	var min, mid, max float64
	if sides[0] < sides[1] {
		min, max = sides[0], sides[1]
	} else {
		min, max = sides[1], sides[0]
	}

	if sides[2] < min {
		min, mid = sides[2], min
	} else if sides[2] < max {
		mid = sides[2]
	} else {
		mid, max = max, sides[2]
	}
	sides[0], sides[1], sides[2] = min, mid, max
}

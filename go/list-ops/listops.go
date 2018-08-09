// Package listops provides a library for interacting with lists.
package listops

type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldr folds IntList beginning at end of list via specified function.
// Returns resulting int.
func (l IntList) Foldr(fn binFunc, initial int) int {
	res := initial
	for i := len(l) - 1; i >= 0; i-- {
		res = fn(l[i], res)
	}
	return res
}

// Foldl folds IntList via beginning at start of list via specified function.
// Returns resulting int.
func (l IntList) Foldl(fn binFunc, initial int) int {
	res := initial
	for _, v := range l {
		res = fn(res, v)
	}
	return res
}

// Filter filters the IntList elements based on the filter function.
// Returns the resulting IntList.
func (l IntList) Filter(fn predFunc) IntList {
	newList := make(IntList, 0)
	for _, v := range l {
		if fn(v) {
			newList = newList.Append(IntList([]int{v}))
		}
	}
	return newList
}

// Length determines how many elements are in the IntList.
// Returns number of elements.
func (l IntList) Length() int {
	var cnt int
	for _, _ = range l {
		cnt++
	}
	return cnt
}

// Map performs an operation on each element of IntList as specified by fn.
// Returns the modified IntList.
func (l IntList) Map(fn unaryFunc) IntList {
	for i, v := range l {
		l[i] = fn(v)
	}
	return l
}

// Reverse reverses the ordering of the elements within the IntList.
// Returns the reversed IntList.
func (l IntList) Reverse() IntList {
	newList := make(IntList, 0)
	for i := (l.Length() - 1); i >= 0; i-- {
		newList = newList.Append(IntList([]int{l[i]}))
	}
	return newList
}

// Append inserts a set of elements onto the end of IntList.
// Returns the updated IntList.
func (l IntList) Append(a IntList) IntList {
	lLen, aLen := l.Length(), a.Length()
	newList := make(IntList, (lLen + aLen))

	i := 0
	for i < lLen {
		newList[i] = l[i]
		i++
	}
	for _, v := range a {
		newList[i] = v
		i++
	}
	return newList
}

// Concat concatenates a IntList set together into one IntList.
// Returns concatenated IntList.
func (l IntList) Concat(a []IntList) IntList {
	for _, v := range a {
		l = l.Append(v)
	}
	return l
}

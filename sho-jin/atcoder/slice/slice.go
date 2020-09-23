package slice

import (
	"errors"
)

// DeleteElement returns a *NEW* slice, that have the same and minimum length and capacity.
// DeleteElement makes a new slice by using easy slice literal.
func DeleteElement(s []int, i int) []int {
	if i < 0 || len(s) <= i {
		panic(errors.New("[index error]"))
	}
	// appendのみの実装
	n := make([]int, 0, len(s)-1)
	n = append(n, s[:i]...)
	n = append(n, s[i+1:]...)
	return n
}

// Concat returns a *NEW* slice, that have the same and minimum length and capacity.
func Concat(s, t []rune) []rune {
	n := make([]rune, 0, len(s)+len(t))
	n = append(n, s...)
	n = append(n, t...)
	return n
}

// SafeDeleteElement returns a *NEW* slice, that have the same and minimum length and capacity.
// SafeDeleteElement makes a new slice by simply appending each element FROM SCRACH.
func SafeDeleteElement(s []int, i int) []int {
	if i < 0 || len(s) <= i {
		panic(errors.New("[index error]"))
	}
	n := make([]int, 0, len(s)-1)
	for idx := 0; idx < len(s); idx++ {
		if idx == i {
			continue
		}
		n = append(n, s[idx])
	}
	return n
}

// SafeConcat returns a *NEW* slice, that have the same and minimum length and capacity.
func SafeConcat(s, t []rune) []rune {
	n := make([]rune, 0, len(s)+len(t))
	for _, e := range s {
		n = append(n, e)
	}
	for _, e := range t {
		n = append(n, e)
	}
	return n
}

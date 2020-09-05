package arithmetic

// IsProductLeq returns whether a*b <= ub or not.
// IsProductLeq panics when it accepts negative integers.
func IsProductLeq(a, b, ub int) bool {
	if a < 0 || b < 0 || ub < 0 {
		panic("IsProductLeq does not accept negative integers")
	}

	return a <= ub/b
}

// IsSumLeq returns whether a*b <= ub or not.
// IsSumLeq panics when it accepts negative integers.
func IsSumLeq(a, b, ub int) bool {
	if a < 0 || b < 0 || ub < 0 {
		panic("IsSumLeq does not accept negative integers")
	}

	return a <= ub-b
}

package arithmetic

import "math"

// IsProductOverflow returns whether a*b <= MAX_INT64 or not.
// IsProductOverflow panics when it accepts negative integers.
func IsProductOverflow(a, b int) bool {
	if a < 0 || b < 0 {
		panic("IsProductOverflow does not accept negative integers")
	}

	return a <= (math.MaxInt64 / b)
}

// IsSumOverflow returns whether a*b <= MAX_INT64 or not.
// IsSumOverflow panics when it accepts negative integers.
func IsSumOverflow(a, b int) bool {
	if a < 0 || b < 0 {
		panic("IsSumOverflow does not accept negative integers")
	}

	return a <= (math.MaxInt64 - b)
}

// IsProductLeq returns whether a*b <= ub or not.
// IsProductLeq panics when it accepts negative integers.
func IsProductLeq(a, b, ub int) bool {
	if a < 0 || b < 0 || ub < 0 {
		panic("IsProductLeq does not accept negative integers")
	}

	return a <= (ub / b)
}

// IsSumLeq returns whether a*b <= ub or not.
// IsSumLeq panics when it accepts negative integers.
func IsSumLeq(a, b, ub int) bool {
	if a < 0 || b < 0 || ub < 0 {
		panic("IsSumLeq does not accept negative integers")
	}

	return a <= (ub - b)
}

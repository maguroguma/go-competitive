package smlpiece

// FloorFrac behaves like math.Floor when you want to use only integer type.
// e.g. math.Floor(float64(a) / float64(b))
func FloorFrac(a, b int) int {
	_abs := func(d int) int {
		if d < 0 {
			return -d
		}
		return d
	}

	if _abs(a)%_abs(b) == 0 {
		return a / b
	}

	if (a < 0 && b < 0) || (a > 0 && b > 0) {
		return a / b
	}
	return a/b - 1
}

// CeilFrac behaves like math.Ceil when you want to use only integer type.
// e.g. math.Ceil(float64(a) / float64(b))
func CeilFrac(a, b int) int {
	_abs := func(d int) int {
		if d < 0 {
			return -d
		}
		return d
	}

	if _abs(a)%_abs(b) == 0 {
		return a / b
	}

	if (a < 0 && b < 0) || (a > 0 && b > 0) {
		return a/b + 1
	}
	return a / b
}

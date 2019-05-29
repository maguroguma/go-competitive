package divisor

// Divisors returns the divisors of an argument integer as map[int]int.
func Divisors(n int) map[int]int {
	res := make(map[int]int)
	// res := map[int]int{}

	for l := 1; l*l <= n; l++ {
		if n%l == 0 {
			res[l] = 1
			res[n/l] = 1
		}
	}

	return res
}

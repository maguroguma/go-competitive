package combination

// CalcComb returns nCr.
func CalcComb(n, r int) int {
	if r > n-r {
		return CalcComb(n, n-r)
	}

	resMul, resDiv := 1, 1
	for i := 0; i < r; i++ {
		resMul *= n - i
		resDiv *= i + 1
	}

	return resMul / resDiv
}

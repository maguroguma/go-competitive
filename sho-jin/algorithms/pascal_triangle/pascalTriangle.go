package pascal_triangle

func Combination(n, k int) int {
	if n < 0 || 100 < n || k < 0 || n < k {
		panic("[Argument Error]")
	}

	var C [101][101]int
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				C[i][j] = 1
			} else {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}

	return C[n][k]
}

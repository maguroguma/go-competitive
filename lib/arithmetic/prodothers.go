package arithmetic

// ProdOthers returns B that B[i] denotes f(A[:i]..., A[i+1:]...).
// Time complexity: O(n)
func ProdOthers(A []int, f func(x, y int) int) (B []int) {
	if len(A) < 2 {
		panic("A must be have more than one element")
	}

	n := len(A)
	L, R := make([]int, n), make([]int, n)

	L[0] = A[0]
	for i := 1; i < n; i++ {
		L[i] = f(L[i-1], A[i])
	}
	R[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		R[i] = f(R[i+1], A[i])
	}

	B = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			B[0] = R[1]
			continue
		}
		if i == n-1 {
			B[n-1] = L[n-2]
			continue
		}

		B[i] = f(L[i-1], R[i+1])
	}

	return B
}

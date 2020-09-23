package slice

// PincerScan shows how to scan from edge to center.
func PincerScan(n int) []int {
	S := make([]int, 0, n)
	for i := 0; i < n/2; i++ {
		S = append(S, i)
		S = append(S, (n-1)-i)
	}
	if n%2 == 1 {
		S = append(S, n/2)
	}
	return S
}

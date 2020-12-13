package combinatorics

type PascalTriangle struct {
	CombTable [][]int
}

// NewPascalTriangle receive maximal n of nCr.
// n should be 60 at most, because Sum(C(n, i)) == 2^n.
// Time complexity: O(n^2)
func NewPascalTriangle(n int) *PascalTriangle {
	pt := new(PascalTriangle)

	C := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]int, n+1)
	}

	C[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			C[i+1][j] += C[i][j]
			C[i+1][j+1] += C[i][j]
		}
	}

	pt.CombTable = C
	return pt
}

// C(n, r) returns the value of nCr.
func (pt *PascalTriangle) C(n, r int) int {
	return pt.CombTable[n][r]
}

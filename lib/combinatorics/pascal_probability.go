package combinatorics

type PascalTriangleProb struct {
	CombProbTable [][]float64
}

// NewPascalTriangleProb receive maximal n of nCr.
// Time complexity: O(n^2)
func NewPascalTriangleProb(n int) *PascalTriangleProb {
	pt := new(PascalTriangleProb)

	C := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]float64, n+1)
	}

	C[0][0] = 1.0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			C[i+1][j] += C[i][j] * 0.5
			C[i+1][j+1] += C[i][j] * 0.5
		}
	}

	pt.CombProbTable = C
	return pt
}

// C(n, r) returns the probability of nCr,
//  the probability of event is 1/2.
func (pt *PascalTriangleProb) CP(n, r int) float64 {
	return pt.CombProbTable[n][r]
}

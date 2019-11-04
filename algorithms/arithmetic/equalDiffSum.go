package arithmetic

import "math"

// ArithmeticSequenceSum returns a sum of an arithmetic sequence.
// a: 初項, d: 公差, n: 項数
func ArithmeticSequenceSum(a, d, n int) int {
	return (2*a + (n-1)*d) * n / 2
}

// GeometricSequenceSum returns a sum of a geometric sequence.
// a: 初項, r: 公比, n: 項数
func GeometricSequenceSum(a, r, n int) int {
	nume := a * (1 - int(math.Pow(float64(r), float64(n))))
	deno := 1 - r
	return nume / deno
}

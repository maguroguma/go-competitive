package csum

type CumulativeSum struct {
	csum []int64
}

func NewCumulativeSum(A []int64) *CumulativeSum {
	cs := new(CumulativeSum)

	n := len(A)
	cs.csum = make([]int64, n+1)
	for i := 0; i < n; i++ {
		cs.csum[i+1] = cs.csum[i] + A[i]
	}

	return cs
}

// RangeSum returns sum of [l, r) elements of original array,
//  that is, Sum(A[l:r+1]...).
func (cs *CumulativeSum) RangeSum(l, r int) int64 {
	return cs.csum[r] - cs.csum[l]
}

package csum

import "testing"

func TestCumulativeSum(t *testing.T) {
	A := []int64{749, 702, 142, 7234, 349, 656}
	n := len(A)

	cs := NewCumulativeSum(A)

	for l := 0; l < n; l++ {
		for r := l + 1; r <= n; r++ {
			actual := int64(0)
			for i := l; i < r; i++ {
				actual += A[i]
			}

			expected := cs.RangeSum(l, r)

			if actual != expected {
				t.Errorf("got %v, want %v", actual, expected)
			}
		}
	}
}

package csum

import "testing"

func TestCumulativeSum2Dim(t *testing.T) {
	A := [][]int64{
		{123, 234, 345, 456},
		{567, 678, 789, 8910},
		{91011, 101112, 111213, 121314},
	}
	h, w := 3, 4

	_bruteforce := func(top, left, bottom, right int) int64 {
		res := int64(0)

		for i := top; i <= bottom; i++ {
			for j := left; j <= right; j++ {
				res += A[i][j]
			}
		}

		return res
	}

	rs := NewRectangleSum(A)

	for top := 0; top < h; top++ {
		for left := 0; left < w; left++ {
			for bottom := top; bottom < h; bottom++ {
				for right := left; right < w; right++ {
					actual := rs.RangeSum(top, left, bottom, right)
					expected := _bruteforce(top, left, bottom, right)

					if actual != expected {
						t.Errorf("got %v, want %v\n", actual, expected)
					}
				}
			}
		}
	}
}

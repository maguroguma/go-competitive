package fenwicktree

import (
	"fmt"
	"testing"
)

// https://judge.yosupo.jp/problem/point_add_range_sum
func TestFenwickTreePointAddAndRangeSum(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	Queries := [][]int{
		{1, 0, 5, 15},
		{1, 2, 4, 7},
		{0, 3, 10},
		{1, 0, 5, 25},
		{1, 0, 3, 6},
	}

	ft := NewFenwickTree(n)
	for i := 0; i < n; i++ {
		ft.Add(i, A[i])
	}

	for _, Q := range Queries {
		if Q[0] == 0 {
			ft.Add(Q[1], Q[2])
		} else {
			sum := ft.RangeSum(Q[1], Q[2])
			if sum != Q[3] {
				t.Errorf("got %v, want %v", sum, Q[3])
			}
		}
	}
}

// original
func TestFenwickTreeLowerBound(t *testing.T) {
	n := 5
	A := []int{0, 1, 2, 3, 4}

	ft := NewFenwickTree(n)
	for i := 0; i < n; i++ {
		ft.Add(i, A[i])
	}

	testcases := []struct {
		w        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2}, {3, 2},
		{4, 3}, {5, 3}, {6, 3},
		{7, 4}, {8, 4}, {9, 4}, {10, 4},
		{11, 5}, {100, 5},
	}

	for id, tc := range testcases {
		testName := fmt.Sprintf("%d test", id)
		t.Run(testName, func(t *testing.T) {
			actual := ft.LowerBound(tc.w)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

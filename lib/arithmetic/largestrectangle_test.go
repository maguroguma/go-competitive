package arithmetic

import (
	"fmt"
	"testing"
)

// https://atcoder.jp/contests/abc189/tasks/abc189_c
func TestLargestRectangle(t *testing.T) {
	testcases := []struct {
		height   []int64
		expected int64
	}{
		{
			[]int64{2, 4, 4, 9, 4, 9},
			20,
		},
		{
			[]int64{200, 4, 4, 9, 4, 9},
			200,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := LargestRectangle(tc.height)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

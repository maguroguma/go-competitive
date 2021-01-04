package inversion_number

import (
	"fmt"
	"testing"
)

func TestInversionNumber(t *testing.T) {
	testcases := []struct {
		A        []int64
		expected int64
	}{
		{
			[]int64{3, 1, 5, 4, 2},
			5,
		},
		{
			[]int64{1, 2, 3, 4, 5, 6},
			0,
		},
		{
			[]int64{7, 6, 5, 4, 3, 2, 1},
			21,
		},
		{
			[]int64{19, 11, 10, 7, 8, 9, 17, 18, 20, 4, 3, 15, 16, 1, 5, 14, 6, 2, 13, 12},
			114,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := InversionNumber(tc.A)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

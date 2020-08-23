package arithmetic

import (
	"fmt"
	"testing"
)

func TestIsSquareNumber(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{0, true}, {1, true}, {4, true}, {9, true}, {12345 * 12345, true},
		{2, false}, {3, false}, {8, false}, {12345*12345 - 1, false},
	}

	for i, tc := range testcases {
		subTest := fmt.Sprintf("%d sub test: %v is %v", i, tc.n, tc.expected)
		t.Run(subTest, func(t *testing.T) {
			actual := IsSquareNumber(tc.n)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

func TestIsCubeNumber(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{0, true}, {1, true}, {8, true}, {27, true}, {12345 * 12345 * 12345, true},
		{2, false}, {7, false}, {26, false}, {12345*12345*12345 - 1, false},
	}

	for i, tc := range testcases {
		subTest := fmt.Sprintf("%d sub test: %v is %v", i, tc.n, tc.expected)
		t.Run(subTest, func(t *testing.T) {
			actual := IsCubeNumber(tc.n)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

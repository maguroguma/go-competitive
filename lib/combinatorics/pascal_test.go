package combinatorics

import (
	"reflect"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	n := 5
	C := [][]int{
		{1, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0},
		{1, 2, 1, 0, 0, 0},
		{1, 3, 3, 1, 0, 0},
		{1, 4, 6, 4, 1, 0},
		{1, 5, 10, 10, 5, 1},
	}

	pt := NewPascalTriangle(n)

	if !reflect.DeepEqual(pt.CombTable, C) {
		t.Errorf("table is wrong: got %v, want %v", pt.CombTable, C)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			actual := pt.C(i, j)
			expected := C[i][j]
			if actual != expected {
				t.Errorf("C(%d, %d) is wrong: got %v, want %v", i, j, actual, expected)
			}
		}
	}
}

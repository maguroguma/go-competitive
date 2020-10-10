package grid

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGridLRUD(t *testing.T) {
	testcases := []struct {
		S          [][]rune
		L, R, U, D [][]int
	}{
		{
			S: [][]rune{
				[]rune("..#"),
				[]rune("#.."),
			},
			L: [][]int{
				{0, 1, -1},
				{-1, 0, 1},
			},
			R: [][]int{
				{1, 0, -1},
				{-1, 1, 0},
			},
			U: [][]int{
				{0, 0, -1},
				{-1, 1, 0},
			},
			D: [][]int{
				{0, 1, -1},
				{-1, 0, 0},
			},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			L, R, U, D := GridLRUD(tc.S)

			ok := reflect.DeepEqual(L, tc.L)
			ok = ok && reflect.DeepEqual(R, tc.R)
			ok = ok && reflect.DeepEqual(U, tc.U)
			ok = ok && reflect.DeepEqual(D, tc.D)

			if !ok {
				t.Errorf("actual: %v, %v, %v, %v", L, R, U, D)
				t.Errorf("expected: %v, %v, %v, %v", tc.L, tc.R, tc.U, tc.D)
			}
		})
	}
}

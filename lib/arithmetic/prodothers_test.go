package arithmetic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProdOthers(t *testing.T) {
	_max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	_min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	var _gcd func(x, y int) int
	_gcd = func(x, y int) int {
		if y == 0 {
			return x
		}
		return _gcd(y, x%y)
	}

	testcases := []struct {
		A        []int
		f        func(x, y int) int
		expected []int
	}{
		{
			[]int{0, 1, 2, 3},
			_max,
			[]int{3, 3, 3, 2},
		},
		{
			[]int{0, 1, 2, 3},
			_min,
			[]int{1, 0, 0, 0},
		},
		{
			[]int{0, 2, 4, 8},
			_gcd,
			[]int{2, 4, 2, 2},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := ProdOthers(tc.A, tc.f)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

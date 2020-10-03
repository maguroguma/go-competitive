package numtheo

import (
	"fmt"
	"reflect"
	"testing"
)

// https://onlinejudge.u-aizu.ac.jp/courses/library/6/NTL/1/NTL_1_A
func TestTrialDivision(t *testing.T) {
	testcases := []struct {
		n int
		P map[int]int
	}{
		{
			n: 12,
			P: map[int]int{2: 2, 3: 1},
		},
		{
			n: 126,
			P: map[int]int{2: 1, 3: 2, 7: 1},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			P := TrialDivision(tc.n)
			if !reflect.DeepEqual(P, tc.P) {
				t.Errorf("got %v, want %v", P, tc.P)
			}
		})
	}
}

// https://onlinejudge.u-aizu.ac.jp/courses/library/6/NTL/1/NTL_1_D
func TestEulerPhi(t *testing.T) {
	testcases := []struct {
		n   int
		num int
	}{
		{n: 6, num: 2}, {n: 1000000, num: 400000},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			num := EulerPhi(tc.n)
			if num != tc.num {
				t.Errorf("got %v, want %v", num, tc.num)
			}
		})
	}
}

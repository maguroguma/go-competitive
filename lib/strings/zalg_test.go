package strings

import (
	"fmt"
	"reflect"
	"testing"
)

// https://judge.yosupo.jp/problem/zalgorithm
func TestZAlgorithm(t *testing.T) {
	testcases := []struct {
		S        []rune
		expected []int
	}{
		{
			[]rune("abcbcba"), []int{7, 0, 0, 0, 0, 0, 1},
		},
		{
			[]rune("mississippi"), []int{11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[]rune("ababacaca"), []int{9, 0, 3, 0, 1, 0, 1, 0, 1},
		},
		{
			[]rune("aaaaa"), []int{5, 4, 3, 2, 1},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := ZAlgorithm(tc.S)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

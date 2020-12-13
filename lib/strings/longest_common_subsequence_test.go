package strings

import (
	"fmt"
	"testing"
)

// https://atcoder.jp/contests/dp/tasks/dp_f
func TestLongestCommonSubsequence(t *testing.T) {
	testcases := []struct {
		S, T     []rune
		expected int
	}{
		{
			[]rune("axyb"),
			[]rune("abyxb"),
			3,
		},
		{
			[]rune("aa"),
			[]rune("xayaz"),
			2,
		},
		{
			[]rune("a"),
			[]rune("z"),
			0,
		},
		{
			[]rune("abracadabra"),
			[]rune("avadakedavra"),
			7,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := LCS(tc.S, tc.T)
			if len(actual) != tc.expected {
				t.Errorf("got %v, want %v", len(actual), tc.expected)
			}
		})
	}
}

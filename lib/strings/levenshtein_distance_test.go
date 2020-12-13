package strings

import (
	"fmt"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	testcases := []struct {
		S, T     []rune
		expected int
	}{
		{
			[]rune("abcdefghi"),
			[]rune("acdefxhij"),
			3,
		},
		{
			[]rune("abc"),
			[]rune("addc"),
			2,
		},
		{
			[]rune("pirikapirirara"),
			[]rune("poporinapeperuto"),
			10,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := LevenshteinDistance(tc.S, tc.T)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

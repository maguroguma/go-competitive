package strings

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubstrings(t *testing.T) {
	testcases := []struct {
		S, T     []rune
		expected int
	}{
		{
			[]rune(string("ABRACADABRA")),
			[]rune(string("ECADADABRBCRDARA")),
			5,
		},
		{
			[]rune(string("UPWJCIRUCAXIIRGL")),
			[]rune(string("SBQNYBSBZDFNEV")),
			0,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := LongestCommonSubstring(tc.S, tc.T)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

package strings

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestPMA(t *testing.T) {
	testcases := []struct {
		text     []rune
		patterns [][]rune
		expected [][]int
	}{
		{
			text: []rune("ABCDABCD"),
			patterns: [][]rune{
				[]rune("A"),
				[]rune("DA"),
				[]rune("ABCDABCD"),
			},
			expected: [][]int{
				{0}, {}, {}, {}, {0, 1}, {}, {}, {2},
			},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			pma := NewPMA(tc.patterns, 'A')
			actual := pma.Match(tc.text)

			for _, A := range actual {
				sort.Sort(sort.IntSlice(A))
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

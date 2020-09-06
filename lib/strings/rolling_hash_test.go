package strings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOffsetHash(t *testing.T) {
	testcases := []struct {
		strs     []string
		expected [][2]int
	}{
		{
			strs: []string{"aaa", "bbb", "ccc", "aaa", "bbb", "ccc"},
			expected: [][2]int{
				{0, 3}, {1, 4}, {2, 5},
			},
		},
	}

	for i, tc := range testcases {
		subTitle := fmt.Sprintf("%d testcase", i)
		t.Run(subTitle, func(t *testing.T) {
			H := make([]*RHash, len(tc.strs))
			for i := 0; i < len(tc.strs); i++ {
				H[i] = NewRHash(tc.strs[i])
			}

			actual := [][2]int{}
			for i := 0; i < len(tc.strs); i++ {
				for j := i + 1; j < len(tc.strs); j++ {
					ihash := H[i].OffsetHash(0, len(tc.strs[i]))
					jhash := H[j].OffsetHash(0, len(tc.strs[j]))
					if ihash == jhash {
						actual = append(actual, [2]int{i, j})
					}
				}
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

func TestSliceHash(t *testing.T) {
	testcases := []struct {
		strs     []string
		expected [][2]int
	}{
		{
			strs: []string{"aaa", "bbb", "ccc", "aaa", "bbb", "ccc"},
			expected: [][2]int{
				{0, 3}, {1, 4}, {2, 5},
			},
		},
	}

	for i, tc := range testcases {
		subTitle := fmt.Sprintf("%d testcase", i)
		t.Run(subTitle, func(t *testing.T) {
			H := make([]*RHash, len(tc.strs))
			for i := 0; i < len(tc.strs); i++ {
				H[i] = NewRHash(tc.strs[i])
			}

			actual := [][2]int{}
			for i := 0; i < len(tc.strs); i++ {
				for j := i + 1; j < len(tc.strs); j++ {
					ihash := H[i].SliceHash(2, len(tc.strs[i]))
					jhash := H[j].SliceHash(2, len(tc.strs[j]))
					if ihash == jhash {
						actual = append(actual, [2]int{i, j})
					}
				}
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

func TestLen(t *testing.T) {
	testcases := []struct {
		str string
	}{
		{"a"},
		{"aa"},
		{"aaa"},
		{"aA"},
		{"a a"},
		{"aBcDeF"},
	}

	for i, tc := range testcases {
		subTitle := fmt.Sprintf("%d testcase", i)
		t.Run(subTitle, func(t *testing.T) {
			h := NewRHash(tc.str)
			actual := h.Len()
			if actual != len(tc.str) {
				t.Errorf("got %v, want %v", actual, len(tc.str))
			}
		})
	}
}

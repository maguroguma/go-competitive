package strings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKMPTable(t *testing.T) {
	testcases := []struct {
		pattern  []rune
		expected []int
	}{
		{
			[]rune("aabaabaaa"),
			[]int{-1, 0, 1, 0, 1, 2, 3, 4, 5, 2},
		},
		{
			[]rune("aabaab"),
			[]int{-1, 0, 1, 0, 1, 2, 3},
		},
		{
			[]rune(""),
			[]int{-1},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("%d test", i)
		t.Run(testName, func(t *testing.T) {
			kmp := NewKMP(tc.pattern)
			actual := kmp.kmpTable
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

func TestKMPSearch(t *testing.T) {
	testcases := []struct {
		pattern  []rune
		text     []rune
		expected []int
	}{
		{
			[]rune("aabaab"),
			[]rune("aaabaabaaa"),
			[]int{1},
		},
		{
			[]rune("AA"),
			[]rune("AAAA"),
			[]int{0, 1, 2},
		},
		{
			[]rune(""),
			[]rune("aabaab"),
			[]int{},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("%d test", i)
		t.Run(testName, func(t *testing.T) {
			kmp := NewKMP(tc.pattern)
			actual := kmp.Search(tc.text)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

func TestKMPStringPeriod(t *testing.T) {
	testcases := []struct {
		pattern  []rune
		expected []int
	}{
		{
			[]rune("abababcaa"),
			[]int{1, 2, 2, 2, 2, 2, 7, 7, 8},
		},
		{
			[]rune(""),
			[]int{},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("%d test", i)
		t.Run(testName, func(t *testing.T) {
			kmp := NewKMP(tc.pattern)
			actual := kmp.Periods()

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

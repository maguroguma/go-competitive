package strings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRLE(t *testing.T) {
	testcases := []struct {
		S    []rune
		comp []rune
		cnts []int
	}{
		{
			[]rune("a"),
			[]rune("a"),
			[]int{1},
		},
		{
			[]rune("zzaaz"),
			[]rune("zaz"),
			[]int{2, 2, 1},
		},
		{
			[]rune("ccff"),
			[]rune("cf"),
			[]int{2, 2},
		},
		{
			[]rune("cbddbb"),
			[]rune("cbdb"),
			[]int{1, 1, 2, 2},
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			comp, cnts := RunLengthEncoding(tc.S)

			if string(comp) != string(tc.comp) || !reflect.DeepEqual(cnts, tc.cnts) {
				t.Errorf("got (%v, %v), want (%v, %v)", string(comp), cnts, string(tc.comp), tc.cnts)
			}

			decode := RunLengthDecoding(comp, cnts)
			if string(decode) != string(tc.S) {
				t.Errorf("got %v, want %v", string(decode), string(tc.S))
			}
		})
	}
}

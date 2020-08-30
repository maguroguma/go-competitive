package segment

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSegment(t *testing.T) {
	testcases := []struct {
		sourceSegments [][2]int
		resultSegments [][2]int
	}{
		{
			[][2]int{
				{-100, 0}, {1, 100}, {-100, 0},
			},
			[][2]int{
				{-100, 0}, {1, 100},
			},
		},
		{
			[][2]int{
				{1, 2}, {0, 1}, {-1, 0},
			},
			[][2]int{
				{-1, 2},
			},
		},
		{
			[][2]int{},
			[][2]int{},
		},
		{
			[][2]int{{0, 100}},
			[][2]int{{0, 100}},
		},
	}

	for i, tc := range testcases {
		subTitle := fmt.Sprintf("%d case", i)
		t.Run(subTitle, func(t *testing.T) {
			actual := MergeSegments(tc.sourceSegments)
			if !reflect.DeepEqual(actual, tc.resultSegments) {
				t.Errorf("got %v, want %v", actual, tc.resultSegments)
			}
		})
	}
}

package bipartite_graph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsBipartiteGraph(t *testing.T) {
	testcases := []struct {
		G      [][]int
		colors []int
		ok     bool
	}{
		{
			G: [][]int{
				{1, 2}, {0, 3}, {0, 3}, {1, 2},
			},
			colors: []int{1, -1, -1, 1},
			ok:     true,
		},
		{
			G: [][]int{
				{1, 2, 3}, {0, 3}, {0, 3}, {0, 1, 2},
			},
			colors: []int{1, -1, 0, 1},
			ok:     false,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			ok, colors := IsBipartiteGraph(tc.G)

			if !reflect.DeepEqual(colors, tc.colors) || ok != tc.ok {
				t.Errorf("got (%v, %v), want (%v, %v)", ok, colors, tc.ok, tc.colors)
			}
		})
	}
}

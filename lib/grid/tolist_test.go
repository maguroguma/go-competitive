package grid

import (
	"reflect"
	"testing"
)

func TestGridToAdjacencyList(t *testing.T) {
	expG := [][]int{
		{1, 2},
		{0, 3},
		{3, 0},
		{2, 1},
	}
	expN := 4

	G, N := GridToAdjacencyList(2, 2)
	if !reflect.DeepEqual(G, expG) || N != expN {
		t.Errorf("got (%v, %v), want (%v, %v)", G, N, expG, expN)
	}
}

package grid

import "testing"

func TestLargestSquare(t *testing.T) {
	G := [][]int{
		{0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
	}
	expected := 2

	actual := LargestSquare(G)
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

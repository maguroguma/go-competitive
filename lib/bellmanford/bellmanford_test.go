package bellmanford

import (
	"os"
	"testing"
)

var (
	n, m int
	L    []Edge
)

func TestMain(m *testing.M) {
	println("before all...")

	initTest()
	code := m.Run()

	println("after all...")
	os.Exit(code)
}

// ABC061-D Sample 3
func initTest() {
	n, m = 6, 5
	Inputs := [][3]int{
		{1, 2, -1000000000},
		{2, 3, -1000000000},
		{3, 4, -1000000000},
		{4, 5, -1000000000},
		{5, 6, -1000000000},
	}

	for _, inp := range Inputs {
		a, b, c := inp[0], inp[1], inp[2]
		a--
		b--

		L = append(L, Edge{a, b, -c})
	}
}

func TestBellmanford(t *testing.T) {
	// v, e = n, m
	dist, _ := Bellmanford(0, n, L)

	if -dist[n-1] != -5000000000 {
		t.Errorf("got %v, want %v", -dist[n-1], -5000000000)
	}
}

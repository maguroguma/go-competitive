package treediameter

import (
	"os"
	"testing"
)

var (
	n int

	G [200000 + 50][]Edge
)

func initTest() {
	n = 4
	Inputs := [][3]int{
		{0, 1, 2},
		{1, 2, 1},
		{1, 3, 3},
	}

	for _, inp := range Inputs {
		s, t, w := inp[0], inp[1], inp[2]
		G[s] = append(G[s], Edge{t, w})
		G[t] = append(G[t], Edge{s, w})
	}
}

func TestMain(m *testing.M) {
	println("before all...")

	initTest()
	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func TestTDiameter(t *testing.T) {
	actual := TreeDiameter(G[:n])

	if actual != 5 {
		t.Errorf("got %v, want %v", actual, 5)
	}
}

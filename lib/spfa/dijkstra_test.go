package spfa

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// https://atcoder.jp/contests/abc143/tasks/abc143_e
// Sample No.3

const (
	INF_BIT60 = 1 << 60
)

var (
	n, l     int
	Queries  [][2]int
	Expected []int

	ds *DijkstraSolver
	A  [300 + 5][300 + 5]int
	G  [300 + 5][]Edge
)

func initTest() {
	n, l = 5, 4

	E := [][3]int{
		{1, 2, 2}, {2, 3, 2}, {3, 4, 3}, {4, 5, 2},
	}
	for _, e := range E {
		a, b, c := e[0], e[1], e[2]
		a--
		b--

		w := EdgeWeight{gas: c}
		G[a] = append(G[a], Edge{to: b, ew: w})
		G[b] = append(G[b], Edge{to: a, ew: w})
	}

	Q := [][2]int{
		{2, 1}, {3, 1}, {4, 1}, {5, 1},
		{1, 2}, {3, 2}, {4, 2}, {5, 2},
		{1, 3}, {2, 3}, {4, 3}, {5, 3},
		{1, 4}, {2, 4}, {3, 4}, {5, 4},
		{1, 5}, {2, 5}, {3, 5}, {4, 5},
	}
	for _, q := range Q {
		s, t := q[0], q[1]
		s--
		t--

		Queries = append(Queries, [2]int{s, t})
	}

	Expected = []int{
		0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 0, 1, 1, 1, 0, 0, 2, 2, 1, 0,
	}
}

func TestMain(m *testing.M) {
	println("before all...")

	initTest()

	code := m.Run()

	println("after all...")

	for i := 0; i < n; i++ {
		fmt.Println(A[i][:n])
	}

	os.Exit(code)
}

func TestDijkstraGeneric(t *testing.T) {
	vinf := V{gas: -1, times: INF_BIT60}
	vzero := V{gas: l, times: 0}
	less := func(l, r V) bool {
		if l.times < r.times {
			return true
		} else if l.times > r.times {
			return false
		} else {
			return l.gas > r.gas
		}
	}
	genNextV := func(cv *Vertex, e Edge) V {
		if l < e.ew.gas {
			return vinf
		}

		if cv.v.gas >= e.ew.gas {
			return V{gas: cv.v.gas - e.ew.gas, times: cv.v.times}
		}

		return V{gas: l - e.ew.gas, times: cv.v.times + 1}
	}

	ds := NewDijkstraSolver(vinf, less, genNextV)

	for i := 0; i < n; i++ {
		dp := ds.Dijkstra([]StartPoint{{id: i, vzero: vzero}}, n, G[:n])

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			A[i][j] = dp[j].times
		}
	}

	Actuals := []int{}
	for _, q := range Queries {
		s, t := q[0], q[1]

		if A[s][t] >= INF_BIT60 {
			Actuals = append(Actuals, -1)
		} else {
			Actuals = append(Actuals, A[s][t])
		}
	}

	if !reflect.DeepEqual(Actuals, Expected) {
		t.Errorf("got %v, want %v", Actuals, Expected)
	}
}

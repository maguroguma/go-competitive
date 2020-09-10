package tsort

import (
	"os"
	"testing"
)

var (
	n int
	A [][]int

	num int
	G   [1000000 + 50][]int
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ABC139-E Sample 2.
func initTest() {
	n = 4
	A = [][]int{
		{2, 3, 4}, {1, 3, 4}, {4, 1, 2}, {3, 1, 2},
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j]--
		}
	}

	for i := 0; i < n; i++ {
		for j := 1; j < len(A[i]); j++ {
			bef := A[i][j-1]
			aft := A[i][j]

			cid := min(i, bef) + n*max(i, bef)
			nid := min(i, aft) + n*max(i, aft)
			G[cid] = append(G[cid], nid)
		}
	}

	num = n * n
}

func TestMain(m *testing.M) {
	println("before all...")

	initTest()
	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func TestTSort(t *testing.T) {
	ok, sorted := TSort(num, G[:num])

	if !ok {
		t.Errorf("graph couldn't be topological sorted")
	}

	actual, _ := LongestPath(sorted, G[:num])
	if actual+1 != 4 {
		t.Errorf("got %v, want %v", actual, 4)
	}
}

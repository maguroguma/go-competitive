package scc

import (
	"os"
	"reflect"
	"testing"
)

var (
	n, m int
	A, B []int
)

// https://atcoder.jp/contests/practice2/tasks/practice2_g
func initTest() {
	n, m = 6, 7
	A = []int{1, 5, 3, 5, 4, 0, 4}
	B = []int{4, 2, 0, 5, 1, 3, 2}
}

func TestMain(m *testing.M) {
	println("before all...")

	initTest()
	code := m.Run()

	println("after all...")

	os.Exit(code)
}

// check only degrades
func TestScc(t *testing.T) {
	expected := [][]int{
		{5}, {1, 4}, {2}, {0, 3},
	}

	scc := NewSccGraph(n)
	for i := 0; i < m; i++ {
		from, to := A[i], B[i]
		scc.AddEdge(from, to)
	}

	actual := scc.Scc()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

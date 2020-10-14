package bit

import (
	"os"
	"testing"
)

var (
	n, m, x int
	C       []int
	A       [][]int
)

const EXPECTED = 1067

func initTest() {
	n, m, x = 8, 5, 22
	C = make([]int, n)
	A = make([][]int, n)
	mat := [][]int{
		{100, 3, 7, 5, 3, 1},
		{164, 4, 5, 2, 7, 8},
		{334, 7, 2, 7, 2, 9},
		{234, 4, 7, 2, 8, 2},
		{541, 5, 4, 3, 3, 6},
		{235, 4, 8, 6, 9, 7},
		{394, 3, 6, 1, 6, 2},
		{872, 8, 4, 3, 7, 2},
	}
	for i, row := range mat {
		C[i] = row[0]
		A[i] = row[1:]
	}
}

func TestMain(m *testing.M) {
	println("before all...")

	initTest()
	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func Test01(t *testing.T) {
	ans := 1 << uint(60)

	BruteForceByBits01(n, func(B []int) {
		p := 0
		X := make([]int, m)
		for i, a := range B {
			if a == 0 {
				continue
			}
			p += C[i]
			for j := range A[i] {
				X[j] += A[i][j]
			}
		}

		for _, xx := range X {
			if xx < x {
				return
			}
		}

		if ans > p {
			ans = p
		}
	})

	if ans != EXPECTED {
		t.Errorf("got %v, want %v", ans, EXPECTED)
	}
}

func TestTF(t *testing.T) {
	ans := 1 << uint(60)

	BruteForceByBitsTF(n, func(B []bool) {
		p := 0
		X := make([]int, m)
		for i, a := range B {
			if !a {
				continue
			}
			p += C[i]
			for j := range A[i] {
				X[j] += A[i][j]
			}
		}

		for _, xx := range X {
			if xx < x {
				return
			}
		}

		if ans > p {
			ans = p
		}
	})

	if ans != EXPECTED {
		t.Errorf("got %v, want %v", ans, EXPECTED)
	}
}

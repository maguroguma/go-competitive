package modint

import (
	"fmt"
	"testing"
)

func TestBasics(t *testing.T) {
	SetMod(Mod1000000007)

	a := Mint(1)
	b := a.Add(10)
	if !(a == 1 && b == 11) {
		t.Errorf("(a, b): got (%v, %v), want (%v, %v)", a, b, 1, 11)
	}

	c := Mint(1)
	c.AddAs(10)
	if c != 11 {
		t.Errorf("got %v, want %v", c, 11)
	}
	d := c.Add(100)
	if !(c == 11 && d == 111) {
		t.Errorf("(c, d): got (%v, %v), want (%v, %v)", c, d, 11, 111)
	}
}

// https://atcoder.jp/contests/abc145/tasks/abc145_d
func TestABC145D(t *testing.T) {
	testcases := []struct {
		x, y     int
		expected Mint
	}{
		{3, 3, 2},
		{2, 2, 0},
		{999999, 999999, 151840682},
	}

	SetMod(Mod1000000007)
	cf := NewCombFactorial(1000000 + 100)

	solver := func(cx, cy Mint) Mint {
		if (cx+cy)%3 != 0 {
			return 0
		}

		n := (cx + cy) / 3
		y := cy - n
		if y < 0 || y > n {
			return 0
		}

		return cf.C(n, y)
	}

	fmt.Println(cf.fact(5))

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			x, y, expected := Mint(tc.x), Mint(tc.y), tc.expected
			actual := solver(x, y)
			if actual != expected {
				t.Errorf("got %v, want %v", actual, expected)
			}
		})
	}
}

// https://atcoder.jp/contests/abc156/tasks/abc156_e
func TestABC156E(t *testing.T) {
	testcases := []struct {
		n, k     Mint
		expected Mint
	}{
		{3, 2, 10},
		{200000, 1000000000, 607923868},
		{15, 6, 22583772},
	}

	SetMod(Mod1000000007)
	cf := NewCombFactorial(1000000 + 100)

	solver := func(n, k Mint) Mint {
		ans := Mint(0)

		mini := n - 1
		if k < n-1 {
			mini = k
		}
		for x := Mint(0); x <= mini; x++ {
			res := cf.H(n.Sub(x), x)
			res.MulAs(cf.C(n, x))
			ans.AddAs(res)
		}

		return ans
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := solver(tc.n, tc.k)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}

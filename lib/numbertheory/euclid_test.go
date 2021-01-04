package numtheo

import (
	"fmt"
	"testing"
)

// https://qiita.com/drken/items/b97ff231e43bce50199a
func TestExtGCD(t *testing.T) {
	testcases := []struct {
		a, b      int
		gcd, x, y int
	}{
		{
			a: 111, b: 30,
			gcd: 3, x: 3, y: -11,
		},
		{
			a: 30, b: 111,
			gcd: 3, x: -11, y: 3,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			gcd, x, y := ExtGCD(tc.a, tc.b)
			if !(gcd == tc.gcd && x == tc.x && y == tc.y) {
				t.Errorf(
					"got (gcd: %v, x: %v, y: %v), want (gcd: %v, x: %v, y: %v)",
					gcd, x, y, tc.gcd, tc.x, tc.y,
				)
			}
		})
	}
}

func TestModInvByExtGCD(t *testing.T) {
	testcases := []struct {
		a, m int
		ia   int
		ok   bool
	}{
		{a: 3, m: 10, ia: 7, ok: true},
		{a: 2, m: 10, ia: -1, ok: false},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			ia, ok := ModInvByExtGCD(tc.a, tc.m)
			if ia != tc.ia || ok != tc.ok {
				t.Errorf("got (%v, %v), want (%v, %v)", ia, ok, tc.ia, tc.ok)
			}
		})
	}
}

// https://atcoder.jp/contests/abc186/tasks/abc186_e
func TestCongruenceEquation(t *testing.T) {
	testcases := []struct {
		a, b, m int
		x       int
		ok      bool
	}{
		{a: 3, b: -4, m: 10, x: 2, ok: true},
		{a: 2, b: -11, m: 1000, x: -1, ok: false},
		{a: 595591169, b: -897581057, m: 998244353, x: 249561088, ok: true},
		{a: 14, b: -6, m: 10000, x: 3571, ok: true},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			x, ok := CongruenceEquation(tc.a, tc.b, tc.m)
			if x != tc.x || ok != tc.ok {
				t.Errorf("got (%v, %v), want (%v, %v)", x, ok, tc.x, tc.ok)
			}
		})
	}
}

func TestGcd(t *testing.T) {
	testcases := []struct {
		a, b int
		g    int
	}{
		{
			a: 30, b: 10, g: 10,
		},
		{
			a: 30, b: 0, g: 30,
		},
		{
			a: 0, b: 30, g: 30,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			g := Gcd(tc.a, tc.b)
			if g != tc.g {
				t.Errorf("got %v, want %v", g, tc.g)
			}
		})
	}
}

// https://onlinejudge.u-aizu.ac.jp/courses/library/6/NTL/1/NTL_1_C
func TestLcm(t *testing.T) {
	testcases := []struct {
		A []int
		l int
	}{
		{
			[]int{3, 4, 6}, 12,
		},
		{
			[]int{1, 2, 3, 5}, 30,
		},
		{
			[]int{10, 30, 0}, 0,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			l := 1
			for _, a := range tc.A {
				l = Lcm(l, a)
			}

			if l != tc.l {
				t.Errorf("got %v, want %v", l, tc.l)
			}
		})
	}
}

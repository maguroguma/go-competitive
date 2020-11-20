package binsearch

import (
	"math"
	"sort"
	"testing"
)

// ABC077-C: Sample-3
// https://atcoder.jp/contests/abc077/tasks/arc084_a
func TestBinarySearch(t *testing.T) {
	expected := int64(87)

	n := int64(6)
	A := []int64{3, 14, 159, 2, 6, 53}
	B := []int64{58, 9, 79, 323, 84, 6}
	C := []int64{2643, 383, 2, 79, 50, 288}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})
	sort.Slice(C, func(i, j int) bool {
		return C[i] < C[j]
	})

	ans := int64(0)
	for i := int64(0); i < n; i++ {
		a := BinarySearch(-1, n, func(mid int64) bool {
			return A[mid] < B[i]
		})
		c := BinarySearch(n, -1, func(mid int64) bool {
			return C[mid] > B[i]
		})

		x := a + 1
		y := n - c

		ans += x * y
	}

	if ans != expected {
		t.Errorf("got %v, want %v", ans, expected)
	}
}

// ABC144-D: Sample-2
// https://atcoder.jp/contests/abc144/tasks/abc144_d
func TestBinarySearchFloat64(t *testing.T) {
	expected := 89.7834636934

	a, b, x := 12, 21, 10
	af, bf, xf := float64(a), float64(b), float64(x)

	ok := BinarySearchFloat64(0.0, 90.0, func(mid float64) bool {
		rad := mid * math.Pi / 180.0

		var V float64
		if math.Tan(rad) >= bf/af {
			S := (1.0 / 2.0) * bf * bf * (1.0 / math.Tan(rad))
			V = S * af
		} else {
			S := (1.0 / 2.0) * af * af * math.Tan(rad)
			V = af*af*bf - S*af
		}

		if V >= xf {
			return true
		} else {
			return false
		}
	})

	if math.Abs(ok-expected) >= 1e-6 {
		t.Errorf("got %v, want %v", ok, expected)
	}
}

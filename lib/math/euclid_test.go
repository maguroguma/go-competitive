package math

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

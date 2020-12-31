package numtheo

import "errors"

// ExtGCD solve an equation: "a*x + b*y = Gcd(a, b)"
// (x, y) satisfies that
// "|x|+|y| is minimized" and "if there are some such ones, x <= y".
// g is Gcd(a, b)
func ExtGCD(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}

	g, s, t := ExtGCD(b, a%b)

	return g, t, s - (a/b)*t
}

// ModInvByExtGCD calculates x^-1 (mod m) by ExtGCD.
// m can not necessarily be prime number.
// ok is false if gcd(x, m) != 1 because x^-1 does not exist.
func ModInvByExtGCD(x, m int) (ix int, ok bool) {
	g, ix, _ := ExtGCD(x, m)

	if g != 1 {
		return -1, false
	}

	if ix < 0 {
		ix += m
	}

	return ix, true
}

// CongruenceEquation calculates x that satisfies an equation a*x == b (mod m).
// No answer exists if ok is false.
func CongruenceEquation(a, b, m int) (x int, ok bool) {
	_mod := func(val, m int) int {
		res := val % m
		if res < 0 {
			res += m
		}
		return res
	}

	a, b = _mod(a, m), _mod(b, m)

	g, ia, _ := ExtGCD(a, m)
	ia = _mod(ia, m)
	if g == 1 {
		return _mod(ia*b, m), true
	}

	if b%g != 0 {
		return -1, false
	}

	a, b, m = a/g, b/g, m/g
	_, ia, _ = ExtGCD(a, m)
	ia = _mod(ia, m)

	return _mod(ia*b, m), true
}

// Gcd returns the Greatest Common Divisor of two natural numbers.
// Gcd only accepts two natural numbers (a, b >= 0).
// Negative number causes panic.
// Gcd uses the Euclidean Algorithm.
func Gcd(a, b int) int {
	if a < 0 || b < 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// Lcm returns the Least Common Multiple of two natural numbers.
// Lcd only accepts two natural numbers (a, b >= 0).
// Negative number causes panic.
// Lcd uses the Euclidean Algorithm indirectly.
func Lcm(a, b int) int {
	if a < 0 || b < 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	if a == 0 || b == 0 {
		return 0
	}

	// a = a'*gcd, b = b'*gcd, a*b = a'*b'*gcd^2
	// a' and b' are relatively prime numbers
	// gcd consists of prime numbers, that are included in a and b
	g := Gcd(a, b)

	// not (a * b / gcd), because of reducing a probability of overflow
	return (a / g) * b
}

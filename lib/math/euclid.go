package math

// ExtGCD solve an equation: "a*x + b*y = Gcd(a, b)"
// g is Gcd(a, b)
func ExtGCD(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}

	g, s, t := ExtGCD(b, a%b)

	return g, t, s - (a/b)*t
}

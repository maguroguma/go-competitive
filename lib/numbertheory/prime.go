package numtheo

import "errors"

// TrialDivision returns the result of prime factorization of integer N.
// Time complexity: O(sqrt(N))
func TrialDivision(n int) map[int]int {
	var i, exp int
	p := map[int]int{}

	if n <= 1 {
		panic(errors.New("[argument error]: TrialDivision only accepts a NATURAL number"))
	}

	for i = 2; i*i <= n; i++ {
		exp = 0
		for n%i == 0 {
			exp++
			n /= i
		}

		if exp == 0 {
			continue
		}
		p[i] = exp
	}
	if n > 1 {
		p[n] = 1
	}

	return p
}

// EulerPhi returns the number of x that satisfies "x in [1, N]" and "Gcd(N, x) == 1".
// Time complexity: O(sqrt(N))
func EulerPhi(n int) int {
	P := TrialDivision(n)
	res := n

	for p := range P {
		res /= p
		res *= p - 1
	}

	return res
}

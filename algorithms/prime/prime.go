package prime

import (
	"errors"
	"math"
)

// TrialDivision returns the result of prime factorization of integer N.
func TrialDivision(n int) map[int]int {
	if n <= 0 {
		panic(errors.New("[argument error]: TrialDivision only accepts a NATURAL number"))
	}
	if n == 1 {
		return map[int]int{1: 1}
	}

	p := map[int]int{}
	sqrt := math.Pow(float64(n), 0.5)
	for i := 2; i <= int(sqrt); i++ {
		exp := 0
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

// IsPrime judges whether an argument integer is a prime number or not.
func IsPrime(n int) bool {
	if n == 1 {
		return false
	}

	sqrt := math.Pow(float64(n), 0.5)
	for i := 2; i <= int(sqrt); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

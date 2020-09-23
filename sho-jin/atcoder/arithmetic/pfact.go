package arithmetic

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

// SieveOfEratosthenes returns the prime numbers less than inter N.
func SieveOfEratosthenes(n int) []int {
	primes := []int{}

	return primes
}

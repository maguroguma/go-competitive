package prime

import "fmt"

// originated from: https://qiita.com/rsk0315_h4x/items/ff3b542a4468679fb409

type Osa_kSieve struct {
	n      int
	minp   []int
	errMsg string
}

// NewOsa_kSieve returns sieve for [0, n).
func NewOsa_kSieve(n int) *Osa_kSieve {
	os := new(Osa_kSieve)
	os.n = n
	os.errMsg = fmt.Sprintf("This Osa_k accept less and equal than %d", os.n)

	// Initialize
	os.minp = make([]int, os.n)
	for i := 0; i < os.n; i++ {
		os.minp[i] = i
	}
	for i := 2; i*i < os.n; i++ {
		if os.minp[i] < i {
			continue
		}

		for j := i * i; j < os.n; j += i {
			if os.minp[j] == j {
				os.minp[j] = i
			}
		}
	}

	return os
}

// Factors returns prime factors consisting of a.
func (os *Osa_kSieve) Factors(a int) []int {
	if a <= 1 || a >= os.n {
		panic(os.errMsg)
	}

	res := []int{}

	for a > 1 {
		res = append(res, os.minp[a])
		a /= os.minp[a]
	}

	return res
}

// IsPrime returns whether a is prime number or not.
func (os *Osa_kSieve) IsPrime(a int) bool {
	return !(os.minp[a] < a)
}

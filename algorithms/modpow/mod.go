package modpow

// CalcNegativeMod can calculate a right residual whether value is positive or negative.
func CalcNegativeMod(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

func modpow(a, e, m int) int {
	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := modpow(a, halfE, m)
		return half * half % m
	}

	return a * modpow(a, e-1, m) % m
}

// CalcModInv returns $a^{-1} mod m$ by Fermat's little theorem.
func CalcModInv(a, m int) int {
	return modpow(a, m-2, m)
}

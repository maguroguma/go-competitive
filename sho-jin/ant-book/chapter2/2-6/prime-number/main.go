package main

import "fmt"

func main() {
	fmt.Println(isPrime(1000000000 + 7))
	fmt.Println(isPrime(617))
	fmt.Println(divisor(1234))
	fmt.Println(primeFactor(100))

	num, primes := sieve(1000000)
	fmt.Println(num)
	fmt.Println(primes[:100])
}

// 入力はすべて正とする
// 素数判定: O(sqrt(n))
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n != 1 // 1の場合は例外
}

// 約数の列挙: O(sqrt(n))
func divisor(n int) []int {
	res := []int{}

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i != (n / i) {
				res = append(res, (n / i))
			}
		}
	}

	return res
}

// 素因数分解: O(sqrt(n))
func primeFactor(n int) map[int]int {
	res := make(map[int]int)

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			res[i]++
			n /= i
		}
	}

	if n != 1 {
		res[n] = 1
	}

	return res
}

// エラトステネスの篩: O(NloglogN)
// n以下の素数の数と素数のスライスを返す
func sieve(n int) (int, []int) {
	// var prime [2000000 + 5]int
	prime := []int{}
	var isPrime [2000000 + 5 + 1]bool

	p := 0
	for i := 0; i <= n; i++ {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for i := 2; i <= n; i++ {
		// iがtrueで残っている場合は素数認定し、自分より大きい倍数を取り除いていく
		if isPrime[i] {
			// prime[p] = i
			prime = append(prime, i)
			p++

			// 倍数の除去
			for j := 2 * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	return p, prime
}

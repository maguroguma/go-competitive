package main

import "fmt"

func main() {
	fmt.Println(segmentSieve(22, 37))
	fmt.Println(segmentSieve(22801763489, 22801787297))
}

// [a, b) の整数に対して篩を行う
// isPrime[i-a] = true ⇔ iが素数
func segmentSieve(a, b int) int {
	var isPrime [2000000 + 5]bool      // 区間の長さ分必要
	var isPrimeSmall [2000000 + 5]bool // 2からsqrt(b)分必要

	for i := 0; i*i < b; i++ {
		isPrimeSmall[i] = true
	}
	for i := 0; i < b-a; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < b; i++ {
		// 2からsqrt(b)までの素数について、素数だったら以下の処理を行う
		if isPrimeSmall[i] {
			// sqrt(b)までの素数以外を取り除く普通のエラトステネスの篩
			for j := 2 * i; j*j < b; j += i {
				isPrimeSmall[j] = false // [2, sqrt(b)) の篩
			}

			// 今見ているiは素数なので、[a, b)におけるiの倍数は取り除く（初期値がややこしいが普通の篩とほとんど同じ）
			tmp := Max(2.0, (a+i-1)/i)
			for j := int(tmp) * i; j < b; j += i {
				isPrime[j-a] = false // [a, b) の篩
			}
		}
	}

	res := 0
	for i := 0; i < (b - a); i++ {
		if isPrime[i] {
			res++
		}
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// エラトステネスの篩: O(NloglogN)
// n以下の素数の数を返す
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

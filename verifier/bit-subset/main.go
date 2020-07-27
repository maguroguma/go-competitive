package main

import (
	"fmt"
)

func main() {
	S := 1<<uint(4) + 1<<uint(2) + 1<<uint(1) + 1

	// Sの部分集合を「降順で」列挙
	cnt := 0
	for T := S; ; T = (T - 1) & S {
		cnt++

		subset := T
		compsub := T ^ S
		bitsum := subset | compsub
		subsetStr := string(ZeroPaddingRuneSlice(subset, 5))
		compsubStr := string(ZeroPaddingRuneSlice(compsub, 5))
		bitsumStr := string(ZeroPaddingRuneSlice(bitsum, 5))

		fmt.Printf("sub: %s, comp: %s, bitsum: %s\n", subsetStr, compsubStr, bitsumStr)

		if T == 0 {
			break
		}
	}
	fmt.Println(cnt)

	cnt = 0
	// Sの部分集合を「昇順で」列挙
	for T := 0; ; T = (T - S) & S {
		cnt++

		subset := T
		compsub := T ^ S
		bitsum := subset | compsub
		subsetStr := string(ZeroPaddingRuneSlice(subset, 5))
		compsubStr := string(ZeroPaddingRuneSlice(compsub, 5))
		bitsumStr := string(ZeroPaddingRuneSlice(bitsum, 5))

		fmt.Printf("sub: %s, comp: %s, bitsum: %s\n", subsetStr, compsubStr, bitsumStr)

		if T == S {
			break
		}
	}
}

// For debugging use.
func ZeroPaddingRuneSlice(n, digitsNum int) []rune {
	sn := fmt.Sprintf("%b", n)

	residualLength := digitsNum - len(sn)
	if residualLength <= 0 {
		return []rune(sn)
	}

	zeros := make([]rune, residualLength)
	for i := 0; i < len(zeros); i++ {
		zeros[i] = '0'
	}

	res := []rune{}
	res = append(res, zeros...)
	res = append(res, []rune(sn)...)

	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

// NextLine reads a line text from stdin, and then returns its string.
func NextLine() string {
	sc.Scan()
	return sc.Text()
}

// NextIntsLine reads a line text, that consists of **ONLY INTEGERS DELIMITED BY SPACES**, from stdin.
// And then returns intergers slice.
func NextIntsLine() []int {
	ints := []int{}
	intsStr := NextLine()
	tmp := strings.Split(intsStr, " ")
	for _, s := range tmp {
		integer, _ := strconv.Atoi(s)
		ints = append(ints, integer)
	}
	return ints
}

// NextRunesLine reads a line text, that consists of **ONLY CHARACTERS ARRANGED CONTINUOUSLY**, from stdin.
// Ant then returns runes slice.
func NextRunesLine() []rune {
	return []rune(NextLine())
}

// Max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Max(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m < integer {
			m = integer
		}
	}
	return m
}

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m > integer {
			m = integer
		}
	}
	return m
}

/*******************************************************************/

var d int
var n []rune

var dp [10010][110]int // dp[i][j]: i桁の正整数のうち、dで割ったあまりがjの正整数の個数

func main() {
	tmp := NextIntsLine()
	d = tmp[0]
	n = NextRunesLine()
	mod := 1000000007
	digitNum := len(n)

	// 初項
	for k := 0; k < 10; k++ {
		dp[1][k%d] += 1
	}
	for i := 2; i < digitNum; i++ {
		for j := 0; j < d; j++ {
			for k := 0; k < 10; k++ {
				dp[i][(j+k)%d] += dp[i-1][j]
				dp[i][(j+k)%d] %= mod
			}
		}
	}

	answer := 0
	savedSum := 0 // 上位桁の数の和を保存
	for i := 0; i < digitNum-1; i++ {
		topDigit := n[i]
		topDigitInt := int(topDigit - '0')
		for j := 0; j < d; j++ {
			for k := 0; k < topDigitInt; k++ {
				diff := k + savedSum // diffを足したとき、dで割ったときのあまりが0となるようなテーブル要素を参照する
				if (diff+j)%d == 0 {
					answer += dp[digitNum-i-1][j]
					answer %= mod
				}
			}
		}
		savedSum += topDigitInt
		savedSum %= d
	}

	fmt.Println(answer - 1)
}

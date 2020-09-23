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

var n int
var p []int
var dp [110][100*100 + 10]bool // dp[i][j]: i番目以降の問題集合から任意個解く問題を選んだとき、j点を取れるかどうか

func main() {
	tmp := NextIntsLine()
	n = tmp[0]
	p = NextIntsLine()

	// 初期化
	for i := 0; i < n+1; i++ {
		dp[i][0] = true
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= 100*n; j++ {
			if j >= p[i] {
				dp[i][j] = dp[i+1][j] || dp[i+1][j-p[i]]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	answer := 0
	for j := 0; j <= 100*n; j++ {
		if dp[0][j] {
			answer++
		}
	}
	fmt.Println(answer)
}

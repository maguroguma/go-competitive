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

/*******************************************************************/

var n, w int
var V, W [110]int
var dp [110][10100]int

func main() {
	tmp := NextIntsLine()
	n, w = tmp[0], tmp[1]
	for i := 0; i < n; i++ {
		tmp = NextIntsLine()
		V[i], W[i] = tmp[0], tmp[1]
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= w; j++ {
			if j >= W[i] {
				dp[i][j] = Max(dp[i+1][j], dp[i+1][j-W[i]]+V[i])
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	fmt.Println(dp[0][w])
}

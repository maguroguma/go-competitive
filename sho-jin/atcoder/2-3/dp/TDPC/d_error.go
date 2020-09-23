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

/* 「配るDP」というカテゴリ */

/*
失敗例！
- 網羅できておらず精度が足りない（？）
- 安易にDPテーブルサイズを取るとMLEしてしまう
*/

var n, d int
var dp [110][80][80][80]float64 // dp[n][x][y][z]: サイコロをn回振ったとき、出た目の積が 2^x * 3^y * 5^z になる確率

func main() {
	tmp := NextIntsLine()
	n, d = tmp[0], tmp[1]

	x, y, z := 0, 0, 0
	for d%2 == 0 {
		d /= 2
		x++
	}
	for d%3 == 0 {
		d /= 3
		y++
	}
	for d%5 == 0 {
		d /= 5
		z++
	}
	if d > 1 {
		fmt.Println(0.0)
		return
	}

	dp[0][0][0][0] = 1.0
	for i := 0; i < n; i++ {
		for j := 0; j <= 75; j++ {
			for k := 0; k <= 75; k++ {
				for l := 0; l <= 75; l++ {
					dp[i+1][j][k][l] += 1.0 / 6.0 * dp[i][j][k][l]
					dp[i+1][j+1][k][l] += 1.0 / 6.0 * dp[i][j][k][l]
					dp[i+1][j][k+1][l] += 1.0 / 6.0 * dp[i][j][k][l]
					dp[i+1][j+2][k][l] += 1.0 / 6.0 * dp[i][j][k][l]
					dp[i+1][j][k][l+1] += 1.0 / 6.0 * dp[i][j][k][l]
					dp[i+1][j+1][k+1][l] += 1.0 / 6.0 * dp[i][j][k][l]
				}
			}
		}
	}

	answer := 0.0
	for j := x; j <= 75; j++ {
		for k := y; k <= 75; k++ {
			for l := z; l <= 75; l++ {
				answer += dp[n][j][k][l]
			}
		}
	}
	fmt.Println(answer)
}

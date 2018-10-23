package main

import (
	"bufio"
	"fmt"
	"math"
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

/*
トーナメント系の問題はDPに限らずありそうなので、基本的な考え方はまとめておきたい
「i回戦でjさんが戦う人の列挙」の部分は本来もっと賢くしないといけない
*/

var k int
var r []int

var dp [12][1050]float64 // dp[i][j]: i回戦でjさんが勝利する確率

// i回戦でのjさんの所属するブロック番号を返す
func getBlockNumber(i, j int) int {
	p := int(math.Pow(2.0, float64(i)))
	bNum := j / p
	if j%p != 0 {
		return bNum + 1
	}
	return bNum
}

// pさんとqさんが対戦したときにpさんが勝つ確率
func pVictory(p, q int) float64 {
	return 1.0 / (1.0 + math.Pow(10.0, (float64(r[q]-r[p]))/400.0))
}

func main() {
	tmp := NextIntsLine()
	k = tmp[0]
	jMax := int(math.Pow(2.0, float64(k)))
	r = append(r, -1.0)
	for i := 0; i < int(math.Pow(2.0, float64(k))); i++ {
		tmp := NextIntsLine()
		r = append(r, tmp[0])
	}

	// 初期化
	for j := 1; j <= jMax; j++ {
		dp[0][j] = 1.0
	}
	for i := 1; i <= k; i++ {
		for j := 1; j <= jMax; j++ {
			jCurBlock := getBlockNumber(i, j)   // i回戦でのjのブロック
			jBefBlock := getBlockNumber(i-1, j) // i-1回戦でのjのブロック
			for k := 1; k <= jMax; k++ {
				kCurBlock := getBlockNumber(i, k)   // i回戦でのkのブロック
				kBefBlock := getBlockNumber(i-1, k) // i-1回戦でのkのブロック
				if jCurBlock == kCurBlock && jBefBlock != kBefBlock {
					dp[i][j] += dp[i-1][j] * dp[i-1][k] * pVictory(j, k)
				}
			}
		}
	}

	for j := 1; j <= jMax; j++ {
		fmt.Println(dp[k][j])
	}
}

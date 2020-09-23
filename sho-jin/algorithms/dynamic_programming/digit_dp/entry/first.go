package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

/*********** I/O ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
)

func init() {
	ReadString = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	// Split sets the split function for the Scanner. The default split function is ScanLines.
	// Split panics if it is called after scanning has started.
	r.Split(bufio.ScanWords)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

// ReadInt returns an integer.
func ReadInt() int {
	return int(readInt64())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(ReadString(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

// ReadIntSlice returns an integer slice that has n integers.
func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
	}
	return b
}

// ReadRuneSlice returns a rune slice.
func ReadRuneSlice() []rune {
	return []rune(ReadString())
}

/*******************************************************************/

/* A以下の非負整数の総数を求めるプログラム */
/* 数値Aを入力すると（0も含むため）A+1を出力する */

const MOD = 1000000000 + 7
const ALPHABET_NUM = 26

var A []rune // 桁DPは与えられるNが大きいために文字列で受け取るケースが多い
// dp[i][j]: 上位i(1~n, 0は存在しない桁)桁まで決めた、A未満が確定している（j==0）
// ときのA以下の非負整数の数
var dp [100000 + 1][2]int

// 上位i桁は、自然言語が意味する通り1-originでカウントする（すなわちiは1桁からスタートし、最上位桁は上位1桁と呼ぶ）
// ただし、DPの都合上、i==0も定義する
// 全体に渡って、A[i]は自然言語的にはAの上位i+1桁の数値に相当することに注意が必要（例えばA[0]はAの上から1桁目の数値を表す）
func main() {
	A = ReadRuneSlice()

	n := len(A)
	// 初期値の決め方は状態や遷移の仕方から適切に決める（今回は数え上げなので1からスタート（？））
	// less==0で1なのは、はじめの桁を決めていないため、Aの最上位桁の値までしか動かせないことからも自然
	// おそらく、less==0で初期値を設定するのは、桁DPのほとんど（すべて？）で成立するはず
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			// xは動かす値の上限
			var x int
			if j == 1 {
				x = 9 // **i桁目の直前で**lessならば9まで自由に動かす
			} else {
				x = int(A[i] - '0') // **i桁目の直前で**lessでないならばAの上位i桁の整数値までしか動かせない
			}

			for d := 0; d <= x; d++ {
				var isLess int
				if j == 1 || d < x {
					// i桁目の時点でlessであるか、そうでなくとも設定する桁の値がAのi桁目の値よりも小さい場合、
					// isLessの方へ加算する形で遷移させる
					isLess = 1
				} else {
					// そうでない場合はisLessフラグが立たない方へ加算する形で遷移させる
					isLess = 0
				}
				dp[i+1][isLess] += dp[i][j] // i+1桁目への遷移
			}
		}
	}

	ans := 0
	for j := 0; j < 2; j++ {
		fmt.Println(dp[n][j]) // j==0のときは（該当するのはAのみであるため）1が出力されるのを確認できる
		ans += dp[n][j]
	}
	fmt.Println(ans)

	fmt.Println("---")

	// この出力もみておくと面白い
	for i := 0; i <= n; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(dp[i][j])
		}
	}
}

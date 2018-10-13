package main

// https://arc061.contest.atcoder.jp/tasks/arc061_a

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int
var input []rune

/*
- 与えられた文字列長を計算する
- 文字列長-1の値を保存
- 深さ優先探索でbit全探索を行う
- bit列を引数として、そこから数式に変換し和を計算する関数を個別実装
*/
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()

	input = []rune(str)
	n = len(str) - 1

	answer := recursion(0, 0)
	fmt.Println(answer)
}

func recursion(i, bits int) int {
	if i == n {
		return bitsToSum(bits)
	}

	on := bits | (1 << byte(i))
	off := bits
	return recursion(i+1, on) + recursion(i+1, off)
}

func bitsToSum(bits int) int {
	sum := 0

	equationStr := ""
	for i := 0; i < n; i++ {
		equationStr += string(input[i])
		bit := 1 & (bits >> byte(i))
		if bit == 1 {
			equationStr += "+"
		}
	}
	equationStr += string(input[n])

	literalStrSlice := strings.Split(equationStr, "+")
	for _, str := range literalStrSlice {
		val, _ := strconv.Atoi(str)
		sum += val
	}

	return sum
}

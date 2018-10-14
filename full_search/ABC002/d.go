package main

/* 最大クリーク問題 */

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

// NextIntsLine reads a line text, that consists of only integers delimited by spaces, from stdin.
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

var n, m int
var X, Y []int
var answer int

func main() {
	tmp := NextIntsLine()
	n, m = tmp[0], tmp[1]
	for i := 0; i < m; i++ {
		tmp = NextIntsLine()
		X = append(X, tmp[0])
		Y = append(Y, tmp[1])
	}

	answer = -1
	recursion(0, 0)
	fmt.Println(answer)
}

func recursion(i, bits int) {
	if i == n {
		number, ok := isOK(bits)
		if ok && number > answer {
			answer = number
		}
		return
	}

	recursion(i+1, bits|(1<<byte(i)))
	recursion(i+1, bits)
}

func isOK(bits int) (int, bool) {
	for i := 0; i < n; i++ {
		ibit := 1 & (bits >> byte(i))
		for j := i + 1; j < n; j++ {
			jbit := 1 & (bits >> byte(j))
			if ibit == 1 && jbit == 1 {
				if !knowEachOther(i+1, j+1) {
					return -1, false
				}
			}
		}
	}

	number := 0
	for i := 0; i < n; i++ {
		bit := 1 & (bits >> byte(i))
		if bit == 1 {
			number++
		}
	}
	return number, true
}

func knowEachOther(a, b int) bool {
	for k := 0; k < m; k++ {
		if a == X[k] && b == Y[k] {
			return true
		}
	}
	return false
}

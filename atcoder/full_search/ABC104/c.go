package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
ボーナスを取らないのなら純粋に配点の高い問題を解くのがよい
*/

var d, g int
var P, C []int
var answer int

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

func main() {
	tmp := NextIntsLine()
	d, g = tmp[0], tmp[1]
	for i := 0; i < d; i++ {
		tmp = NextIntsLine()
		p, c := tmp[0], tmp[1]
		P = append(P, p)
		C = append(C, c)
	}

	answer = 1000000
	recursion(0, 0)
	fmt.Println(answer)
}

func recursion(i, bits int) {
	if i == d {
		tmp := solvedProbNum(bits)
		if tmp < answer {
			answer = tmp
		}
		return
	}

	recursion(i+1, bits|(1<<byte(i)))
	recursion(i+1, bits)
}

func solvedProbNum(bits int) int {
	score := 0
	probNum := 0
	for i := 0; i < d; i++ {
		bit := 1 & (bits >> byte(i))
		if bit == 1 {
			score += P[i]*(i+1)*100 + C[i]
			probNum += P[i]
		}
	}

	if score >= g {
		return probNum
	}

	residual := g - score
	maxId := -1
	for i := d - 1; i >= 0; i-- {
		bit := 1 & (bits >> byte(i))
		if bit == 0 {
			maxId = i
			break
		}
	}
	if maxId == -1 {
		return 1000001
	}

	q := residual / (100 * (maxId + 1))
	r := (residual / 100) % (maxId + 1)
	if r > 0 {
		q++
	}
	if q >= P[maxId] {
		return 1000001
	} else {
		return probNum + q
	}
}

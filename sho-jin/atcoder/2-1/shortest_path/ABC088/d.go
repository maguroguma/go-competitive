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

// -----------------------------------

var h, w int
var S [][]rune
var memo [60][60]int

const inf = 1000000

type node [3]int

func main() {
	tmp := NextIntsLine()
	h, w = tmp[0], tmp[1]
	for i := 0; i < h; i++ {
		row := NextRunesLine()
		S = append(S, row)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			memo[i][j] = inf
		}
	}

	delta := [4][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{0, -1},
		[2]int{-1, 0},
	}

	queue := []node{}
	queue = append(queue, node{0, 0, 0})
	memo[0][0] = 0
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		ny, nx := now[0], now[1]

		if ny == h-1 && nx == w-1 {
			break
		}

		for _, d := range delta {
			dy, dx := ny+d[0], nx+d[1]
			if 0 <= dx && dx < w && 0 <= dy && dy < h {
				if memo[dy][dx] == inf && S[dy][dx] == '.' {
					queue = append(queue, node{dy, dx, now[2] + 1})
					memo[dy][dx] = now[2] + 1
				}
			}
		}
	}

	if memo[h-1][w-1] == inf {
		fmt.Println(-1)
	} else {
		x := memo[h-1][w-1] + 1
		b := 0
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if S[i][j] == '#' {
					b++
				}
			}
		}
		answer := h*w - x - b
		fmt.Println(answer)
	}
}

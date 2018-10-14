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

var h, w int
var C [][]rune
var delta [4][2]int
var answer bool
var sx, sy int

func main() {
	tmp := NextIntsLine()
	h, w = tmp[0], tmp[1]
	for i := 0; i < h; i++ {
		row := []rune(NextLine())
		C = append(C, row)
	}

	delta = [4][2]int{
		[2]int{1, 0},
		[2]int{0, 1},
		[2]int{-1, 0},
		[2]int{0, -1},
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if C[i][j] == 's' {
				sx, sy = j, i
			}
		}
	}

	answer = false
	dfs(sx, sy)
	if answer {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(x, y int) {
	C[y][x] = '#'
	for _, d := range delta {
		xx := x + d[0]
		yy := y + d[1]
		if 0 <= xx && xx < w && 0 <= yy && yy < h {
			if C[yy][xx] == 'g' {
				answer = true
			} else if C[yy][xx] == '.' {
				dfs(xx, yy)
			}
		}
	}
}

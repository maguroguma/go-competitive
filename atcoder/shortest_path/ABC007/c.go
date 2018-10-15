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

var r, c, sy, sx, gy, gx int
var C [][]rune
var memo [100][100]int

const INF = 1000000

type next [3]int

var delta [4][2]int

func main() {
	tmp := NextIntsLine()
	r, c = tmp[0], tmp[1]
	tmp = NextIntsLine()
	sy, sx = tmp[0]-1, tmp[1]-1
	tmp = NextIntsLine()
	gy, gx = tmp[0]-1, tmp[1]-1
	for i := 0; i < r; i++ {
		row := NextRunesLine()
		C = append(C, row)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			memo[i][j] = INF
		}
	}

	delta = [4][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{0, -1},
		[2]int{-1, 0},
	}

	queue := []next{}
	queue = append(queue, next{sx, sy, 0})
	memo[sy][sx] = 0
	for {
		if len(queue) == 0 {
			break
		}

		now := queue[0]
		queue = queue[1:]
		nx, ny := now[0], now[1]
		if nx == gx && ny == gy {
			break
		}

		for _, d := range delta {
			dy, dx := ny+d[0], nx+d[1]
			if memo[dy][dx] == INF && C[dy][dx] == '.' {
				queue = append(queue, next{dx, dy, now[2] + 1})
				memo[dy][dx] = now[2] + 1
			}
		}
	}

	fmt.Println(memo[gy][gx])
}

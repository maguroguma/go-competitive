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

// -----------------------

var h, w int
var C [][]rune
var sx, sy, gx, gy int

type next struct {
	x, y int
}

// 踏んだときの壁を破った回数を記録
var memo [510][510]int

func main() {
	tmp := NextIntsLine()
	h, w = tmp[0], tmp[1]
	for i := 0; i < h; i++ {
		row := NextRunesLine()
		C = append(C, row)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			memo[i][j] = -1
			if C[i][j] == 's' {
				sy, sx = i, j
			} else if C[i][j] == 'g' {
				gy, gx = i, j
			}
		}
	}

	delta := [4][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{0, -1},
		[2]int{-1, 0},
	}

	queue := []*next{}
	queue = append(queue, &next{x: sx, y: sy})
	memo[sy][sx] = 0
	ok := false
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		nx, ny := now.x, now.y
		times := memo[ny][nx]
		if nx == gx && ny == gy {
			ok = true
			break
		}

		for _, d := range delta {
			dx, dy := nx+d[0], ny+d[1]
			// すでに踏んだことがあっても、壁を壊した回数が小さかったら踏み直して更新する
			if 0 <= dx && dx < w && 0 <= dy && dy < h && (memo[dy][dx] == -1 || times < memo[dy][dx]) {
				if C[dy][dx] == '.' || C[dy][dx] == 'g' {
					queue = append(queue, &next{x: dx, y: dy})
					memo[dy][dx] = times
				} else if C[dy][dx] == '#' && times < 2 {
					queue = append(queue, &next{x: dx, y: dy})
					memo[dy][dx] = times + 1
				}
			}
		}
	}

	if ok {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

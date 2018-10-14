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

var A [][]rune
var delta [4][2]int
var copiedA [10][10]rune

func main() {
	for i := 0; i < 10; i++ {
		tmp := NextRunesLine()
		A = append(A, tmp)
	}

	delta = [4][2]int{
		[2]int{1, 0},
		[2]int{0, 1},
		[2]int{-1, 0},
		[2]int{0, -1},
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			isOcean := true
			if A[i][j] == 'o' {
				isOcean = false
			}
			A[i][j] = 'o'
			c := count()
			if c == 1 {
				fmt.Println("YES")
				return
			}
			if isOcean {
				A[i][j] = 'x'
			}
		}
	}
	fmt.Println("NO")
	return
}

func count() int {
	number := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			copiedA[i][j] = A[i][j]
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if copiedA[i][j] == 'o' {
				dfs(j, i)
				number++
			}
		}
	}
	return number
}

func dfs(x, y int) {
	copiedA[y][x] = 'x'

	for _, d := range delta {
		dx, dy := d[0], d[1]
		xx := x + dx
		yy := y + dy
		if 0 <= xx && xx < 10 && 0 <= yy && yy < 10 && copiedA[yy][xx] == 'o' {
			dfs(xx, yy)
		}
	}
}

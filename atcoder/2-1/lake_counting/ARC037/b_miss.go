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

var n, m int
var count int
var nodes []int
var U, V []int

func main() {
	tmp := NextIntsLine()
	n, m = tmp[0], tmp[1]
	for i := 0; i < n; i++ {
		nodes = append(nodes, 0)
	}
	for i := 0; i < m; i++ {
		tmp := NextIntsLine()
		U = append(U, tmp[0]-1)
		V = append(V, tmp[1]-1)
	}

	count = 0
	for i := 0; i < n; i++ {
		if nodes[i] == 0 {
			if dfs(i, -1) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func dfs(current, prev int) bool {
	nodes[current] = 1
	retVal := true
	for i := 0; i < m; i++ {
		if U[i] == current {
			if nodes[V[i]] == 0 {
				retVal = retVal && dfs(V[i], current)
			} else {
				if V[i] != prev {
					retVal = false
				}
			}
		}
	}
	return retVal
}

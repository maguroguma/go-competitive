package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

/*******************************************************************/

var n, m int

type bridge struct {
	start, end int
}
type Interface []*bridge

func (b Interface) Len() int {
	return len(b)
}
func (b Interface) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b Interface) Less(i, j int) bool {
	return b[i].end < b[j].end
}

func main() {
	tmp := NextIntsLine()
	n, m = tmp[0], tmp[1]
	bridgeSlice := []*bridge{}
	for i := 0; i < m; i++ {
		tmp = NextIntsLine()
		bridgeSlice = append(bridgeSlice, &bridge{start: tmp[0], end: tmp[1]})
	}
	sort.Sort(Interface(bridgeSlice))

	answer := 0
	prev := -1
	for _, b := range bridgeSlice {
		// すでに侵略不可となっている場合はスキップ
		if b.start <= prev {
			continue
		}
		// 橋を壊す
		prev = b.end - 1
		answer++
	}
	fmt.Println(answer)
}

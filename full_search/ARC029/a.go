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

var n int
var T []int

func main() {
	tmp := NextIntsLine()
	n = tmp[0]
	for i := 0; i < n; i++ {
		tmp = NextIntsLine()
		T = append(T, tmp[0])
	}

	min := 1000000
	for i := 0; i < (1 << byte(n)); i++ {
		time := calcTime(i)
		if time < min {
			min = time
		}
	}
	fmt.Println(min)
}

func calcTime(bits int) int {
	a, b := 0, 0
	for i := 0; i < n; i++ {
		bit := 1 & (bits >> byte(i))
		if bit == 1 {
			a += T[i]
		} else {
			b += T[i]
		}
	}

	if a < b {
		return b
	} else {
		return a
	}
}

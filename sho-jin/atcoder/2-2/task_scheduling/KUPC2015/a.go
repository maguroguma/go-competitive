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

/*******************************************************************/

var T int
var S [][]rune

func main() {
	tmp := NextIntsLine()
	T = tmp[0]
	for i := 0; i < T; i++ {
		row := NextRunesLine()
		S = append(S, row)
	}

	answers := []int{}
	for _, s := range S {
		a := 0
		for i := 0; i < len(s)-4; i++ {
			if s[i] == 't' && s[i+1] == 'o' && s[i+2] == 'k' && s[i+3] == 'y' && s[i+4] == 'o' {
				a++
				i += 4
			} else if s[i] == 'k' && s[i+1] == 'y' && s[i+2] == 'o' && s[i+3] == 't' && s[i+4] == 'o' {
				a++
				i += 4
			}
		}
		answers = append(answers, a)
	}

	for _, a := range answers {
		fmt.Println(a)
	}
}

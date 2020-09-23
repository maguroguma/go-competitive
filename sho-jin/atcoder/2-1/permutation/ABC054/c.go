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

// GeneratePermutation returns n! in a [][]int style.
func GeneratePermutation(n int) [][]int {
	interim, residual := []int{}, []int{}
	for i := 1; i < n+1; i++ {
		residual = append(residual, i)
	}

	return recursion(interim, residual)
}

// recursion finally returns only leaf node of a tree diagram
func recursion(interim, residual []int) [][]int {
	if len(residual) == 0 {
		return [][]int{interim}
	}

	permutation := [][]int{}
	for i, r := range residual {
		copiedInterim := make([]int, len(interim))
		copy(copiedInterim, interim)
		copiedResidual := deleteElement(residual, i)

		copiedInterim = append(copiedInterim, r)
		p := recursion(copiedInterim, copiedResidual)
		permutation = append(permutation, p...)
	}
	return permutation
}

func deleteElement(s []int, i int) []int {
	newS := []int{}
	for j, e := range s {
		if j == i {
			continue
		}
		newS = append(newS, e)
	}
	return newS
}

// -----------------

var n, m int
var adjMatrix [15][15]bool

func main() {
	tmp := NextIntsLine()
	n, m = tmp[0], tmp[1]
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			adjMatrix[i][j] = false
		}
	}
	for i := 0; i < m; i++ {
		row := NextIntsLine()
		j, k := row[0], row[1]
		adjMatrix[j][k] = true
		adjMatrix[k][j] = true
	}

	answer := 0
	patterns := GeneratePermutation(n)
	for _, p := range patterns {
		if p[0] != 1 {
			continue
		}

		ok := true
		for i := 0; i < n-1; i++ {
			prev, next := p[i], p[i+1]
			if !adjMatrix[prev][next] {
				ok = false
				break
			}
		}
		if ok {
			answer++
		}
	}

	fmt.Println(answer)
}

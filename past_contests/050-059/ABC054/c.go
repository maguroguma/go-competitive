package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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

// Max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Max(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m < integer {
			m = integer
		}
	}
	return m
}

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m > integer {
			m = integer
		}
	}
	return m
}

// PowInt is integer version of math.Pow
func PowInt(a, e int) int {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}
	fa := float64(a)
	fe := float64(e)
	fanswer := math.Pow(fa, fe)
	return int(fanswer)
}

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	fa := float64(a)
	fanswer := math.Abs(fa)
	return int(fanswer)
}

// DeleteElement returns a *NEW* slice, that have the same and minimum length and capacity.
// DeleteElement makes a new slice by using easy slice literal.
func DeleteElement(s []int, i int) []int {
	if i < 0 || len(s) <= i {
		panic(errors.New("[index error]"))
	}
	// appendのみの実装
	n := make([]int, 0, len(s)-1)
	n = append(n, s[:i]...)
	n = append(n, s[i+1:]...)
	return n
}

// Concat returns a *NEW* slice, that have the same and minimum length and capacity.
func Concat(s, t []rune) []rune {
	n := make([]rune, 0, len(s)+len(t))
	n = append(n, s...)
	n = append(n, t...)
	return n
}

// sort package (snippets)
//sort.Sort(sort.IntSlice(s))
//sort.Sort(sort.Reverse(sort.IntSlice(s)))
//sort.Sort(sort.Float64Slice(s))
//sort.Sort(sort.StringSlice(s))

// copy function
//a = []int{0, 1, 2}
//b = make([]int, len(a))
//copy(b, a)

// GeneratePermutation returns n! in a [][]int style.
// Each pattern consists of 1 to n integers.
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

/*******************************************************************/

var n, m int
var A, B []int

func main() {
	tmp := NextIntsLine()
	n, m = tmp[0], tmp[1]
	A, B = []int{}, []int{}
	for i := 0; i < m; i++ {
		tmp = NextIntsLine()
		A = append(A, tmp[0])
		B = append(B, tmp[1])
	}
	adjMatrix := [10][10]int{}
	for i := 0; i < m; i++ {
		start := A[i]
		end := B[i]
		adjMatrix[start][end] = 1
		adjMatrix[end][start] = 1
	}

	permutations := GeneratePermutation(n)
	answer := 0
	for _, p := range permutations {
		if p[0] != 1 {
			continue
		}
		ok := true
		for i := 0; i < len(p)-1; i++ {
			start := p[i]
			end := p[i+1]
			if adjMatrix[start][end] != 1 {
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

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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

/*******************************************************************/

var n int   // 最大100000
var A []int // 各要素の最大は1000000000

func main() {
	tmp := NextIntsLine()
	n = tmp[0]
	for i := 0; i < n; i++ {
		tmp = NextIntsLine()
		A = append(A, tmp[0])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	answer1, answer2 := 0, 0
	// 最初が下りのパターンの最大値
	coords := map[int]int{2: 0, 1: 0, -1: 0, -2: 0} // それぞれ+2, +1, -1, -2の数
	if n%2 == 0 {
		coords[1] = 1
		coords[-1] = 1
		coords[2], coords[-2] = n/2-1, n/2-1
	} else {
		coords[1] = 2
		coords[-1] = 0
		coords[2] = n/2 - 1
		coords[-2] = n / 2
	}
	i := 0
	for _, c := range []int{2, 1, -1, -2} {
		for j := 0; j < coords[c]; j++ {
			answer1 += c * A[i]
			i++
		}
	}
	// 最初が上りのパターンの最大値
	coords = map[int]int{2: 0, 1: 0, -1: 0, -2: 0}
	if n%2 == 0 {
		coords[1] = 1
		coords[-1] = 1
		coords[2], coords[-2] = n/2-1, n/2-1
	} else {
		coords[1] = 0
		coords[-1] = 2
		coords[2] = n / 2
		coords[-2] = n/2 - 1
	}
	i = 0
	for _, c := range []int{2, 1, -1, -2} {
		for j := 0; j < coords[c]; j++ {
			answer2 += c * A[i]
			i++
		}
	}

	if answer1 > answer2 {
		fmt.Println(answer1)
	} else {
		fmt.Println(answer2)
	}
}

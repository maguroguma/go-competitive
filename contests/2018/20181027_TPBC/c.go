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

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}
func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}
func (is intSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	} else {
		return b - a
	}
}

/*******************************************************************/

/*
8割ぐらいしかテストケース通せなかったやつ
網羅できているようで、多分しきれていない誤答
こういうのを惜しかったと思わず、忘れて正しい解法の考え方を身につけること
*/

var n int   // 最大100000
var A []int // 各要素の最大は1000000000

func main() {
	tmp := NextIntsLine()
	n = tmp[0]
	for i := 0; i < n; i++ {
		tmp = NextIntsLine()
		A = append(A, tmp[0])
	}

	sort.Sort(intSlice(A))
	answer := 0
	left, right := -1, -1
	for i := 0; i < n/2; i++ {
		// 小さい方
		current := A[i]
		// 最初のみ
		if left == -1 {
			left = current
			right = current
		}

		labs := abs(left, current)
		rabs := abs(right, current)
		if labs > rabs {
			answer += labs
			left = current
		} else {
			answer += rabs
			right = current
		}

		// 大きい方
		current = A[n-1-i]
		labs = abs(left, current)
		rabs = abs(right, current)
		if labs > rabs {
			answer += labs
			left = current
		} else {
			answer += rabs
			right = current
		}
	}
	// 奇数ならば中央も処理する
	if n%2 == 1 {
		current := A[n/2]
		labs := abs(left, current)
		rabs := abs(right, current)
		if labs > rabs {
			answer += labs
			left = current
		} else {
			answer += rabs
			right = current
		}
	}
	fmt.Println(answer)
}

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

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

/*******************************************************************/

var n int
var A []int

func main() {
	//tmp := NextIntsLine()
	tmp := readLine()
	n, _ = strconv.Atoi(tmp)
	//	A = NextIntsLine()
	A = make([]int, 0, 1000000)
	str := readLine()
	tmp2 := strings.Split(str, " ")
	for _, s := range tmp2 {
		integer, _ := strconv.Atoi(s)
		A = append(A, integer)
	}

	S := make([]int, len(A))
	S[0] = A[0]
	for i := 1; i < len(A); i++ {
		sum := S[i-1]
		S[i] = sum + A[i]
	}
	// 最初を正とする場合と負とする場合の両方を試す
	answers := []int{}
	for _, firstSign := range []int{1, -1} {
		comp, answer := 0, 0
		if (firstSign == 1 && S[0] <= 0) || (firstSign == -1 && S[0] >= 0) {
			comp = firstSign - S[0]
			answer = AbsInt(comp)
		}

		for i := 1; i < len(S); i++ {
			var befSign int
			if S[i-1]+comp < 0 {
				befSign = -1
			} else {
				befSign = 1
			}
			if (befSign == -1 && S[i]+comp > 0) || (befSign == 1 && S[i]+comp < 0) {
				continue
			}
			x := -befSign - (S[i] + comp)
			answer += AbsInt(x)
			comp += x
		}
		answers = append(answers, answer)
	}

	fmt.Println(Min(answers...))
}

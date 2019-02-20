package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

// readLine can read long line string (at least 10^5)
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

// NextLine reads a line text from stdin, and then returns its string.
func NextLine() string {
	return readLine()
}

/*
var sc = bufio.NewScanner(os.Stdin)

// NextLine reads a line text from stdin, and then returns its string.
func NextLine() string {
	sc.Scan()
	return sc.Text()
}
*/

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

// NextStringsLine reads a line text, that consists of **STRINGS DELIMITED BY SPACES**, from stdin.
// And then returns strings slice.
func NextStringsLine() []string {
	str := NextLine()
	return strings.Split(str, " ")
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

// UpperRune is rune version of `strings.ToUpper()`.
func UpperRune(r rune) rune {
	str := strings.ToUpper(string(r))
	return []rune(str)[0]
}

// LowerRune is rune version of `strings.ToLower()`.
func LowerRune(r rune) rune {
	str := strings.ToLower(string(r))
	return []rune(str)[0]
}

// ToggleRune returns a upper case if an input is a lower case, v.v.
func ToggleRune(r rune) rune {
	var str string
	if 'a' <= r && r <= 'z' {
		str = strings.ToUpper(string(r))
	} else if 'A' <= r && r <= 'Z' {
		str = strings.ToLower(string(r))
	} else {
		str = string(r)
	}
	return []rune(str)[0]
}

// ToggleString iteratively calls ToggleRune, and returns the toggled string.
func ToggleString(s string) string {
	inputRunes := []rune(s)
	outputRunes := make([]rune, 0, len(inputRunes))
	for _, r := range inputRunes {
		outputRunes = append(outputRunes, ToggleRune(r))
	}
	return string(outputRunes)
}

// Strtoi is a wrapper of `strconv.Atoi()`.
// If `strconv.Atoi()` returns an error, Strtoi calls panic.
func Strtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(errors.New("[argument error]: Strtoi only accepts integer string"))
	} else {
		return i
	}
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

type IntHeap []int

// 5つのメソッドを実装する
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's LENGTH, NOT JUST ITS CONTENTS.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntReverseHeap []int

// 5つのメソッドを実装する
func (h IntReverseHeap) Len() int           { return len(h) }
func (h IntReverseHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntReverseHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's LENGTH, NOT JUST ITS CONTENTS.
func (h *IntReverseHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntReverseHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*******************************************************************/

var n int
var A []int

func main() {
	tmp := NextIntsLine()
	n = tmp[0]
	A = NextIntsLine()

	First, Second, Third := &IntHeap{}, []int{}, &IntReverseHeap{}
	heap.Init(First)
	heap.Init(Third)
	for i := 0; i < n; i++ {
		heap.Push(First, A[i])
	}
	for i := n; i < 2*n; i++ {
		Second = append(Second, A[i])
	}
	for i := 2 * n; i < 3*n; i++ {
		heap.Push(Third, A[i])
	}

	for len(Second) > 0 {
		a := (*First)[0]
		btop := Second[0]
		btail := Second[len(Second)-1]
		c := (*Third)[0]

		leftCase := btop - a
		rightCase := -1 * (btail - c)
		which := Max(leftCase, 0, rightCase)
		if which == leftCase {
			heap.Pop(First)
			Second = Second[1:]
			heap.Push(First, btop)
		} else if which == rightCase {
			heap.Pop(Third)
			Second = Second[:len(Second)-1]
			heap.Push(Third, btail)
		} else {
			if leftCase < rightCase {
				Second = Second[1:]
			} else {
				Second = Second[:len(Second)-1]
			}
		}
	}

	firstSum, lastSum := 0, 0
	for First.Len() > 0 {
		firstSum += First.Pop().(int)
	}
	for Third.Len() > 0 {
		lastSum += Third.Pop().(int)
	}
	fmt.Println(firstSum - lastSum)
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*********** I/O ***********/

var (
	ReadString func() string // ReadString returns a WORD string.
	stdout     *bufio.Writer
)

func init() {
	ReadString = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	// Split sets the split function for the Scanner. The default split function is ScanLines.
	// Split panics if it is called after scanning has started.
	r.Split(bufio.ScanWords)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

func ReadInt() int {
	return int(readInt64())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(ReadString(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
	}
	return b
}

func ReadLengthAndSlice() (int, []int) {
	n := ReadInt()
	return n, ReadIntSlice(n)
}

/*********** Arithmetic ***********/

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

/*********** Utilities ***********/

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

// LowerBound returns an index of a slice whose value is EQUAL TO AND LARGER THAN A KEY VALUE.
func LowerBound(s []int, key int) int {
	isLarger := func(index, key int) bool {
		if s[index] >= key {
			return true
		} else {
			return false
		}
	}

	left, right := -1, len(s)

	for right-left > 1 {
		mid := left + (right-left)/2
		if isLarger(mid, key) {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}

// UpperBound returns an index of a slice whose value is EQUAL TO AND SMALLER THAN A KEY VALUE.
func UpperBound(s []int, key int) int {
	isSmaller := func(index, key int) bool {
		if s[index] <= key {
			return true
		} else {
			return false
		}
	}

	left, right := -1, len(s)

	for right-left > 1 {
		mid := left + (right-left)/2
		if isSmaller(mid, key) {
			left = mid
		} else {
			right = mid
		}
	}

	return left
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

/*******************************************************************/

var n, m int
var P []int

var parents [100005]int

func main() {
	//	tmp := NextIntsLine(true)
	//	n, m = tmp[0], tmp[1]
	//	P = NextIntsLine(false)
	//	for i := 0; i < n; i++ {
	//		P[i]--
	//	}
	n = ReadInt()
	m = ReadInt()
	P = ReadIntSlice(n)
	for i := 0; i < n; i++ {
		P[i]--
	}

	initialize()

	for i := 0; i < m; i++ {
		//tmp = NextIntsLine(true)
		tmp := ReadIntSlice(2)
		x, y := tmp[0]-1, tmp[1]-1
		unite(x, y)
	}

	ans := 0
	for i := 0; i < n; i++ {
		if same(P[i], i) {
			ans++
		}
	}
	fmt.Println(ans)
}

func initialize() {
	for i := 0; i < n; i++ {
		parents[i] = i
	}
}

func unite(x, y int) {
	xp := root(x)
	yp := root(y)
	if xp == yp {
		return
	}
	parents[xp] = yp
}

func same(x, y int) bool {
	return root(x) == root(y)
}

func root(x int) int {
	if parents[x] == x {
		return x
	} else {
		parents[x] = root(parents[x])
		return parents[x]
	}
}

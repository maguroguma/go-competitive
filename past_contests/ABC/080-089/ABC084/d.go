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

/*
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
*/

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

/*******************************************************************/

var q int
var L, R []int
var isPrime [100001]int
var isNear2017 [100001]int
var sums [100001]int

func main() {
	tmp := NextIntsLine()
	q = tmp[0]
	for i := 0; i < q; i++ {
		tmp = NextIntsLine()
		l, r := tmp[0], tmp[1]
		L = append(L, l)
		R = append(R, r)
	}

	for i := 1; i <= 100000; i++ {
		if i > 2 && i%2 == 0 {
			continue
		}
		if sub(i) {
			isPrime[i] = 1
		}
	}

	for i := 1; i <= 100000; i++ {
		if i > 2 && i%2 == 0 {
			continue
		}
		if isPrime[i] == 1 && isPrime[(i+1)/2] == 1 {
			isNear2017[i] = 1
		}
	}

	summation := 0
	for i := 1; i <= 100000; i++ {
		summation += isNear2017[i]
		sums[i] = summation
	}

	/*
		for i := 1; i <= 11; i++ {
			fmt.Println(isPrime[i], isNear2017[i], sums[i])
		}
	*/

	for i := 0; i < q; i++ {
		l, r := L[i], R[i]
		fmt.Println(sums[r] - sums[l-1])
	}
}

func sub(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

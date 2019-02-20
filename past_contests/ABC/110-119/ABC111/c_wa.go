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
	// ReadString returns a WORD string.
	ReadString func() string
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

// ReadInt returns an integer.
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

// ReadIntSlice returns an integer slice that has n integers.
func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
	}
	return b
}

// ReadRuneSlice returns a rune slice.
func ReadRuneSlice() []rune {
	return []rune(ReadString())
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

// Gcd returns the Greatest Common Divisor of two natural numbers.
// Gcd only accepts two natural numbers (a, b >= 1).
// 0 or negative number causes panic.
// Gcd uses the Euclidean Algorithm.
func Gcd(a, b int) int {
	if a <= 0 || b <= 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}
	if a < b {
		a, b = b, a
	}

	// Euclidean Algorithm
	for b > 0 {
		div := a % b
		a, b = b, div
	}

	return a
}

// Lcm returns the Least Common Multiple of two natural numbers.
// Lcd only accepts two natural numbers (a, b >= 1).
// 0 or negative number causes panic.
// Lcd uses the Euclidean Algorithm indirectly.
func Lcm(a, b int) int {
	if a <= 0 || b <= 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	// a = a'*gcd, b = b'*gcd, a*b = a'*b'*gcd^2
	// a' and b' are relatively prime numbers
	// gcd consists of prime numbers, that are included in a and b
	gcd := Gcd(a, b)

	// not (a * b / gcd), because of reducing a probability of overflow
	return (a / gcd) * b
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

/*********** Binary Search ***********/

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

/*********** Factorization, Prime Number ***********/

// TrialDivision returns the result of prime factorization of integer N.
func TrialDivision(n int) map[int]int {
	if n <= 0 {
		panic(errors.New("[argument error]: TrialDivision only accepts a NATURAL number"))
	}
	if n == 1 {
		return map[int]int{1: 1}
	}

	p := map[int]int{}
	sqrt := math.Pow(float64(n), 0.5)
	for i := 2; i <= int(sqrt); i++ {
		exp := 0
		for n%i == 0 {
			exp++
			n /= i
		}

		if exp == 0 {
			continue
		}
		p[i] = exp
	}
	if n > 1 {
		p[n] = 1
	}

	return p
}

// IsPrime judges whether an argument integer is a prime number or not.
func IsPrime(n int) bool {
	if n == 1 {
		return false
	}

	sqrt := math.Pow(float64(n), 0.5)
	for i := 2; i <= int(sqrt); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

/********** sort package (snippets) **********/
//sort.Sort(sort.IntSlice(s))
//sort.Sort(sort.Reverse(sort.IntSlice(s)))
//sort.Sort(sort.Float64Slice(s))
//sort.Sort(sort.StringSlice(s))

/********** copy function (snippets) **********/
//a = []int{0, 1, 2}
//b = make([]int, len(a))
//copy(b, a)

/********** I/O usage **********/

//str := ReadString()
//i := ReadInt()
//X := ReadIntSlice(n)
//S := ReadRuneSlice()

/*******************************************************************/

const MOD = 1e9 + 7
const ALPHABET_NUM = 26

var n int
var V []int

func main() {
	n = ReadInt()
	V = ReadIntSlice(n)

	//oddMemo, evenMemo := [1e5 + 1]int{}, [1e5 + 1]int{}
	oddMemo, evenMemo := make([]int, 1e5+1), make([]int, 1e5+1)
	for i, v := range V {
		if (i+1)%2 == 1 {
			oddMemo[v]++
		} else {
			evenMemo[v]++
		}
	}

	ans := 0
	//	ans += n/2 + (n - Max(oddMemo...))
	//	ans += n/2 + (n - Max(evenMemo...))
	maxOddCount, maxEvenCount := 0, 0
	semiMaxOddCount, semiMaxEvenCount := 0, 0
	maxOddInt, maxEvenInt := -1, -1
	for i := 0; i < len(oddMemo); i++ {
		//		maxOddCount = Max(maxOddCount, oddMemo[i])
		//		maxEvenCount = Max(maxEvenCount, evenMemo[i])
		if maxOddCount <= oddMemo[i] {
			semiMaxOddCount = maxOddCount
			maxOddCount = oddMemo[i]
			maxOddInt = i
		}
		if maxEvenCount <= evenMemo[i] {
			semiMaxEvenCount = maxEvenCount
			maxEvenCount = evenMemo[i]
			maxEvenInt = i
		}

		//		if oddMemo[i] == n/2 && evenMemo[i] == n/2 {
		//			fmt.Println(n / 2)
		//			return
		//		}
	}
	//	ans += (n/2 - maxOddCount)
	//	ans += (n/2 - maxEvenCount)

	if maxOddInt == maxEvenInt {
		if maxOddCount == n/2 && maxEvenCount == n/2 {
			fmt.Println(n / 2)
			return
		} else if semiMaxOddCount < semiMaxEvenCount {
			ans += (n/2 - maxOddCount)
			ans += (n/2 - semiMaxEvenCount)
		} else {
			ans += (n/2 - semiMaxOddCount)
			ans += (n/2 - maxEvenCount)
		}
	} else {
		ans += (n/2 - maxOddCount)
		ans += (n/2 - maxEvenCount)
	}

	fmt.Println(ans)
}

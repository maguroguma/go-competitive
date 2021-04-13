/*
URL:
https://atcoder.jp/contests/abc198/tasks/abc198_d
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	println          = fmt.Println
	yes, no, invalid = "Yes", "No", "UNSOLVABLE"

	A, B, C []rune

	memo map[rune]int
	R    []rune
	I    map[rune]int
)

func main() {
	defer stdout.Flush()

	A, B, C = readrs(), readrs(), readrs()

	memo := make(map[rune]int)
	for _, r := range A {
		memo[r] = 1
	}
	for _, r := range B {
		memo[r] = 1
	}
	for _, r := range C {
		memo[r] = 1
	}
	if len(memo) > 10 {
		println(invalid)
		return
	}

	R = []rune{}
	for k := range memo {
		R = append(R, k)
	}
	sort.Slice(R, func(i, j int) bool {
		return R[i] < R[j]
	})
	I = make(map[rune]int)
	for i, r := range R {
		I[r] = i
	}

	numbers := []rune{}
	for i := 0; i < 10; i++ {
		numbers = append(numbers, '0'+rune(i))
	}

	// patterns := PermutationPatterns(numbers, len(R))
	// for _, pat := range patterns {
	// 	X := convert(A, pat)
	// 	Y := convert(B, pat)
	// 	Z := convert(C, pat)
	// 	if X[0] == '0' || Y[0] == '0' || Z[0] == '0' {
	// 		continue
	// 	}

	// 	x, y, z := Atoi(X), Atoi(Y), Atoi(Z)
	// 	if x+y == z {
	// 		println(x)
	// 		println(y)
	// 		println(z)
	// 		return
	// 	}
	// }

	ok, gx, gy, gz := false, -1, -1, -1
	PermutationPatterns(numbers, len(R), func(seq []rune) {
		if ok {
			return
		}

		X := convert(A, seq)
		Y := convert(B, seq)
		Z := convert(C, seq)
		if X[0] == '0' || Y[0] == '0' || Z[0] == '0' {
			return
		}

		x, y, z := Atoi(X), Atoi(Y), Atoi(Z)
		if x+y == z {
			ok, gx, gy, gz = true, x, y, z
			return
		}
	})

	if ok {
		println(gx)
		println(gy)
		println(gz)
		return
	}

	println(invalid)
}

func Atoi(A []rune) int {
	res := 0

	base := 1
	for i := 0; i < len(A)-1; i++ {
		base *= 10
	}

	for _, r := range A {
		if !('0' <= r && r <= '9') {
			panic(fmt.Sprintf("%v: cannot convert to an integer", string(A)))
		}
		res += int(r-'0') * base
		base /= 10
	}

	return res
}

func convert(A []rune, table []rune) []rune {
	res := []rune{}
	for _, r := range A {
		idx := I[r]
		res = append(res, table[idx])
	}
	return res
}

func PermutationPatterns(N []rune, k int, fn func(seq []rune)) {
	var rec func(curSeq []rune, flags []bool, k int, fn func(seq []rune))
	rec = func(curSeq []rune, flags []bool, k int, fn func(seq []rune)) {
		if len(curSeq) == k {
			fn(curSeq)
		}

		for i, isUsed := range flags {
			if isUsed {
				continue
			}

			curSeq = append(curSeq, N[i])
			flags[i] = true

			rec(curSeq, flags, k, fn)
			curSeq = curSeq[:len(curSeq)-1]
			flags[i] = false
		}
	}

	flags := make([]bool, len(N))
	rec([]rune{}, flags, k, fn)
}

// func rec(curSeq, N []rune, flags []bool, k int, fn func(seq []rune)) {
// 	if len(curSeq) == k {
// 		fn(curSeq)
// 	}

// 	for i, isUsed := range flags {
// 		if isUsed {
// 			continue
// 		}

// 		curSeq = append(curSeq, N[i])
// 		flags[i] = true

// 		rec(curSeq, N, flags, k, fn)
// 		curSeq = curSeq[:len(curSeq)-1]
// 		flags[i] = false
// 	}
// }

/*******************************************************************/

/********** common constants **********/

const (
	MOD = 1000000000 + 7
	// MOD          = 998244353
	ALPH_N  = 26
	INF_I64 = math.MaxInt64
	INF_B60 = 1 << 60
	INF_I32 = math.MaxInt32
	INF_B30 = 1 << 30
	NIL     = -1
	EPS     = 1e-10
)

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	reads = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

// mod can calculate a right residual whether value is positive or negative.
func mod(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func min(integers ...int) int {
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

// max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func max(integers ...int) int {
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

// chmin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func chmin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

// chmax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func chmax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
}

// sum returns multiple integers sum.
func sum(integers ...int) int {
	var s int
	s = 0

	for _, i := range integers {
		s += i
	}

	return s
}

// abs is integer version of math.Abs
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// pow is integer version of math.Pow
// pow calculate a power by Binary Power (二分累乗法(O(log e))).
func pow(a, e int) int {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}

	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := pow(a, halfE)
		return half * half
	}

	return a * pow(a, e-1)
}

/********** FAU standard libraries **********/

//fmt.Sprintf("%b\n", 255) 	// binary expression

/********** I/O usage **********/

//str := reads()
//i := readi()
//X := readis(n)
//S := readrs()
//a := readf()
//A := readfs(n)

//str := ZeroPaddingRuneSlice(num, 32)
//str := PrintIntsLine(X...)

/*********** Input ***********/

var (
	// reads returns a WORD string.
	reads  func() string
	stdout *bufio.Writer
)

func newReadString(ior io.Reader, sf bufio.SplitFunc) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+9)) // for Codeforces
	r.Split(sf)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

// readi returns an integer.
func readi() int {
	return int(_readInt64())
}
func readi2() (int, int) {
	return int(_readInt64()), int(_readInt64())
}
func readi3() (int, int, int) {
	return int(_readInt64()), int(_readInt64()), int(_readInt64())
}
func readi4() (int, int, int, int) {
	return int(_readInt64()), int(_readInt64()), int(_readInt64()), int(_readInt64())
}

// readll returns as integer as int64.
func readll() int64 {
	return _readInt64()
}
func readll2() (int64, int64) {
	return _readInt64(), _readInt64()
}
func readll3() (int64, int64, int64) {
	return _readInt64(), _readInt64(), _readInt64()
}
func readll4() (int64, int64, int64, int64) {
	return _readInt64(), _readInt64(), _readInt64(), _readInt64()
}

func _readInt64() int64 {
	i, err := strconv.ParseInt(reads(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

// readis returns an integer slice that has n integers.
func readis(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = readi()
	}
	return b
}

// readlls returns as int64 slice that has n integers.
func readlls(n int) []int64 {
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		b[i] = readll()
	}
	return b
}

// readf returns an float64.
func readf() float64 {
	return float64(_readFloat64())
}

func _readFloat64() float64 {
	f, err := strconv.ParseFloat(reads(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// ReadFloatSlice returns an float64 slice that has n float64.
func readfs(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = readf()
	}
	return b
}

// readrs returns a rune slice.
func readrs() []rune {
	return []rune(reads())
}

/*********** Output ***********/

// PrintIntsLine returns integers string delimited by a space.
func PrintIntsLine(A ...int) string {
	res := []rune{}

	for i := 0; i < len(A); i++ {
		str := strconv.Itoa(A[i])
		res = append(res, []rune(str)...)

		if i != len(A)-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
}

// PrintIntsLine returns integers string delimited by a space.
func PrintInts64Line(A ...int64) string {
	res := []rune{}

	for i := 0; i < len(A); i++ {
		str := strconv.FormatInt(A[i], 10) // 64bit int version
		res = append(res, []rune(str)...)

		if i != len(A)-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
}

// Printf is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func printf(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}

/*********** Debugging ***********/

// debugf is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func debugf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// ZeroPaddingRuneSlice returns binary expressions of integer n with zero padding.
// For debugging use.
func ZeroPaddingRuneSlice(n, digitsNum int) []rune {
	sn := fmt.Sprintf("%b", n)

	residualLength := digitsNum - len(sn)
	if residualLength <= 0 {
		return []rune(sn)
	}

	zeros := make([]rune, residualLength)
	for i := 0; i < len(zeros); i++ {
		zeros[i] = '0'
	}

	res := []rune{}
	res = append(res, zeros...)
	res = append(res, []rune(sn)...)

	return res
}

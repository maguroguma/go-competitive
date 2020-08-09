/*
URL:
https://codeforces.com/contest/1391/problem/D
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

/*******************************************************************/

const (
	// General purpose
	MOD = 1000000000 + 7
	// MOD          = 998244353
	ALPHABET_NUM = 26
	INF_INT64    = math.MaxInt64
	INF_BIT60    = 1 << 60
	INF_INT32    = math.MaxInt32
	INF_BIT30    = 1 << 30
	NIL          = -1

	// for dijkstra, prim, and so on
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

var (
	n, m int

	R [][]rune
	M [][]int
	A []int

	dp   [1000000 + 50][8 + 10]int
	ok   [8 + 10][8 + 10]bool
	done [8 + 10][8 + 10]bool
)

func main() {
	n, m = ReadInt2()
	for i := 0; i < n; i++ {
		row := ReadRuneSlice()
		R = append(R, row)
	}

	if n >= 4 {
		fmt.Println(-1)
		return
	}
	if n == 1 {
		fmt.Println(0)
		return
	}

	M = sub(R)
	PrintfDebug("%v\n", M)

	A = subsub(M)
	PrintfDebug("%v\n", A)

	for i := 0; i < m; i++ {
		for S := 0; S < 1<<uint(n); S++ {
			dp[i][S] = INF_BIT30
		}
	}
	for S := 0; S < 1<<uint(n); S++ {
		dp[0][S] = PopCount(S^A[0], n)
	}

	for i := 1; i < m; i++ {
		for pmask := 0; pmask < 1<<uint(n); pmask++ {
			for cmask := 0; cmask < 1<<uint(n); cmask++ {
				if check(pmask, cmask) {
					ChMin(&dp[i][cmask], dp[i-1][pmask]+PopCount(cmask^A[i], n))
				}
			}
		}
	}

	ans := INF_BIT30
	for S := 0; S < 1<<uint(n); S++ {
		ChMin(&ans, dp[m-1][S])
	}
	fmt.Println(ans)
}

func check(pmask, cmask int) bool {
	if done[pmask][cmask] {
		return ok[pmask][cmask]
	}

	done[pmask][cmask] = true
	for i := 0; i < n-1; i++ {
		cnt := 0
		cnt += NthBit(pmask, i)
		cnt += NthBit(cmask, i)
		cnt += NthBit(pmask, i+1)
		cnt += NthBit(cmask, i+1)

		if cnt%2 == 0 {
			ok[pmask][cmask] = false
			return ok[pmask][cmask]
		}
	}
	ok[pmask][cmask] = true
	return ok[pmask][cmask]
}

func sub(R [][]rune) [][]int {
	res := [][]int{}
	for j := 0; j < m; j++ {
		row := []int{}
		for i := 0; i < n; i++ {
			if R[i][j] == '1' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		res = append(res, row)
	}
	return res
}

func subsub(M [][]int) []int {
	res := []int{}
	for i := 0; i < len(M); i++ {
		mask := 0
		for j := 0; j < len(M[i]); j++ {
			if M[i][j] == 1 {
				mask += 1 << uint(j)
			}
		}
		res = append(res, mask)
	}
	return res
}

// NthBit returns nth bit value of an argument.
// n starts from 0.
func NthBit(num int, nth int) int {
	return num >> uint(nth) & 1
}

// OnBit returns the integer that has nth ON bit.
// If an argument has nth ON bit, OnBit returns the argument.
func OnBit(num int, nth int) int {
	return num | (1 << uint(nth))
}

// OffBit returns the integer that has nth OFF bit.
// If an argument has nth OFF bit, OffBit returns the argument.
func OffBit(num int, nth int) int {
	return num & ^(1 << uint(nth))
}

// PopCount returns the number of ON bit of an argument.
func PopCount(num int, ub int) int {
	res := 0

	for i := 0; i < ub; i++ {
		if ((num >> uint(i)) & 1) == 1 {
			res++
		}
	}

	return res
}

// ChMin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

/*******************************************************************/

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

/********** FAU standard libraries **********/

//fmt.Sprintf("%b\n", 255) 	// binary expression

/********** I/O usage **********/

//str := ReadString()
//i := ReadInt()
//X := ReadIntSlice(n)
//S := ReadRuneSlice()
//a := ReadFloat64()
//A := ReadFloat64Slice(n)

//str := ZeroPaddingRuneSlice(num, 32)
//str := PrintIntsLine(X...)

/*********** Input ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
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

// ReadInt returns an integer.
func ReadInt() int {
	return int(readInt64())
}
func ReadInt2() (int, int) {
	return int(readInt64()), int(readInt64())
}
func ReadInt3() (int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64())
}
func ReadInt4() (int, int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64()), int(readInt64())
}

// ReadInt64 returns as integer as int64.
func ReadInt64() int64 {
	return readInt64()
}
func ReadInt64_2() (int64, int64) {
	return readInt64(), readInt64()
}
func ReadInt64_3() (int64, int64, int64) {
	return readInt64(), readInt64(), readInt64()
}
func ReadInt64_4() (int64, int64, int64, int64) {
	return readInt64(), readInt64(), readInt64(), readInt64()
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

// ReadInt64Slice returns as int64 slice that has n integers.
func ReadInt64Slice(n int) []int64 {
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt64()
	}
	return b
}

// ReadFloat64 returns an float64.
func ReadFloat64() float64 {
	return float64(readFloat64())
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(ReadString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// ReadFloatSlice returns an float64 slice that has n float64.
func ReadFloat64Slice(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = ReadFloat64()
	}
	return b
}

// ReadRuneSlice returns a rune slice.
func ReadRuneSlice() []rune {
	return []rune(ReadString())
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

// PrintfBufStdout is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func PrintfBufStdout(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}

/*********** Debugging ***********/

// PrintfDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintfDebug(format string, a ...interface{}) {
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

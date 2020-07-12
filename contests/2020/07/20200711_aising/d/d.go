/*
URL:
https://atcoder.jp/contests/aising2020/tasks/aising2020_d
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

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

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

var (
	n    int
	X    []rune
	revX []rune

	R       [2][200000 + 50]int
	RS      [2]int
	answers []int
)

func main() {
	n = ReadInt()
	X = ReadRuneSlice()
	revX = Reverse(X)

	pc := 0
	for _, x := range X {
		if x == '1' {
			pc++
		}
	}

	if pc == 0 {
		// 全部0なので、一つ反転しても一回で終わるはず
		reigai()
		return
	}

	if pc >= 2 {
		R[0][0] = 1 % (pc - 1)
		R[1][0] = 1 % (pc + 1)
		for i := 1; i < n; i++ {
			R[0][i] = (R[0][i-1] * 2) % (pc - 1)
			R[1][i] = (R[1][i-1] * 2) % (pc + 1)
		}
	} else {
		R[1][0] = 1 % (pc + 1)
		for i := 1; i < n; i++ {
			R[1][i] = (R[1][i-1] * 2) % (pc + 1)
		}
	}

	for i := 0; i < n; i++ {
		if revX[i] == '1' {
			if pc >= 2 {
				RS[0] += R[0][i]
				RS[0] %= (pc - 1)
				RS[1] += R[1][i]
				RS[1] %= (pc + 1)
			} else {
				RS[1] += R[1][i]
				RS[1] %= (pc + 1)
			}
		}
	}

	for i := 0; i < n; i++ {
		// 反転
		var use int
		if revX[i] == '1' {
			revX[i] = '0'
			use = 0
		} else {
			revX[i] = '1'
			use = 1
		}

		if use == 0 {
			if pc-1 > 0 {
				answers = append(answers, sub(NegativeMod(RS[0]-R[0][i], pc-1)))
			} else {
				answers = append(answers, 0)
			}
		} else {
			answers = append(answers, sub(NegativeMod(RS[1]+R[1][i], pc+1)))
		}

		// もとに戻す
		if revX[i] == '1' {
			revX[i] = '0'
		} else {
			revX[i] = '1'
		}
	}

	revAns := ReverseInts(answers)
	for i := 0; i < len(revAns); i++ {
		// fmt.Println(revAns[i])
		PrintfBufStdout("%d\n", revAns[i])
	}
	stdout.Flush()
}

// NegativeMod can calculate a right residual whether value is positive or negative.
func NegativeMod(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

func sub(cur int) int {
	ans := 1
	for cur > 0 {
		pc := PopCountInstant(cur)
		cur %= pc
		ans++
	}
	return ans
}

// 今のrevXについて計算する
// func sub(idx, mm int) int {
// 	resi := 0

// 	for i := 0; i < n; i++ {
// 		if revX[i] == '1' {
// 			resi += R[idx][i]
// 			resi %= mm
// 		}
// 	}

// 	ans := 1
// 	resi %= mm
// 	// 負になったら無限ループするので注意
// 	for resi > 0 {
// 		pc := PopCountInstant(resi)
// 		resi %= pc
// 		ans++
// 	}

// 	return ans
// }

func PopCountInstant(val int) int {
	str := fmt.Sprintf("%b", val)

	res := 0
	for _, r := range str {
		if r == '1' {
			res++
		}
	}

	return res
}

func reigai() {
	for i := 0; i < n; i++ {
		// fmt.Println(1)
		PrintfBufStdout("1\n")
	}
	stdout.Flush()
}

func Reverse(A []rune) []rune {
	res := []rune{}

	n := len(A)
	for i := n - 1; i >= 0; i-- {
		res = append(res, A[i])
	}

	return res
}

func ReverseInts(A []int) []int {
	res := []int{}

	n := len(A)
	for i := n - 1; i >= 0; i-- {
		res = append(res, A[i])
	}

	return res
}

/*******************************************************************/

/*********** I/O ***********/

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

/*********** Debugging ***********/

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

// Strtoi is a wrapper of strconv.Atoi().
// If strconv.Atoi() returns an error, Strtoi calls panic.
func Strtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(errors.New("[argument error]: Strtoi only accepts integer string"))
	} else {
		return i
	}
}

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

// PrintfDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintfDebug(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// PrintfBufStdout is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func PrintfBufStdout(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}

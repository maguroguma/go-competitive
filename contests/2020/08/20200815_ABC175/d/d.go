/*
URL:
https://atcoder.jp/contests/abc175/tasks/abc175_d
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

var (
	n, k int
	P, C []int

	roopScore [5000 + 50]int
	roopNum   [5000 + 50]int
	R         [5000 + 50][5000 + 50]int
)

func main() {
	n, k = ReadInt2()
	P = ReadIntSlice(n)
	C = ReadIntSlice(n)

	for i := 0; i < n; i++ {
		P[i]--
	}
	// PrintfDebug("P: %v\n", P)

	for i := 0; i < n; i++ {
		score, num := dfs(i, i, 0, 0)
		roopScore[i], roopNum[i] = score, num
	}
	// PrintfDebug("roopScore: %v\n", roopScore[:n])
	// PrintfDebug("roopNum: %v\n", roopNum[:n])
	// for i := 0; i < n; i++ {
	// 	PrintfDebug("R[%d]: %v\n", i, R[i][:n+1])
	// }

	ans := -INF_INT64

	// ループを可能な限り繰り返す場合
	for i := 0; i < n; i++ {
		roopCnt := k / roopNum[i]

		rs := 0
		if roopScore[i] > 0 && roopCnt > 0 {
			rs = roopScore[i] * roopCnt
		} else {
			// ループが負の場合はスキップ、ループできない場合もスキップ
			continue
		}

		resiCnt := k % roopNum[i]
		resiS := 0
		for j := 1; j <= resiCnt; j++ {
			ChMax(&resiS, R[i][j])
		}

		ChMax(&ans, rs+resiS)
	}

	// ループを繰り返さない場合
	for i := 0; i < n; i++ {
		resiS := -INF_INT64
		maxCnt := Min(k, roopNum[i])
		for j := 1; j <= maxCnt; j++ {
			ChMax(&resiS, R[i][j])
		}
		ChMax(&ans, resiS)
	}

	// ループをギリギリまで繰り返す場合
	for i := 0; i < n; i++ {
		// 最大ループ可能回数-1
		roopCnt := (k / roopNum[i]) - 1

		rs := 0
		if roopScore[i] > 0 && roopCnt > 0 {
			rs = roopScore[i] * roopCnt
		} else {
			// ループが負の場合はスキップ、ループできない場合もスキップ
			continue
		}

		// resiCnt := k % roopNum[i]
		maxCnt := Min(k, roopNum[i])
		resiS := 0
		// for j := 1; j <= resiCnt; j++ {
		for j := 1; j <= maxCnt; j++ {
			ChMax(&resiS, R[i][j])
		}

		ChMax(&ans, rs+resiS)
	}

	fmt.Println(ans)
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

// ChMax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
}

func dfs(oid, cid, cscore, cnum int) (int, int) {
	cscore += C[P[cid]]
	cnum++
	R[oid][cnum] = cscore

	if P[cid] == oid {
		return cscore, cnum
	}

	return dfs(oid, P[cid], cscore, cnum)
}

/*******************************************************************/

/********** common constants **********/

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

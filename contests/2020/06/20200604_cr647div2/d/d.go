/*
URL:
https://codeforces.com/contest/1362/problem/D
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
	MOD          = 1000000000 + 7
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
	n, m  int
	A, B  []int
	T     []int
	G     [500000 + 50][]int
	blogs []Blog

	D [500000 + 50]int
)

func main() {
	n, m = ReadInt2()
	for i := 0; i < m; i++ {
		a, b := ReadInt2()
		a--
		b--
		A = append(A, a)
		B = append(B, b)

		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	T = ReadIntSlice(n)

	// nより大きいトピックがあったらNG
	for i := 0; i < n; i++ {
		if T[i] > n {
			fmt.Println(-1)
			return
		}
	}

	// 同じトピックが隣接していたらアウト
	for i := 0; i < n; i++ {
		for _, t := range G[i] {
			if T[i] == T[t] {
				fmt.Println(-1)
				return
			}
		}
	}

	// 孤立したブログは1以外ありえない
	for i := 0; i < n; i++ {
		if len(G[i]) == 0 && T[i] != 1 {
			fmt.Println(-1)
			return
		}
	}

	// トピックが飛んでいたらアウト
	memo := make(map[int]int)
	for i := 0; i < n; i++ {
		memo[T[i]] = 1
	}
	maxt := Max(T...)
	if len(memo) != maxt {
		fmt.Println(-1)
		return
	}

	for i := 0; i < n; i++ {
		tid, bid := T[i], i
		blogs = append(blogs, Blog{tid: tid, bid: bid})
	}

	sort.SliceStable(blogs, func(i, j int) bool {
		// if blogs[i].tid < blogs[j].tid {
		// 	return true
		// } else if blogs[i].tid > blogs[j].tid {
		// 	return false
		// } else {
		// 	return blogs[i].bid < blogs[j].bid
		// }
		return blogs[i].tid < blogs[j].tid
	})

	// for i := 0; i < n; i++ {
	// 	D[i] = 0
	// }
	answers := []int{}
	for i, blog := range blogs {
		// 以前のものよりトピックが飛んでいたらアウト
		if i >= 1 && blogs[i].tid-blogs[i-1].tid > 1 {
			fmt.Println(-1)
			return
		}

		// 隣接するものの最大をみる
		// maxi := 0
		// for _, t := range G[blog.bid] {
		// 	ChMax(&maxi, D[t])
		// }
		// if maxi+1 != blog.tid {
		// 	fmt.Println(-1)
		// 	return
		// }
		// blog.tid未満がすべて存在するかどうか調べる
		memo := make(map[int]int)
		for _, t := range G[blog.bid] {
			if 1 <= D[t] && D[t] < blog.tid {
				memo[D[t]] = 1
			}
		}
		if len(memo) != blog.tid-1 {
			fmt.Println(-1)
			return
		}

		answers = append(answers, blog.bid+1)
		D[blog.bid] = blog.tid
	}
	fmt.Println(PrintIntsLine(answers...))
}

type Blog struct {
	tid, bid int
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

/*
URL:
https://codeforces.com/contest/1388/problem/D
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
	n int
	A []int64
	B []int

	G    [200000 + 50][]int
	done [200000 + 50]bool
)

func main() {
	n = ReadInt()
	A = ReadInt64Slice(n)
	B = ReadIntSlice(n)

	for i := 0; i < n; i++ {
		b := B[i]

		if b != -1 {
			G[i] = append(G[i], b-1)
		}
	}

	_, T := tsort(n, G[:n])

	ans := int64(0)
	P := []int{}

	for _, cid := range T {
		if A[cid] <= 0 {
			continue
		}

		ans += A[cid]
		done[cid] = true
		P = append(P, cid+1)

		for _, nid := range G[cid] {
			A[nid] += A[cid]
		}
	}

	revT := Reverse(T)
	for _, cid := range revT {
		if done[cid] {
			continue
		}

		ans += A[cid]
		P = append(P, cid+1)
	}

	fmt.Println(ans)
	fmt.Println(PrintIntsLine(P...))
}

func Reverse(A []int) []int {
	res := []int{}

	n := len(A)
	for i := n - 1; i >= 0; i-- {
		res = append(res, A[i])
	}

	return res
}

// O(|E| + |V|)
// ノードIDは0-based
// ok, ans := tsort(v, edges[:])
// https://onlinejudge.u-aizu.ac.jp/problems/GRL_4_B
func tsort(nn int, edges [][]int) (bool, []int) {
	res := []int{}

	degin := make([]int, nn)
	for s := 0; s < nn; s++ {
		for _, t := range edges[s] {
			degin[t]++
		}
	}

	st := []int{}
	for nid := 0; nid < nn; nid++ {
		if degin[nid] == 0 {
			st = append(st, nid)
		}
	}

	for len(st) > 0 {
		cid := st[len(st)-1]
		res = append(res, cid)
		st = st[:len(st)-1]

		for _, nid := range edges[cid] {
			degin[nid]--
			if degin[nid] == 0 {
				st = append(st, nid)
			}
		}
	}

	if len(res) != nn {
		return false, nil
	}

	return true, res
}

// トポロジカルソート済みリストから最長経路の長さを計算する
// l, dp := longestPath(res, edges[:])
// https://atcoder.jp/contests/abc139/tasks/abc139_e
func longestPath(tsortedNodes []int, edges [][]int) (maxLength int, dp []int) {
	_chmax := func(updatedValue *int, target int) bool {
		if *updatedValue < target {
			*updatedValue = target
			return true
		}
		return false
	}

	dp = make([]int, len(tsortedNodes)+1)

	for i := 0; i < len(tsortedNodes); i++ {
		cid := tsortedNodes[i]
		for _, nid := range edges[cid] {
			_chmax(&dp[nid], dp[cid]+1)
		}
	}

	maxLength = 0
	for i := 0; i < len(tsortedNodes); i++ {
		_chmax(&maxLength, dp[i])
	}

	return maxLength, dp
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

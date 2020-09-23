/*
URL:
https://atcoder.jp/contests/indeednow-quala/tasks/indeednow_2015_quala_4
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
	h, w int
	C    [][]int8

	forward, backward map[Board]int
	steps             [][2]int
)

func main() {
	h, w = ReadInt2()
	for i := 0; i < h; i++ {
		row := ReadIntSlice(w)
		tmp := make([]int8, w)
		for i := 0; i < w; i++ {
			tmp[i] = int8(row[i])
		}

		C = append(C, tmp)
	}
	iy, ix := 0, 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if C[i][j] == 0 {
				iy, ix = i, j
				break
			}
		}
	}

	// 目標の盤面を作成する
	G := make([][]int8, h)
	for i := 0; i < h; i++ {
		G[i] = make([]int8, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			G[i][j] = int8(w*i + j + 1)
		}
	}
	G[h-1][w-1] = 0
	// for i := 0; i < h; i++ {
	// 	PrintfDebug("%v\n", G[i])
	// }

	forward, backward = make(map[Board]int), make(map[Board]int)
	// forward[makeBoard(C)] = 0
	// backward[makeBoard(G)] = 0
	steps = [][2]int{
		[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1},
	}
	dfs(C, 0, iy, ix, -1, -1, forward)
	dfs(G, 0, h-1, w-1, -1, -1, backward)

	ans := 30
	for f, fs := range forward {
		if bs, ok := backward[f]; ok {
			val := fs + bs
			ChMin(&ans, val)
		}
	}
	for b, bs := range backward {
		if fs, ok := forward[b]; ok {
			val := bs + fs
			ChMin(&ans, val)
		}
	}
	fmt.Println(ans)
}

func dfs(B [][]int8, cur int, cy, cx, by, bx int, memo map[Board]int) {
	if cur >= 14 {
		return
	}
	// 現在の盤面でマップを更新する
	bb := makeBoard(B)
	if _, ok := memo[bb]; ok {
		if memo[bb] > cur {
			memo[bb] = cur
		} else {
			return
		}
	} else {
		memo[bb] = cur
	}

	for _, s := range steps {
		ny, nx := cy+s[0], cx+s[1]

		if ny == by && nx == bx {
			continue
		}

		if 0 <= ny && ny < h && 0 <= nx && nx < w {
			// 新しい盤面を作る
			NB := make([][]int8, h)
			for i := 0; i < h; i++ {
				NB[i] = make([]int8, w)
			}
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					NB[i][j] = B[i][j]
				}
			}

			NB[cy][cx], NB[ny][nx] = NB[ny][nx], NB[cy][cx]
			dfs(NB, cur+1, ny, nx, cy, cx, memo)
		}
	}
}

func makeBoard(B [][]int8) Board {
	res := Board{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			res.B[i][j] = B[i][j]
		}
	}
	return res
}

type Board struct {
	B [6][6]int8
}

func solve() {
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

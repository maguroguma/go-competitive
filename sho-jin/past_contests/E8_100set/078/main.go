/*
URL:
https://atcoder.jp/contests/joi2011ho/tasks/joi2011ho1
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
	m, n int
	k    int
	C    [][]rune
)

func main() {
	m, n = readi2()
	k = readi()
	for i := 0; i < m; i++ {
		row := ReadRuneSlice()
		C = append(C, row)
	}

	J, O, I := make([][]int, m), make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		J[i] = make([]int, n)
		O[i] = make([]int, n)
		I[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if C[i][j] == 'J' {
				J[i][j] = 1
			} else if C[i][j] == 'O' {
				O[i][j] = 1
			} else {
				I[i][j] = 1
			}
		}
	}

	js := NewRectangleSum(J)
	os := NewRectangleSum(O)
	is := NewRectangleSum(I)

	for i := 0; i < k; i++ {
		top, left, bottom, right := readi4()
		top, left, bottom, right = top-1, left-1, bottom-1, right-1
		a := js.GetSum(top, left, bottom, right)
		b := os.GetSum(top, left, bottom, right)
		c := is.GetSum(top, left, bottom, right)
		fmt.Println(a, b, c)
	}
}

type RectangleSum struct {
	matrix [][]int
	recSum [][]int
}

// NewRectangleSum は2次元累積和を計算するための構造体のポインタを返す
func NewRectangleSum(m [][]int) *RectangleSum {
	rs := new(RectangleSum)
	rs.matrix = m

	h, w := len(m), len(m[0])
	for y := 0; y < h; y++ {
		tmp := make([]int, w)
		rs.recSum = append(rs.recSum, tmp)
	}

	// 1行ずつスキャンする
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rs.recSum[y][x] = rs.matrix[y][x] // 同じ座標の値を加算
			if y > 0 {
				rs.recSum[y][x] += rs.recSum[y-1][x] // 1マス上の座標と原点座標がなす長方形の和を加算
			}
			if x > 0 {
				rs.recSum[y][x] += rs.recSum[y][x-1] // 1マス左の座標と原点座標がなす長方形の和を加算
			}
			if y > 0 && x > 0 {
				rs.recSum[y][x] -= rs.recSum[y-1][x-1] // 過剰に加算した部分（左上のマスと原点座標がなす長方形の和）を減算
			}
		}
	}

	return rs
}

// GetSum は2次元累積和の初期化と逆の要領で、グリッド内の任意の長方形の和を計算し返す
func (rs *RectangleSum) GetSum(top, left, bottom, right int) int {
	res := rs.recSum[bottom][right]
	if left > 0 {
		res -= rs.recSum[bottom][left-1]
	}
	if top > 0 {
		res -= rs.recSum[top-1][right]
	}
	if left > 0 && top > 0 {
		res += rs.recSum[top-1][left-1]
	}
	return res
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

// readi returns an integer.
func readi() int {
	return int(readInt64())
}
func readi2() (int, int) {
	return int(readInt64()), int(readInt64())
}
func readi3() (int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64())
}
func readi4() (int, int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64()), int(readInt64())
}

// readll returns as integer as int64.
func readll() int64 {
	return readInt64()
}
func readll2() (int64, int64) {
	return readInt64(), readInt64()
}
func readll3() (int64, int64, int64) {
	return readInt64(), readInt64(), readInt64()
}
func readll4() (int64, int64, int64, int64) {
	return readInt64(), readInt64(), readInt64(), readInt64()
}

func readInt64() int64 {
	i, err := strconv.ParseInt(ReadString(), 0, 64)
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

/*
URL:
https://atcoder.jp/contests/joisc2007/tasks/joisc2007_mall
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
	n, m int
	a, b int
	A    [][]int
)

func main() {
	defer stdout.Flush()

	m, n = readi2()
	b, a = readi2()
	for i := 0; i < n; i++ {
		row := readis(m)
		A = append(A, row)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == -1 {
				A[i][j] = INF_B30
			}
		}
	}

	rs := NewRectangleSum(A)

	ans := INF_B30
	for i := 0; i+a-1 < n; i++ {
		for j := 0; j+b-1 < m; j++ {
			val := rs.GetSum(i, j, i+a-1, j+b-1)
			ChMin(&ans, val)
		}
	}
	fmt.Println(ans)
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
	MOD = 1000000000 + 7
	// MOD          = 998244353
	ALPH_N  = 26
	INF_I64 = math.MaxInt64
	INF_B60 = 1 << 60
	INF_I32 = math.MaxInt32
	INF_B30 = 1 << 30
	NIL     = -1
)

// modi can calculate a right residual whether value is positive or negative.
func modi(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// modll can calculate a right residual whether value is positive or negative.
func modll(val, m int64) int64 {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	reads = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
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

/*
URL:
https://onlinejudge.u-aizu.ac.jp/courses/library/7/DPL/3/DPL_3_B
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

var (
	println = fmt.Println

	h, w int
	G    [][]int32
)

func main() {
	defer stdout.Flush()

	h, w = readi2()
	for i := 0; i < h; i++ {
		row := readis(w)
		tmp := make([]int32, w)
		for j := 0; j < w; j++ {
			tmp[j] = int32(row[j])
		}
		G = append(G, tmp)
	}

	_, _, U, _ := GridLRUD(G)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			U[i][j]++
		}
	}
	// debugf("U: %v\n", U)

	ans := 0
	for i := 0; i < h; i++ {
		row := U[i]
		tmp := make([]int, w)
		for j := 0; j < w; j++ {
			tmp[j] = int(row[j])
		}

		res := LargestRectangle(tmp)
		chmax(&ans, res)
	}

	println(ans)
}

// GridLRUD returns matrices that say how many cells you can move from S[i][j].
func GridLRUD(S [][]int32) (L, R, U, D [][]int32) {
	// const BLOCK_CELL, EMPTY_CELL = '#', '.'
	const BLOCK_CELL, EMPTY_CELL = 1, 0

	h, w := len(S), len(S[0])
	T := [][]int32{}

	wall := make([]int32, w+2)
	for i := 0; i < len(wall); i++ {
		wall[i] = BLOCK_CELL
	}

	T = append(T, wall)
	for i := 0; i < h; i++ {
		row := []int32{BLOCK_CELL}
		row = append(row, S[i]...)
		row = append(row, BLOCK_CELL)
		T = append(T, row)
	}
	T = append(T, wall)

	L, R, U, D =
		make([][]int32, h+2), make([][]int32, h+2), make([][]int32, h+2), make([][]int32, h+2)
	for i := 0; i < h+2; i++ {
		L[i], R[i], U[i], D[i] =
			make([]int32, w+2), make([]int32, w+2), make([]int32, w+2), make([]int32, w+2)
	}
	for i := 0; i < h+2; i++ {
		for j := 0; j < w+2; j++ {
			if T[i][j] == EMPTY_CELL {
				continue
			}
			// for block
			U[i][j], D[i][j], L[i][j], R[i][j] = -1, -1, -1, -1
		}
	}

	for y := 1; y <= h; y++ {
		for x := 1; x <= w; x++ {
			if T[y][x] == BLOCK_CELL {
				continue
			}
			U[y][x] = U[y-1][x] + 1
			L[y][x] = L[y][x-1] + 1
		}
	}
	for y := h; y >= 1; y-- {
		for x := w; x >= 1; x-- {
			if T[y][x] == BLOCK_CELL {
				continue
			}
			D[y][x] = D[y+1][x] + 1
			R[y][x] = R[y][x+1] + 1
		}
	}

	cut := func(G [][]int32) [][]int32 {
		res := [][]int32{}
		for i := 1; i <= h; i++ {
			res = append(res, G[i][1:w+1])
		}
		return res
	}

	L, R, U, D = cut(L), cut(R), cut(U), cut(D)

	return
}

// LargestRectangle calculates an area of largest rectangle in a histgram.
// Time complexity: O(len(height))
func LargestRectangle(height []int) int {
	H := make([]int, len(height))
	copy(H, height)

	_max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	_top := func(S []int) int { return S[len(S)-1] }

	st := []int{}
	H = append(H, 0)
	left := make([]int, len(H))
	res := int(0)

	for i := int(0); i < int(len(H)); i++ {
		for len(st) > 0 && H[_top(st)] >= H[i] {
			res = _max(res, (i-left[_top(st)]-1)*H[_top(st)])
			st = st[:len(st)-1]
		}

		if len(st) == 0 {
			left[i] = -1
		} else {
			left[i] = _top(st)
		}

		st = append(st, i)
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

/*
URL:
https://atcoder.jp/contests/hhkb2020/tasks/hhkb2020_e
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
	h, w int
	n, m int
	S    [][]rune

	P [4000000 + 1000]int
)

func main() {
	defer stdout.Flush()

	h, w = readi2()
	for i := 0; i < h; i++ {
		row := readrs()
		S = append(S, row)
	}

	E := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if S[y][x] == '#' {
				continue
			}
			E++
		}
	}

	P[0] = 1
	for i := 1; i <= 2000*2000+10; i++ {
		P[i] = P[i-1] * 2
		P[i] %= MOD
	}

	L, R, U, D := GridLRUD(S)

	ans := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if S[y][x] == '#' {
				continue
			}

			u, d, l, r := U[y][x], D[y][x], L[y][x], R[y][x]
			num := u + d + l + r + 1

			nokori := E - num

			A := P[num]
			A = mod(A-1, MOD)
			B := P[nokori]

			pat := A * B
			pat %= MOD

			ans += pat
			ans %= MOD
		}
	}

	fmt.Println(ans)
}

// GridLRUD returns matrices that say how many cells you can move from S[i][j].
func GridLRUD(S [][]rune) (L, R, U, D [][]int) {
	const BLOCK_CELL, EMPTY_CELL = '#', '.'

	h, w := len(S), len(S[0])
	T := [][]rune{}

	wall := make([]rune, w+2)
	for i := 0; i < len(wall); i++ {
		wall[i] = BLOCK_CELL
	}

	T = append(T, wall)
	for i := 0; i < h; i++ {
		row := []rune{BLOCK_CELL}
		row = append(row, S[i]...)
		row = append(row, BLOCK_CELL)
		T = append(T, row)
	}
	T = append(T, wall)

	L, R, U, D =
		make([][]int, h+2), make([][]int, h+2), make([][]int, h+2), make([][]int, h+2)
	for i := 0; i < h+2; i++ {
		L[i], R[i], U[i], D[i] =
			make([]int, w+2), make([]int, w+2), make([]int, w+2), make([]int, w+2)
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

	cut := func(G [][]int) [][]int {
		res := [][]int{}
		for i := 1; i <= h; i++ {
			res = append(res, G[i][1:w+1])
		}
		return res
	}

	L, R, U, D = cut(L), cut(R), cut(U), cut(D)

	return
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

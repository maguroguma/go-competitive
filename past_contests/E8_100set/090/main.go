/*
URL:
https://atcoder.jp/contests/s8pc-5/tasks/s8pc_5_b
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
	n, m    int
	X, Y, R []float64
)

const ir = 100000.0

func main() {
	defer stdout.Flush()

	n, m = readi2()
	for i := 0; i < n; i++ {
		x, y, r := readi3()
		xf, yf, rf := float64(x), float64(y), float64(r)
		X = append(X, xf)
		Y = append(Y, yf)
		R = append(R, rf)
	}
	for i := 0; i < m; i++ {
		x, y := readi2()
		xf, yf := float64(x), float64(y)
		X = append(X, xf)
		Y = append(Y, yf)
	}

	ok := BinarySearch(0.0, ir, func(r float64) bool {
		for i := 0; i < n+m; i++ {
			for j := i + 1; j < n+m; j++ {
				var x1, y1, r1, x2, y2, r2 float64
				x1, y1, x2, y2 = X[i], Y[i], X[j], Y[j]
				if i < n {
					r1 = R[i]
				} else {
					r1 = r
				}
				if j < n {
					r2 = R[j]
				} else {
					r2 = r
				}

				if isOverlap(x1, y1, r1, x2, y2, r2) {
					return false
				}
			}
		}

		return true
	})

	// if ok >= ir-1.0 {
	if m == 0 || n == 0 {
		for i := 0; i < n; i++ {
			ok = math.Min(ok, R[i])
		}
		PrintfBufStdout("%v\n", ok)
	} else {
		PrintfBufStdout("%v\n", ok)
	}
}

func isOverlap(x1, y1, r1, x2, y2, r2 float64) bool {
	d := Distance(x1, y1, x2, y2)

	// 統合を含めるとWAになる
	return d <= r1+r2
	// return d < r1+r2
}

func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func UnitVector(dx, dy float64) (ex, ey float64) {
	dist := math.Sqrt(dx*dx + dy*dy)
	return dx / dist, dy / dist
}

func Rotate90(cx, cy float64) (nx, ny float64) {
	return -cy, cx
}

func RotateN(cx, cy, radi float64) (nx, ny float64) {
	nx = math.Cos(radi)*cx - math.Sin(radi)*cy
	ny = math.Sin(radi)*cx + math.Cos(radi)*cy
	return nx, ny
}

func Midpoint(x1, y1, x2, y2 float64) (mx, my float64) {
	return (x1 + x2) / 2.0, (y1 + y2) / 2.0
}

func BinarySearch(initOK, initNG float64, isOK func(mid float64) bool) (ok float64) {
	ng := initNG
	ok = initOK
	for i := 0; i < 1000; i++ {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
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

// modint can calculate a right residual whether value is positive or negative.
func modint(val, m int) int {
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

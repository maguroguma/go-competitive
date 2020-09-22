/*
URL:
https://atcoder.jp/contests/joisc2013-day1/tasks/joisc2013_joi_poster
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
	n, w, h int
	X, Y    []int

	wf, hf float64
)

func main() {
	defer stdout.Flush()

	n, w, h = readi3()
	for i := 0; i < n; i++ {
		x, y := readi2()
		X = append(X, x)
		Y = append(Y, y)
	}

	wf, hf = float64(w), float64(h)

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					memo := map[int]int{}
					memo[i]++
					memo[j]++
					memo[k]++
					memo[l]++
					if len(memo) != 4 {
						continue
					}

					if check(i, j, k, l) {
						ans++
					}
				}
			}
		}
	}

	fmt.Println(ans)
}

type Coord struct {
	x, y float64
}

const (
	e = 1e-10
)

func check(i, j, k, l int) bool {
	c1 := Coord{float64(X[i]), float64(Y[i])}
	a := Coord{float64(X[j]), float64(Y[j])}
	c2 := Coord{float64(X[k]), float64(Y[k])}
	b := Coord{float64(X[l]), float64(Y[l])}

	r1 := Distance(c1.x, c1.y, a.x, a.y)
	r2 := Distance(c2.x, c2.y, b.x, b.y)

	// 誤差で死ぬ！
	dist := Distance(c1.x, c1.y, c2.x, c2.y)
	// if !(dist < math.Abs(r1-r2) && r2 > r1) {
	// 	return false
	// }
	if !(dist+e < math.Abs(r1-r2) && r2 > r1) {
		return false
	}

	// 整数に閉じたやり方
	// D := (X[i]-X[k])*(X[i]-X[k]) + (Y[i]-Y[k])*(Y[i]-Y[k])
	// R1 := (X[i]-X[j])*(X[i]-X[j]) + (Y[i]-Y[j])*(Y[i]-Y[j])
	// R2 := (X[k]-X[l])*(X[k]-X[l]) + (Y[k]-Y[l])*(Y[k]-Y[l])
	// if !(R1+R2-D >= 0 && 4*R1*R2 < (R1+R2-D)*(R1+R2-D) && r2 > r1) {
	// 	return false
	// }

	left := c2.x - r2
	right := c2.x + r2
	up := c2.y + r2
	down := c2.y - r2

	if !(left >= 0.0 && right <= wf && up <= hf && down >= 0.0) {
		return false
	}

	return true
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

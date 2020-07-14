/*
URL:
https://atcoder.jp/contests/abc151/tasks/abc151_f
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
	n    int
	X, Y []float64
)

func main() {
	n = ReadInt()
	for i := 0; i < n; i++ {
		x, y := ReadInt2()
		xf, yf := float64(x), float64(y)
		X = append(X, xf)
		Y = append(Y, yf)
	}

	// res := calcTwoIntersections(0.0, 0.0, 4.0, 0.0, 1.0)
	// PrintfDebug("%v\n", res)
	// res = calcTwoIntersections(0.0, 0.0, 2.0, 2.0, 1.414)
	// PrintfDebug("%v\n", res)

	ans := BinarySearch(10000000.0, 0.0, func(r float64) bool {
		A := [][2]float64{}
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				x1, y1, x2, y2 := X[i], Y[i], X[j], Y[j]

				// 2点間の距離
				d := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
				if d > 2.0*r {
					// 2円は交点を生じない
					continue
				}

				// 2点の中点と、交点との距離
				h := calcH(d, r)

				// 2点間のベクトル
				dx := x1 - x2
				dy := y1 - y2
				// 単位ベクトル化
				dist := math.Sqrt(dx*dx + dy*dy)
				dx /= dist
				dy /= dist
				// 90度回転
				dx, dy = -dy, dx
				// 中点の座標
				mx, my := (x1+x2)/2.0, (y1+y2)/2.0
				// 交点を2つ求める
				x3, y3 := mx+h*dx, my+h*dy
				x4, y4 := mx-h*dx, my-h*dy
				// Aに追加
				A = append(A, [2]float64{x3, y3})
				A = append(A, [2]float64{x4, y4})
			}
		}

		for _, P := range A {
			cx, cy := P[0], P[1]
			num := 0
			for i := 0; i < n; i++ {
				x, y := X[i], Y[i]
				dist := math.Sqrt((cx-x)*(cx-x) + (cy-y)*(cy-y))
				// 交点を作る2円を確実に含めるように微小値を加える
				if dist <= (r + 0.0000001) {
					num++
				}
			}

			if num >= n {
				return true
			}
		}
		return false
	})

	fmt.Println(ans)
}

func calcH(d, r float64) float64 {
	return math.Sqrt(r*r - (d/2.0)*(d/2.0))
}

func calcTwoIntersections(x1, y1, x2, y2, h float64) [][2]float64 {
	// 2点間のベクトル
	dx := x1 - x2
	dy := y1 - y2
	// 単位ベクトル化
	dist := math.Sqrt(dx*dx + dy*dy)
	dx /= dist
	dy /= dist
	// 90度回転
	dx, dy = -dy, dx
	// 中点の座標
	mx, my := (x1+x2)/2.0, (y1+y2)/2.0
	// 交点を2つ求める
	x3, y3 := mx+h*dx, my+h*dy
	x4, y4 := mx-h*dx, my-h*dy

	return [][2]float64{
		[2]float64{x3, y3}, [2]float64{x4, y4},
	}
}

func BinarySearch(initOK, initNG float64, isOK func(mid float64) bool) (ok float64) {
	ng := initNG
	ok = initOK
	// for int(math.Abs(float64(ok-ng))) > 1 {
	for i := 0; i < 50; i++ {
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

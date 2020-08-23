/*
URL:
https://atcoder.jp/contests/abc176/tasks/abc176_e
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

type Coord struct {
	y, x int
}

var (
	h, w, m int
	H, W    []int
	HH, WW  []int

	memo map[Coord]int
	Y, X [300000 + 50][]int

	xn []V
	yn []V
)

type V struct {
	x   int
	num int
}

// func main() {
// 	h, w, m = ReadInt3()
// 	memo = make(map[Coord]int)
// 	for i := 0; i < m; i++ {
// 		y, x := ReadInt2()
// 		y--
// 		x--

// 		H = append(H, y)
// 		W = append(W, x)

// 		Y[y] = append(Y[y], x)
// 		X[x] = append(X[x], y)

// 		memo[Coord{y: y, x: x}] = 1
// 	}

// 	for i := 0; i < m; i++ {
// 		y := H[i]
// 		HH = append(HH, y)
// 		x := W[i]
// 		WW = append(WW, x)
// 	}
// 	sort.Sort(sort.IntSlice(HH))
// 	sort.Sort(sort.IntSlice(WW))
// 	hc, _ := RunLengthEncoding(HH)
// 	wc, _ := RunLengthEncoding(WW)
// 	// PrintfDebug("%v\n", hc)
// 	// PrintfDebug("%v\n", wc)

// 	// for _, x := range W {
// 	// 	xn = append(xn, V{x: x, num: len(X[x])})
// 	// }
// 	for _, x := range wc {
// 		xn = append(xn, V{x: x, num: len(X[x])})
// 	}
// 	sort.SliceStable(xn, func(i, j int) bool {
// 		return xn[i].num > xn[j].num
// 	})

// 	ans := 0
// 	// すべてのy座標を試す
// 	// for _, y := range H {
// 	for _, y := range hc {
// 		ynum := len(Y[y])

// 		for i := 0; i < Min(len(xn), 20); i++ {
// 			all := ynum + xn[i].num
// 			co := Coord{y: y, x: xn[i].x}
// 			if _, ok := memo[co]; ok {
// 				all--
// 			}
// 			ChMax(&ans, all)
// 		}
// 	}

// 	fmt.Println(ans)
// }

func main() {
	h, w, m = ReadInt3()
	memo = make(map[Coord]int)
	for i := 0; i < m; i++ {
		y, x := ReadInt2()
		y--
		x--

		H = append(H, y)
		W = append(W, x)

		Y[y] = append(Y[y], x)
		X[x] = append(X[x], y)

		memo[Coord{y: y, x: x}] = 1
	}

	for i := 0; i < m; i++ {
		y := H[i]
		HH = append(HH, y)
		x := W[i]
		WW = append(WW, x)
	}
	sort.Sort(sort.IntSlice(HH))
	sort.Sort(sort.IntSlice(WW))
	hc, _ := RunLengthEncoding(HH)
	wc, _ := RunLengthEncoding(WW)
	// PrintfDebug("%v\n", hc)
	// PrintfDebug("%v\n", wc)

	for _, x := range wc {
		xn = append(xn, V{x: x, num: len(X[x])})
	}
	sort.SliceStable(xn, func(i, j int) bool {
		return xn[i].num > xn[j].num
	})
	for _, y := range hc {
		yn = append(yn, V{x: y, num: len(Y[y])})
	}
	sort.SliceStable(yn, func(i, j int) bool {
		return yn[i].num > yn[j].num
	})

	ans := 0

	for i := 0; i < m; i++ {
		y, x := H[i], W[i]
		num := len(X[x]) + len(Y[y]) - 1
		ChMax(&ans, num)
	}

	for i := 0; i < Min(len(yn), 2000); i++ {
		for j := 0; j < Min(len(xn), 2000); j++ {
			num := yn[i].num + xn[j].num
			co := Coord{y: yn[i].x, x: xn[j].x}
			if _, ok := memo[co]; ok {
				num--
			}
			ChMax(&ans, num)
			// if ChMax(&ans, num) {
			// 	PrintfDebug("%v, %v\n", co.y, co.x)
			// }
		}
	}
	fmt.Println(ans)

	// すべてのy座標を試す
	// for _, y := range H {
	// for _, y := range hc {
	// 	ynum := len(Y[y])

	// 	for i := 0; i < Min(len(xn), 20); i++ {
	// 		all := ynum + xn[i].num
	// 		co := Coord{y: y, x: xn[i].x}
	// 		if _, ok := memo[co]; ok {
	// 			all--
	// 		}
	// 		ChMax(&ans, all)
	// 	}
	// }

	// fmt.Println(ans)
}

// RunLengthEncoding returns encoded slice of an input.
func RunLengthEncoding(S []int) ([]int, []int) {
	runes := []int{}
	lengths := []int{}

	l := 0
	for i := 0; i < len(S); i++ {
		// 1文字目の場合保持
		if i == 0 {
			l = 1
			continue
		}

		if S[i-1] == S[i] {
			// 直前の文字と一致していればインクリメント
			l++
		} else {
			// 不一致のタイミングで追加し、長さをリセットする
			runes = append(runes, S[i-1])
			lengths = append(lengths, l)
			l = 1
		}
	}
	runes = append(runes, S[len(S)-1])
	lengths = append(lengths, l)

	return runes, lengths
}

// RunLengthDecoding decodes RLE results.
func RunLengthDecoding(S []int, L []int) []int {
	if len(S) != len(L) {
		panic("S, L are not RunLengthEncoding results")
	}

	res := []int{}

	for i := 0; i < len(S); i++ {
		for j := 0; j < L[i]; j++ {
			res = append(res, S[i])
		}
	}

	return res
}

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int) int {
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

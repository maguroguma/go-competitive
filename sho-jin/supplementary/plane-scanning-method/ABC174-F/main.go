/*
URL:
https://atcoder.jp/contests/abc174/tasks/abc174_f
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

var (
	n, q int
	C    []int

	P [500000 + 50]int
	// E []Segment
	E [][4]int
)

// type Segment struct {
// 	c    int
// 	qid  int
// 	l, r int
// }

func main() {
	n, q = ReadInt2()
	C = ReadIntSlice(n)

	// E = make([]Segment, 0, 2000000)

	for i := 0; i <= n; i++ {
		P[i] = -1
	}
	for i := 0; i < n; i++ {
		c := C[i]

		if P[c] == -1 {
			P[c] = i
			continue
		}
		// seg := Segment{c: 0, l: P[c], r: i}
		// E = append(E, seg)
		E = append(E, [4]int{P[c], i, 0, 0})
		P[c] = i
	}
	// PrintfDebug("%v\n", P[:n+1])

	for i := 0; i < q; i++ {
		l, r := ReadInt2()
		// seg := Segment{c: 1, l: l - 1, r: r - 1, qid: i}
		// E = append(E, seg)
		E = append(E, [4]int{l - 1, r - 1, 1, i})
	}

	sort.SliceStable(E, func(i, j int) bool {
		// if E[i].l > E[j].l {
		// 	return true
		// } else if E[i].l < E[j].l {
		// 	return false
		// } else {
		// 	return E[i].c < E[j].c
		// }
		if E[i][0] > E[j][0] {
			return true
		} else if E[i][0] < E[j][0] {
			return false
		} else {
			return E[i][2] < E[j][2]
		}
	})

	answers := make([]int, q)
	ft := NewFenwickTree(500000 + 50)
	for _, seg := range E {
		// if seg.c == 0 {
		if seg[2] == 0 {
			// 点の追加
			// ft.Add(seg.r, 1)
			ft.Add(seg[1], 1)
		} else {
			// クエリ処理
			// ans := ft.Sum(seg.r)
			// answers[seg.qid] = (seg.r - seg.l + 1) - ans
			ans := ft.Sum(seg[1])
			answers[seg[3]] = (seg[1] - seg[0] + 1) - ans
		}
	}

	for i := 0; i < q; i++ {
		PrintfBufStdout("%d\n", answers[i])
	}
	stdout.Flush()
}

// Public methods
// ft := NewFenwickTree(200000 + 5)
// s := ft.Sum(i) 						// Sum of [1, i] (1-based)
// ft.Add(i, x) 							// Add x to i (1-based)
// idx := ft.LowerBound(w) 		// minimum idx such that bit.Sum(idx) >= w

type FenwickTree struct {
	dat     []int
	n       int
	minPow2 int
}

// n(>=1) is number of elements of original data
func NewFenwickTree(n int) *FenwickTree {
	ft := new(FenwickTree)

	ft.dat = make([]int, n+1)
	ft.n = n

	ft.minPow2 = 1
	for {
		if (ft.minPow2 << 1) > n {
			break
		}
		ft.minPow2 <<= 1
	}

	return ft
}

// Sum of [1, i](1-based)
func (ft *FenwickTree) Sum(i int) int {
	s := 0

	for i > 0 {
		s += ft.dat[i]
		i -= i & (-i)
	}

	return s
}

// Add x to i(1-based)
func (ft *FenwickTree) Add(i, x int) {
	for i <= ft.n {
		ft.dat[i] += x
		i += i & (-i)
	}
}

// LowerBound returns minimum i such that bit.Sum(i) >= w.
func (ft *FenwickTree) LowerBound(w int) int {
	if w <= 0 {
		return 0
	}

	x := 0
	for k := ft.minPow2; k > 0; k /= 2 {
		if x+k <= ft.n && ft.dat[x+k] < w {
			w -= ft.dat[x+k]
			x += k
		}
	}

	return x + 1
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

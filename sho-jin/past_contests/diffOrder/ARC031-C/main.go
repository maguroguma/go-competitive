/*
URL:
https://atcoder.jp/contests/arc031/tasks/arc031_3
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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

func Reverse(A []int) []int {
	res := []int{}

	n := len(A)
	for i := n - 1; i >= 0; i-- {
		res = append(res, A[i])
	}

	return res
}

var (
	n int
	B []int

	A    []int
	revA []int
	R, L []int
	idx  int
)

func main() {
	n = ReadInt()
	B = ReadIntSlice(n)

	if n <= 2 {
		fmt.Println(0)
		return
	}

	C := []Item{}
	for i := 0; i < n; i++ {
		C = append(C, Item{idx: i + 1, val: B[i]})
	}
	sort.SliceStable(C, func(i, j int) bool {
		return C[i].val < C[j].val
	})

	bit := NewBIT(200000 + 50)
	for i := 1; i <= n; i++ {
		bit.Add(i, 1)
	}

	ans := 0
	for _, c := range C {
		ans += Min(bit.Sum(c.idx-1), bit.Sum(n)-bit.Sum(c.idx))
		bit.Add(c.idx, -1)
	}
	fmt.Println(ans)
}

type Item struct {
	idx, val int
}

// func main() {
// 	n = ReadInt()
// 	B = ReadIntSlice(n)

// 	if n <= 2 {
// 		fmt.Println(0)
// 		return
// 	}

// 	A = make([]int, len(B))
// 	copy(A, B)

// 	for i := 0; i < n; i++ {
// 		if A[i] == n {
// 			A[i] = 0
// 			idx = i
// 			break
// 		}
// 	}
// 	revA = Reverse(A)

// 	R, L = make([]int, n), make([]int, n)

// 	bit := NewBIT(200000 + 50)
// 	for i := 0; i < n; i++ {
// 		if A[i] == 0 {
// 			continue
// 		}
// 		R[i] = i - bit.Sum(A[i])
// 		bit.Add(A[i], 1)
// 	}
// 	bit = NewBIT(200000 + 50)
// 	for i := 0; i < n; i++ {
// 		if revA[i] == 0 {
// 			continue
// 		}
// 		L[i] = i - bit.Sum(revA[i])
// 		bit.Add(revA[i], 1)
// 	}
// 	L = Reverse(L)
// 	PrintfDebug("R: %v\n", R)
// 	PrintfDebug("L: %v\n", L)

// 	prefR, suffL := make([]int, n), make([]int, n)
// 	prefR[0] = R[0]
// 	for i := 1; i < n; i++ {
// 		prefR[i] = prefR[i-1] + R[i]
// 	}
// 	suffL[n-1] = L[n-1]
// 	for i := n - 2; i >= 0; i-- {
// 		suffL[i] = suffL[i+1] + L[i]
// 	}
// 	PrintfDebug("prefR: %v\n", prefR)
// 	PrintfDebug("suffL: %v\n", suffL)

// 	ans := Min(prefR[n-2]+AbsInt(idx-(n-1)), suffL[1]+AbsInt(idx-0))
// 	for i := 1; i < n-1; i++ {
// 		ChMin(&ans, prefR[i-1]+suffL[i+1]+AbsInt(idx-i))
// 	}
// 	fmt.Println(ans)
// }

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

// Public methods
// bit := NewBIT(200000 + 5)
// s := bit.Sum(i) 						// Sum of [1, i] (1-based)
// bit.Add(i, x) 							// Add x to i (1-based)
// idx := bit.LowerBound(w) 	// minimum idx such that bit.Sum(idx) >= w

type BinaryIndexedTree struct {
	bit     []int
	n       int
	minPow2 int
}

// n(>=1) is number of elements of original data
func NewBIT(n int) *BinaryIndexedTree {
	newBit := new(BinaryIndexedTree)

	newBit.bit = make([]int, n+1)
	newBit.n = n

	newBit.minPow2 = 1
	for {
		if (newBit.minPow2 << 1) > n {
			break
		}
		newBit.minPow2 <<= 1
	}

	return newBit
}

// Sum of [1, i](1-based)
func (b *BinaryIndexedTree) Sum(i int) int {
	s := 0

	for i > 0 {
		s += b.bit[i]
		i -= i & (-i)
	}

	return s
}

// Add x to i(1-based)
func (b *BinaryIndexedTree) Add(i, x int) {
	for i <= b.n {
		b.bit[i] += x
		i += i & (-i)
	}
}

// LowerBound returns minimum i such that bit.Sum(i) >= w.
func (b *BinaryIndexedTree) LowerBound(w int) int {
	if w <= 0 {
		return 0
	}

	x := 0
	for k := b.minPow2; k > 0; k /= 2 {
		if x+k <= b.n && b.bit[x+k] < w {
			w -= b.bit[x+k]
			x += k
		}
	}

	return x + 1
}

func solve() {
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

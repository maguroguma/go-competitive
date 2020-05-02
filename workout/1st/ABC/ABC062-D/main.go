/*
URL:
https://atcoder.jp/contests/abc062/tasks/arc074_b
*/

package main

import (
	"bufio"
	"container/heap"
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
	MOD          = 1000000000 + 7
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
	n int
	A []int

	used    [300000 + 50]bool
	befSums [300000 + 50]int
	aftSums [300000 + 50]int
)

func main() {
	n = ReadInt()
	A = ReadIntSlice(n * 3)

	temp := make(BeforePQ, 0)
	befPQ := &temp
	heap.Init(befPQ)
	temp2 := make(AfterPQ, 0)
	aftPQ := &temp2
	heap.Init(aftPQ)

	befSum, aftSum := 0, 0

	for i := 0; i < n; i++ {
		heap.Push(befPQ, &Before{pri: A[i], id: i})
		befSum += A[i]
		used[i] = true
	}
	befSums[n-1] = befSum
	for i := n; i < 2*n; i++ {
		bp := heap.Pop(befPQ).(*Before)
		if bp.pri < A[i] {
			// 交換
			befSum -= bp.pri
			befSum += A[i]
			heap.Push(befPQ, &Before{pri: A[i], id: i})
		} else {
			// もとに戻す
			heap.Push(befPQ, bp)
		}

		befSums[i] = befSum
	}

	for i := 2 * n; i < 3*n; i++ {
		heap.Push(aftPQ, &After{pri: A[i], id: i})
		aftSum += A[i]
		used[i] = true
	}
	aftSums[2*n-1] = aftSum
	for i := 2*n - 1; i >= n; i-- {
		ap := heap.Pop(aftPQ).(*After)
		if ap.pri > A[i] {
			// 交換
			aftSum -= ap.pri
			aftSum += A[i]
			heap.Push(aftPQ, &After{pri: A[i], id: i})
		} else {
			// もとに戻す
			heap.Push(aftPQ, ap)
		}

		aftSums[i-1] = aftSum
	}

	PrintfDebug("aftSums: %v\n", aftSums[:3*n])
	PrintfDebug("befSums: %v\n", befSums[:3*n])

	ans := -INF_BIT60
	for i := n - 1; i < 2*n; i++ {
		ChMax(&ans, befSums[i]-aftSums[i])
	}
	fmt.Println(ans)
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

type After struct {
	pri int
	id  int
}
type AfterPQ []*After

func (pq AfterPQ) Len() int           { return len(pq) }
func (pq AfterPQ) Less(i, j int) bool { return pq[i].pri > pq[j].pri } // <: ASC, >: DESC
func (pq AfterPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *AfterPQ) Push(x interface{}) {
	item := x.(*After)
	*pq = append(*pq, item)
}
func (pq *AfterPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// how to use
// temp := make(AfterPQ, 0, 100000+1)
// pq := &temp
// heap.Init(pq)
// heap.Push(pq, &After{pri: intValue})
// popped := heap.Pop(pq).(*After)

type Before struct {
	pri int
	id  int
}
type BeforePQ []*Before

func (pq BeforePQ) Len() int           { return len(pq) }
func (pq BeforePQ) Less(i, j int) bool { return pq[i].pri < pq[j].pri } // <: ASC, >: DESC
func (pq BeforePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *BeforePQ) Push(x interface{}) {
	item := x.(*Before)
	*pq = append(*pq, item)
}
func (pq *BeforePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// how to use
// temp := make(BeforePQ, 0, 100000+1)
// pq := &temp
// heap.Init(pq)
// heap.Push(pq, &Before{pri: intValue})
// popped := heap.Pop(pq).(*Before)

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

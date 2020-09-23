/*
URL:
https://atcoder.jp/contests/past202004-open/tasks/past202004_l
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

var (
	n, k, d int
	A       []int

	dp [200000 + 50]int
)

func main() {
	n, k, d = ReadInt3()
	A = ReadIntSlice(n)

	for i := n - 1; i >= 0; i-- {
		if i+d < n {
			dp[i] = dp[i+d] + 1
		} else {
			dp[i] = 1
		}
	}
	PrintfDebug("dp: %v\n", dp[:n])

	f := func(lv, rv T) T {
		if lv.val < rv.val {
			return lv
		} else if lv.val > rv.val {
			return rv
		} else {
			if lv.id < rv.id {
				return lv
			} else {
				return rv
			}
		}
	}
	ti := T{id: 200000 + 50, val: 1000000000 + 50}
	st := NewSegmentTree(n, f, ti)
	for i := 0; i < n; i++ {
		st.Set(i, T{id: i, val: A[i]})
	}
	st.Build()

	answers := []int{}
	left := 0
	nokori := k
	for nokori > 0 {
		right := BinarySearch(-1, n, func(mid int) bool {
			return dp[mid] >= nokori
		})

		if right-left+1 <= 0 {
			fmt.Println(-1)
			return
		}

		t := st.Query(left, right+1)
		answers = append(answers, t.val)
		left = t.id + d
		nokori--
	}

	fmt.Println(PrintIntsLine(answers...))
}

func BinarySearch(initOK, initNG int, isOK func(mid int) bool) (ok int) {
	ng := initNG
	ok = initOK
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

// type T int // (T, f): Monoid
type T struct {
	id, val int
}

type SegmentTree struct {
	sz   int              // minimum power of 2
	data []T              // elements in T
	f    func(lv, rv T) T // T <> T -> T
	ti   T                // identity element of Monoid
}

func NewSegmentTree(
	n int, f func(lv, rv T) T, ti T,
) *SegmentTree {
	st := new(SegmentTree)
	st.ti = ti
	st.f = f

	st.sz = 1
	for st.sz < n {
		st.sz *= 2
	}

	st.data = make([]T, 2*st.sz-1)
	for i := 0; i < 2*st.sz-1; i++ {
		st.data[i] = st.ti
	}

	return st
}

func (st *SegmentTree) Set(k int, x T) {
	st.data[k+(st.sz-1)] = x
}

func (st *SegmentTree) Build() {
	for i := st.sz - 2; i >= 0; i-- {
		st.data[i] = st.f(st.data[2*i+1], st.data[2*i+2])
	}
}

func (st *SegmentTree) Update(k int, x T) {
	k += st.sz - 1
	st.data[k] = x

	for k > 0 {
		k = (k - 1) / 2
		st.data[k] = st.f(st.data[2*k+1], st.data[2*k+2])
	}
}

func (st *SegmentTree) Query(a, b int) T {
	return st.query(a, b, 0, 0, st.sz)
}

func (st *SegmentTree) query(a, b, k, l, r int) T {
	if r <= a || b <= l {
		return st.ti
	}

	if a <= l && r <= b {
		return st.data[k]
	}

	lv := st.query(a, b, 2*k+1, l, (l+r)/2)
	rv := st.query(a, b, 2*k+2, (l+r)/2, r)
	return st.f(lv, rv)
}

func (st *SegmentTree) Get(k int) T {
	return st.data[k+(st.sz-1)]
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

/*
URL:
https://atcoder.jp/contests/past202004-open/tasks/past202004_n
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
	n, q       int
	X, Y, D, C []int
	A, B       []int
)

type Query struct {
	x, y   int
	y1, y2 int
	cost   int
	qid    int
	c      int // 1: 加算, 2: 座標, 3: 減算
}

func main() {
	n, q = ReadInt2()
	ytmp := []int{}
	Q := []Query{}
	for i := 0; i < n; i++ {
		x, y, d, c := ReadInt4()
		X = append(X, x)
		Y = append(Y, y)
		D = append(D, d)
		C = append(C, c)

		ytmp = append(ytmp, y)
		ytmp = append(ytmp, y+d)
	}
	for i := 0; i < q; i++ {
		a, b := ReadInt2()
		A = append(A, a)
		B = append(B, b)

		ytmp = append(ytmp, b)
	}

	_, to, _ := ZaAtsu1Dim(ytmp, 0)
	for i := 0; i < n; i++ {
		x, y, d, c := X[i], Y[i], D[i], C[i]
		y1 := to[y]
		y2 := to[y+d]

		Q = append(Q, Query{x: x, y1: y1, y2: y2, cost: c, c: 1})
		Q = append(Q, Query{x: x + d, y1: y1, y2: y2, cost: -c, c: 3})
	}
	for i := 0; i < q; i++ {
		a, b := A[i], B[i]
		y := to[b]

		Q = append(Q, Query{x: a, y: y, qid: i, c: 2})
	}

	// x座標でクエリをソート
	sort.SliceStable(Q, func(i, j int) bool {
		// return Q[i].x < Q[j].x
		if Q[i].x < Q[j].x {
			return true
		} else if Q[i].x > Q[j].x {
			return false
		} else {
			return Q[i].c < Q[j].c
		}
	})

	f := func(lv, rv T) T {
		return T(int(lv) + int(rv))
	}
	g := func(to T, from E) T {
		return T(int(to) + int(from))
	}
	h := func(to, from E) E {
		return E(int(to) + int(from))
	}
	p := func(e E, length int) E {
		return E(int(e) * length)
	}
	lst := NewLazySegmentTree(200000+50, f, g, h, p, 0, 0)

	answers := make([]int, q)
	for i := 0; i < len(Q); i++ {
		query := Q[i]
		if query.c == 1 || query.c == 3 {
			// RAQする
			lst.Update(query.y1, query.y2+1, E(query.cost))
		} else {
			// 点クエリする
			cost := int(lst.Query(query.y, query.y+1))
			answers[query.qid] = cost
		}
	}

	for i := 0; i < q; i++ {
		fmt.Println(answers[i])
	}
}

// Assumption: T == E
type T int // (T, f): Monoid
type E int // (E, h): Operator Monoid

type LazySegmentTree struct {
	sz   int
	data []T
	lazy []E
	f    func(lv, rv T) T        // T <> T -> T
	g    func(to T, from E) T    // T <> E -> T (assignment operator)
	h    func(to, from E) E      // E <> E -> E (assignment operator)
	p    func(e E, length int) E // E <> N -> E
	ti   T
	ei   E
}

func NewLazySegmentTree(
	n int,
	f func(lv, rv T) T, g func(to T, from E) T,
	h func(to, from E) E, p func(e E, length int) E,
	ti T, ei E,
) *LazySegmentTree {
	lst := new(LazySegmentTree)
	lst.f, lst.g, lst.h, lst.p = f, g, h, p
	lst.ti, lst.ei = ti, ei

	lst.sz = 1
	for lst.sz < n {
		lst.sz *= 2
	}

	lst.data = make([]T, 2*lst.sz-1)
	lst.lazy = make([]E, 2*lst.sz-1)
	for i := 0; i < 2*lst.sz-1; i++ {
		lst.data[i] = lst.ti
		lst.lazy[i] = lst.ei
	}

	return lst
}

func (lst *LazySegmentTree) Set(k int, x T) {
	lst.data[k+(lst.sz-1)] = x
}

func (lst *LazySegmentTree) Build() {
	for i := lst.sz - 2; i >= 0; i-- {
		lst.data[i] = lst.f(lst.data[2*i+1], lst.data[2*i+2])
	}
}

func (lst *LazySegmentTree) propagate(k, length int) {
	if lst.lazy[k] != lst.ei {
		if k < lst.sz-1 {
			lst.lazy[2*k+1] = lst.h(lst.lazy[2*k+1], lst.lazy[k])
			lst.lazy[2*k+2] = lst.h(lst.lazy[2*k+2], lst.lazy[k])
		}
		lst.data[k] = lst.g(lst.data[k], lst.p(lst.lazy[k], length))
		lst.lazy[k] = lst.ei
	}
}

func (lst *LazySegmentTree) Update(a, b int, x E) T {
	return lst.update(a, b, x, 0, 0, lst.sz)
}

func (lst *LazySegmentTree) update(a, b int, x E, k, l, r int) T {
	lst.propagate(k, r-l)

	if r <= a || b <= l {
		return lst.data[k]
	}

	if a <= l && r <= b {
		lst.lazy[k] = lst.h(lst.lazy[k], x)
		lst.propagate(k, r-l)
		return lst.data[k]
	}

	lv := lst.update(a, b, x, 2*k+1, l, (l+r)/2)
	rv := lst.update(a, b, x, 2*k+2, (l+r)/2, r)
	lst.data[k] = lst.f(lv, rv)
	return lst.data[k]
}

func (lst *LazySegmentTree) Query(a, b int) T {
	return lst.query(a, b, 0, 0, lst.sz)
}

func (lst *LazySegmentTree) query(a, b, k, l, r int) T {
	lst.propagate(k, r-l)

	if r <= a || b <= l {
		return lst.ti
	}

	if a <= l && r <= b {
		return lst.data[k]
	}

	lv := lst.query(a, b, 2*k+1, l, (l+r)/2)
	rv := lst.query(a, b, 2*k+2, (l+r)/2, r)
	return lst.f(lv, rv)
}

func (lst *LazySegmentTree) Get(k int) T {
	return lst.Query(k, k+1)
}

// ZaAtsu1Dim returns 3 values.
// pressed: pressed slice of the original slice
// orgToPress: map for translating original value to pressed value
// pressToOrg: reverse resolution of orgToPress
// O(nlogn)
func ZaAtsu1Dim(org []int, initVal int) (pressed []int, orgToPress, pressToOrg map[int]int) {
	pressed = make([]int, len(org))
	copy(pressed, org)
	sort.Sort(sort.IntSlice(pressed))

	orgToPress = make(map[int]int)
	for i := 0; i < len(org); i++ {
		if i == 0 {
			orgToPress[pressed[0]] = initVal
			continue
		}

		if pressed[i-1] != pressed[i] {
			initVal++
			orgToPress[pressed[i]] = initVal
		}
	}

	for i := 0; i < len(org); i++ {
		pressed[i] = orgToPress[org[i]]
	}

	pressToOrg = make(map[int]int)
	for k, v := range orgToPress {
		pressToOrg[v] = k
	}

	return
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

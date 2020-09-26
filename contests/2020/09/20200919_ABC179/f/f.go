/*
URL:
https://atcoder.jp/contests/abc179/tasks/abc179_f
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
	n, q int
)

func main() {
	defer stdout.Flush()

	n, q = readi2()

	f := func(lv, rv T) T {
		return T(math.Min(float64(lv), float64(rv)))
	}
	g := func(to T, from E) T {
		return T(from)
	}
	h := func(to, from E) E {
		return from
	}
	p := func(e E, length int) E {
		return e
	}
	ti := T(1<<31 - 1)
	ei := E(1<<31 - 1)
	col := NewLazySegmentTree(n, f, g, h, p, ti, ei)
	row := NewLazySegmentTree(n, f, g, h, p, ti, ei)
	for i := 0; i < n; i++ {
		col.Set(i, T(n))
		row.Set(i, T(n))
	}
	col.Build()
	row.Build()

	ans := 0
	for i := 0; i < q; i++ {
		com, x := readi2()
		if com == 1 {
			// colに対して作用
			idx := int(row.Query(x-1, x))
			ans += idx - 2
			col.Update(0, idx, E(x))
		} else {
			// rowに対して作用
			idx := int(col.Query(x-1, x))
			ans += idx - 2
			row.Update(0, idx, E(x))
		}
	}

	fmt.Println((n-2)*(n-2) - ans)
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

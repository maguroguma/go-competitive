/*
URL:
https://yukicoder.me/problems/no/1170
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

/*******************************************************************/

var (
	n, a, b int
	X       []int
)

func main() {
	n, a, b = ReadInt3()
	X = ReadIntSlice(n)

	f := func(lv, rv T) T {
		return lv + rv
	}
	g := func(to T, from E) T {
		return to + T(from)
	}
	h := func(to, from E) E {
		return to + from
	}
	p := func(e E, length int) E {
		return e * E(length)
	}
	ti := T(0)
	ei := E(0)
	lst := NewLazySegmentTree(n, f, g, h, p, ti, ei)
	for i := 0; i < n; i++ {
		lst.Set(i, 1)
	}
	lst.Build()

	// D := make([]int, n)
	for i := 0; i < n; i++ {
		x := X[i]

		T := X[i+1:]
		mini := BinarySearch(len(T), -1, func(mid int) bool {
			return T[mid]-x >= a
		})
		maxi := BinarySearch(-1, len(T), func(mid int) bool {
			return T[mid]-x <= b
		})
		PrintfDebug("mini: %d, maxi: %d\n", mini+i+1, maxi+i+1)

		if maxi < mini {
			// continue
		} else {
			// D[i] = maxi - mini + 1
			// val := lst.Query(i, i+1)
			lst.Update(mini+(i+1), maxi+(i+1)+1, E(1))
		}

		// if i == 0 {
		// 	continue
		// }

		T = X[:i]
		maxi = BinarySearch(-1, len(T), func(mid int) bool {
			return x-T[mid] >= a
		})
		mini = BinarySearch(len(T), -1, func(mid int) bool {
			return x-T[mid] <= b
		})

		if maxi < mini {
		} else {
			lst.Update(mini, maxi+1, E(1))
		}
	}

	ff := func(lv, rv T) T {
		return T(math.Max(float64(lv), float64(rv)))
	}
	gg := func(to T, from E) T {
		return T(from)
	}
	hh := func(to, from E) E {
		return from
	}
	pp := func(e E, length int) E {
		return e
	}
	titi := T(0)
	eiei := E(0)
	lst2 := NewLazySegmentTree(n, ff, gg, hh, pp, titi, eiei)

	for i := 0; i < n; i++ {
		// fmt.Println(D[i] + int(lst.Query(i, i+1)))
		// fmt.Println(int(lst.Query(i, i+1)))
		lst2.Set(i, lst.Query(i, i+1))
	}
	lst2.Build()

	for i := 0; i < n; i++ {
		ans := 0

		x := X[i]

		T := X[i+1:]
		mini := BinarySearch(len(T), -1, func(mid int) bool {
			return T[mid]-x >= a
		})
		maxi := BinarySearch(-1, len(T), func(mid int) bool {
			return T[mid]-x <= b
		})
		PrintfDebug("mini: %d, maxi: %d\n", mini+i+1, maxi+i+1)

		if maxi < mini {
		} else {
			// lst.Update(mini+(i+1), maxi+(i+1)+1, E(1))
			ChMax(&ans, int(lst2.Query(mini+(i+1), maxi+(i+1)+1)))
		}

		T = X[:i]
		maxi = BinarySearch(-1, len(T), func(mid int) bool {
			return x-T[mid] >= a
		})
		mini = BinarySearch(len(T), -1, func(mid int) bool {
			return x-T[mid] <= b
		})

		if maxi < mini {
		} else {
			// lst.Update(mini, maxi+1, E(1))
			ChMax(&ans, int(lst2.Query(mini, maxi+1)))
		}

		ChMax(&ans, int(lst2.Query(i, i+1)))

		fmt.Println(ans)
	}
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

/*******************************************************************/

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

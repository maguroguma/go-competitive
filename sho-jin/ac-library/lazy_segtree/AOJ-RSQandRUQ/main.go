/*
URL:
https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/2/DSL_2_I
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

	v := make([]S, n)
	for i := 0; i < n; i++ {
		v[i] = S{0, 1}
	}
	e := func() S {
		return S{0, 0}
		// なぜか↓でも通る
		// return S{0, 1000}
	}
	merger := func(a, b S) S { return S{a.val + b.val, a.size + b.size} }
	mapper := func(f F, x S) S {
		if f.val == -(1 << uint(60)) {
			return x
		}
		return S{f.val * x.size, x.size}
	}
	composer := func(f, g F) F {
		if f.val == -(1 << uint(60)) {
			return g
		}
		return F{f.val}
	}
	id := func() F { return F{-(1 << uint(60))} }

	lseg := NewLazySegtree(v, e, merger, mapper, composer, id)

	for i := 0; i < q; i++ {
		if readi() == 0 {
			s, t, x := readi3()
			lseg.RangeApply(s, t+1, F{x})
		} else {
			s, t := readi2()
			res := lseg.Prod(s, t+1)
			printf("%d\n", res.val)
		}
	}
}

// originated from: https://qiita.com/EmptyBox_0/items/2f8e3cf7bd44e0f789d5#lazysegtree
// docs: https://atcoder.github.io/ac-library/production/document_ja/lazysegtree.html

// type of monoid
type S struct {
	val  int
	size int
}

// type of mapping
type F struct {
	val int
}

type E func() S
type Merger func(a, b S) S
type Mapper func(f F, x S) S
type Composer func(f, g F) F
type Id func() F
type Compare func(v S) bool

func NewLazySegtree(
	v []S, e E, merger Merger, mapper Mapper, composer Composer, id Id,
) *LazySegtree {
	lseg := new(LazySegtree)
	lseg.n = len(v)
	lseg.log = lseg._ceilPow2(lseg.n)
	lseg.size = 1 << uint(lseg.log)
	lseg.d = make([]S, 2*lseg.size)
	lseg.e = e
	lseg.lz = make([]F, lseg.size)
	lseg.merger = merger
	lseg.mapper = mapper
	lseg.composer = composer
	lseg.id = id
	for i := range lseg.d {
		lseg.d[i] = lseg.e()
	}
	for i := range lseg.lz {
		lseg.lz[i] = lseg.id()
	}
	for i := 0; i < lseg.n; i++ {
		lseg.d[lseg.size+i] = v[i]
	}
	for i := lseg.size - 1; i >= 1; i-- {
		lseg._update(i)
	}
	return lseg
}

type LazySegtree struct {
	n        int
	size     int
	log      int
	d        []S
	lz       []F
	e        E
	merger   Merger
	mapper   Mapper
	composer Composer
	id       Id
}

// Set sets a[p] = x
// Time complexity: O(logn)
func (lseg *LazySegtree) Set(p int, x S) {
	p += lseg.size
	for i := lseg.log; i <= 1; i-- {
		lseg._push(p >> uint(i))
	}
	lseg.d[p] = x
	for i := 1; i <= lseg.log; i++ {
		lseg._update(p >> uint(i))
	}
}

// Get returns a[p]
// Time complexity: O(logn)
func (lseg *LazySegtree) Get(p int) S {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg._push(p >> uint(i))
	}
	return lseg.d[p]
}

// Prod returns op(a[l:r]...)
// Time complexity: O(logn)
func (lseg *LazySegtree) Prod(l, r int) S {
	if l == r {
		return lseg.e()
	}
	l += lseg.size
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		if (l>>uint(i))<<uint(i) != l {
			lseg._push(l >> uint(i))
		}
		if (r>>uint(i))<<uint(i) != r {
			lseg._push(r >> uint(i))
		}
	}
	sml, smr := lseg.e(), lseg.e()
	for l < r {
		if (l & 1) == 1 {
			sml = lseg.merger(sml, lseg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = lseg.merger(lseg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return lseg.merger(sml, smr)
}

// AllProd returns op(a...)
// Time complexity: O(1)
func (lseg *LazySegtree) AllProd() S {
	return lseg.d[1]
}

// Apply sets a[p] = f(a[p])
// Time complexity: O(logn)
func (lseg *LazySegtree) Apply(p int, f F) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg._push(p >> uint(i))
	}
	lseg.d[p] = lseg.mapper(f, lseg.d[p])
	for i := 1; i <= lseg.log; i++ {
		lseg._update(p >> uint(i))
	}
}

// RangeApply sets a[i] = f(a[i]) (i in [l, r-1])
// Time complexity: O(logn)
func (lseg *LazySegtree) RangeApply(l int, r int, f F) {
	if l == r {
		return
	}
	l += lseg.size
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		if (l>>uint(i))<<uint(i) != l {
			lseg._push(l >> uint(i))
		}
		if (r>>uint(i))<<uint(i) != r {
			lseg._push((r - 1) >> uint(i))
		}
	}
	l2, r2 := l, r
	for l < r {
		if l&1 == 1 {
			lseg._allApply(l, f)
			l++
		}
		if r&1 == 1 {
			r--
			lseg._allApply(r, f)
		}
		l >>= 1
		r >>= 1
	}
	l, r = l2, r2
	for i := 1; i <= lseg.log; i++ {
		if (l>>uint(i))<<uint(i) != l {
			lseg._update(l >> uint(i))
		}
		if (r>>uint(i))<<uint(i) != r {
			lseg._update((r - 1) >> uint(i))
		}
	}
}

// Time complexity: O(logn)
func (lseg *LazySegtree) MaxRight(l int, cmp Compare) int {
	if l == lseg.n {
		return lseg.n
	}
	l += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg._push(l >> uint(i))
	}
	sm := lseg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !cmp(lseg.merger(sm, lseg.d[l])) {
			for l < lseg.size {
				lseg._push(l)
				l = 2 * l
				if cmp(lseg.merger(sm, lseg.d[l])) {
					sm = lseg.merger(sm, lseg.d[l])
					l++
				}
			}
			return l - lseg.size
		}
		sm = lseg.merger(sm, lseg.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return lseg.n
}

// Time complexity: O(logn)
func (lseg *LazySegtree) MinLeft(r int, cmp Compare) int {
	if r == 0 {
		return 0
	}
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg._push(r - 1>>uint(i))
	}
	sm := lseg.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !cmp(lseg.merger(lseg.d[r], sm)) {
			for r < lseg.size {
				lseg._push(r)
				r = 2*r + 1
				if cmp(lseg.merger(lseg.d[r], sm)) {
					sm = lseg.merger(lseg.d[r], sm)
					r--
				}
			}
			return r + 1 - lseg.size
		}
		sm = lseg.merger(lseg.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}

func (lseg *LazySegtree) _allApply(k int, f F) {
	lseg.d[k] = lseg.mapper(f, lseg.d[k])
	if k < lseg.size {
		lseg.lz[k] = lseg.composer(f, lseg.lz[k])
	}
}

func (lseg *LazySegtree) _update(k int) {
	lseg.d[k] = lseg.merger(lseg.d[2*k], lseg.d[2*k+1])
}

func (lseg *LazySegtree) _push(k int) {
	lseg._allApply(2*k, lseg.lz[k])
	lseg._allApply(2*k+1, lseg.lz[k])
	lseg.lz[k] = lseg.id()
}

func (lseg *LazySegtree) _ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
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
	EPS     = 1e-10
)

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	reads = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

// mod can calculate a right residual whether value is positive or negative.
func mod(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func min(integers ...int) int {
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

// max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func max(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m < integer {
			m = integer
		}
	}
	return m
}

// chmin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func chmin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

// chmax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func chmax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
}

// sum returns multiple integers sum.
func sum(integers ...int) int {
	var s int
	s = 0

	for _, i := range integers {
		s += i
	}

	return s
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

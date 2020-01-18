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

/*********** I/O ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
)

func init() {
	ReadString = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	// r.Buffer(make([]byte, 1024), int(1e+11)) // for AtCoder
	r.Buffer(make([]byte, 1024), int(1e+9)) // for Codeforces
	// Split sets the split function for the Scanner. The default split function is ScanLines.
	// Split panics if it is called after scanning has started.
	r.Split(bufio.ScanWords)

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

// PrintDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintDebug(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
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

/*
ASCII code

ASCII   10進数  ASCII   10進数  ASCII   10進数
!       33      "       34      #       35
$       36      %       37      &       38
'       39      (       40      )       41
*       42      +       43      ,       44
-       45      .       46      /       47
0       48      1       49      2       50
3       51      4       52      5       53
6       54      7       55      8       56
9       57      :       58      ;       59
<       60      =       61      >       62
?       63      @       64      A       65
B       66      C       67      D       68
E       69      F       70      G       71
H       72      I       73      J       74
K       75      L       76      M       77
N       78      O       79      P       80
Q       81      R       82      S       83
T       84      U       85      V       86
W       87      X       88      Y       89
Z       90      [       91      \       92
]       93      ^       94      _       95
`       96      a       97      b       98
c       99      d       100     e       101
f       102     g       103     h       104
i       105     j       106     k       107
l       108     m       109     n       110
o       111     p       112     q       113
r       114     s       115     t       116
u       117     v       118     w       119
x       120     y       121     z       122
{       123     |       124     }       125
~       126             127
*/

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

var n, q int

func main() {
	n, q = ReadInt2()
	f := func(le, re *Element) *Element {
		// return &Element{v: int(math.Min(float64(le.v), float64(re.v)))}
		return &Element{v: le.v + re.v}
	}
	g := func(to, from *Element) *Element {
		// return &Element{v: from.v}
		return &Element{v: to.v + from.v}
	}
	h := func(to, from *Element) *Element {
		// return &Element{v: from.v}
		return &Element{v: to.v + from.v}
	}
	p := func(e *Element, l int) *Element {
		// return &Element{v: e.v}
		return &Element{v: e.v * l}
	}
	lst := NewLazySegmentTree(n, f, g, h, p, 0, 0)

	for i := 0; i < q; i++ {
		c := ReadInt()
		if c == 0 {
			s, t, x := ReadInt3()
			lst.Update(s-1, t, &Element{v: x})
		} else {
			s, t := ReadInt2()
			e := lst.Query(s-1, t)
			fmt.Println(e.v)
		}
	}
}

// OperatorMonoid = Monoid
type Element struct {
	v int
}

type LazySegmentTree struct {
	sz         int
	data, lazy []*Element
	compose    func(le, re *Element) *Element
	lazyToData func(to, from *Element) *Element
	lazyToLazy func(to, from *Element) *Element
	prop       func(e *Element, l int) *Element
	ident      int
	opIdent    int
}

func NewLazySegmentTree(
	n int,
	f func(le, re *Element) *Element, g func(to, from *Element) *Element,
	h func(to, from *Element) *Element, p func(e *Element, l int) *Element,
	ident int, opIdent int,
) *LazySegmentTree {
	lst := new(LazySegmentTree)
	lst.compose, lst.lazyToData, lst.lazyToLazy, lst.prop = f, g, h, p
	lst.ident, lst.opIdent = ident, opIdent

	lst.sz = 1
	for lst.sz < n {
		lst.sz *= 2
	}

	lst.data = make([]*Element, 2*lst.sz-1)
	lst.lazy = make([]*Element, 2*lst.sz-1)
	for i := 0; i < 2*lst.sz-1; i++ {
		lst.data[i] = &Element{v: lst.ident}
		lst.lazy[i] = &Element{v: lst.opIdent}
	}

	return lst
}

func (lst *LazySegmentTree) Set(k int, x *Element) {
	lst.data[k+(lst.sz-1)] = x
}

func (lst *LazySegmentTree) Build() {
	for i := lst.sz - 2; i >= 0; i-- {
		lst.data[i] = lst.compose(lst.data[2*i+1], lst.data[2*i+2])
	}
}

func (lst *LazySegmentTree) propagate(k, length int) {
	if lst.lazy[k].v != lst.opIdent {
		if k < lst.sz-1 {
			lst.lazy[2*k+1] = lst.lazyToLazy(lst.lazy[2*k+1], lst.lazy[k])
			lst.lazy[2*k+2] = lst.lazyToLazy(lst.lazy[2*k+2], lst.lazy[k])
		}
		lst.data[k] = lst.lazyToData(lst.data[k], lst.prop(lst.lazy[k], length))
		lst.lazy[k] = &Element{v: lst.opIdent}
	}
}

func (lst *LazySegmentTree) Update(a, b int, x *Element) *Element {
	return lst.update(a, b, x, 0, 0, lst.sz)
}

func (lst *LazySegmentTree) update(a, b int, x *Element, k, l, r int) *Element {
	lst.propagate(k, r-l)

	if r <= a || b <= l {
		return lst.data[k]
	}

	if a <= l && r <= b {
		lst.lazy[k] = lst.lazyToLazy(lst.lazy[k], x)
		lst.propagate(k, r-l)
		return lst.data[k]
	}

	le := lst.update(a, b, x, 2*k+1, l, (l+r)/2)
	re := lst.update(a, b, x, 2*k+2, (l+r)/2, r)
	lst.data[k] = lst.compose(le, re)
	return lst.data[k]
}

func (lst *LazySegmentTree) Query(a, b int) *Element {
	return lst.query(a, b, 0, 0, lst.sz)
}

func (lst *LazySegmentTree) query(a, b, k, l, r int) *Element {
	lst.propagate(k, r-l)

	if r <= a || b <= l {
		return &Element{v: lst.ident}
	}

	if a <= l && r <= b {
		return lst.data[k]
	}

	le := lst.query(a, b, 2*k+1, l, (l+r)/2)
	re := lst.query(a, b, 2*k+2, (l+r)/2, r)
	return lst.compose(le, re)
}

func (lst *LazySegmentTree) Get(k int) *Element {
	return lst.Query(k, k+1)
}

/*
- まずは全探索を検討しましょう
- MODは最後にとりましたか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
- いきなりオーバーフローはしていませんか？
- MOD取る系はint64必須ですよ？
*/

/*******************************************************************/

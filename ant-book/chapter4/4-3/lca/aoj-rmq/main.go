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

func main() {
	n := ReadInt()
	for i := 0; i < n; i++ {
		k := ReadInt()
		for j := 0; j < k; j++ {
			c := ReadInt()
			G[i] = append(G[i], c)
			G[c] = append(G[c], i)
		}
	}

	root = 0
	initialize()

	f := func(lv, rv T) T {
		t := T{}
		if lv.v > rv.v {
			t.v = rv.v
			t.idx = rv.idx
		} else {
			t.v = lv.v
			t.idx = lv.idx
		}
		return t
	}
	ti := T{v: 1 << 30, idx: -1}
	st := NewSegmentTree(2*n-1, f, ti)
	for i := 0; i < 2*n-1; i++ {
		st.Set(i, T{v: depth[i], idx: i})
	}
	st.Build()

	q := ReadInt()
	for i := 0; i < q; i++ {
		u, v := ReadInt2()
		l := int(math.Min(float64(id[u]), float64(id[v])))
		r := int(math.Max(float64(id[u]), float64(id[v])))
		t := st.Query(l, r+1)
		fmt.Println(vs[t.idx])
	}
}

const MAX_V = 100000 + 5

var G [MAX_V][]int
var root int

var vs [MAX_V*2 - 1]int    // DFSでの訪問順
var depth [MAX_V*2 - 1]int // 根からの深さ
var id [MAX_V]int          // 各頂点がvsにはじめて登場するインデックス

func dfs(v, p, d int, k *int) {
	id[v] = *k
	vs[*k] = v
	depth[*k] = d
	(*k)++

	for _, to := range G[v] {
		if to != p {
			dfs(to, v, d+1, k)
			vs[*k] = v
			depth[*k] = d
			(*k)++
		}
	}
}

// 初期化
func initialize() {
	// vs, depth, idを初期化する
	k := 0
	dfs(root, -1, 0, &k)
	// RMQを初期化する（最小値ではなく、最小値のインデックスを返すようにする）
	// rmq_init(depth, V*2-1)
}

// u, vのLCAを求める
func lca(u, v int) int {
	// return vs[query(Min(id[u], id[v]), Max(id[u], id[v]) + 1)]
	return -1
}

// type T int // (T, f): Monoid
type T struct {
	v, idx int
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

/*
- まずは全探索を検討しましょう
- MODは最後にとりましたか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
- いきなりオーバーフローはしていませんか？
- MOD取る系はint64必須ですよ？
*/

/*******************************************************************/

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

type Edge struct {
	to, weight int
}

var h, w int
var A [][]int

var G [3000][]Edge

func main() {
	h, w = ReadInt2()
	for i := 0; i < h; i++ {
		row := ReadIntSlice(w)
		A = append(A, row)
	}

	// グリッドをグラフへ
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			id := i*w + j
			if i-1 >= 0 {
				uid := (i-1)*w + j
				G[id] = append(G[id], Edge{to: uid, weight: A[i-1][j]})
				G[uid] = append(G[uid], Edge{to: id, weight: A[i][j]})
			}
			if i+1 < h {
				did := (i+1)*w + j
				G[id] = append(G[id], Edge{to: did, weight: A[i+1][j]})
				G[did] = append(G[did], Edge{to: id, weight: A[i][j]})
			}
			if j-1 >= 0 {
				lid := i*w + (j - 1)
				G[id] = append(G[id], Edge{to: lid, weight: A[i][j-1]})
				G[lid] = append(G[lid], Edge{to: id, weight: A[i][j]})
			}
			if j+1 < w {
				rid := i*w + (j + 1)
				G[id] = append(G[id], Edge{to: rid, weight: A[i][j+1]})
				G[rid] = append(G[rid], Edge{to: id, weight: A[i][j]})
			}
		}
	}

	sid, mid, gid := (h-1)*w+0, (h-1)*w+(w-1), (w - 1)
	smemo, mmemo, gmemo := dijkstra(sid), dijkstra(mid), dijkstra(gid)
	ans := INF_BIT30
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			id := i*w + j
			// dp := dijkstra(id)
			// tmp := dp[sid] + dp[mid] + dp[gid] + A[i][j]
			tmp := smemo[id] + mmemo[id] + gmemo[id] - (2 * A[i][j])
			ChMin(&ans, tmp)
		}
	}

	fmt.Println(ans)
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

func dijkstra(sid int) []int {
	dp, colors := make([]int, 3000), make([]int, 3000)
	for i := 0; i < 3000; i++ {
		dp[i] = INF_BIT30
	}

	temp := make(NodePQ, 0)
	pq := &temp
	heap.Init(pq)

	dp[sid], colors[sid] = 0, GRAY
	heap.Push(pq, &Node{pri: 0, id: sid})
	for pq.Len() > 0 {
		cn := heap.Pop(pq).(*Node)

		if colors[cn.id] == BLACK {
			continue
		}
		colors[cn.id] = BLACK

		for _, e := range G[cn.id] {
			if colors[e.to] == BLACK {
				continue
			}
			if dp[e.to] > dp[cn.id]+e.weight {
				dp[e.to] = dp[cn.id] + e.weight
				heap.Push(pq, &Node{pri: dp[e.to], id: e.to})
				colors[e.to] = GRAY
			}
		}
	}

	return dp
}

type Node struct {
	pri int
	id  int
}
type NodePQ []*Node

func (pq NodePQ) Len() int           { return len(pq) }
func (pq NodePQ) Less(i, j int) bool { return pq[i].pri < pq[j].pri } // <: ASC, >: DESC
func (pq NodePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *NodePQ) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}
func (pq *NodePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// how to use
// temp := make(NodePQ, 0, 100000+1)
// pq := &temp
// heap.Init(pq)
// heap.Push(pq, &Node{pri: intValue})
// popped := heap.Pop(pq).(*Node)

/*
- まずは全探索を検討しましょう
- MODは最後にとりましたか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
- いきなりオーバーフローはしていませんか？
- MOD取る系はint64必須ですよ？
*/

/*******************************************************************/

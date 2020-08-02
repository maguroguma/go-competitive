// https://atcoder.jp/contests/arc056/tasks/arc056_b

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

func init() {
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

var (
	n, m, s int

	G [200000 + 50][]Edge
)

func main() {
	n, m, s = ReadInt3()
	s--
	for i := 0; i < m; i++ {
		u, v := ReadInt2()
		u--
		v--
		G[u] = append(G[u], Edge{to: v, cost: 1})
		G[v] = append(G[v], Edge{to: u, cost: 1})
	}

	dp := Dijkstra(s, n, G[:n])

	for i := 0; i < n; i++ {
		if int(dp[i]) == i {
			fmt.Println(i + 1)
		}
	}
}

func Dijkstra(sid, n int, AG [][]Edge) []V {
	// データをすべて初期化
	dp, colors := InitAll(n)

	// 始点の設定（問題によっては複数始点もありうる）
	pq := InitStartPoint(sid, dp, colors)

	// アルゴリズム本体
	for pq.Len() > 0 {
		pop := pq.pop()
		colors[pop.id] = BLACK
		if Less(dp[pop.id], pop.v) {
			continue
		}

		// 次のノードへの遷移
		for _, e := range AG[pop.id] {
			if colors[e.to] == BLACK {
				continue
			}

			// 値の更新
			nv := GenNextV(pop, e, dp)

			if Less(nv, dp[e.to]) {
				dp[e.to] = nv
				pq.push(&Vertex{id: e.to, v: nv})
				colors[e.to] = GRAY
			}
		}
	}

	return dp
}

// InitAll returns initialized dp and colors slices.
func InitAll(n int) (dp []V, colors []int) {
	dp, colors = make([]V, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = DijkstraVInf()
		colors[i] = WHITE
	}

	return dp, colors
}

// DijkstraVInf returns a infinite value for DP.
func DijkstraVInf() V {
	// return V{num: INF_BIT60, nokori: -1}

	return V(-1)
}

// InitStartPoint returns initialized priority queue, and update dp and colors slices.
// *This function update arguments (side effects).*
func InitStartPoint(sid int, dp []V, colors []int) *VertexPQ {
	pq := NewVertexPQ()

	// for d := 0; d < 4; d++ {
	// 	cid := toid(y1, x1, d)
	// 	pq.push(&Vertex{id: cid, v: V{num: 0, nokori: 0}})

	// 	dp[cid].num, dp[cid].nokori = 0, 0
	// 	colors[cid] = GRAY
	// }

	pq.push(&Vertex{id: sid, v: V(sid)})
	dp[sid] = V(sid)
	colors[sid] = GRAY

	return pq
}

// Less returns whether l is strictly less than r.
// This function is also used by priority queue.
func Less(l, r V) bool {
	// if l.num < r.num {
	// 	return true
	// } else if l.num > r.num {
	// 	return false
	// } else {
	// 	return l.nokori > r.nokori
	// }

	return l > r
}

// GenNextV returns next value considered by transition.
func GenNextV(cv *Vertex, e Edge, dp []V) V {
	// prevd := tod(cv.id)
	// nextd := tod(to)
	// nv := V{num: dp[cv.id].num, nokori: dp[cv.id].nokori}
	// if prevd != nextd {
	// 	// 方向転換
	// 	nv.num++
	// 	nv.nokori = k
	// } else if nv.nokori == 0 {
	// 	// 次へ進むがk-1にする
	// 	nv.num++
	// 	nv.nokori = k - 1
	// } else {
	// 	nv.nokori--
	// }

	_min := func(integers ...int) int {
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

	return V(_min(int(dp[cv.id]), e.to))
}

// DP value type
// type V struct {
// 	num, nokori int
// }
type V int

type Edge struct {
	to   int
	cost int
}

type Vertex struct {
	id int
	v  V
}
type VertexPQ []*Vertex

func NewVertexPQ() *VertexPQ {
	temp := make(VertexPQ, 0)
	pq := &temp
	heap.Init(pq)

	return pq
}
func (pq *VertexPQ) push(target *Vertex) {
	heap.Push(pq, target)
}
func (pq *VertexPQ) pop() *Vertex {
	return heap.Pop(pq).(*Vertex)
}

func (pq VertexPQ) Len() int { return len(pq) }
func (pq VertexPQ) Less(i, j int) bool {
	return Less(pq[i].v, pq[j].v)
}
func (pq VertexPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *VertexPQ) Push(x interface{}) {
	item := x.(*Vertex)
	*pq = append(*pq, item)
}
func (pq *VertexPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// type (
// 	Edge struct {
// 		to   int
// 		cost int
// 	}
// 	Vertex struct {
// 		pri int
// 		id  int
// 	}
// )

// const INF_DIJK = 1 << 60

// // 「各ノードについて「そこに到達するまでに経由するノードの最小ID」を最大化したパスの最小ID」を計算する
// func dijkstra(sid, n int, AG [][]Edge) ([]int, []int) {
// 	dp := make([]int, n)
// 	colors, parents := make([]int, n), make([]int, n)
// 	for i := 0; i < n; i++ {
// 		dp[i] = -1
// 		colors[i], parents[i] = WHITE, -1
// 	}

// 	temp := make(VertexPQ, 0, 100000+5)
// 	pq := &temp
// 	heap.Init(pq)
// 	heap.Push(pq, &Vertex{pri: sid, id: sid}) // priorityは単調減少させていく
// 	dp[sid] = sid
// 	colors[sid] = GRAY

// 	for pq.Len() > 0 {
// 		pop := heap.Pop(pq).(*Vertex)

// 		colors[pop.id] = BLACK

// 		// if pop.pri > dp[pop.id] {
// 		if pop.pri < dp[pop.id] {
// 			continue
// 		}

// 		for _, e := range AG[pop.id] {
// 			if colors[e.to] == BLACK {
// 				continue
// 			}

// 			// if dp[e.to] > dp[pop.id]+e.cost {
// 			if dp[e.to] < Min(dp[pop.id], e.to) {
// 				// dp[e.to] = dp[pop.id] + e.cost
// 				dp[e.to] = Min(dp[pop.id], e.to)
// 				heap.Push(pq, &Vertex{pri: dp[e.to], id: e.to})
// 				colors[e.to], parents[e.to] = GRAY, pop.id
// 			}
// 		}
// 	}

// 	return dp, parents
// }

// type VertexPQ []*Vertex

// func (pq VertexPQ) Len() int           { return len(pq) }
// func (pq VertexPQ) Less(i, j int) bool { return pq[i].pri > pq[j].pri } // <: ASC, >: DESC
// func (pq VertexPQ) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }
// func (pq *VertexPQ) Push(x interface{}) {
// 	item := x.(*Vertex)
// 	*pq = append(*pq, item)
// }
// func (pq *VertexPQ) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	*pq = old[0 : n-1]
// 	return item
// }

// // Min returns the min integer among input set.
// // This function needs at least 1 argument (no argument causes panic).
// func Min(integers ...int) int {
// 	m := integers[0]
// 	for i, integer := range integers {
// 		if i == 0 {
// 			continue
// 		}
// 		if m > integer {
// 			m = integer
// 		}
// 	}
// 	return m
// }

// how to use
// temp := make(VertexPQ, 0, 100000+1)
// pq := &temp
// heap.Init(pq)
// heap.Push(pq, &Vertex{pri: intValue})
// popped := heap.Pop(pq).(*Vertex)

/*
- まずは全探索を検討しましょう
- MODは最後にとりましたか？
- 負のMODはちゃんと関数を使って処理していますか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
- いきなりオーバーフローはしていませんか？
- MOD取る系はint64必須ですよ？
- 後ろ・逆・ゴールから考えましたか？
- 3者のうち真ん中に着目しましたか？
*/

/*******************************************************************/

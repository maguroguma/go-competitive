/*
URL:
https://atcoder.jp/contests/joi2017yo/tasks/joi2017yo_f
*/

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var (
	n, m, x int32
	F       []byte
	A, B    []int16
	D       []uint8

	// G [4500000 + 50][]Edge
	G [][]Edge
)

const (
	COLD, OK, HOT = 0, 1, 2
	CC, HC        = 0, 1
)

func main() {
	defer stdout.Flush()

	nn, mm, xx := readi3()
	n, m, x = int32(nn), int32(mm), int32(xx)
	tmp := readis(int(n))
	F = make([]byte, n)
	for i := 0; i < int(n); i++ {
		F[i] = byte(tmp[i])
	}
	A, B, D = make([]int16, m), make([]int16, m), make([]uint8, m)
	for i := 0; i < int(m); i++ {
		aa, bb, dd := readi3()
		a, b, d := int16(aa), int16(bb), uint8(dd)
		a--
		b--

		A[i], B[i], D[i] = a, b, d
	}

	N := int32(2 * (x + 1) * n)
	G = make([][]Edge, N)

	toid := func(cf, cx, id int32) int32 {
		return (cf*(x+1)+cx)*n + id
	}

	for i := 0; i < int(m); i++ {
		a, b, d := int32(A[i]), int32(B[i]), int32(D[i])
		fa, fb := F[a], F[b]

		if fa == OK && fb == OK {
			for f := int32(0); f <= 1; f++ {
				for y := int32(0); y <= x; y++ {
					nx := int32(min(y+d, x))

					// a -> b
					from := toid(f, y, a)
					to := toid(f, nx, b)
					G[from] = append(G[from], Edge{to: to, w: Weight(d)})
					// b -> a
					from = toid(f, y, b)
					to = toid(f, nx, a)
					G[from] = append(G[from], Edge{to: to, w: Weight(d)})
				}
			}

			continue
		}

		if fa == OK {
			var FB int32
			if fb == COLD {
				FB = CC
			} else if fb == HOT {
				FB = HC
			}
			for y := int32(0); y <= x; y++ {
				nx := min(y+d, x)

				// a -> b
				for f := int32(0); f <= 1; f++ {
					if f != FB && nx >= x {
						from := toid(f, y, a)
						to := toid(FB, 0, b)
						G[from] = append(G[from], Edge{to: to, w: Weight(d)})
					} else if f == FB {
						from := toid(FB, y, a)
						to := toid(FB, 0, b)
						G[from] = append(G[from], Edge{to: to, w: Weight(d)})
					}
				}
				// b -> a: 制約はない
				from := toid(FB, 0, b)
				to := toid(FB, min(d, x), a)
				G[from] = append(G[from], Edge{to: to, w: Weight(d)})
			}

			continue
		}

		if fb == OK {
			var FA int32
			if fa == COLD {
				FA = CC
			} else if fa == HOT {
				FA = HC
			}
			for y := int32(0); y <= x; y++ {
				nx := min(y+d, x)

				// a -> b: 制約はない
				from := toid(FA, 0, a)
				to := toid(FA, min(d, x), b)
				G[from] = append(G[from], Edge{to: to, w: Weight(d)})
				// b -> a
				for f := int32(0); f <= 1; f++ {
					if f != FA && nx >= x {
						from = toid(f, y, b)
						to = toid(FA, 0, a)
						G[from] = append(G[from], Edge{to: to, w: Weight(d)})
					} else if f == FA {
						from = toid(FA, y, b)
						to = toid(FA, 0, a)
						G[from] = append(G[from], Edge{to: to, w: Weight(d)})
					}
				}
			}

			continue
		}

		if fa == fb {
			var FF int32
			if fa == COLD {
				FF = CC
			} else {
				FF = HC
			}

			// a -> b
			from := toid(FF, 0, a)
			to := toid(FF, 0, b)
			G[from] = append(G[from], Edge{to: to, w: Weight(d)})

			// b -> a
			from = toid(FF, 0, b)
			to = toid(FF, 0, a)
			G[from] = append(G[from], Edge{to: to, w: Weight(d)})

			continue
		}

		if d >= x {
			var aid, bid int32

			if fa == COLD {
				aid = toid(CC, 0, a)
			} else if fa == HOT {
				aid = toid(HC, 0, a)
			}

			if fb == COLD {
				bid = toid(CC, 0, b)
			} else if fb == HOT {
				bid = toid(HC, 0, b)
			}

			G[aid] = append(G[aid], Edge{to: bid, w: Weight(d)})
			G[bid] = append(G[bid], Edge{to: aid, w: Weight(d)})
		}
	}

	vinf := Value(math.MaxUint16)
	less := func(l, r Value) bool { return l < r }
	estimate := func(cid int32, cv Value, e Edge) Value {
		return Value(int32(cv) + int32(e.w))
	}
	SS := []StartPoint{{id: toid(CC, 0, 0), vzero: Value(0)}}

	ds := NewDijkstraSolver(vinf, less, estimate)
	dp := ds.Dijkstra(SS, N, G[:N])

	ans := int32(INF_B30)
	for f := int32(0); f <= 1; f++ {
		for y := int32(0); y <= x; y++ {
			chmin(&ans, int32(dp[toid(f, y, n-1)]))
		}
	}

	fmt.Println(ans)
}

// type Value and Weight should be modified according to problems.

// DP value type
// type Value struct {
// 	v int32
// }
type Value uint16

// weight of edge
// type Weight struct {
// 	v int32
// }
type Weight uint8

// edge of graph
type Edge struct {
	to int32
	w  Weight
}

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	id    int32
	vzero Value
}

// Less returns l < r, and shared with pq.
type Less func(l, r Value) bool

// Estimate returns next value considered by transition.
type Estimate func(cid int32, cv Value, e Edge) Value

func NewDijkstraSolver(vinf Value, less Less, estimate Estimate) *DijkstraSolver {
	ds := new(DijkstraSolver)

	// shared with priority queue
	__less = less

	ds.vinf, ds.less, ds.estimate = vinf, less, estimate

	return ds
}

// verified by [ABC143-E](https://atcoder.jp/contests/abc143/tasks/abc143_e)
func (ds *DijkstraSolver) Dijkstra(S []StartPoint, n int32, AG [][]Edge) []Value {
	// initialize data
	dp, colors := ds._initAll(n)

	// configure about start points (some problems have multi start points)
	pq := ds._initStartPoint(S, dp, colors)

	// body of dijkstra algorithm
	for pq.Len() > 0 {
		pop := pq.pop()
		colors[pop.id] = BLACK
		if ds.less(dp[pop.id], pop.v) {
			continue
		}

		// to next node
		for _, e := range AG[pop.id] {
			if colors[e.to] == BLACK {
				continue
			}

			// update optimal value of the next node
			nv := ds.estimate(pop.id, pop.v, e)

			if ds.less(nv, dp[e.to]) {
				dp[e.to] = nv
				pq.push(&Vertex{id: e.to, v: nv})
				colors[e.to] = GRAY
			}
		}
	}

	return dp
}

type DijkstraSolver struct {
	vinf     Value
	less     Less
	estimate Estimate
}

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

// _initAll returns initialized dp and colors slices.
func (ds *DijkstraSolver) _initAll(n int32) (dp []Value, colors []byte) {
	dp, colors = make([]Value, n), make([]byte, n)
	for i := int32(0); i < n; i++ {
		dp[i] = ds.vinf
		colors[i] = WHITE
	}

	return dp, colors
}

// _initStartPoint returns initialized priority queue, and update dp and colors slices.
// *This function update arguments (side effects).*
func (ds *DijkstraSolver) _initStartPoint(S []StartPoint, dp []Value, colors []byte) *VertexPQ {
	pq := NewVertexPQ()

	for _, sp := range S {
		pq.push(&Vertex{id: sp.id, v: sp.vzero})
		dp[sp.id] = sp.vzero
		colors[sp.id] = GRAY
	}

	return pq
}

// Less function is shared with a priority queue.
var __less Less

// Definitions of a priority queue
type Vertex struct {
	id int32
	v  Value
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
	return __less(pq[i].v, pq[j].v)
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
func min(integers ...int32) int32 {
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
func chmin(updatedValue *int32, target int32) bool {
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

// abs is integer version of math.Abs
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

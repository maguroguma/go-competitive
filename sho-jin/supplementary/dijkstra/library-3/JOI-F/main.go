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
	n, m, x int
	F       []int
	A, B, D []int

	G [6000000][]Edge
)

const (
	COLD, OK, HOT = 0, 1, 2
	CC, HC        = 0, 1
)

func toid(cf, cx, id int) int {
	return (cf*(x+1)+cx)*n + id
}
func tof(nid int) int {
	return nid / n / (x + 1)
}
func tod(nid int) int {
	return (nid / n) % (x + 1)
}
func toi(nid int) int {
	return nid % n
}

func main() {
	defer stdout.Flush()

	n, m, x = readi3()
	F = readis(n)
	A, B, D = make([]int, m), make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		a, b, d := readi3()
		a--
		b--

		A[i], B[i], D[i] = a, b, d
	}

	for i := 0; i < m; i++ {
		a, b, d := A[i], B[i], D[i]

		G[a] = append(G[a], Edge{to: b, w: Weight{d}})
		G[b] = append(G[b], Edge{to: a, w: Weight{d}})
	}

	vinf := Value{INF_B60}
	less := func(l, r Value) bool { return l.v < r.v }
	transit := func(cv *Vertex, AG [][]Edge) []*Vertex {
		res := []*Vertex{}

		id := toi(cv.vid)
		cd := tod(cv.vid)
		cf := tof(cv.vid)
		for _, e := range AG[id] {
			nid := e.to
			nf := F[e.to]
			var ff int
			if nf == OK {
				ff = cf
			} else if nf == COLD {
				ff = CC
			} else {
				ff = HC
			}

			var nx int

			if nf == OK {
				nx = min(cd+e.w.v, x)
			} else {
				if cf != ff {
					dd := min(cd+e.w.v, x)
					if dd < x {
						continue
					}
				}
				nx = 0
			}

			next := toid(ff, nx, nid)

			nval := Value{cv.v.v + e.w.v}
			nv := &Vertex{next, nval}

			res = append(res, nv)
		}

		return res
	}
	SS := []StartPoint{{vid: toid(CC, 0, 0), vzero: Value{0}}}
	N := 2 * (x + 1) * n

	ds := NewDijkstraSolver(vinf, less, transit)
	dp := ds.Dijkstra(SS, N, G[:N])

	ans := INF_B60
	for f := 0; f <= 1; f++ {
		for y := 0; y <= x; y++ {
			chmin(&ans, dp[toid(f, y, n-1)].v)
		}
	}

	fmt.Println(ans)
}

// type Value and Weight should be modified according to problems.

// DP value type
type Value struct {
	v int
}

// weight of edge
type Weight struct {
	v int
}

// edge of graph
type Edge struct {
	to int
	w  Weight
}

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	vid   int
	vzero Value
}

// Less returns l < r, and shared with pq.
type Less func(l, r Value) bool

// Transit calculates all possible transition.
type Transit func(cv *Vertex, AG [][]Edge) []*Vertex

// func NewDijkstraSolver(vinf Value, less Less, estimate Estimate) *DijkstraSolver {
func NewDijkstraSolver(vinf Value, less Less, transit Transit) *DijkstraSolver {
	ds := new(DijkstraSolver)

	// shared with priority queue
	__less = less

	ds.vinf, ds.less, ds.transit = vinf, less, transit

	return ds
}

// verified by [ABC143-E](https://atcoder.jp/contests/abc143/tasks/abc143_e)
func (ds *DijkstraSolver) Dijkstra(S []StartPoint, n int, AG [][]Edge) []Value {
	// initialize data
	dp, colors := ds._initAll(n)

	// configure about start points (some problems have multi start points)
	pq := ds._initStartPoint(S, dp, colors)

	// body of dijkstra algorithm
	for pq.Len() > 0 {
		pop := pq.pop()
		colors[pop.vid] = BLACK
		if ds.less(dp[pop.vid], pop.v) {
			continue
		}

		// to next nodes
		estimates := ds.transit(pop, AG)
		for _, es := range estimates {
			if colors[es.vid] == BLACK {
				continue
			}

			if ds.less(es.v, dp[es.vid]) {
				dp[es.vid] = es.v
				pq.push(es)
				colors[es.vid] = GRAY
			}
		}
	}

	return dp
}

type DijkstraSolver struct {
	vinf    Value
	less    Less
	transit Transit
}

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

// _initAll returns initialized dp and colors slices.
func (ds *DijkstraSolver) _initAll(n int) (dp []Value, colors []int) {
	dp, colors = make([]Value, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = ds.vinf
		colors[i] = WHITE
	}

	return dp, colors
}

// _initStartPoint returns initialized priority queue, and update dp and colors slices.
// *This function update arguments (side effects).*
func (ds *DijkstraSolver) _initStartPoint(S []StartPoint, dp []Value, colors []int) *VertexPQ {
	pq := NewVertexPQ()

	for _, sp := range S {
		pq.push(&Vertex{vid: sp.vid, v: sp.vzero})
		dp[sp.vid] = sp.vzero
		colors[sp.vid] = GRAY
	}

	return pq
}

// Less function is shared with a priority queue.
var __less Less

// Definitions of a priority queue
type Vertex struct {
	vid int
	v   Value
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

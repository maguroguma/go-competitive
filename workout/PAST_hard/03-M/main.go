/*
URL:
https://atcoder.jp/contests/past202005-open/tasks/past202005_m
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
	n, m int
	s    int
	k    int
	T    []int

	G  [100000 + 5][]Edge
	A  [20][20]int
	dp [1<<16 + 5][20]int
)

func main() {
	defer stdout.Flush()

	n, m = readi2()
	for i := 0; i < m; i++ {
		u, v := readi2()
		u--
		v--

		G[u] = append(G[u], Edge{to: v, ew: EdgeWeight{cost: 1}})
		G[v] = append(G[v], Edge{to: u, ew: EdgeWeight{cost: 1}})
	}

	s = readi() - 1
	k = readi()
	T = readis(k)
	for i := 0; i < k; i++ {
		T[i]--
	}

	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			A[i][j] = INF_B60
		}
	}

	for S := 0; S < 1<<uint(k); S++ {
		for i := 0; i < k; i++ {
			dp[S][i] = INF_B60
		}
	}

	vinf := V{v: INF_B60}
	less := func(l, r V) bool { return l.v < r.v }
	nextv := func(cv *Vertex, e Edge) V {
		return V{v: cv.v.v + e.ew.cost}
	}
	ds := NewDijkstraSolver(vinf, less, nextv)
	for i, tid := range T {
		dist := ds.Dijkstra([]StartPoint{{id: tid, vzero: V{v: 0}}}, n, G[:n])
		for j := 0; j < k; j++ {
			nid := T[j]
			A[i][j] = dist[nid].v
		}
	}

	dist := ds.Dijkstra([]StartPoint{{id: s, vzero: V{v: 0}}}, n, G[:n])
	for i := 0; i < k; i++ {
		nid := T[i]
		dp[OnBit(0, i)][i] = dist[nid].v
	}

	for S := 1; S < 1<<uint(k); S++ {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if NthBit(S, j) == 1 {
					continue
				}
				ChMin(&dp[OnBit(S, j)][j], dp[S][i]+A[i][j])
			}
		}
	}

	ans := INF_B60
	for i := 0; i < k; i++ {
		ChMin(&ans, dp[1<<uint(k)-1][i])
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

// NthBit returns nth bit value of an argument.
// n starts from 0.
func NthBit(num int, nth int) int {
	return num >> uint(nth) & 1
}

// OnBit returns the integer that has nth ON bit.
// If an argument has nth ON bit, OnBit returns the argument.
func OnBit(num int, nth int) int {
	return num | (1 << uint(nth))
}

// OffBit returns the integer that has nth OFF bit.
// If an argument has nth OFF bit, OffBit returns the argument.
func OffBit(num int, nth int) int {
	return num & ^(1 << uint(nth))
}

// PopCount returns the number of ON bit of an argument.
func PopCount(num int, ub int) int {
	res := 0

	for i := 0; i < ub; i++ {
		if ((num >> uint(i)) & 1) == 1 {
			res++
		}
	}

	return res
}

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

// DP value type
type V struct {
	// {{
	v int
	// }}
}

// weight of edge
type EdgeWeight struct {
	// {{
	cost int
	// }}
}

// edge of graph
type Edge struct {
	to int
	ew EdgeWeight
}

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	id    int
	vzero V
}

type DijkstraSolver struct {
	vinf  V
	Less  func(l, r V) bool          // Less returns l < r, and shared with pq.
	NextV func(cv *Vertex, e Edge) V // NextV returns next value considered by transition.
}

func NewDijkstraSolver(
	vinf V, Less func(l, r V) bool, NextV func(cv *Vertex, e Edge) V,
) *DijkstraSolver {
	ds := new(DijkstraSolver)

	// shared with priority queue
	__less = Less

	ds.vinf, ds.Less, ds.NextV = vinf, Less, NextV

	return ds
}

// verified by [ABC143-E](https://atcoder.jp/contests/abc143/tasks/abc143_e)
func (ds *DijkstraSolver) Dijkstra(S []StartPoint, n int, AG [][]Edge) []V {
	// initialize data
	dp, colors := ds.initAll(n)

	// configure about start points (some problems have multi start points)
	pq := ds.initStartPoint(S, dp, colors)

	// body of dijkstra algorithm
	for pq.Len() > 0 {
		pop := pq.pop()
		colors[pop.id] = BLACK
		if ds.Less(dp[pop.id], pop.v) {
			continue
		}

		// to next node
		for _, e := range AG[pop.id] {
			if colors[e.to] == BLACK {
				continue
			}

			// update optimal value of the next node
			nv := ds.NextV(pop, e)

			if ds.Less(nv, dp[e.to]) {
				dp[e.to] = nv
				pq.push(&Vertex{id: e.to, v: nv})
				colors[e.to] = GRAY
			}
		}
	}

	return dp
}

// initAll returns initialized dp and colors slices.
func (ds *DijkstraSolver) initAll(n int) (dp []V, colors []int) {
	dp, colors = make([]V, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = ds.vinf
		colors[i] = WHITE
	}

	return dp, colors
}

// initStartPoint returns initialized priority queue, and update dp and colors slices.
// *This function update arguments (side effects).*
func (ds *DijkstraSolver) initStartPoint(S []StartPoint, dp []V, colors []int) *VertexPQ {
	pq := NewVertexPQ()

	for _, sp := range S {
		pq.push(&Vertex{id: sp.id, v: sp.vzero})
		dp[sp.id] = sp.vzero
		colors[sp.id] = GRAY
	}

	return pq
}

// Less function is shared with a priority queue.
var __less func(l, r V) bool

// Definitions of a priority queue
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

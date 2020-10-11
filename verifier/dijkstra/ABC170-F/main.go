/*
URL:
https://atcoder.jp/contests/abc170/tasks/abc170_f
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
	h, w, k        int
	x1, y1, x2, y2 int
	C              [][]rune

	steps [4][2]int
	N     int
	G     [4000000 + 50][]Edge
)

func toid(y, x, d int) int {
	return (y*w+x)*4 + d
}
func toy(id int) int {
	// return id / w / 4
	return id / (w * 4)
}
func tox(id int) int {
	return id / 4 % w
}
func tod(id int) int {
	return id % 4
}

func main() {
	h, w, k = readi3()
	y1, x1, y2, x2 = readi4()
	y1--
	x1--
	y2--
	x2--
	for i := 0; i < h; i++ {
		row := readrs()
		C = append(C, row)
	}

	steps = [4][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	N = h * w * 4
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for d := 0; d < 4; d++ {
				cid := toid(i, j, d)

				dy, dx := steps[d][0], steps[d][1]
				ny, nx := i+dy, j+dx
				if 0 <= ny && ny < h && 0 <= nx && nx < w && C[ny][nx] == '.' {
					nid := toid(ny, nx, d)
					w := Weight{cost: 1}
					G[cid] = append(G[cid], Edge{to: nid, w: w})
				}

				// 同じグリッドで異なる方向を向く
				for nd := 0; nd < 4; nd++ {
					if d == nd {
						continue
					}
					nid := toid(i, j, nd)
					w := Weight{cost: 1}
					G[cid] = append(G[cid], Edge{to: nid, w: w})
				}
			}
		}
	}

	vinf := Value{num: INF_B60, nokori: -1}
	vinit := Value{num: 0, nokori: 0}
	less := func(l, r Value) bool {
		if l.num < r.num {
			return true
		} else if l.num > r.num {
			return false
		} else {
			return l.nokori > r.nokori
		}
	}
	estimate := func(cid int, cv Value, e Edge) Value {
		prevd := tod(cid)
		nextd := tod(e.to)
		nv := Value{num: cv.num, nokori: cv.nokori}
		if prevd != nextd {
			// 方向転換
			nv.num++
			nv.nokori = k
		} else if nv.nokori == 0 {
			// 次へ進むがk-1にする
			nv.num++
			nv.nokori = k - 1
		} else {
			nv.nokori--
		}

		return nv
	}
	ds := NewDijkstraSolver(vinf, less, estimate)
	S := []StartPoint{}
	for d := 0; d < 4; d++ {
		S = append(S, StartPoint{id: toid(y1, x1, d), vzero: vinit})
	}
	dp := ds.Dijkstra(S, N, G[:N])

	ans := INF_B60
	for d := 0; d < 4; d++ {
		id := toid(y2, x2, d)
		chmin(&ans, dp[id].num)
	}

	if ans >= INF_B60 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

// type Value and Weight should be modified according to problems.

// DP value type
type Value struct {
	num, nokori int
}

// weight of edge
type Weight struct {
	cost int
}

// edge of graph
type Edge struct {
	to int
	w  Weight
}

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	id    int
	vzero Value
}

// Less returns l < r, and shared with pq.
type Less func(l, r Value) bool

// Estimate returns next value considered by transition.
type Estimate func(cid int, cv Value, e Edge) Value

func NewDijkstraSolver(vinf Value, less Less, estimate Estimate) *DijkstraSolver {
	ds := new(DijkstraSolver)

	// shared with priority queue
	__less = less

	ds.vinf, ds.less, ds.estimate = vinf, less, estimate

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
	id int
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

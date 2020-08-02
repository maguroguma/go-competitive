/*
URL:
https://atcoder.jp/contests/abc170/tasks/abc170_f
*/

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

/*******************************************************************/

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

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
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

var (
	h, w, k        int
	x1, y1, x2, y2 int
	C              [][]rune

	steps [4][2]int
	N     int
	G     [4000000 + 50][]int
)

func toid(y, x, d int) int {
	return (y*w+x)*4 + d
}
func toy(id int) int {
	return id / w / 4
}
func tox(id int) int {
	return id / 4 % w
}
func tod(id int) int {
	return id % 4
}

func main() {
	h, w, k = ReadInt3()
	y1, x1, y2, x2 = ReadInt4()
	y1--
	x1--
	y2--
	x2--
	for i := 0; i < h; i++ {
		row := ReadRuneSlice()
		C = append(C, row)
	}

	steps = [4][2]int{
		[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0},
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
					G[cid] = append(G[cid], nid)
				}

				// 同じグリッドで異なる方向を向く
				for nd := 0; nd < 4; nd++ {
					if d == nd {
						continue
					}
					nid := toid(i, j, nd)
					G[cid] = append(G[cid], nid)
				}
			}
		}
	}

	dp := Dijkstra(N, G[:N])

	ans := INF_BIT60
	for d := 0; d < 4; d++ {
		id := toid(y2, x2, d)
		ChMin(&ans, dp[id].num)
	}

	if ans >= INF_BIT60 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func Dijkstra(n int, AG [][]int) []V {
	// データをすべて初期化
	dp, colors := InitAll(n)

	// 始点の設定（問題によっては複数始点もありうる）
	pq := InitStartPoint(dp, colors)

	// アルゴリズム本体
	for pq.Len() > 0 {
		pop := pq.pop()
		colors[pop.id] = BLACK
		if Less(dp[pop.id], pop.v) {
			continue
		}

		// 次のノードへの遷移
		for _, to := range AG[pop.id] {
			if colors[to] == BLACK {
				continue
			}

			// 値の更新
			nv := GenNextV(pop, to, dp)

			if Less(nv, dp[to]) {
				dp[to] = nv
				pq.push(&Vertex{id: to, v: nv})
				colors[to] = GRAY
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
	return V{num: INF_BIT60, nokori: -1}
}

// InitStartPoint returns initialized priority queue, and update dp and colors slices.
// *This function update arguments (side effects).*
func InitStartPoint(dp []V, colors []int) *VertexPQ {
	pq := NewVertexPQ()
	for d := 0; d < 4; d++ {
		cid := toid(y1, x1, d)
		pq.push(&Vertex{id: cid, v: V{num: 0, nokori: 0}})

		dp[cid].num, dp[cid].nokori = 0, 0
		colors[cid] = GRAY
	}

	return pq
}

// Less returns whether l is strictly less than r.
// This function is also used by priority queue.
func Less(l, r V) bool {
	if l.num < r.num {
		return true
	} else if l.num > r.num {
		return false
	} else {
		return l.nokori > r.nokori
	}
}

// GenNextV returns next value considered by transition.
func GenNextV(cv *Vertex, to int, dp []V) V {
	prevd := tod(cv.id)
	nextd := tod(to)
	nv := V{num: dp[cv.id].num, nokori: dp[cv.id].nokori}
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

// DP value type
type V struct {
	num, nokori int
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

/*******************************************************************/

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

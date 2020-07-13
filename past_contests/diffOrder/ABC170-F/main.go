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

var (
	h, w, k        int
	x1, y1, x2, y2 int
	C              [][]rune

	steps [4][2]int
	N     int
	G     [4000000 + 50][]int

	// dp [][][]*Vertex
	dp []*Vertex
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

				// 同じ方向のまま隣のグリッドへ移動
				// for _, s := range steps {
				// 	dy, dx := s[0], s[1]
				// 	ny, nx := i+dy, j+dx
				// 	if 0 <= ny && ny < h && 0 <= nx && nx < w && C[ny][nx] == '.' {
				// 		nid := toid(ny, nx, d)
				// 		G[cid] = append(G[cid], nid)
				// 	}
				// }

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

	dp := make([]*Vertex, N)
	colors := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = &Vertex{id: i, num: INF_BIT60, nokori: -1}
		colors[i] = WHITE
	}

	pq := NewVertexPQ()
	for d := 0; d < 4; d++ {
		cid := toid(y1, x1, d)
		pq.push(&Vertex{id: cid, num: 0, nokori: 0})

		dp[cid].num, dp[cid].nokori = 0, 0
		colors[cid] = GRAY
	}

	for pq.Len() > 0 {
		pop := pq.pop()
		colors[pop.id] = BLACK
		if Less(dp[pop.id], pop) {
			continue
		}

		prevd := tod(pop.id)

		for _, to := range G[pop.id] {
			if colors[to] == BLACK {
				continue
			}

			nextd := tod(to)
			nv := &Vertex{id: to, num: dp[pop.id].num, nokori: dp[pop.id].nokori}
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

			if Less(nv, dp[to]) {
				// dp[to] = nv
				dp[to].num, dp[to].nokori = nv.num, nv.nokori
				pq.push(nv)
				colors[to] = GRAY
			}
		}
	}

	// for i := 0; i < N; i++ {
	// 	PrintfDebug("(y, x, d) = (%d, %d, %d): %v\n", toy(i), tox(i), tod(i), dp[i])
	// }

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

// Kiriage returns Ceil(a/b)
// a >= 0, b > 0
func Kiriage(a, b int) int {
	return (a + (b - 1)) / b
}

// func Large(l, r *Vertex) bool {
// 	if l.num > r.num {
// 		return true
// 	} else if l.num < r.num {
// 		return false
// 	} else {
// 		return l.nokori < r.nokori
// 	}
// }

func Less(l, r *Vertex) bool {
	if l.num < r.num {
		return true
	} else if l.num > r.num {
		return false
	} else {
		return l.nokori > r.nokori
		// return l.nokori >= r.nokori
	}
}

type Vertex struct {
	// pri int
	// y, x        int
	// dir         int
	id          int
	num, nokori int
}
type VertexPQ []*Vertex

// Interfaces
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
	// return pq[i].pri < pq[j].pri
	if pq[i].num < pq[j].num {
		return true
	} else if pq[i].num > pq[j].num {
		return false
	} else {
		return pq[i].nokori > pq[j].nokori
	}
} // <: ASC, >: DESC
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

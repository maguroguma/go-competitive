/*
URL:
https://atcoder.jp/contests/abc176/tasks/abc176_d
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
	h, w           int
	sy, sx, gy, gx int
	S              [][]rune

	G [][]Edge
)

func main() {
	h, w = ReadInt2()
	sy, sx = ReadInt2()
	gy, gx = ReadInt2()
	sy--
	sx--
	gy--
	gx--
	for i := 0; i < h; i++ {
		row := ReadRuneSlice()
		S = append(S, row)
	}

	steps := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	toid := func(i, j int) int { return w*i + j }
	// toy := func(id int) int { return id / w }
	// tox := func(id int) int { return id % w }
	worps := [][2]int{}
	for i := -2; i <= 2; i++ {
		for j := -2; j <= 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			worps = append(worps, [2]int{i, j})
		}
	}

	N := h * w
	G = make([][]Edge, N)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if S[i][j] == '#' {
				continue
			}

			cid := toid(i, j)

			for _, s := range steps {
				dy, dx := s[0], s[1]
				ny, nx := i+dy, j+dx
				if 0 <= ny && ny < h && 0 <= nx && nx < w && S[ny][nx] == '.' {
					nid := toid(ny, nx)

					G[cid] = append(G[cid], Edge{to: nid, cost: 0})
				}
			}

			for _, s := range worps {
				dy, dx := s[0], s[1]
				ny, nx := i+dy, j+dx
				if 0 <= ny && ny < h && 0 <= nx && nx < w && S[ny][nx] == '.' {
					nid := toid(ny, nx)

					G[cid] = append(G[cid], Edge{to: nid, cost: 1})
				}
			}
		}
	}

	sid := toid(sy, sx)
	gid := toid(gy, gx)
	dp, _ := dijkstra(sid, N, G[:N])
	if dp[gid] >= INF_DIJK {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[gid])
	}
}

type (
	Edge struct {
		to   int
		cost int
	}
	Vertex struct {
		pri int
		id  int
	}
)

const INF_DIJK = 1 << 60

func dijkstra(sid, n int, AG [][]Edge) (dp, parents []int) {
	dp = make([]int, n)
	parents = make([]int, n)
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = INF_DIJK
		colors[i], parents[i] = WHITE, -1
	}

	temp := make(VertexPQ, 0, 100000+5)
	pq := &temp
	heap.Init(pq)
	heap.Push(pq, &Vertex{pri: 0, id: sid})
	dp[sid] = 0
	colors[sid] = GRAY

	for pq.Len() > 0 {
		pop := heap.Pop(pq).(*Vertex)

		colors[pop.id] = BLACK

		if pop.pri > dp[pop.id] {
			continue
		}

		for _, e := range AG[pop.id] {
			if colors[e.to] == BLACK {
				continue
			}

			if dp[e.to] > dp[pop.id]+e.cost {
				dp[e.to] = dp[pop.id] + e.cost
				heap.Push(pq, &Vertex{pri: dp[e.to], id: e.to})
				colors[e.to], parents[e.to] = GRAY, pop.id
			}
		}
	}

	return dp, parents
}

type VertexPQ []*Vertex

func (pq VertexPQ) Len() int           { return len(pq) }
func (pq VertexPQ) Less(i, j int) bool { return pq[i].pri < pq[j].pri } // <: ASC, >: DESC
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

// how to use
// temp := make(VertexPQ, 0, 100000+1)
// pq := &temp
// heap.Init(pq)
// heap.Push(pq, &Vertex{pri: intValue})
// popped := heap.Pop(pq).(*Vertex)

/*******************************************************************/

/********** common constants **********/

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

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
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

/*********** Input ***********/

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

// PrintfBufStdout is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func PrintfBufStdout(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}

/*********** Debugging ***********/

// PrintfDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintfDebug(format string, a ...interface{}) {
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

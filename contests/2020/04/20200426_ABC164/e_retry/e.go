/*
URL:
https://atcoder.jp/contests/abc164/tasks/abc164_e
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

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int) int {
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

var (
	N, M, S          int
	U, V, A, B, C, D []int

	G [MT + 5][]Edge
)

const (
	MN = 50
	MM = 5000
	MT = 260000
)

func main() {
	N, M, S = ReadInt3()
	for i := 0; i < M; i++ {
		u, v, a, b := ReadInt4()
		u--
		v--
		U = append(U, u)
		V = append(V, v)
		A = append(A, a)
		B = append(B, b)
	}
	for i := 0; i < N; i++ {
		c, d := ReadInt2()
		C = append(C, c)
		D = append(D, d)
	}

	for coin := 0; coin <= MM; coin++ {
		for i := 0; i < M; i++ {
			u, v, a, b := U[i], V[i], A[i], B[i]

			if coin-a >= 0 {
				G[u+N*coin] = append(G[u+N*coin], Edge{to: v + (coin-a)*N, cost: b})
				G[v+N*coin] = append(G[v+N*coin], Edge{to: u + (coin-a)*N, cost: b})
			}
		}

		for i := 0; i < N; i++ {
			c, d := C[i], D[i]
			tocoin := Min(coin+c, MM)

			if coin == tocoin {
				continue
			}

			G[i+N*coin] = append(G[i+N*coin], Edge{to: i + N*tocoin, cost: d})
		}
	}

	dp, _ := dijkstra(Min(S, MM)*N, MT, G[:MT])

	for i := 1; i < N; i++ {
		ans := INF_DIJK
		for j := 0; j <= MM; j++ {
			ChMin(&ans, dp[i+N*j])
		}
		fmt.Println(ans)
	}
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

func dijkstra(sid, n int, AG [][]Edge) ([]int, []int) {
	dp := make([]int, n)
	colors, parents := make([]int, n), make([]int, n)
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

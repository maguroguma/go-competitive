/*
URL:
https://atcoder.jp/contests/abc143/tasks/abc143_e
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

/*******************************************************************/

var (
	n, m, l int
	q       int

	A [300 + 5][300 + 5]int
	G [][]EdgeWeight
)

func main() {
	n, m, l = ReadInt3()
	for i := 0; i < n; i++ {
		row := make([]EdgeWeight, n)
		for j := 0; j < n; j++ {
			row[j] = EdgeWeight{gas: INF_BIT30}
		}
		G = append(G, row)
	}
	for i := 0; i < m; i++ {
		a, b, c := ReadInt3()
		a--
		b--

		ew := EdgeWeight{gas: c}
		G[a][b] = ew
		G[b][a] = ew
	}

	vinf := V{gas: -1, times: INF_BIT60}
	einf := EdgeWeight{gas: INF_BIT30}
	vzero := V{gas: l, times: 0}
	less := func(l, r V) bool {
		if l.times < r.times {
			return true
		} else if l.times > r.times {
			return false
		} else {
			return l.gas > r.gas
		}
	}
	genNextV := func(cv V, e EdgeWeight) V {
		if l < e.gas {
			return vinf
		}

		if cv.gas >= e.gas {
			return V{gas: cv.gas - e.gas, times: cv.times}
		}

		return V{gas: l - e.gas, times: cv.times + 1}
	}
	ds := NewDijkstraSolver(vinf, einf, less, genNextV)

	for i := 0; i < n; i++ {
		dp := ds.Dijkstra([]StartPoint{{id: i, vzero: vzero}}, n, G)

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			A[i][j] = dp[j].times
		}
	}

	q = ReadInt()
	for i := 0; i < q; i++ {
		s, t := ReadInt2()
		s--
		t--

		if A[s][t] >= INF_BIT60 {
			fmt.Println(-1)
		} else {
			fmt.Println(A[s][t])
		}
	}
}

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

// DP value type
type V struct {
	// {{
	gas, times int
	// }}
}

// weight of edges
type EdgeWeight struct {
	// {{
	gas int
	// }}
}

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	id    int
	vzero V
}

type DijkstraSolver struct {
	vinf  V
	ewinf EdgeWeight
	Less  func(l, r V) bool          // Less returns l < r.
	NextV func(cv V, e EdgeWeight) V // NextV returns next value considered by transition.
}

func NewDijkstraSolver(
	vinf V, ewinf EdgeWeight, Less func(l, r V) bool, NextV func(cv V, ew EdgeWeight) V,
) *DijkstraSolver {
	ds := new(DijkstraSolver)

	ds.vinf, ds.ewinf = vinf, ewinf
	ds.Less, ds.NextV = Less, NextV

	return ds
}

// verified by [ABC143-E](https://atcoder.jp/contests/abc143/tasks/abc143_e)
func (ds *DijkstraSolver) Dijkstra(S []StartPoint, n int, AG [][]EdgeWeight) []V {
	// initialize data
	dp, colors := ds.initAll(n)

	// configure about start points (some problems have multi start points)
	ds.initStartPoint(S, dp, colors)

	// body of dijkstra algorithm (O(n^2))
	for {
		minv, u := ds.vinf, -1

		// find next optimal node
		for i := 0; i < n; i++ {
			if ds.Less(dp[i], minv) && colors[i] != BLACK {
				u = i
				minv = dp[i]
			}
		}
		if u == -1 {
			break
		}

		colors[u] = BLACK

		// update all nodes v from node u
		for v := 0; v < n; v++ {
			if colors[v] != BLACK && AG[u][v] != ds.ewinf {
				nv := ds.NextV(dp[u], AG[u][v])
				if ds.Less(nv, dp[v]) {
					dp[v] = nv
					colors[v] = GRAY
				}
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
func (ds *DijkstraSolver) initStartPoint(S []StartPoint, dp []V, colors []int) {
	for _, sp := range S {
		dp[sp.id] = sp.vzero
		colors[sp.id] = GRAY
	}
}

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
	// WHITE = 0
	// GRAY  = 1
	// BLACK = 2
)

/*******************************************************************/

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

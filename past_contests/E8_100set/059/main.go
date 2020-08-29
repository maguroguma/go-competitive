/*
URL:
https://atcoder.jp/contests/joi2014yo/tasks/joi2014yo_e
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

var (
	n, k int

	C [5000 + 5]int
	R [5000 + 5]int
	// M [5000 + 5][5000 + 5]int
	M [][]int32
	g [5000 + 5][]int

	// G [5000 + 5][]Edge
)

func main() {
	n, k = readi2()

	M = make([][]int32, n)
	for i := 0; i < n; i++ {
		M[i] = make([]int32, n)
	}

	for i := 0; i < n; i++ {
		c, r := readi2()
		C[i], R[i] = c, r
	}
	for i := 0; i < k; i++ {
		s, t := readi2()
		s--
		t--
		g[s] = append(g[s], t)
		g[t] = append(g[t], s)
	}

	for i := 0; i < n; i++ {
		bfs(i)
	}

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i == j {
	// 			continue
	// 		}
	// 		if M[i][j] >= INF_BIT60 {
	// 			continue
	// 		}
	// 		G[i] = append(G[i], Edge{to: j, cost: M[i][j]})
	// 	}
	// }

	less := func(l, r V) bool { return l < r }
	// f := func(cv *Vertex, e Edge) V {
	// 	return cv.v + V(e.cost)
	// }
	f := func(cv V, e int) V {
		return cv + V(e)
	}
	// ds := NewDijkstraSolver(INF_BIT60, less, f)
	ds := NewDijkstraSolver(INF_BIT30, INF_BIT30, less, f)
	dp := ds.Dijkstra([]StartPoint{{0, 0}}, n, M)
	fmt.Println(dp[n-1])
}

func bfs(sid int) {
	visited := make([]bool, n)
	dp := make([]int, n)

	Q := []int{sid}
	visited[sid], dp[sid] = true, 0
	for len(Q) > 0 {
		cid := Q[0]
		Q = Q[1:]

		for _, nid := range g[cid] {
			if visited[nid] {
				continue
			}

			visited[nid] = true
			dp[nid] = dp[cid] + 1
			Q = append(Q, nid)
		}
	}

	for i := 0; i < n; i++ {
		if dp[i] <= R[sid] {
			M[sid][i] = int32(C[sid])
		} else {
			M[sid][i] = int32(INF_BIT30)
		}
	}
}

// DP value type
// type V struct {
// }
type V int

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	id    int
	vinit V
}

type DijkstraSolver struct {
	vinf     V
	einf     int
	Less     func(l, r V) bool   // Less returns whether l is strictly less than r, and is also used for priority queue.
	GenNextV func(cv V, e int) V // GenNextV returns next value considered by transition.
}

func NewDijkstraSolver(
	vinf V, einf int, Less func(l, r V) bool, GenNextV func(cv V, e int) V,
) *DijkstraSolver {
	ds := new(DijkstraSolver)

	ds.vinf, ds.einf = vinf, einf
	ds.Less, ds.GenNextV = Less, GenNextV

	return ds
}

// verified by [ABC143-E](https://atcoder.jp/contests/abc143/tasks/abc143_e)
func (ds *DijkstraSolver) Dijkstra(S []StartPoint, n int, AG [][]int32) []V {
	// initialize data
	dp, colors := ds.InitAll(n)

	// configure about start points (some problems have multi start points)
	ds.InitStartPoint(S, dp, colors)

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
			if colors[v] != BLACK && AG[u][v] != int32(ds.einf) {
				nv := ds.GenNextV(dp[u], int(AG[u][v]))
				if ds.Less(nv, dp[v]) {
					dp[v] = nv
					colors[v] = GRAY
				}
			}
		}
	}

	return dp
}

// InitAll returns initialized dp and colors slices.
func (ds *DijkstraSolver) InitAll(n int) (dp []V, colors []int) {
	dp, colors = make([]V, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = ds.vinf
		colors[i] = WHITE
	}

	return dp, colors
}

// InitStartPoint returns initialized priority queue, and update dp and colors slices.
// *This function update arguments (side effects).*
func (ds *DijkstraSolver) InitStartPoint(S []StartPoint, dp []V, colors []int) {
	for _, sp := range S {
		dp[sp.id] = sp.vinit
		colors[sp.id] = GRAY
	}
}

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

// readi returns an integer.
func readi() int {
	return int(readInt64())
}
func readi2() (int, int) {
	return int(readInt64()), int(readInt64())
}
func readi3() (int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64())
}
func readi4() (int, int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64()), int(readInt64())
}

// readll returns as integer as int64.
func readll() int64 {
	return readInt64()
}
func readll2() (int64, int64) {
	return readInt64(), readInt64()
}
func readll3() (int64, int64, int64) {
	return readInt64(), readInt64(), readInt64()
}
func readll4() (int64, int64, int64, int64) {
	return readInt64(), readInt64(), readInt64(), readInt64()
}

func readInt64() int64 {
	i, err := strconv.ParseInt(ReadString(), 0, 64)
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

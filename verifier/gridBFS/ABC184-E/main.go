/*
URL:
https://atcoder.jp/contests/abc184/tasks/abc184_e
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var (
	h, w int
	A    [][]rune

	R    [30][]int
	warp [30]bool

	G        [][]int
	sid, gid int
)

func tovid(y, x int) int {
	return y*w + x
}

var toy func(int) int
var tox func(int) int

func main() {
	defer stdout.Flush()

	h, w = readi2()
	for i := 0; i < h; i++ {
		row := readrs()
		A = append(A, row)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if A[i][j] != 'S' && A[i][j] != 'G' && A[i][j] != '.' && A[i][j] != '#' {
				v := A[i][j] - 'a'
				R[v] = append(R[v], tovid(i, j))
			} else if A[i][j] == 'S' {
				sid = tovid(i, j)
			} else if A[i][j] == 'G' {
				gid = tovid(i, j)
			}
		}
	}

	toy = func(id int) int { return id / w }
	tox = func(id int) int { return id % w }

	// N := h * w
	// G = make([][]int, N)
	// for i := 0; i < h; i++ {
	// 	for j := 0; j < w; j++ {
	// 		if A[i][j] == '#' {
	// 			continue
	// 		}

	// 		cid := toid(i, j)

	// 		for _, s := range steps {
	// 			dy, dx := s[0], s[1]
	// 			ny, nx := i+dy, j+dx
	// 			if 0 <= ny && ny < h && 0 <= nx && nx < w && A[ny][nx] != '#' {
	// 				nid := toid(ny, nx)

	// 				G[cid] = append(G[cid], nid)
	// 			}
	// 		}
	// 	}
	// }

	G, N := GridToAdjacencyList(h, w)

	dp, visited := SSSPByBFS(sid, N, G, 1)
	if visited[gid] {
		fmt.Println(dp[gid])
	} else {
		fmt.Println(-1)
	}
}

// GridToAdjacencyList provides basic function
//  for converting 2-dimensional grid to graph as adjacency list format.
// Grid: size is H*W, and no wall.
func GridToAdjacencyList(h, w int) (G [][]int, N int) {
	steps := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	toid := func(i, j int) int { return w*i + j }
	isInGrid := func(y, x int) bool {
		return 0 <= y && y < h && 0 <= x && x < w
	}

	N = h * w
	G = make([][]int, N)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if A[i][j] == '#' {
				continue
			}

			cid := toid(i, j)

			for _, s := range steps {
				dy, dx := s[0], s[1]
				ny, nx := i+dy, j+dx
				if isInGrid(ny, nx) && A[ny][nx] != '#' {
					nid := toid(ny, nx)

					G[cid] = append(G[cid], nid)
				}
			}
		}
	}

	return
}

// verified by https://codeforces.com/contest/1320/problem/B
func SSSPByBFS(sid int, n int, AG [][]int, edgeCost int) (dp []int, visited []bool) {
	const INF_SSSP = 1 << uint(30)

	dp = make([]int, n)
	visited = make([]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = INF_SSSP
		visited[i] = false
	}

	Q := []Node{}
	dp[sid] = 0
	visited[sid] = true
	Q = append(Q, Node{id: sid, cost: dp[sid]})

	for len(Q) > 0 {
		cnode := Q[0]
		Q = Q[1:]

		for _, nid := range AG[cnode.id] {
			if visited[nid] {
				continue
			}

			dp[nid] = cnode.cost + edgeCost
			visited[nid] = true

			if nid == gid {
				return
			}

			Q = append(Q, Node{id: nid, cost: dp[nid]})
		}
		// ワープ移動
		i, j := toy(cnode.id), tox(cnode.id)
		if A[i][j] != 'S' && A[i][j] != 'G' && A[i][j] != '.' && A[i][j] != '#' {
			v := A[i][j] - 'a'
			if warp[v] {
				continue
			}

			warp[v] = true
			for _, to := range R[v] {
				if visited[to] {
					continue
				}

				dp[to] = cnode.cost + 1
				visited[to] = true
				Q = append(Q, Node{id: to, cost: dp[to]})
			}
		}
	}

	return dp, visited
}

type Node struct {
	id, cost int
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

// pow is integer version of math.Pow
// pow calculate a power by Binary Power (二分累乗法(O(log e))).
func pow(a, e int) int {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}

	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := pow(a, halfE)
		return half * half
	}

	return a * pow(a, e-1)
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

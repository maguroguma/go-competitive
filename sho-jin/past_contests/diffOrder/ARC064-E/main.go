/*
URL:
https://atcoder.jp/contests/arc064/tasks/arc064_c
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

func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func UnitVector(dx, dy float64) (ex, ey float64) {
	dist := math.Sqrt(dx*dx + dy*dy)
	return dx / dist, dy / dist
}

func Rotate90(cx, cy float64) (nx, ny float64) {
	return -cy, cx
}

func RotateN(cx, cy, radi float64) (nx, ny float64) {
	nx = math.Cos(radi)*cx - math.Sin(radi)*cy
	ny = math.Sin(radi)*cx + math.Cos(radi)*cy
	return nx, ny
}

func Midpoint(x1, y1, x2, y2 float64) (mx, my float64) {
	return (x1 + x2) / 2.0, (y1 + y2) / 2.0
}

var (
	xs, ys, xt, yt float64
	n              int
	X, Y, R        []float64

	G [][]Weight
	N int
)

func main() {
	defer stdout.Flush()

	F := readis(4)
	xs, ys, xt, yt = float64(F[0]), float64(F[1]), float64(F[2]), float64(F[3])
	n = readi()
	N = n + 2
	X, Y, R = make([]float64, N), make([]float64, N), make([]float64, N)
	for i := 0; i < n; i++ {
		x, y, r := readf(), readf(), readf()
		X[i], Y[i], R[i] = x, y, r
	}
	X[n], Y[n], R[n] = xs, ys, 0.0
	X[n+1], Y[n+1], R[n+1] = xt, yt, 0.0

	winf := Weight{1e15}
	G = make([][]Weight, N)
	for i := 0; i < N; i++ {
		G[i] = make([]Weight, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			G[i][j] = winf
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			d := Distance(X[i], Y[i], X[j], Y[j])
			r := R[i] + R[j]
			dist := math.Max(d-r, 0.0)
			w := Weight{dist}

			G[i][j], G[j][i] = w, w
		}
	}

	vinf := Value{1e15}
	less := func(l, r Value) bool { return l.v < r.v }
	estimate := func(cv Value, w Weight) Value {
		return Value{cv.v + w.dist}
	}
	SS := []StartPoint{{id: n, vzero: Value{0.0}}}

	ds := NewDijkstraSolver(vinf, winf, less, estimate)
	dp := ds.Dijkstra(SS, N, G)

	fmt.Println(dp[n+1].v)
}

// type Value and Weight should be modified according to problems.

// DP value type
type Value struct {
	v float64
}

// weight of edges
type Weight struct {
	dist float64
}

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	id    int
	vzero Value
}

// Less returns l < r, and shared with pq.
type Less func(l, r Value) bool

// Estimate returns next value considered by transition.
type Estimate func(cv Value, w Weight) Value

func NewDijkstraSolver(vinf Value, winf Weight, less Less, estimate Estimate) *DijkstraSolver {
	ds := new(DijkstraSolver)

	ds.vinf, ds.winf = vinf, winf
	ds.less, ds.estimate = less, estimate

	return ds
}

// verified by [ABC143-E](https://atcoder.jp/contests/abc143/tasks/abc143_e)
func (ds *DijkstraSolver) Dijkstra(S []StartPoint, n int, AG [][]Weight) []Value {
	// initialize data
	dp, colors := ds._initAll(n)

	// configure about start points (some problems have multi start points)
	ds._initStartPoint(S, dp, colors)

	// body of dijkstra algorithm (O(n^2))
	for {
		minv, u := ds.vinf, -1

		// find next optimal node
		for i := 0; i < n; i++ {
			if ds.less(dp[i], minv) && colors[i] != BLACK {
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
			if colors[v] != BLACK && AG[u][v] != ds.winf {
				nv := ds.estimate(dp[u], AG[u][v])
				if ds.less(nv, dp[v]) {
					dp[v] = nv
					colors[v] = GRAY
				}
			}
		}
	}

	return dp
}

type DijkstraSolver struct {
	vinf     Value
	winf     Weight
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
func (ds *DijkstraSolver) _initStartPoint(S []StartPoint, dp []Value, colors []int) {
	for _, sp := range S {
		dp[sp.id] = sp.vzero
		colors[sp.id] = GRAY
	}
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

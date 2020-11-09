/*
URL:
https://atcoder.jp/contests/joisc2008/tasks/joisc2008_sheet
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
	n, w, h int
	A       [][]int

	X, Y       [1000 + 50][]int
	L, R, U, D [1000 + 50]int

	G [1000 + 50][]Edge
)

func main() {
	defer stdout.Flush()

	n, w, h = readi3()
	for i := 0; i < h; i++ {
		row := readis(w)
		A = append(A, row)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			A[i][j]--
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			v := A[i][j]
			if v == -1 {
				continue
			}
			X[v] = append(X[v], j)
			Y[v] = append(Y[v], i)
		}
	}

	for i := 0; i < n; i++ {
		L[i], R[i], U[i], D[i] = -1, -1, -1, -1
	}

	for i := 0; i < n; i++ {
		if len(X[i]) == 0 {
			continue
		}

		maxX, minX := max(X[i]...), min(X[i]...)
		maxY, minY := max(Y[i]...), min(Y[i]...)
		L[i], R[i], U[i], D[i] = minX, maxX, maxY, minY
	}

	for i := 0; i < n; i++ {
		if L[i] == -1 {
			continue
		}

		l, r, u, d := L[i], R[i], U[i], D[i]
		memo := make(map[int]bool)
		for y := d; y <= u; y++ {
			for x := l; x <= r; x++ {
				if A[y][x] == i {
					continue
				}
				memo[A[y][x]] = true
			}
		}

		for k := range memo {
			G[i] = append(G[i], Edge{k, 1})
		}
	}

	_, ids := TSort(n, G[:n])

	for i := 0; i < len(ids); i++ {
		ids[i]++
	}

	fmt.Println(PrintIntsLine(ids...))
}

type Edge struct {
	to   int
	cost int
}

// TSort returns a node ids list in topological order.
// node id is 0-based.
// Time complexity: O(|E| + |V|)
func TSort(nn int, AG [][]Edge) (ok bool, tsortedIDs []int) {
	tsortedIDs = []int{}

	inDegrees := make([]int, nn)
	for s := 0; s < nn; s++ {
		for _, e := range AG[s] {
			id := e.to
			inDegrees[id]++
		}
	}

	stack := []int{}
	for nid := 0; nid < nn; nid++ {
		if inDegrees[nid] == 0 {
			stack = append(stack, nid)
		}
	}

	for len(stack) > 0 {
		cid := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		tsortedIDs = append(tsortedIDs, cid)

		for _, e := range AG[cid] {
			nid := e.to
			inDegrees[nid]--
			if inDegrees[nid] == 0 {
				stack = append(stack, nid)
			}
		}
	}

	if len(tsortedIDs) != nn {
		return false, nil
	}

	return true, tsortedIDs
}

// LongestPath returns a length of longest path of a given graph.
// Time complexity: O(|E| + |V|)
func LongestPath(tsortedIDs []int, AG [][]Edge) (maxLength int, dp []int) {
	_chmax := func(updatedValue *int, target int) bool {
		if *updatedValue < target {
			*updatedValue = target
			return true
		}
		return false
	}

	dp = make([]int, len(tsortedIDs))

	for i := 0; i < len(tsortedIDs); i++ {
		cid := tsortedIDs[i]
		for _, e := range AG[cid] {
			nid := e.to
			_chmax(&dp[nid], dp[cid]+e.cost)
		}
	}

	maxLength = 0
	for i := 0; i < len(tsortedIDs); i++ {
		_chmax(&maxLength, dp[i])
	}

	return maxLength, dp
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

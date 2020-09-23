/*
URL:
https://atcoder.jp/contests/dp/tasks/dp_v
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
	n, m  int
	edges [][]int
)

func main() {
	n, m = ReadInt2()
	for i := 0; i < n-1; i++ {
		x, y := ReadInt2()
		x--
		y--

		e := []int{x, y}
		edges = append(edges, e)
	}

	f := func(l, r T) T { return T(int(l) * int(r) % m) }
	g := func(t T, idx int) T { return t + 1 }
	s := NewReRooting(n, edges, 1, f, g)
	for i := 0; i < n; i++ {
		fmt.Println(s.Query(i) - 1)
	}
}

type T int

type ReRooting struct {
	NodeCount int

	Identity    T
	Operate     func(l, r T) T
	OperateNode func(t T, idx int) T

	Adjacents         [][]int
	IndexForAdjacents [][]int

	Res []T
	DP  [][]T
}

func NewReRooting(
	nodeCount int, edges [][]int, identity T, operate func(l, r T) T, operateNode func(t T, idx int) T,
) *ReRooting {
	s := new(ReRooting)

	s.NodeCount = nodeCount
	s.Identity = identity
	s.Operate = operate
	s.OperateNode = operateNode

	s.Adjacents = make([][]int, nodeCount)
	s.IndexForAdjacents = make([][]int, nodeCount)
	for _, e := range edges {
		s.IndexForAdjacents[e[0]] = append(s.IndexForAdjacents[e[0]], len(s.Adjacents[e[1]]))
		s.IndexForAdjacents[e[1]] = append(s.IndexForAdjacents[e[1]], len(s.Adjacents[e[0]]))
		s.Adjacents[e[0]] = append(s.Adjacents[e[0]], e[1])
		s.Adjacents[e[1]] = append(s.Adjacents[e[1]], e[0])
	}

	s.DP = make([][]T, len(s.Adjacents))
	s.Res = make([]T, len(s.Adjacents))

	for i := 0; i < len(s.Adjacents); i++ {
		s.DP[i] = make([]T, len(s.Adjacents[i]))
	}

	if s.NodeCount > 1 {
		s.Initialize()
	} else {
		s.Res[0] = s.OperateNode(s.Identity, 0)
	}

	return s
}

func (s *ReRooting) Query(node int) T {
	return s.Res[node]
}

func (s *ReRooting) Initialize() {
	parents, order := make([]int, s.NodeCount), make([]int, s.NodeCount)

	// #region InitOrderedTree
	index := 0
	stack := []int{}
	stack = append(stack, 0)
	parents[0] = -1
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order[index] = node
		index++
		for i := 0; i < len(s.Adjacents[node]); i++ {
			adjacent := s.Adjacents[node][i]
			if adjacent == parents[node] {
				continue
			}
			stack = append(stack, adjacent)
			parents[adjacent] = node
		}
	}
	// endregion

	// #region fromLeaf
	for i := len(order) - 1; i >= 1; i-- {
		node := order[i]
		parent := parents[node]

		accum := s.Identity
		parentIndex := -1
		for j := 0; j < len(s.Adjacents[node]); j++ {
			if s.Adjacents[node][j] == parent {
				parentIndex = j
				continue
			}
			accum = s.Operate(accum, s.DP[node][j])
		}
		s.DP[parent][s.IndexForAdjacents[node][parentIndex]] = s.OperateNode(accum, node)
	}
	// endregion

	// #region toLeaf
	for i := 0; i < len(order); i++ {
		node := order[i]
		accum := s.Identity
		accumsFromTail := make([]T, len(s.Adjacents[node]))
		accumsFromTail[len(accumsFromTail)-1] = s.Identity
		for j := len(accumsFromTail) - 1; j >= 1; j-- {
			accumsFromTail[j-1] = s.Operate(s.DP[node][j], accumsFromTail[j])
		}
		for j := 0; j < len(accumsFromTail); j++ {
			s.DP[s.Adjacents[node][j]][s.IndexForAdjacents[node][j]] = s.OperateNode(s.Operate(accum, accumsFromTail[j]), node)
			accum = s.Operate(accum, s.DP[node][j])
		}
		s.Res[node] = s.OperateNode(accum, node)
	}
	// endregion
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

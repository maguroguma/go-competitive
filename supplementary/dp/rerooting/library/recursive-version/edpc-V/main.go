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
	n, m int
	G    [100000 + 50][]int
)

func main() {
	n, m = ReadInt2()
	for i := 0; i < n-1; i++ {
		x, y := ReadInt2()
		x--
		y--

		G[x] = append(G[x], y)
		G[y] = append(G[y], x)
	}

	f := func(x, y T) T { return T(int(x) * int(y) % m) }
	g := func(v T, p int) T { return v + 1 }
	s := NewReRooting(n, G[:n], 1, f, g)
	for i := 0; i < n; i++ {
		fmt.Println(s.Query(i) - 1)
	}
}

type T int

type ReRooting struct {
	n int
	G [][]int

	ti      T
	dp, res []T
	merge   func(l, r T) T
	addNode func(t T, idx int) T
}

func NewReRooting(
	n int, AG [][]int, ti T, merge func(l, r T) T, addNode func(t T, idx int) T,
) *ReRooting {
	s := new(ReRooting)
	s.n, s.G, s.ti, s.merge, s.addNode = n, AG, ti, merge, addNode
	s.dp, s.res = make([]T, n), make([]T, n)

	s.Solve()

	return s
}

func (s *ReRooting) Solve() {
	s.inOrder(0, -1)
	s.reroot(0, -1, s.ti)
}

func (s *ReRooting) Query(idx int) T {
	return s.res[idx]
}

func (s *ReRooting) inOrder(cid, pid int) T {
	res := s.ti

	for _, nid := range G[cid] {
		if nid == pid {
			continue
		}

		res = s.merge(res, s.inOrder(nid, cid))
	}
	res = s.addNode(res, cid)
	s.dp[cid] = res

	return s.dp[cid]
}

func (s *ReRooting) reroot(cid, pid int, parentValue T) {
	childValues := []T{}
	nexts := []int{}
	for _, nid := range G[cid] {
		if nid == pid {
			continue
		}
		childValues = append(childValues, s.dp[nid])
		nexts = append(nexts, nid)
	}

	// result of cid
	rootValue := s.ti
	for _, v := range childValues {
		rootValue = s.merge(rootValue, v)
	}
	rootValue = s.merge(rootValue, parentValue)
	rootValue = s.addNode(rootValue, cid)
	s.res[cid] = rootValue

	// for children
	accum := s.merge(s.ti, parentValue)
	length := len(childValues)
	if length == 0 {
		return
	}
	if length == 1 {
		s.reroot(nexts[0], cid, s.addNode(accum, cid))
		return
	}

	// cid has more than one child
	R, L := make([]T, length), make([]T, length)
	L[0] = s.merge(s.ti, childValues[0])
	for i := 1; i < length; i++ {
		L[i] = s.merge(L[i-1], childValues[i])
	}
	R[length-1] = s.merge(s.ti, childValues[length-1])
	for i := length - 2; i >= 0; i-- {
		R[i] = s.merge(R[i+1], childValues[i])
	}

	for i, nid := range nexts {
		if i == 0 {
			s.reroot(nid, cid, s.addNode(s.merge(accum, R[1]), cid))
		} else if i == length-1 {
			s.reroot(nid, cid, s.addNode(s.merge(accum, L[length-2]), cid))
		} else {
			s.reroot(nid, cid, s.addNode(s.merge(accum, s.merge(L[i-1], R[i+1])), cid))
		}
	}
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

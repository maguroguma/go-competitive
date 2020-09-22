/*
URL:
https://atcoder.jp/contests/abc040/tasks/abc040_d
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

var (
	n, m, q int
	A, B, Y []int
	V, W    []int
)

type Person struct {
	key            int
	id, city, year int
}
type PersonList []*Person

func (l PersonList) Len() int {
	return len(l)
}
func (l PersonList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l PersonList) Less(i, j int) bool {
	return l[i].year > l[j].year
}

// how to use
// L := make(PersonList, 0, 200000+5)
// L = append(L, &Person{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

type Edge struct {
	key        int
	s, t, year int
}
type EdgeList []*Edge

func (l EdgeList) Len() int {
	return len(l)
}
func (l EdgeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l EdgeList) Less(i, j int) bool {
	return l[i].year > l[j].year
}

// how to use
// L := make(EdgeList, 0, 200000+5)
// L = append(L, &Edge{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

func main() {
	n, m = ReadInt2()
	for i := 0; i < m; i++ {
		a, b, y := ReadInt3()
		a--
		b--
		A = append(A, a)
		B = append(B, b)
		Y = append(Y, y)
	}
	q = ReadInt()
	for i := 0; i < q; i++ {
		v, w := ReadInt2()
		v--
		V = append(V, v)
		W = append(W, w)
	}

	E := make(EdgeList, 0)
	for i := 0; i < m; i++ {
		a, b, y := A[i], B[i], Y[i]
		E = append(E, &Edge{s: a, t: b, year: y})
	}
	P := make(PersonList, 0)
	for i := 0; i < q; i++ {
		v, w := V[i], W[i]
		P = append(P, &Person{id: i, city: v, year: w})
	}
	sort.Stable(E)
	sort.Stable(P)

	uf := NewUnionFind(n)
	answers := make([]int, q)
	j := 0
	for i := 0; i < len(P); i++ {
		p := P[i]
		for j < len(E) && E[j].year > p.year {
			uf.Unite(E[j].s, E[j].t)
			j++
		}

		answers[p.id] = uf.CcSize(p.city)
	}

	for i := 0; i < q; i++ {
		// fmt.Println(answers[i])
		PrintfBufStdout("%d\n", answers[i])
	}
	stdout.Flush()
}

// 0-based
// uf := NewUnionFind(n)
// uf.Root(x) 			// Get root node of the node x
// uf.Unite(x, y) 	// Unite node x and node y
// uf.Same(x, y) 		// Judge x and y are in the same connected component.
// uf.CcSize(x) 		// Get size of the connected component including node x
// uf.CcNum() 			// Get number of connected components

// UnionFind provides disjoint set algorithm.
// Node id starts from 0 (0-based setting).
type UnionFind struct {
	parents []int
}

// NewUnionFind returns a pointer of a new instance of UnionFind.
func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.parents = make([]int, n)

	for i := 0; i < n; i++ {
		uf.parents[i] = -1
	}

	return uf
}

// Root method returns root node of an argument node.
// Root method is a recursive function.
func (uf *UnionFind) Root(x int) int {
	if uf.parents[x] < 0 {
		return x
	}

	// route compression
	uf.parents[x] = uf.Root(uf.parents[x])
	return uf.parents[x]
}

// Unite method merges a set including x and a set including y.
func (uf *UnionFind) Unite(x, y int) bool {
	xp := uf.Root(x)
	yp := uf.Root(y)

	if xp == yp {
		return false
	}

	// merge: xp -> yp
	// merge larger set to smaller set
	if uf.CcSize(xp) > uf.CcSize(yp) {
		xp, yp = yp, xp
	}
	// update set size
	uf.parents[yp] += uf.parents[xp]
	// finally, merge
	uf.parents[xp] = yp

	return true
}

// Same method returns whether x is in the set including y or not.
func (uf *UnionFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

// CcSize method returns the size of a set including an argument node.
func (uf *UnionFind) CcSize(x int) int {
	return -uf.parents[uf.Root(x)]
}

// CcNum method returns the number of connected components.
// Time complextity is O(n)
func (uf *UnionFind) CcNum() int {
	res := 0
	for i := 0; i < len(uf.parents); i++ {
		if uf.parents[i] < 0 {
			res++
		}
	}
	return res
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

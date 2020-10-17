/*
URL:
https://atcoder.jp/contests/arc104/tasks/arc104_c
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
	n    int
	A, B []int
	cnts [200 + 10]int

	Q []Seg
	E [200 + 10]int
)

type Seg struct {
	s, t int
	dir  int
}

const (
	STAY, RIGHT, LEFT = 0, 1, 2
)

func main() {
	defer stdout.Flush()

	n = readi()
	for i := 0; i < n; i++ {
		a, b := readi2()
		A = append(A, a)
		B = append(B, b)

		if a >= 1 {
			cnts[a]++
		}
		if b >= 1 {
			cnts[b]++
		}
	}

	// 必要条件
	for i := 0; i < n; i++ {
		if A[i] == -1 || B[i] == -1 {
			continue
		}

		if A[i] >= B[i] {
			// debugf("AがB以上\n")
			fmt.Println("No")
			return
		}
	}
	for i := 1; i <= 2*n; i++ {
		if cnts[i] >= 2 {
			// debugf("同じ階で乗り降りする人たちがいた\n")
			fmt.Println("No")
			return
		}
	}

	undefNum := 0
	for i := 0; i < n; i++ {
		a, b := A[i], B[i]
		if a == -1 && b == -1 {
			// 最後に処理する
			undefNum++
		} else if a == -1 {
			if b-1 <= 0 {
				fmt.Println("No")
				return
			}

			seg := Seg{b - 1, b, LEFT}
			Q = append(Q, seg)
		} else if b == -1 {
			if a+1 > 2*n {
				fmt.Println("No")
				return
			}

			seg := Seg{a, a + 1, RIGHT}
			Q = append(Q, seg)
		} else {
			seg := Seg{a, b, STAY}
			Q = append(Q, seg)
		}
	}

	for i := 0; i < n+5; i++ {
		// check
		for j := 0; j < len(Q); j++ {
			if Q[j].dir == STAY {
				continue
			}

			ok := true
			for k := 0; k < len(Q); k++ {
				if j == k {
					continue
				}
				// セグメントj, kの衝突を調べる
				if isCross(Q[j], Q[k]) {
					ok = false
					break
				}
			}

			if ok {
				Q[j].dir = STAY
			}
		}

		// 未確定をすべて伸ばす
		for j := 0; j < len(Q); j++ {
			// STAY以外は伸ばす
			if Q[j].dir == STAY {
				continue
			}

			if Q[j].dir == RIGHT {
				Q[j].t++
				if Q[j].t > 2*n {
					fmt.Println("No")
					return
				}
			} else {
				Q[j].s--
				if Q[j].s <= 0 {
					fmt.Println("No")
					return
				}
			}
		}
	}

	// 確定済みのものについて、正しいかチェックする
	for i := 0; i < len(Q); i++ {
		for j := i + 1; j < len(Q); j++ {
			if isCross(Q[i], Q[j]) {
				fmt.Println("No")
				return
			}
		}
	}

	for i := 0; i < len(Q); i++ {
		E[Q[i].s] = 1
		E[Q[i].t] = 1
	}

	// 両端が不定のものは、未定義の場所が長さlあるとき、(l-1)個埋められる
	// それらで処理し切ることができるならYes
	tmp := []int{}
	for i := 1; i <= 2*n; i++ {
		tmp = append(tmp, E[i])
	}
	comp, cnts := RunLengthEncoding(tmp)
	cap := 0
	for i := 0; i < len(comp); i++ {
		if comp[i] == 0 {
			cap += cnts[i] - 1
		}
	}

	if undefNum <= cap {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isCross(A, B Seg) bool {
	a, b := A.s, A.t
	c, d := B.s, B.t

	return (a <= c && c <= b) || (c <= a && a <= d)
}

// RunLengthEncoding returns encoded slice of an input.
func RunLengthEncoding(S []int) ([]int, []int) {
	runes := []int{}
	lengths := []int{}

	l := 0
	for i := 0; i < len(S); i++ {
		// 1文字目の場合保持
		if i == 0 {
			l = 1
			continue
		}

		if S[i-1] == S[i] {
			// 直前の文字と一致していればインクリメント
			l++
		} else {
			// 不一致のタイミングで追加し、長さをリセットする
			runes = append(runes, S[i-1])
			lengths = append(lengths, l)
			l = 1
		}
	}
	runes = append(runes, S[len(S)-1])
	lengths = append(lengths, l)

	return runes, lengths
}

// RunLengthDecoding decodes RLE results.
func RunLengthDecoding(S []int, L []int) []int {
	if len(S) != len(L) {
		panic("S, L are not RunLengthEncoding results")
	}

	res := []int{}

	for i := 0; i < len(S); i++ {
		for j := 0; j < L[i]; j++ {
			res = append(res, S[i])
		}
	}

	return res
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

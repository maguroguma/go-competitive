/*
URL:
https://yukicoder.me/problems/no/430
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
	text     []rune
	m        int
	patterns [][]rune
)

func main() {
	defer stdout.Flush()

	text = readrs()
	m = readi()
	for i := 0; i < m; i++ {
		patterns = append(patterns, readrs())
	}

	pma := NewPMA(patterns, 'A')
	M := pma.Match(text)

	ans := 0
	for _, A := range M {
		ans += len(A)
	}

	fmt.Println(ans)
}

const (
	_PMA_CHAR_SIZE = 26
)

type PMA struct {
	root *pmaNode
	base rune
}

type pmaNode struct {
	next    [_PMA_CHAR_SIZE]*pmaNode
	accept  []int
	failure *pmaNode
}

// NewPMA returns Pattern Matching Automaton by pattern strings (dictionary).
func NewPMA(patterns [][]rune, base rune) *PMA {
	root := new(pmaNode)
	root.failure = root

	pma := new(PMA)
	pma.root = root
	pma.base = base

	// build Trie by pattern strings
	var curNode *pmaNode
	for i := 0; i < len(patterns); i++ {
		curNode = root
		for _, r := range patterns[i] {
			c := r - base
			if curNode.next[c] == nil {
				curNode.next[c] = new(pmaNode)
			}
			curNode = curNode.next[c]
		}
		curNode.accept = append(curNode.accept, i)
	}

	// Aho-Corasick method
	que := []*pmaNode{}
	// Initialize queue
	for i := 0; i < _PMA_CHAR_SIZE; i++ {
		if root.next[i] == nil {
			root.next[i] = root
		} else {
			root.next[i].failure = root
			que = append(que, root.next[i])
		}
	}
	// BFS
	for len(que) > 0 {
		curNode = que[0]
		que = que[1:]

		for i := 0; i < _PMA_CHAR_SIZE; i++ {
			if curNode.next[i] == nil {
				continue
			}

			beforeFailNode := curNode.failure
			for beforeFailNode.next[i] == nil {
				beforeFailNode = beforeFailNode.failure
			}

			curNode.next[i].failure = beforeFailNode.next[i]
			curNode.next[i].accept = pma._setUnion(curNode.next[i].accept, beforeFailNode.next[i].accept)

			que = append(que, curNode.next[i])
		}
	}

	return pma
}

// Match returns all matched patterns in the text.
// res[i] has word indices that matches the last character in i index (0-index).
// res[i] can have multiple word indices, and in that case,
//  its order is not deterministic.
func (pma *PMA) Match(text []rune) [][]int {
	res := make([][]int, len(text))

	curNode := pma.root
	for i, r := range text {
		c := r - pma.base

		// use the failure link
		for curNode.next[c] == nil {
			curNode = curNode.failure // like an epsilon transition
		}
		curNode = curNode.next[c] // consume a character

		// check pattern match
		tmp := make([]int, len(curNode.accept))
		copy(tmp, curNode.accept)
		res[i] = tmp
	}

	return res
}

func (pma *PMA) _setUnion(A, B []int) []int {
	res := []int{}

	memo := map[int]bool{}
	for _, a := range A {
		memo[a] = true
	}
	for _, b := range B {
		memo[b] = true
	}

	for k := range memo {
		res = append(res, k)
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

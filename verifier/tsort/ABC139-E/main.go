/*
URL:
https://atcoder.jp/contests/abc139/tasks/abc139_e
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
	n int
	A [][]int

	G [1000000 + 50][]int
)

func main() {
	n = readi()
	for i := 0; i < n; i++ {
		row := readis(n - 1)
		for i := 0; i < len(row); i++ {
			row[i]--
		}
		A = append(A, row)
	}

	for i := 0; i < n; i++ {
		for j := 1; j < len(A[i]); j++ {
			bef := A[i][j-1]
			aft := A[i][j]

			cid := Min(i, bef) + n*Max(i, bef)
			nid := Min(i, aft) + n*Max(i, aft)
			G[cid] = append(G[cid], nid)
		}
	}

	num := n * n
	ok, sorted := TSort(num, G[:num])

	if !ok {
		fmt.Println(-1)
		return
	} else {
		ans, _ := LongestPath(sorted, G[:num])
		fmt.Println(ans + 1)
	}
}

// TSort returns a node ids list in topological order.
// node id is 0-based.
// Time complexity: O(|E| + |V|)
func TSort(nn int, AG [][]int) (ok bool, tsortedIDs []int) {
	tsortedIDs = []int{}

	inDegrees := make([]int, nn)
	for s := 0; s < nn; s++ {
		for _, t := range AG[s] {
			inDegrees[t]++
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

		for _, nid := range AG[cid] {
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
// This function assumes that all costs of edges are 1.
// Time complexity: O(|E| + |V|)
func LongestPath(tsortedIDs []int, AG [][]int) (maxLength int, dp []int) {
	_chmax := func(updatedValue *int, target int) bool {
		if *updatedValue < target {
			*updatedValue = target
			return true
		}
		return false
	}

	dp = make([]int, len(tsortedIDs)+1)

	for i := 0; i < len(tsortedIDs); i++ {
		cid := tsortedIDs[i]
		for _, nid := range AG[cid] {
			_chmax(&dp[nid], dp[cid]+1)
		}
	}

	maxLength = 0
	for i := 0; i < len(tsortedIDs); i++ {
		_chmax(&maxLength, dp[i])
	}

	return maxLength, dp
}

// Max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Max(integers ...int) int {
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

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int) int {
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
)

// modi can calculate a right residual whether value is positive or negative.
func modi(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// modll can calculate a right residual whether value is positive or negative.
func modll(val, m int64) int64 {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	reads = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
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

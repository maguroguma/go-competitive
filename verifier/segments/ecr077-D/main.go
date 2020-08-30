/*
URL:
https://codeforces.com/contest/1260/problem/D
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	m, n, k, t int
	A          []int
	L          []Trap
)

type Trap struct {
	l, r, d int
}

func main() {
	m, n, k, t = readi4()
	A = readis(m)
	// L = make(TrapList, 0)
	for i := 0; i < k; i++ {
		l, r, d := readi3()
		// L = append(L, &Trap{key: l, l: l, r: r, d: d})
		L = append(L, Trap{l, r, d})
	}
	maxAgility := Max(A...)

	// 区間の左端で昇順ソート
	// sort.Stable(byKey{L})
	sort.Slice(L, func(i, j int) bool { return L[i].l < L[j].l })

	// m は中央を意味する何らかの値
	isOK := func(m int) bool {
		if C(m) {
			return true
		}
		return false
	}

	ng, ok := -1, maxAgility+1
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	minAgility := ok

	num := 0
	for i := 0; i < m; i++ {
		if A[i] >= minAgility {
			num++
		}
	}
	fmt.Println(num)
}

func C(m int) bool {
	// segments := []Trap{}
	// l, r := 0, -1
	// for i := 0; i < len(L); i++ {
	// 	t := L[i]
	// 	if t.d <= m {
	// 		continue
	// 	}

	// 	if r == -1 {
	// 		l, r = t.l, t.r
	// 		continue
	// 	}

	// 	if r >= t.l-1 {
	// 		// マージして継続
	// 		// r = t.r
	// 		ChMax(&r, t.r)
	// 	} else {
	// 		// マージせず中断して追加
	// 		segments = append(segments, Trap{l: l, r: r})
	// 		l, r = t.l, t.r
	// 	}
	// }
	// if r != -1 {
	// 	segments = append(segments, Trap{l: l, r: r})
	// }

	S := [][2]int{}
	for i := 0; i < len(L); i++ {
		if L[i].d <= m {
			continue
		}
		S = append(S, [2]int{L[i].l, L[i].r})
	}

	segments := MergeSegments(S)

	time := 1 + n
	for _, seg := range segments {
		// time += 2 * (seg.r - seg.l + 1)
		time += 2 * (seg[1] - seg[0] + 1)
	}
	// fmt.Printf("time: %d\n", time)

	return time <= t
}

// MergeSegments returns merged source segments.
// Segments(Sections) are closed section.
// e.g.: [l, r], [l', r'] are merged if r >= l'
func MergeSegments(srcSegments [][2]int) [][2]int {
	_chmax := func(updatedValue *int, target int) bool {
		if *updatedValue < target {
			*updatedValue = target
			return true
		}
		return false
	}

	res := [][2]int{}
	isInitialized := false

	// current segment
	curL, curR := 0, 0

	// sort asc by LEFT coordinate
	sort.Slice(srcSegments, func(i, j int) bool {
		return srcSegments[i][0] < srcSegments[j][0]
	})

	for i := 0; i < len(srcSegments); i++ {
		seg := srcSegments[i]

		if !isInitialized {
			curL, curR = seg[0], seg[1]
			isInitialized = true
			continue
		}

		if curR >= seg[0] {
			// merge and continue
			_chmax(&curR, seg[1])
		} else {
			// do not merge, and add it to result
			res = append(res, [2]int{curL, curR})
			curL, curR = seg[0], seg[1]
		}
	}

	if isInitialized {
		res = append(res, [2]int{curL, curR})
	}

	return res
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

// modint can calculate a right residual whether value is positive or negative.
func modint(val, m int) int {
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

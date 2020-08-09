/*
URL:
https://codeforces.com/contest/1393/problem/C
*/

package main

import (
	"bufio"
	"container/heap"
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
	t int

	n    int
	A    []int
	cnts []int
)

func main() {
	t = ReadInt()
	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		A = ReadIntSlice(n)
		cnts = make([]int, n+5)

		solve()
	}
}

func solve() {
	for _, a := range A {
		cnts[a]++
	}

	ok := BinarySearch(0, n, func(x int) bool {
		B := []*Element{}

		pq := NewElementPQ()
		for i := 1; i <= n; i++ {
			if cnts[i] >= 1 {
				pq.push(&Element{val: i, num: cnts[i]})
			}
		}

		C := []int{}
		for pq.Len() > 0 {
			resi := x + 1

			for resi > 0 {
				if pq.Len() == 0 {
					// return false
					break
				}

				pop := pq.pop()
				C = append(C, pop.val)
				pop.num--
				if pop.num > 0 {
					B = append(B, pop)
				}
				resi--
			}

			for _, e := range B {
				pq.push(e)
			}
			B = []*Element{}
		}

		bef := make([]int, n+5)
		for i := 1; i <= n; i++ {
			bef[i] = -1
		}

		// PrintfDebug("C: %v, x: %d\n", C, x)
		for i := 0; i < len(C); i++ {
			if bef[C[i]] == -1 {
				bef[C[i]] = i
				continue
			}

			dist := i - bef[C[i]] - 1
			// PrintfDebug("dist: %d\n", dist)
			if dist < x {
				return false
			}
			bef[C[i]] = i
		}
		// PrintfDebug("bef: %v\n", bef)
		// PrintfDebug("OK\n")

		return true
	})

	fmt.Println(ok)
}

func BinarySearch(initOK, initNG int, isOK func(mid int) bool) (ok int) {
	ng := initNG
	ok = initOK
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

type Element struct {
	pri      int
	val, num int
}
type ElementPQ []*Element

// Interfaces
func NewElementPQ() *ElementPQ {
	temp := make(ElementPQ, 0)
	pq := &temp
	heap.Init(pq)

	return pq
}
func (pq *ElementPQ) push(target *Element) {
	heap.Push(pq, target)
}
func (pq *ElementPQ) pop() *Element {
	return heap.Pop(pq).(*Element)
}

func (pq ElementPQ) Len() int { return len(pq) }
func (pq ElementPQ) Less(i, j int) bool {
	if pq[i].num > pq[j].num {
		return true
	} else if pq[i].num < pq[j].num {
		return false
	} else {
		return pq[i].val < pq[j].val
	}
} // <: ASC, >: DESC
func (pq ElementPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *ElementPQ) Push(x interface{}) {
	item := x.(*Element)
	*pq = append(*pq, item)
}
func (pq *ElementPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// func solve() {
// 	for _, a := range A {
// 		cnts[a]++
// 	}
// 	PrintfDebug("cnts: %v\n", cnts)

// 	C := []int{}
// 	for i := 1; i <= n; i++ {
// 		C = append(C, cnts[i])
// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(C)))

// 	if C[0] != C[1] {
// 		// 個数について単独首位がある
// 		cnt := C[0]
// 		ans := (n - cnt) / (cnt - 1)
// 		fmt.Println(ans)
// 		return
// 	}

// 	comp, cc := RunLengthEncoding(C)
// 	fv, sv := comp[0], comp[1]
// 	fk, sk := cc[0], cc[1]
// 	PrintfDebug("fv: %d, fk: %d, sv: %d, sk: %d\n", fv, fk, sv, sk)

// 	ans := fk - 1
// 	if fv-sv >= 2 {
// 		fmt.Println(ans)
// 	} else {
// 		fmt.Println(ans + sk)
// 	}
// }

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

// ChMax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
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

/*
URL:
https://atcoder.jp/contests/agc040/tasks/agc040_b
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

type Segment struct {
	id, l, r int
}

var (
	n int

	A    []Segment
	L, R []Segment
	M    []Segment
)

func main() {
	n = ReadInt()
	for i := 0; i < n; i++ {
		l, r := ReadInt2()
		A = append(A, Segment{id: i, l: l, r: r})
		L = append(L, Segment{id: i, l: l, r: r})
		R = append(R, Segment{id: i, l: l, r: r})
		M = append(M, Segment{id: i, l: l, r: r})
	}

	// Lは左端の降順, Rは右端の昇順, Mは長さの降順
	sort.SliceStable(L, func(i, j int) bool { return L[i].l > L[j].l })
	sort.SliceStable(R, func(i, j int) bool { return R[i].r < R[j].r })
	sort.SliceStable(M, func(i, j int) bool { return M[i].r-M[i].l > M[j].r-M[j].l })

	if L[0].id == R[0].id {
		ans := (L[0].r - L[0].l + 1) + (M[0].r - M[0].l + 1)
		fmt.Println(ans)
		return
	}

	// lid, rid := L[0].id, R[0].id
	ans := Max(R[0].r-L[0].l+1, 0) + (M[0].r - M[0].l + 1) // グループ片方を最小にするパターン

	maxL, minR := L[0].l, R[0].r
	// BはAからlid, ridの区間を除外したもの
	B := []Item{}
	for i := 0; i < n; i++ {
		// if i == lid || i == rid {
		// 	continue
		// }
		B = append(B, Item{a: Max(A[i].r-maxL+1, 0), b: Max(minR-A[i].l+1, 0)})
	}
	sort.SliceStable(B, func(i, j int) bool {
		if B[i].a < B[j].a {
			return true
		} else if B[i].a > B[j].a {
			return false
		} else {
			return B[i].b > B[j].b
		}
	})

	// minS, minT := A[lid].r-A[lid].l+1, A[rid].r-A[rid].l+1

	// T := make([]int, len(B)+1)
	// // i番目以降を集合Tに含める時のminTをすべて計算
	// T[len(B)] = minT
	// for i := len(B) - 1; i >= 0; i-- {
	// 	T[i] = Min(T[i+1], B[i].b)
	// }

	// PrintfDebug("lid, rid: %d, %d\n", lid, rid)
	// PrintfDebug("B: %v\n", B)
	// PrintfDebug("T: %v\n", T)

	// ChMax(&ans, minS+T[0])
	// for i := 0; i < len(B); i++ {
	// 	// iを含めてiまでを集合Sに含める
	// 	minS = Min(minS, B[i].a)
	// 	t := T[i+1]

	// 	PrintfDebug("minS: %d\n", minS)
	// 	ChMax(&ans, minS+t)
	// }

	PrintfDebug("B: %v\n", B)
	T := make([]int, n)
	T[n-1] = B[n-1].a
	for i := n - 2; i >= 0; i-- {
		T[i] = Min(T[i+1], B[i].a)
	}
	minS := INF_BIT60
	for i := 0; i < n-1; i++ {
		minS = Min(minS, B[i].b)
		t := T[i+1]
		ChMax(&ans, minS+t)
	}

	fmt.Println(ans)
}

type Item struct {
	a, b int
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

// ChMin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
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

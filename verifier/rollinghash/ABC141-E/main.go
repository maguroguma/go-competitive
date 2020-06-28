/*
URL:
https://atcoder.jp/contests/abc141/tasks/abc141_e
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
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
	n int
	S []rune
)

func main() {
	n = ReadInt()
	S = ReadRuneSlice()

	InitRollingHashConfing()

	ans := 0
	rlh := NewRollingHash(string(S))
	for i := 0; i < len(S); i++ {
		for j := i; j < len(S); j++ {
			for ans < j-i && j+ans < n && rlh.Slice(i, ans+1) == rlh.Slice(j, ans+1) {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

// rolling hash
// reference: https://atcoder.jp/contests/abc141/submissions/7717102

const (
	MASK30       uint64 = (1 << 30) - 1
	MASK31       uint64 = (1 << 31) - 1
	R_MOD        uint64 = (1 << 61) - 1
	POSITIVIZER  uint64 = R_MOD * ((1 << 3) - 1)
	MAX_S_LENGTH        = 200000 + 50
)

var (
	Base    uint64
	PowMemo []uint64
)

type RollingHash struct {
	hash []uint64
}

func InitRollingHashConfing() {
	rand.Seed(time.Now().UnixNano())

	Base = uint64(rand.Int31n(math.MaxInt32-129)) + uint64(129)
	PowMemo = make([]uint64, MAX_S_LENGTH)
	PowMemo[0] = 1
	for i := 1; i < len(PowMemo); i++ {
		PowMemo[i] = CalcMod(Mul(PowMemo[i-1], Base))
	}
}

func NewRollingHash(s string) *RollingHash {
	rlh := new(RollingHash)

	rlh.hash = make([]uint64, len(s)+1)
	for i := 0; i < len(s); i++ {
		rlh.hash[i+1] = CalcMod(Mul(rlh.hash[i], Base) + uint64(s[i]))
	}

	return rlh
}

func (rlh *RollingHash) Slice(begin, length int) uint64 {
	return CalcMod(rlh.hash[begin+length] + POSITIVIZER - Mul(rlh.hash[begin], PowMemo[length]))
}

func Mul(l, r uint64) uint64 {
	var lu uint64 = l >> 31
	var ld uint64 = l & MASK31
	var ru uint64 = r >> 31
	var rd uint64 = r & MASK31
	var middleBit uint64 = ld*ru + lu*rd

	return ((lu * ru) << 1) + ld*rd + ((middleBit & MASK30) << 31) + (middleBit >> 30)
}

func CalcMod(val uint64) uint64 {
	val = (val & R_MOD) + (val >> 61)
	if val > R_MOD {
		val -= R_MOD
	}
	return val
}

func solve() {
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

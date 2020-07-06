/*
URL:
https://atcoder.jp/contests/abc173/tasks/abc173_e
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

var (
	n, k int
	A    []int

	posi, nega, zero int
	P, N             []int
)

func main() {
	n, k = ReadInt2()
	A = ReadIntSlice(n)

	if n == k {
		sub()
		return
	}

	for i := 0; i < n; i++ {
		if A[i] > 0 {
			posi++
		} else if A[i] < 0 {
			nega++
		} else {
			zero++
		}
	}

	if posi == 0 && k%2 == 1 {
		subsub()
		return
	}

	if posi+nega < k {
		fmt.Println(0)
		return
	}

	for i := 0; i < n; i++ {
		if A[i] > 0 {
			P = append(P, A[i])
		} else if A[i] < 0 {
			N = append(N, A[i])
		}
	}
	lp, ln := len(P), len(N)

	sort.Slice(P, func(i, j int) bool {
		return P[i] > P[j]
	})
	sort.Slice(N, func(i, j int) bool {
		return N[i] < N[j]
	})

	ans := 1
	nokori := k
	pid, nid := 0, 0
	for nokori > 0 {
		if pid >= lp {
			ans *= N[nid]
			ans = NegativeMod(ans, MOD)

			nid++
			nokori--
			continue
		}
		if nid >= ln {
			ans *= P[pid]
			ans = NegativeMod(ans, MOD)

			pid++
			nokori--
			continue
		}

		if P[pid] > -N[nid] {
			ans *= P[pid]
			pid++
		} else {
			ans *= N[nid]
			nid++
		}
		ans = NegativeMod(ans, MOD)
		nokori--
	}

	if nid%2 == 0 {
		fmt.Println(ans)
		return
	}

	if pid >= lp && nid >= ln {
		fmt.Println(0)
		return
	}

	if pid >= lp || nid == 0 {
		ans *= ModInv(P[pid-1], MOD)
		ans = NegativeMod(ans, MOD)
		ans *= N[nid]
		ans = NegativeMod(ans, MOD)
		fmt.Println(ans)
		return
	}
	if nid >= ln || pid == 0 {
		ans *= ModInv(N[nid-1], MOD)
		ans = NegativeMod(ans, MOD)
		ans *= P[pid]
		ans = NegativeMod(ans, MOD)
		fmt.Println(ans)
		return
	}

	// 答えは現在負数で正負をそれぞれ少なくとも1つ含んでいる
	p1, n1 := P[pid-1], -N[nid-1]
	p2, n2 := P[pid], -N[nid]
	// if n2/p1 > p2/n1
	if n1*n2 > p1*p2 {
		ans *= ModInv(p1, MOD)
		ans = NegativeMod(ans, MOD)
		ans *= -n2
		ans = NegativeMod(ans, MOD)
	} else {
		ans *= -1
		ans = NegativeMod(ans, MOD)
		ans *= ModInv(n1, MOD)
		ans = NegativeMod(ans, MOD)
		ans *= p2
		ans = NegativeMod(ans, MOD)
	}
	fmt.Println(ans)
}

// ModInv returns $a^{-1} mod m$ by Fermat's little theorem.
// O(1), but C is nearly equal to 30 (when m is 1000000000+7).
func ModInv(a, m int) int {
	return modpow(a, m-2, m)
}

func modpow(a, e, m int) int {
	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := modpow(a, halfE, m)
		return half * half % m
	}

	return a * modpow(a, e-1, m) % m
}

func subsub() {
	B := make([]int, n)
	for i := 0; i < n; i++ {
		B[i] = A[i]
	}

	sort.Slice(B, func(i, j int) bool {
		return AbsInt(B[i]) < AbsInt(B[j])
	})

	ans := 1
	for i := 0; i < k; i++ {
		ans *= B[i]
		ans = NegativeMod(ans, MOD)
	}

	fmt.Println(ans)
}

func sub() {
	ans := 1
	for i := 0; i < n; i++ {
		ans *= A[i]
		ans = NegativeMod(ans, MOD)
	}
	fmt.Println(ans)
}

// NegativeMod can calculate a right residual whether value is positive or negative.
func NegativeMod(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

/*
URL:
https://atcoder.jp/contests/abc160/tasks/abc160_f
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

type Value struct {
	num, size int
}

var (
	n int

	cf      *CombFactorial
	G       [200000 + 50][]int
	dp      [200000 + 50]Value
	answers [200000 + 50]Value
)

func main() {
	n = ReadInt()
	for i := 0; i < n-1; i++ {
		a, b := ReadInt2()
		a--
		b--

		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	cf = NewCombFactorial(300000 + 50)

	dfs(0, -1)
	PrintfDebug("%v\n", dp[:n])

	rdfs(0, -1, Value{num: 1, size: 0})
	for i := 0; i < n; i++ {
		fmt.Println(answers[i].num)
	}
}

func rdfs(cid, pid int, dpar Value) {
	// cidの計算
	num, size := 1, 1
	for _, nid := range G[cid] {
		if nid == pid {
			num *= dpar.num
			num %= MOD
			num *= cf.invFactorial[dpar.size]
			num %= MOD
			size += dpar.size
			continue
		}
		num *= dp[nid].num
		num %= MOD
		num *= cf.invFactorial[dp[nid].size]
		num %= MOD
		size += dp[nid].size
	}
	num *= cf.factorial[size-1]
	num %= MOD
	answers[cid].num, answers[cid].size = num, size

	// 子方向に対してdparを計算しながら進行する
	cdp := []Value{Value{num: 1, size: 0}}
	nexts := []int{-1}
	for _, nid := range G[cid] {
		if nid == pid {
			continue
		}
		cdp = append(cdp, dp[nid])
		nexts = append(nexts, nid)
	}
	cdp = append(cdp, Value{num: 1, size: 0})
	nexts = append(nexts, -1)

	l := len(nexts)
	L := make([]Value, l)
	L[0], L[l-1] = Value{num: 1, size: 0}, Value{num: 1, size: 0}
	R := make([]Value, l)
	R[0], R[l-1] = Value{num: 1, size: 0}, Value{num: 1, size: 0}
	for i := 1; i <= l-2; i++ {
		L[i] = merge(L[i-1], cdp[i])
	}
	for i := l - 2; i >= 1; i-- {
		R[i] = merge(R[i+1], cdp[i])
	}

	for i := 1; i <= l-2; i++ {
		nid := nexts[i]
		v := merge(L[i-1], R[i+1])
		v = merge(v, dpar)

		v.size++
		// v.num *= cf.factorial[v.size-1]
		// v.num %= MOD
		rdfs(nid, cid, v)
	}
}

func merge(left, right Value) Value {
	num, size := 1, 0

	num *= left.num
	num %= MOD
	num *= right.num
	num %= MOD
	num *= cf.invFactorial[left.size]
	num %= MOD
	num *= cf.invFactorial[right.size]
	num %= MOD

	size += left.size
	size += right.size

	num *= cf.factorial[size]
	num %= MOD

	return Value{num: num, size: size}
}

func dfs(cid, pid int) {
	size := 1
	num := 1

	for _, nid := range G[cid] {
		if nid == pid {
			continue
		}

		dfs(nid, cid)

		num *= dp[nid].num
		num %= MOD
		num *= cf.invFactorial[dp[nid].size]
		num %= MOD

		size += dp[nid].size
	}

	num *= cf.factorial[size-1]
	num %= MOD

	dp[cid].num, dp[cid].size = num, size
}

// cf := NewCombFactorial(2000000) // maxNum == "maximum n" * 2 (for H(n,r))
// res := cf.C(n, r) 	// 組み合わせ
// res := cf.H(n, r) 	// 重複組合せ
// res := cf.P(n, r) 	// 順列

type CombFactorial struct {
	factorial, invFactorial []int
	maxNum                  int
}

func NewCombFactorial(maxNum int) *CombFactorial {
	cf := new(CombFactorial)
	cf.maxNum = maxNum
	cf.factorial = make([]int, maxNum+50)
	cf.invFactorial = make([]int, maxNum+50)
	cf.initCF()

	return cf
}
func (c *CombFactorial) modInv(a int) int {
	return c.modpow(a, MOD-2)
}
func (c *CombFactorial) modpow(a, e int) int {
	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := c.modpow(a, halfE)
		return half * half % MOD
	}

	return a * c.modpow(a, e-1) % MOD
}
func (c *CombFactorial) initCF() {
	var i int

	for i = 0; i <= c.maxNum; i++ {
		if i == 0 {
			c.factorial[i] = 1
			c.invFactorial[i] = c.modInv(c.factorial[i])
			continue
		}

		num := i * c.factorial[i-1]
		num %= MOD
		c.factorial[i] = num
		c.invFactorial[i] = c.modInv(c.factorial[i])
	}
}
func (c *CombFactorial) C(n, r int) int {
	var res int

	res = 1
	res *= c.factorial[n]
	res %= MOD
	res *= c.invFactorial[r]
	res %= MOD
	res *= c.invFactorial[n-r]
	res %= MOD

	return res
}
func (c *CombFactorial) P(n, r int) int {
	var res int

	res = 1
	res *= c.factorial[n]
	res %= MOD
	res *= c.invFactorial[n-r]
	res %= MOD

	return res
}
func (c *CombFactorial) H(n, r int) int {
	return c.C(n-1+r, r)
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

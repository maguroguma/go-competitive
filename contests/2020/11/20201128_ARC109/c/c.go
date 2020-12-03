/*
URL:
https://atcoder.jp/contests/arc109/tasks/arc109_c
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

var (
	n, k int
	S    []rune

	T []rune

	memo [30][30]bool

	dp   [100 + 5][100 + 5]rune
	inv2 int
)

const (
	r, p, s = 'R', 'P', 'S'
)

func main() {
	defer stdout.Flush()

	memo[r-'A'][s-'A'] = true
	memo[s-'A'][p-'A'] = true
	memo[p-'A'][r-'A'] = true

	n, k = readi2()
	S = readrs()

	for i := 0; i < n; i++ {
		S = append(S, S[i])
	}
	debugf("S: %v\n", string(S))

	for c := 0; c < 2; c++ {
		for i := 0; i < n; i++ {
			T = append(T, S[i])
		}
	}

	for c := 0; c < k; c++ {
		for i := 0; i+1 < 2*n; i += 2 {
			if S[i] == S[i+1] {
				T[i/2] = S[i]
			} else {
				a, b := S[i], S[i+1]
				if memo[a-'A'][b-'A'] {
					T[i/2] = a
				} else {
					T[i/2] = b
				}
			}
		}
		// S = T

		S = []rune{}
		for ct := 0; ct < 2; ct++ {
			for i := 0; i < n; i++ {
				S = append(S, T[i])
			}
		}
	}

	printf("%c\n", S[0])
}

// func main() {
// 	defer stdout.Flush()

// 	memo[r-'A'][s-'A'] = true
// 	memo[s-'A'][p-'A'] = true
// 	memo[p-'A'][r-'A'] = true

// 	n, k = readi2()
// 	S = readrs()

// 	if n == 1 {
// 		printf("%c\n", S[0])
// 		return
// 	}
// 	if n == 2 {
// 		L, R := S[0], S[1]
// 		var ans rune
// 		if L == R {
// 			ans = L
// 		} else {
// 			if memo[L-'A'][R-'A'] {
// 				ans = L
// 			} else {
// 				ans = R
// 			}
// 		}
// 		printf("%c\n", ans)
// 		return
// 	}

// 	inv2 = ModInv(2, n)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			// if i != j {
// 			// 	dp[i][j] = 'x'
// 			// } else {
// 			// 	dp[i][j] = S[i]
// 			// }
// 			dp[i][j] = 'x'
// 		}
// 	}

// 	ans := rec(0, (modpow(2, k, n)-1)%n, k)

// 	// debugf("dp[2][0]: %c\n", dp[2][0])

// 	printf("%c\n", ans)
// }

// l, r は mod n
func rec(l, r int, ck int) rune {
	if dp[l][r] != 'x' {
		return dp[l][r]
	}
	// if mod(r-l, n) == 0 {
	if r == l {
		// r %= n
		// l %= n
		// if ck == 0 {
		dp[l][r] = S[l]
		debugf("dp[%d][%d]: %c\n", l, r, dp[l][r])
		return dp[l][r]
	}

	m := mod(l+r, n)
	m *= inv2
	m %= n

	L := rec(l, (m-1)%n, ck-1)
	R := rec(m, r, ck-1)

	if L == R {
		dp[l][r] = L
	} else {
		if memo[L-'A'][R-'A'] {
			dp[l][r] = L
		} else {
			dp[l][r] = R
		}
	}

	return dp[l][r]
}

const (
	// MAX_LOG  = 70
	MAX_LOG  = 120
	MAX_NODE = 200000
)

type DoublingSolver struct {
	next [MAX_LOG + 5][MAX_NODE + 50]int
	N    int
}

func NewDoublingSolver(A []int, n int) *DoublingSolver {
	ds := new(DoublingSolver)
	ds.N = n

	for v := 0; v < ds.N; v++ {
		ds.next[0][v] = A[v]
	}

	for d := 0; d+1 < MAX_LOG; d++ {
		for v := 0; v < ds.N; v++ {
			ds.next[d+1][v] = ds.next[d][ds.next[d][v]]
		}
	}

	return ds
}

func (ds *DoublingSolver) Jump(sid, k int) (gid int) {
	gid = sid

	for d := 0; d+1 < MAX_LOG; d++ {
		if (k >> uint(d) & 1) == 1 {
			gid = ds.next[d][gid]
		}
	}

	return gid
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

// abs is integer version of math.Abs
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// pow is integer version of math.Pow
// pow calculate a power by Binary Power (二分累乗法(O(log e))).
func pow(a, e int) int {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}

	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := pow(a, halfE)
		return half * half
	}

	return a * pow(a, e-1)
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

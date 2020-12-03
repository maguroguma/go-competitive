/*
URL:
https://atcoder.jp/contests/dwango2015-prelims/tasks/dwango2015_prelims_3
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

var (
	n int

	dp [100 + 5]float64
	C  [100 + 5][100 + 5][100 + 5]float64
)

func main() {
	defer stdout.Flush()

	n = readi()

	for i := 0; i <= n; i++ {
		dp[i] = -1.0
	}
	dp[1] = 0.0

	// debugf("P(1, 1, 1): %v\n", P(1, 1, 1))
	initP()

	fmt.Println(rec(n))
}

func rec(i int) float64 {
	if dp[i] >= 0.0 {
		return dp[i]
	}

	// dp[i] = 0.0

	// i -> iの確率
	pii := 0.0
	pii += P(i, 0, 0)
	pii += P(0, i, 0)
	pii += P(0, 0, i)
	if i%3 == 0 {
		pii += P(i/3, i/3, i/3)
	}
	// debugf("pii: %v\n", pii)

	// e := 0.0
	// for j := 1; j < i; j++ {
	// 	pij := 0.0
	// 	// jが単独最小
	// 	for k := j + 1; k+j <= i; k++ {
	// 		l := i - j - k
	// 		// if l != 0 && l <= j {
	// 		if 0 < l && l <= j {
	// 			continue
	// 		}
	// 		debugf("%d %d %d\n", j, k, l)
	// 		pij += P(j, k, l)
	// 		pij += P(j, l, k)
	// 		pij += P(k, j, l)
	// 		pij += P(l, j, k)
	// 		pij += P(k, l, j)
	// 		pij += P(l, k, j)
	// 	}
	// 	// jが2つで並ぶ
	// 	k := j
	// 	l := i - j - k
	// 	if l > j || l == 0 {
	// 		pij += P(j, k, l)
	// 		pij += P(l, j, k)
	// 		pij += P(k, l, j)
	// 	}
	// 	e += pij * (rec(j) + 1.0)
	// }

	e := 0.0
	for j := 0; j < i; j++ {
		for k := 0; k < i; k++ {
			l := i - (j + k)
			if l < 0 {
				continue
			}
			A := []int{j, k, l}
			sort.Sort(sort.IntSlice(A))

			if A[0] == 0 && A[1] == 0 {
				continue
			}
			if A[0] == A[2] {
				continue
			}
			nx := A[0]
			if nx == 0 {
				nx = A[1]
			}
			e += P(j, k, l) * (rec(nx) + 1.0)
		}
	}

	dp[i] = (pii + e) / (1 - pii)

	return dp[i]
}

func initP() {
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 0; k <= 100; k++ {
				C[i][j][k] = -1.0
			}
		}
	}
	C[0][0][0] = 1.0

	var dfs func(i, j, k int) float64
	dfs = func(i, j, k int) float64 {
		if C[i][j][k] >= 0.0 {
			return C[i][j][k]
		}

		C[i][j][k] = 0.0
		if i-1 >= 0 {
			C[i][j][k] += dfs(i-1, j, k) / 3.0
		}
		if j-1 >= 0 {
			C[i][j][k] += dfs(i, j-1, k) / 3.0
		}
		if k-1 >= 0 {
			C[i][j][k] += dfs(i, j, k-1) / 3.0
		}

		return C[i][j][k]
	}

	dfs(100, 100, 100)
}

// func P(r, s, p int) float64 {
// 	return C[r][s][p]
// }

func P(r, s, p int) float64 {
	res := 1.0

	N := r + s + p
	for i := 1; i <= N; i++ {
		res *= float64(i) / 3.0
	}

	for i := 1; i <= r; i++ {
		res /= float64(i)
	}
	for i := 1; i <= p; i++ {
		res /= float64(i)
	}
	for i := 1; i <= s; i++ {
		res /= float64(i)
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

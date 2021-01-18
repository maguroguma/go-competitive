/*
URL:
https://atcoder.jp/contests/past201912-open/tasks/past201912_o
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
	println = fmt.Println

	n int
	A [][]int

	B  [][]int
	dp [180000 + 50]float64
	D  [180000 + 50]float64
	I  [180000 + 50]int
)

func main() {
	defer stdout.Flush()

	n = readi()
	for i := 0; i < n; i++ {
		A = append(A, readis(6))
	}

	comp := NewCompress()
	for i := 0; i < n; i++ {
		comp.Add(A[i]...)
	}
	comp.Build()

	B = make([][]int, n)
	for i := 0; i < n; i++ {
		B[i] = make([]int, 6)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			B[i][j] = comp.Get(A[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			I[B[i][j]] = i
		}
	}
	// debugf("I: %v\n", I[:6*n])

	// ここからDP

	N := 6 * n
	f := 1.0 / 6.0
	for i := 0; i < N; i++ {
		D[i] = 1.0
	}
	dp[N-1] = 1.0
	idx := I[N-1]
	D[idx] += f*(dp[N-1]+1.0) - f
	cur := math.Max(1.0, D[idx])

	for i := N - 2; i >= 0; i-- {
		dp[i] = cur
		idx := I[i]
		D[idx] += f*(dp[i]+1.0) - f
		cur = math.Max(cur, D[idx])
	}

	// debugf("dp[:N]: %v\n", dp[:N])

	ans := 0.0
	for i := 0; i < n; i++ {
		tmp := 0.0
		for j := 0; j < 6; j++ {
			tmp += (dp[B[i][j]] + 1.0) * f
		}
		ans = math.Max(ans, tmp)
	}

	// for i := 0; i < N; i++ {
	// 	ans += dp[i] + f
	// }

	// println(ans / 6.0)
	println(ans)
}

// NewCompress returns a compress algorithm.
func NewCompress() *Compress {
	c := new(Compress)
	c.xs = []int{}
	c.cs = []int{}

	return c
}

// Add can add any number of elements.
// Time complexity: O(1)
func (c *Compress) Add(X ...int) {
	c.xs = append(c.xs, X...)
}

// Build compresses input elements by sorting.
// Time complexity: O(NlogN)
func (c *Compress) Build() {
	sort.Slice(c.xs, func(i, j int) bool {
		return c.xs[i] < c.xs[j]
	})

	if len(c.xs) == 0 {
		panic("Compress doesn't have any elements")
	}

	c.cs = append(c.cs, c.xs[0])
	for i := 1; i < len(c.xs); i++ {
		if c.xs[i-1] == c.xs[i] {
			continue
		}
		c.cs = append(c.cs, c.xs[i])
	}
}

// Get returns index that is equal to by binary search.
// Results are in [0, len(c.cs)).
// Time complexity: O(logN)
func (c *Compress) Get(x int) int {
	_abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var ng, ok = int(-1), int(len(c.cs))
	for _abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if c.cs[mid] >= x {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

// InvGet returns original value that equals to i (compressed values).
// InvGet accepts [0, len(c.cs))
// Time complexity: O(1)
func (c *Compress) InvGet(i int) int {
	if !(0 <= i && i < int(len(c.cs))) {
		panic("i is out of range")
	}
	return c.cs[i]
}

// Kind returns number of different values, that is len(c.cs).
// Time complexity: O(1)
func (c *Compress) Kind() int {
	return len(c.cs)
}

type Compress struct {
	xs []int // sorted original values
	cs []int // sorted and compressed original values
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

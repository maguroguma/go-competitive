/*
URL:
https://atcoder.jp/contests/arc114/tasks/arc114_a
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
	X []int
)

func main() {
	defer stdout.Flush()

	n = readi()
	X = readis(n)
	sort.Sort(sort.IntSlice(X))

	P := []int{}
	for i := 2; i <= 50; i++ {
		if IsPrime(i) {
			P = append(P, i)
		}
	}

	ans := INF_I64
	BruteForceByBits01(len(P), func(bits []int) {
		val := 1
		for i := 0; i < len(bits); i++ {
			if bits[i] == 1 {
				if IsProductOverflow(val, P[i]) {
					return
				}
				val *= P[i]
			}
		}

		for _, x := range X {
			if Gcd(x, val) == 1 {
				return
			}
		}

		chmin(&ans, val)
	})

	println(ans)
}

// Gcd returns the Greatest Common Divisor of two natural numbers.
// Gcd only accepts two natural numbers (a, b >= 0).
// Negative number causes panic.
// Gcd uses the Euclidean Algorithm.
func Gcd(a, b int) int {
	if a < 0 || b < 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// Lcm returns the Least Common Multiple of two natural numbers.
// Lcd only accepts two natural numbers (a, b >= 0).
// Negative number causes panic.
// Lcd uses the Euclidean Algorithm indirectly.
func Lcm(a, b int) int {
	if a < 0 || b < 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	if a == 0 || b == 0 {
		return 0
	}

	// a = a'*gcd, b = b'*gcd, a*b = a'*b'*gcd^2
	// a' and b' are relatively prime numbers
	// gcd consists of prime numbers, that are included in a and b
	g := Gcd(a, b)

	// not (a * b / gcd), because of reducing a probability of overflow
	return (a / g) * b
}

// IsProductOverflow returns whether a*b > MAX_INT64 or not.
// IsProductOverflow panics when it accepts negative integers.
func IsProductOverflow(a, b int) bool {
	if a < 0 || b < 0 {
		panic("IsProductOverflow does not accept negative integers")
	}

	return a > (math.MaxInt64 / b)
}

// IsSumOverflow returns whether a+b > MAX_INT64 or not.
// IsSumOverflow panics when it accepts negative integers.
func IsSumOverflow(a, b int) bool {
	if a < 0 || b < 0 {
		panic("IsSumOverflow does not accept negative integers")
	}

	return a > (math.MaxInt64 - b)
}

// IsProductLeq returns whether a*b <= ub or not.
// IsProductLeq panics when it accepts negative integers.
func IsProductLeq(a, b, ub int) bool {
	if a < 0 || b < 0 || ub < 0 {
		panic("IsProductLeq does not accept negative integers")
	}

	return a <= (ub / b)
}

// IsSumLeq returns whether a+b <= ub or not.
// IsSumLeq panics when it accepts negative integers.
func IsSumLeq(a, b, ub int) bool {
	if a < 0 || b < 0 || ub < 0 {
		panic("IsSumLeq does not accept negative integers")
	}

	return a <= (ub - b)
}

// IsPrime judges whether an argument integer is a prime number or not.
func IsPrime(n int) bool {
	var i int

	if n == 1 {
		return false
	}

	for i = 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// originated from ktateish@AtCoder
// e.g.: https://atcoder.jp/contests/abc167/submissions/13042361

// BruteForceByBits01(bitsNum, fn) calls fn with []int for each n-bit 0/1 pattern
func BruteForceByBits01(bitsNum int, fn func(bits []int)) {
	// e.g.
	// BruteForceByBits01(10, func(b []int) { fmt.Println(b) }
	// -> [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	//    [1, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	//    [0, 1, 0, 0, 0, 0, 0, 0, 0, 0]
	//    [1, 1, 0, 0, 0, 0, 0, 0, 0, 0]
	//    [0, 0, 1, 0, 0, 0, 0, 0, 0, 0]
	//    ...
	N := 1 << uint(bitsNum)
	a := make([]int, bitsNum)
	for i := 0; i < N; i++ {
		for j := 0; j < bitsNum; j++ {
			k := bitsNum - j - 1
			if i&(1<<uint(j)) == 0 {
				a[k] = 0
			} else {
				a[k] = 1
			}
		}
		fn(a)
	}
}

// BruteForceByBitsTF(bitsNum, fn) calls fn with []bool for each n-bit true/false pattern
func BruteForceByBitsTF(bitsNum int, fn func(bitFlags []bool)) {
	// e.g.
	// BruteForceByBitsTF(10, func(b []bool) { fmt.Println(b) }
	// -> [false, false, false, false, false, false, false, false, false, false]
	//    [true, false, false, false, false, false, false, false, false, false]
	//    [false, true, false, false, false, false, false, false, false, false]
	//    ...
	N := 1 << uint(bitsNum)
	a := make([]bool, bitsNum)
	for i := 0; i < N; i++ {
		for j := 0; j < bitsNum; j++ {
			k := bitsNum - j - 1
			if i&(1<<uint(j)) == 0 {
				a[k] = false
			} else {
				a[k] = true
			}
		}
		fn(a)
	}
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

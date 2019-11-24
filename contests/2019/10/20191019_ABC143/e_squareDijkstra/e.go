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

/*********** I/O ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
)

func init() {
	ReadString = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	// r.Buffer(make([]byte, 1024), int(1e+11)) // for AtCoder
	r.Buffer(make([]byte, 1024), int(1e+9)) // for Codeforces
	// Split sets the split function for the Scanner. The default split function is ScanLines.
	// Split panics if it is called after scanning has started.
	r.Split(bufio.ScanWords)

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

/*********** DP sub-functions ***********/

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

// NthBit returns nth bit value of an argument.
// n starts from 0.
func NthBit(num, nth int) int {
	return num >> uint(nth) & 1
}

// OnBit returns the integer that has nth ON bit.
// If an argument has nth ON bit, OnBit returns the argument.
func OnBit(num, nth int) int {
	return num | (1 << uint(nth))
}

// OffBit returns the integer that has nth OFF bit.
// If an argument has nth OFF bit, OffBit returns the argument.
func OffBit(num, nth int) int {
	return num & ^(1 << uint(nth))
}

// PopCount returns the number of ON bit of an argument.
func PopCount(num int) int {
	res := 0

	for i := 0; i < 70; i++ {
		if ((num >> uint(i)) & 1) == 1 {
			res++
		}
	}

	return res
}

/*********** Arithmetic ***********/

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

// Sum returns multiple integers sum.
func Sum(integers ...int) int {
	s := 0

	for _, i := range integers {
		s += i
	}

	return s
}

// PowInt is integer version of math.Pow
// PowInt calculate a power by Binary Power (二分累乗法(O(log e))).
func PowInt(a, e int) int {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}

	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := PowInt(a, halfE)
		return half * half
	}

	return a * PowInt(a, e-1)
}

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Gcd returns the Greatest Common Divisor of two natural numbers.
// Gcd only accepts two natural numbers (a, b >= 1).
// 0 or negative number causes panic.
// Gcd uses the Euclidean Algorithm.
func Gcd(a, b int) int {
	if a <= 0 || b <= 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}
	if a < b {
		a, b = b, a
	}

	// Euclidean Algorithm
	for b > 0 {
		div := a % b
		a, b = b, div
	}

	return a
}

// Lcm returns the Least Common Multiple of two natural numbers.
// Lcd only accepts two natural numbers (a, b >= 1).
// 0 or negative number causes panic.
// Lcd uses the Euclidean Algorithm indirectly.
func Lcm(a, b int) int {
	if a <= 0 || b <= 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	// a = a'*gcd, b = b'*gcd, a*b = a'*b'*gcd^2
	// a' and b' are relatively prime numbers
	// gcd consists of prime numbers, that are included in a and b
	gcd := Gcd(a, b)

	// not (a * b / gcd), because of reducing a probability of overflow
	return (a / gcd) * b
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

const MOD = 1000000000 + 7
const ALPHABET_NUM = 26
const INF_INT64 = math.MaxInt64
const INF_BIT60 = 1 << 60

var m, l int
var q int
var answers [MAX][MAX]int

func main() {
	n, m, l = ReadInt3()
	for i := 0; i < n; i++ {
		M[i][i] = 0
		for j := i + 1; j < n; j++ {
			M[i][j] = INFTY
			M[j][i] = INFTY
		}
	}
	for i := 0; i < m; i++ {
		a, b, c := ReadInt3()
		a--
		b--
		M[a][b] = c
		M[b][a] = c
	}
	// 答えを前計算しておく
	for i := 0; i < n; i++ {
		res := dijkstra(i)
		for j := i + 1; j < n; j++ {
			answers[i][j] = res[j]
			answers[j][i] = answers[i][j]
		}
	}
	q = ReadInt()
	for i := 0; i < q; i++ {
		s, t := ReadInt2()
		s--
		t--
		fmt.Println(answers[s][t])
	}
}

const (
	MAX   = 500 + 5
	INFTY = 1 << 30
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

var n int
var M [MAX][MAX]int

type Vertex struct {
	times, gas int
}

// O(|V|^2)
// 各ノードの最短路の長さを返す
func dijkstra(s int) []int {
	var minv Vertex
	d, color := make([]Vertex, MAX), make([]int, MAX)

	for i := 0; i < n; i++ {
		d[i] = Vertex{times: INFTY, gas: -1}
		color[i] = WHITE
	}

	// d[0].times, d[0].gas = 0, l
	// color[0] = GRAY
	d[s].times, d[s].gas = 0, l
	color[s] = GRAY
	for {
		minv = Vertex{times: INFTY, gas: -1}
		u := -1
		for i := 0; i < n; i++ {
			if color[i] == BLACK {
				continue
			}
			if minv.times > d[i].times {
				u, minv = i, d[i]
			} else if (minv.times == d[i].times) && (minv.gas < d[i].gas) {
				u, minv = i, d[i]
			}
		}
		if u == -1 {
			break
		}

		color[u] = BLACK

		for v := 0; v < n; v++ {
			// vへの移動を検討
			if color[v] != BLACK && M[u][v] != INFTY {
				dist := M[u][v]
				resGas := d[u].gas
				curTimes := d[u].times

				if resGas >= dist {
					// 補給の必要なし
					if d[v].times > curTimes {
						d[v].times, d[v].gas = curTimes, resGas-dist
					} else if d[v].times == curTimes && d[v].gas < resGas-dist {
						d[v].gas = resGas - dist
					}
				} else {
					// 補給が必要
					// 補給しても通れない場合はパス
					if l < dist {
						continue
					}

					if d[v].times > curTimes+1 {
						d[v].times, d[v].gas = curTimes+1, l-dist
					} else if d[v].times == curTimes+1 && d[v].gas < l-dist {
						d[v].gas = l - dist
					}
				}
			}
		}
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] != BLACK {
			res[i] = -1
		} else {
			res[i] = d[i].times
		}
	}
	return res
}

/*
- MODは最後にとりましたか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
*/

/*******************************************************************/

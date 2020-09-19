/*
URL:
https://atcoder.jp/contests/joi2014ho/tasks/joi2014ho1
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var (
	m, n int
	S    [][]rune
	T    [][]rune

	V [1000 + 5][1000 + 5][5]int
	M map[rune]int
)

const (
	J, O, I = 0, 1, 2
)

func main() {
	defer stdout.Flush()

	M = make(map[rune]int)
	M['J'], M['O'], M['I'] = J, O, I

	m, n = readi2()
	for i := 0; i < m; i++ {
		row := readrs()
		S = append(S, row)
	}
	for i := 0; i < 2; i++ {
		row := readrs()
		T = append(T, row)
	}

	B := [][]rune{
		{'j', 'j'},
		{'j', 'j'},
	}

	ans := 0
	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			B[0][0], B[0][1] = S[i][j], S[i][j+1]
			B[1][0], B[1][1] = S[i+1][j], S[i+1][j+1]
			if check(B) {
				ans++
			}
		}
	}
	debugf("ans: %v\n", ans)

	for _, r := range []rune{'J', 'O', 'I'} {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i-1 >= 0 && j-1 >= 0 {
					// 左上
					for k := 0; k < 2; k++ {
						for l := 0; l < 2; l++ {
							B[k][l] = S[i-1+k][j-1+l]
						}
					}

					if check(B) {
						V[i][j][M[r]]--
					}

					B[1][1] = r
					if check(B) {
						V[i][j][M[r]]++
					}
				}
				if i-1 >= 0 && j+1 < n {
					// 右上
					for k := 0; k < 2; k++ {
						for l := 0; l < 2; l++ {
							B[k][l] = S[i-1+k][j+l]
						}
					}

					if check(B) {
						V[i][j][M[r]]--
					}

					B[1][0] = r
					if check(B) {
						V[i][j][M[r]]++
					}
				}
				if i+1 < m && j-1 >= 0 {
					// 左下
					for k := 0; k < 2; k++ {
						for l := 0; l < 2; l++ {
							B[k][l] = S[i+k][j-1+l]
						}
					}

					if check(B) {
						V[i][j][M[r]]--
					}

					B[0][1] = r
					if check(B) {
						V[i][j][M[r]]++
					}
				}
				if i+1 < m && j+1 < n {
					// 右下
					for k := 0; k < 2; k++ {
						for l := 0; l < 2; l++ {
							B[k][l] = S[i+k][j+l]
						}
					}

					if check(B) {
						V[i][j][M[r]]--
					}

					B[0][0] = r
					if check(B) {
						V[i][j][M[r]]++
					}
				}
			}
		}
	}

	maxi := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 3; k++ {
				ChMax(&maxi, V[i][j][k])
			}
		}
	}

	fmt.Println(ans + maxi)
}

func check(U [][]rune) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if U[i][j] != T[i][j] {
				return false
			}
		}
	}
	return true
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
)

// modi can calculate a right residual whether value is positive or negative.
func modi(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// modll can calculate a right residual whether value is positive or negative.
func modll(val, m int64) int64 {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	reads = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
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

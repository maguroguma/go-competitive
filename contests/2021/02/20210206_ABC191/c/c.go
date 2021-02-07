/*
URL:
https://atcoder.jp/contests/abc191/tasks/abc191_c
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
	println = fmt.Println

	h, w int
	S    [][]rune

	L, R, U, D [20][20]int
)

func main() {
	defer stdout.Flush()

	h, w = readi2()
	for i := 0; i < h; i++ {
		row := readrs()
		S = append(S, row)
	}

	steps := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if S[i][j] == '#' {
				wn := 0

				for _, s := range steps {
					ny, nx := i+s[0], j+s[1]
					if S[ny][nx] == '.' {
						wn++
					}
				}

				if wn == 0 {
					continue
				}

				// 向きをカウントする
				for _, s := range steps {
					ny, nx := i+s[0], j+s[1]
					if S[ny][nx] == '.' {
						if i < ny {
							// 下向き
							D[i][j]++
						} else if i > ny {
							// 上向き
							U[i][j]++
						} else if j < nx {
							// 右向き
							R[j][i]++
						} else {
							// 左向き
							L[j][i]++
						}
					}
				}
			}
		}
	}
	// debugf("L: %v\n", L)
	// debugf("R: %v\n", R)
	// debugf("U: %v\n", U)
	// debugf("D: %v\n", D)

	ans := 0
	for i := 0; i < 15; i++ {
		// if D[i] > 0 {
		// 	ans++
		// }
		// if U[i] > 0 {
		// 	ans++
		// }
		comp, _ := RunLengthEncoding(D[i][:15])
		for _, v := range comp {
			if v == 1 {
				ans++
			}
		}
		comp, _ = RunLengthEncoding(U[i][:15])
		for _, v := range comp {
			if v == 1 {
				ans++
			}
		}
	}
	for j := 0; j < 15; j++ {
		// if R[j] > 0 {
		// 	ans++
		// }
		// if L[j] > 0 {
		// 	ans++
		// }
		comp, _ := RunLengthEncoding(R[j][:15])
		for _, v := range comp {
			if v == 1 {
				ans++
			}
		}
		comp, _ = RunLengthEncoding(L[j][:15])
		for _, v := range comp {
			if v == 1 {
				ans++
			}
		}
	}
	println(ans)

	// steps := [][]int{
	// 	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	// }

	// black := 0
	// minus := 0
	// for i := 0; i < h; i++ {
	// 	for j := 0; j < w; j++ {
	// 		if S[i][j] == '#' {
	// 			black++

	// 			for _, s := range steps {
	// 				ny, nx := i+s[0], j+s[1]
	// 				if 0 <= ny && ny < h && 0 <= nx && nx < w && S[ny][nx] == '#' {
	// 					minus++
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// ans := black*4 - minus
	// println(ans)
}

// RunLengthEncoding returns encoded slice of an input.
// Time complexity: O(|S|)
func RunLengthEncoding(S []int) (comp []int, cnts []int) {
	comp = []int{}
	cnts = []int{}

	l := 0
	for i := 0; i < len(S); i++ {
		if i == 0 {
			l = 1
			continue
		}

		if S[i-1] == S[i] {
			l++
		} else {
			comp = append(comp, S[i-1])
			cnts = append(cnts, l)
			l = 1
		}
	}
	comp = append(comp, S[len(S)-1])
	cnts = append(cnts, l)

	return
}

// RunLengthDecoding decodes RLE results.
// Time complexity: O(|S|)
func RunLengthDecoding(comp []int, cnts []int) (S []int) {
	if len(comp) != len(cnts) {
		panic("S, L are not RunLengthEncoding results")
	}

	S = []int{}

	for i := 0; i < len(comp); i++ {
		for j := 0; j < cnts[i]; j++ {
			S = append(S, comp[i])
		}
	}

	return
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

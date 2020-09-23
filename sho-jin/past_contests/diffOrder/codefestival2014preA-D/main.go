/*
URL:
https://atcoder.jp/contests/code-festival-2014-quala/tasks/code_festival_qualA_d
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

var (
	A    []rune
	a, k int
)

func main() {
	A = ReadRuneSlice()
	a, _ = strconv.Atoi(string(A))
	k = ReadInt()

	// if len(A) <= k || k == 10 {
	// 	fmt.Println(0)
	// 	return
	// }
	// if check() {
	// 	fmt.Println(0)
	// 	return
	// }
	if isOK(A) {
		fmt.Println(0)
		return
	}

	ans := INF_BIT60
	for i := -1; i < len(A); i++ {
		for j := 0; j < 10; j++ {
			for l := 0; l < 10; l++ {
				// i未満までAをコピー
				val := []rune{}
				for m := 0; m < i+1; m++ {
					val = append(val, A[m])
				}

				// jを追加
				if len(val) < len(A) {
					val = append(val, rune(j)+'0')
				}
				// Aの長さに一致するまでlを追加
				for len(val) < len(A) {
					val = append(val, rune(l)+'0')
				}

				if val[0] == '0' {
					val = val[1:]
				}

				// 条件チェック
				if isOK(val) {
					v, _ := strconv.Atoi(string(val))
					ChMin(&ans, AbsInt(a-v))
				}
			}
		}
	}
	fmt.Println(ans)
}

func isOK(B []rune) bool {
	memo := make([]int, 10)
	for _, b := range B {
		memo[b-'0'] = 1
	}

	cnt := 0
	for i := 0; i < 10; i++ {
		cnt += memo[i]
	}

	return cnt <= k
}

func check() bool {
	memo := [10]int{}
	for i := 0; i < len(A); i++ {
		memo[A[i]-'0'] = 1
	}

	return count(memo) <= k
}

func copyMemo(org [10]int) [10]int {
	res := [10]int{}
	for i := 0; i < 10; i++ {
		res[i] = org[i]
	}
	return res
}

func count(memo [10]int) int {
	res := 0
	for i := 0; i < 10; i++ {
		res += memo[i]
	}
	return res
}

func used(digit int, memo [10]int) bool {
	return memo[digit] == 1
}

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

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

// func main() {
// 	A = ReadRuneSlice()
// 	a, _ := strconv.Atoi(string(A))
// 	k = ReadInt()

// 	if len(A) <= k || k == 10 {
// 		fmt.Println(0)
// 		return
// 	}

// 	if check() {
// 		fmt.Println(0)
// 		return
// 	}

// 	ans := INF_BIT60
// 	for i := -1; i < len(A); i++ {
// 		memo := [10]int{}
// 		val := []rune{}
// 		for j := 0; j <= i; j++ {
// 			val = append(val, A[j])
// 			memo[A[j]-'0'] = 1
// 		}
// 		if count(memo) > k {
// 			break
// 		}

// 		if len(val) == len(A) {
// 			fmt.Println(0)
// 			return
// 		}

// 		for j := 0; j < 10; j++ {
// 			memo2 := copyMemo(memo)
// 			val2 := []rune(string(val))

// 			if !used(j, memo2) && count(memo2) >= k {
// 				continue
// 			}
// 			memo2[j] = 1
// 			val2 = append(val2, rune(j)+'0')

// 			if len(val2) == len(A) {
// 				v, _ := strconv.Atoi(string(val2))
// 				ChMin(&ans, AbsInt(a-v))
// 				continue
// 			}

// 			for l := 0; l < 10; l++ {
// 				memo3 := copyMemo(memo2)
// 				val3 := []rune(string(val2))

// 				if !used(l, memo3) && count(memo3) >= k {
// 					continue
// 				}
// 				memo3[l] = 1

// 				// for m := len(val3); m < len(A); m++ {
// 				// 	val3 = append(val3, rune(l)+'0')
// 				// }
// 				for len(val3) < len(A) {
// 					val3 = append(val3, rune(l)+'0')
// 				}

// 				v, _ := strconv.Atoi(string(val3))
// 				ChMin(&ans, AbsInt(a-v))
// 			}
// 		}
// 	}
// 	fmt.Println(ans)
// }

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

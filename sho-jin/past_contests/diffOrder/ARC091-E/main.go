/*
URL:
https://atcoder.jp/contests/arc091/tasks/arc091_c
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
	n, a, b int
)

func main() {
	n, a, b = ReadInt3()

	if a+b > n+1 || a*b < n {
		fmt.Println(-1)
		return
	}

	AB := [][]int{}
	for i := 1; i <= a; i++ {
		tmp := []int{}
		if i > 1 {
			tmp = append(tmp, b*i)
		} else {
			for j := b; j >= 1; j-- {
				tmp = append(tmp, j)
			}
		}

		AB = append(AB, tmp)
	}

	nokori := n - b - (a - 1)
OUTER:
	for nokori > 0 {
		for i := 1; i < len(AB); i++ {
			last := AB[i][len(AB[i])-1]
			AB[i] = append(AB[i], last-1)
			nokori--

			if nokori <= 0 {
				break OUTER
			}
		}
	}
	PrintfDebug("%v\n", AB)

	// a*bがでかいとTLEする
	// nokori := a*b - n
	// OUTER:
	// for nokori > 0 {
	// 	for _, B := range AB {
	// 		for i := 1; i < len(B); i++ {
	// 			B[i] = 0
	// 			nokori--

	// 			if nokori <= 0 {
	// 				break OUTER
	// 			}
	// 		}
	// 	}
	// }

	ans := []int{}
	for _, B := range AB {
		for _, e := range B {
			if e == 0 {
				continue
			}
			ans = append(ans, e)
		}
	}

	answers, _, _ := ZaAtsu1Dim(ans, 1)

	// チェック
	// lis := LIS(answers)
	// lds := LIS(Reverse(answers))
	// PrintfDebug("LIS: %d, LDS: %d\n", lis, lds)

	fmt.Println(PrintIntsLine(answers...))
}

// ZaAtsu1Dim returns 3 values.
// pressed: pressed slice of the original slice
// orgToPress: map for translating original value to pressed value
// pressToOrg: reverse resolution of orgToPress
// O(nlogn)
func ZaAtsu1Dim(org []int, initVal int) (pressed []int, orgToPress, pressToOrg map[int]int) {
	pressed = make([]int, len(org))
	copy(pressed, org)
	sort.Sort(sort.IntSlice(pressed))

	orgToPress = make(map[int]int)
	for i := 0; i < len(org); i++ {
		if i == 0 {
			orgToPress[pressed[0]] = initVal
			continue
		}

		if pressed[i-1] != pressed[i] {
			initVal++
			orgToPress[pressed[i]] = initVal
		}
	}

	for i := 0; i < len(org); i++ {
		pressed[i] = orgToPress[org[i]]
	}

	pressToOrg = make(map[int]int)
	for k, v := range orgToPress {
		pressToOrg[v] = k
	}

	return
}

// func main() {
// 	patterns := FactorialPatterns([]int{1, 2, 3, 4, 5, 6})
// 	for _, P := range patterns {
// 		lis := LIS(P)
// 		lds := LIS(Reverse(P))

// 		fmt.Printf("a: %d, b: %d - %v\n", lis, lds, P)
// 	}

// 	// n, a, b = ReadInt3()

// 	if n == 1 {
// 		if a == 1 && b == 1 {
// 			fmt.Println(1)
// 		} else {
// 			fmt.Println(-1)
// 		}

// 		return
// 	}

// 	if n == 2 {
// 		if a == 2 && b == 1 {
// 			fmt.Println(1, 2)
// 		} else if a == 1 && b == 2 {
// 			fmt.Println(2, 1)
// 		} else {
// 			fmt.Println(-1)
// 		}

// 		return
// 	}

// 	if (a == b && a+b == n) || (AbsInt(a-b) == 1 && a+b == n) {
// 		ans := []int{}
// 		for i := n - a + 1; i <= n; i++ {
// 			ans = append(ans, i)
// 		}
// 		ans = append(ans, 1)
// 		for i := n - a; i >= 2; i-- {
// 			ans = append(ans, i)
// 		}

// 		// if a >= b {
// 		// 	fmt.Println(PrintIntsLine(ans...))
// 		// } else {
// 		// 	rev := Reverse(ans)
// 		// 	fmt.Println(PrintIntsLine(rev...))
// 		// }

// 		fmt.Println(PrintIntsLine(ans...))
// 		return
// 	}

// 	if a+b != n+1 {
// 		fmt.Println(-1)
// 		return
// 	}

// 	ans := []int{}
// 	for i := n - a + 1; i <= n; i++ {
// 		ans = append(ans, i)
// 	}
// 	for i := n - a; i >= 1; i-- {
// 		ans = append(ans, i)
// 	}
// 	fmt.Println(PrintIntsLine(ans...))

// }

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Reverse(A []int) []int {
	res := []int{}

	n := len(A)
	for i := n - 1; i >= 0; i-- {
		res = append(res, A[i])
	}

	return res
}

// LIS returns a length of Longest Increasing Subsequence of the argument slice A.
// O( Nlog(N) )
func LIS(A []int) int {
	// data structure
	const (
		INIT_VAL = 1 << 60
		MAX_N    = 200000 + 50
	)
	var dp = make([]int, MAX_N)

	// binary search sub function
	sub := func(a int) int {
		isOK := func(m, a int) bool {
			if dp[m] < a {
				return true
			}
			return false
		}
		ng, ok := len(dp), -1
		for int(math.Abs(float64(ok-ng))) > 1 {
			mid := (ok + ng) / 2
			if isOK(mid, a) {
				ok = mid
			} else {
				ng = mid
			}
		}
		bIdx := ok
		return bIdx
	}

	// main algorithm
	for i := 0; i < len(dp); i++ {
		dp[i] = INIT_VAL
	}
	for i := 0; i < len(A); i++ {
		idx := sub(A[i])
		dp[idx+1] = A[i]
	}
	return sub(INIT_VAL) + 1
}

// FactorialPatterns returns all patterns of n! of elems([]int).
func FactorialPatterns(elems []int) [][]int {
	newResi := make([]int, len(elems))
	copy(newResi, elems)

	return factRec([]int{}, newResi)
}

// DFS function for FactorialPatterns.
func factRec(pattern, residual []int) [][]int {
	if len(residual) == 0 {
		return [][]int{pattern}
	}

	res := [][]int{}
	for i, e := range residual {
		newPattern := make([]int, len(pattern))
		copy(newPattern, pattern)
		newPattern = append(newPattern, e)

		newResi := []int{}
		newResi = append(newResi, residual[:i]...)
		newResi = append(newResi, residual[i+1:]...)

		res = append(res, factRec(newPattern, newResi)...)
	}

	return res
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

/*
URL:
https://atcoder.jp/contests/abc179/tasks/abc179_e
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
	n, x, m int

	A [1000000 + 50]int

	D    [100000 + 50]int
	z    int
	toz  int
	roop int
)

func main() {
	n, x, m = readi3()

	if n <= 200000 {
		ans := excep()
		fmt.Println(ans)
		return
	}

	D[x] = 1
	cx := x
	for {
		y := (cx * cx) % m

		if y == 0 {
			// 例外処理
			ans := excep()
			fmt.Println(ans)
			return
		}

		if D[y] > 0 {
			z = y
			break
		}

		D[y]++
		cx = y
	}

	// zまでの遷移数
	toz = 0
	cx = x
	for cx != z {
		toz++
		y := (cx * cx) % m
		cx = y
	}
	// zからzまでの遷移数
	roop = 1
	cx = (z * z) % m
	for cx != z {
		roop++
		y := (cx * cx) % m
		cx = y
	}

	ans := 0
	// 最初の分
	cx = x
	for i := 0; i < toz; i++ {
		ans += cx
		y := (cx * cx) % m
		cx = y
	}
	// ループ分
	base := 0
	cx = z
	for i := 0; i < roop; i++ {
		base += cx
		y := (cx * cx) % m
		cx = y
	}

	resi := n - toz
	roopNum := resi / roop
	nokori := resi % roop

	ans += base * roopNum

	cx = z
	for i := 0; i < nokori; i++ {
		ans += cx
		y := (cx * cx) % m
		cx = y
	}

	fmt.Println(ans)
}

func excep() int {
	A[1] = x
	for i := 2; i <= 1000000; i++ {
		A[i] = (A[i-1] * A[i-1]) % m
	}

	ans := 0
	for i := 1; i <= Min(n, 1000000); i++ {
		ans += A[i]
	}

	return ans
}

// func main() {
// 	defer stdout.Flush()

// 	n, x, m = readi3()
// 	A[1] = x
// 	for i := 2; i <= 1000000; i++ {
// 		A[i] = (A[i-1] * A[i-1]) % m
// 	}

// 	ans := 0
// 	for i := 1; i <= Min(n, 1000000); i++ {
// 		ans += A[i]
// 	}
// 	fmt.Println(ans)
// }

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

/*
URL:
https://atcoder.jp/contests/abc173/tasks/abc173_e
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
	n, k int
	A    []int

	posi, nega, zero int
)

func main() {
	n, k = ReadInt2()
	A = ReadIntSlice(n)

	if n == k {
		ans := 1
		for i := 0; i < n; i++ {
			ans *= A[i]
			ans = NegativeMod(ans, MOD)
		}
		fmt.Println(ans)
		return
	}

	Count()

	if posi+nega < k {
		// 0を選ばざるを得ない
		fmt.Println(0)
		return
	}

	if !isPositive() {
		// 0以外にしようとすると負になってしまう
		if zero > 0 {
			// 0にできるので0にする
			fmt.Println(0)
		} else {
			// 負の最小を目指す
			res := solveNegativeMax()
			fmt.Println(res)
		}
		return
	}

	// 正の最大値を目指す
	solve()
}

func solve() {
	P, N := []int{}, []int{}

	// 0は捨ててしまう
	for i := 0; i < n; i++ {
		if A[i] < 0 {
			N = append(N, AbsInt(A[i]))
		} else if A[i] > 0 {
			P = append(P, AbsInt(A[i]))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(N)))
	sort.Sort(sort.Reverse(sort.IntSlice(P)))

	ans := 1
	pid, nid := 0, 0
	nokori := k
	for nokori > 0 {
		if pid >= len(P) {
			// 正がもうない
			ans *= N[nid]
			ans %= MOD

			nid++
			nokori--
			continue
		}
		if nid >= len(N) {
			// 負がもうない
			ans *= P[pid]
			ans %= MOD

			pid++
			nokori--
			continue
		}

		if P[pid] > N[nid] {
			ans *= P[pid]
			ans %= MOD

			pid++
		} else {
			ans *= N[nid]
			ans %= MOD

			nid++
		}

		nokori--
	}

	if nid%2 == 0 {
		// この状態が最大値（負を1つも含まない場合も、このケースに含まれる）
		fmt.Println(ans)
		return
	}

	// if nid == 0 {
	// 	// 負が一つも選ばれていない
	// }
	if pid == 0 {
		// 正が一つも選ばれていない -> 負の最小と正の最大を交換
		nmin := N[nid-1]
		pmax := P[pid]

		ans *= ModInv(nmin, MOD)
		ans %= MOD
		ans *= pmax
		ans %= MOD

		fmt.Println(ans)
		return
	}

	nmin := N[nid-1]
	pmax := P[pid]
	tmp1 := float64(pmax) / float64(nmin)

	pmin := P[pid-1]
	nmax := N[nid]
	tmp2 := float64(nmax) / float64(pmin)

	if tmp1 > tmp2 {
		ans *= ModInv(nmin, MOD)
		ans %= MOD
		ans *= pmax
		ans %= MOD
	} else {
		ans *= ModInv(pmin, MOD)
		ans %= MOD
		ans *= nmax
		ans %= MOD
	}

	fmt.Println(ans)
}

func solveNegativeMax() int {
	// Aに0は存在しない！

	// 絶対値の小さいものをk個選ぶ
	B := []Number{}
	for i := 0; i < n; i++ {
		var s int
		if A[i] < 0 {
			s = -1
		} else {
			s = 1
		}
		B = append(B, Number{absVal: AbsInt(A[i]), sign: s})
	}

	sort.SliceStable(B, func(i, j int) bool {
		return B[i].absVal < B[j].absVal
	})

	res := 1
	for i := 0; i < k; i++ {
		val := B[i].absVal
		if B[i].sign < 0 {
			val *= -1
		}

		res *= val
		res = NegativeMod(res, MOD)
	}

	return res
}

type Number struct {
	absVal int
	sign   int
}

func isPositive() bool {
	return !(nega == n && k%2 == 1)
}

// func isPositive() bool {
// 	nokori := k

// 	for i := 0; i < nega/2; i++ {
// 		if nokori-2 >= 0 {
// 			nokori -= 2
// 		} else {
// 			break
// 		}
// 	}

// 	if nokori <= 0 {
// 		// 負だけ選んで正にできる
// 		return true
// 	}

// 	if nokori <= posi {
// 		// 残りを正で消化して正にできる
// 		return true
// 	}

// 	// 0以外を選ぶと負にならざるを得ない
// 	return false
// }

func Count() {
	posi, nega, zero = 0, 0, 0
	for i := 0; i < n; i++ {
		if A[i] < 0 {
			nega++
		} else if A[i] > 0 {
			posi++
		} else {
			zero++
		}
	}
}

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// NegativeMod can calculate a right residual whether value is positive or negative.
func NegativeMod(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
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

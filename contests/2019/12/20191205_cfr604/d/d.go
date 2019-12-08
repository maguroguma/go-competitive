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

// PrintDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintDebug(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
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

var a, b, c, d int

func main() {
	a, b, c, d = ReadInt4()

	oneNeeds := a + 1
	twoNeeds := d + 1

	// 1, 2が足りないとNO
	if a-1 > b {
		fmt.Println("NO")
		return
	}
	if d-1 > c {
		fmt.Println("NO")
		return
	}

	oneDiff := b - oneNeeds
	twoDiff := c - twoNeeds

	if oneDiff >= 0 && twoDiff >= 0 {
		// 両方とも非負
		diff := AbsInt(oneDiff - twoDiff)
		if diff < 2 {
			fmt.Println("YES")
			sub(oneDiff, twoDiff)
		} else {
			fmt.Println("NO")
		}
	} else {
		S := []int{}
		if oneDiff == -1 && twoDiff == 0 {
			fmt.Println("YES")
			for i := 0; i < a; i++ {
				S = append(S, 0, 1)
			}
			S = append(S, 2)
			for i := 0; i < d; i++ {
				S = append(S, 3, 2)
			}
			fmt.Println(PrintIntsLine(S...))
		} else if oneDiff == 0 && twoDiff == -1 {
			fmt.Println("YES")
			S = append(S, 1)
			for i := 0; i < a; i++ {
				S = append(S, 0, 1)
			}
			for i := 0; i < d; i++ {
				S = append(S, 2, 3)
			}
			fmt.Println(PrintIntsLine(S...))
		} else if oneDiff == -1 && twoDiff == -1 {
			fmt.Println("YES")
			for i := 0; i < a; i++ {
				S = append(S, 0, 1)
			}
			for i := 0; i < d; i++ {
				S = append(S, 2, 3)
			}
			fmt.Println(PrintIntsLine(S...))
		} else {
			fmt.Println("NO")
		}
	}
}

func sub(oneDiff, twoDiff int) {
	S := []int{}

	if oneDiff > twoDiff {
		// 101から
		for i := 0; i < oneDiff+twoDiff+2; i++ {
			if i%2 == 0 {
				S = append(S, 1, 0, 1)
			} else {
				S = append(S, 2, 3, 2)
			}
		}
	} else {
		// 232から
		for i := 0; i < oneDiff+twoDiff+2; i++ {
			if i%2 == 0 {
				S = append(S, 2, 3, 2)
			} else {
				S = append(S, 1, 0, 1)
			}
		}
	}

	pressed, _ := RunLengthEncoding(S)

	fmt.Println(PrintIntsLine(pressed...))
}

// RunLengthEncoding returns encoded slice of an input.
func RunLengthEncoding(S []int) ([]int, []int) {
	runes := []int{}
	lengths := []int{}

	l := 0
	for i := 0; i < len(S); i++ {
		// 1文字目の場合保持
		if i == 0 {
			l = 1
			continue
		}

		if S[i-1] == S[i] {
			// 直前の文字と一致していればインクリメント
			l++
		} else {
			// 不一致のタイミングで追加し、長さをリセットする
			runes = append(runes, S[i-1])
			lengths = append(lengths, l)
			l = 1
		}
	}
	runes = append(runes, S[len(S)-1])
	lengths = append(lengths, l)

	return runes, lengths
}

// RunLengthDecoding decodes RLE results.
func RunLengthDecoding(S []int, L []int) []int {
	if len(S) != len(L) {
		panic("S, L are not RunLengthEncoding results")
	}

	res := []int{}

	for i := 0; i < len(S); i++ {
		for j := 0; j < L[i]; j++ {
			res = append(res, S[i])
		}
	}

	return res
}

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
- MODは最後にとりましたか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
*/

/*******************************************************************/

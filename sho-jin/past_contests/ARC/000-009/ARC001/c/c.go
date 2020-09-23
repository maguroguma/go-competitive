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

/*
ASCII code

ASCII   10進数  ASCII   10進数  ASCII   10進数
!       33      "       34      #       35
$       36      %       37      &       38
'       39      (       40      )       41
*       42      +       43      ,       44
-       45      .       46      /       47
0       48      1       49      2       50
3       51      4       52      5       53
6       54      7       55      8       56
9       57      :       58      ;       59
<       60      =       61      >       62
?       63      @       64      A       65
B       66      C       67      D       68
E       69      F       70      G       71
H       72      I       73      J       74
K       75      L       76      M       77
N       78      O       79      P       80
Q       81      R       82      S       83
T       84      U       85      V       86
W       87      X       88      Y       89
Z       90      [       91      \       92
]       93      ^       94      _       95
`       96      a       97      b       98
c       99      d       100     e       101
f       102     g       103     h       104
i       105     j       106     k       107
l       108     m       109     n       110
o       111     p       112     q       113
r       114     s       115     t       116
u       117     v       118     w       119
x       120     y       121     z       122
{       123     |       124     }       125
~       126             127
*/

/*******************************************************************/

const (
	// General purpose
	MOD          = 1000000000 + 7
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

var (
	C        [][]rune
	F        [][]rune
	up, down [][]Coord
	emptyRow []int
)

func main() {
	for i := 0; i < 8; i++ {
		row := ReadRuneSlice()
		C = append(C, row)
		crow := make([]rune, 8)
		F = append(F, crow)
	}

	for i := 0; i < 8; i++ {
		if !haveQween(C[i]) {
			emptyRow = append(emptyRow, i)
		}
	}
	if len(emptyRow) >= 6 {
		fmt.Println("No Answer")
		return
	}

	base := []int{}
	for i := 0; i < 8; i++ {
		base = append(base, i)
	}
	// patterns := PermutationPatterns(base, 5)
	patterns := DuplicatePatterns(base, 5)
	PrintDebug("%v\n", len(patterns))

	// buildField(patterns[0])
	// for i := 0; i < 8; i++ {
	// 	PrintDebug("%s\n", string(F[i]))
	// }

	initUpDown()

	for _, p := range patterns {
		buildField(p)

		// for i := 0; i < 8; i++ {
		// 	PrintDebug("%s\n", string(F[i]))
		// }
		// PrintDebug("--- --- ---\n")

		if isOK() {
			for i := 0; i < 8; i++ {
				fmt.Println(string(F[i]))
			}
			return
		}
	}
	fmt.Println("No Answer")
}

func isOK() bool {
	// 行
	for i := 0; i < 8; i++ {
		rnum := 0
		for j := 0; j < 8; j++ {
			if F[i][j] == 'Q' {
				rnum++
			}
		}
		if rnum != 1 {
			return false
		}
	}
	// 列
	for j := 0; j < 8; j++ {
		cnum := 0
		for i := 0; i < 8; i++ {
			if F[i][j] == 'Q' {
				cnum++
			}
		}
		if cnum != 1 {
			return false
		}
	}
	// 右肩上がり
	for i := 0; i < 15; i++ {
		num := 0
		for _, c := range up[i] {
			if F[c.y][c.x] == 'Q' {
				num++
			}
		}
		if num >= 2 {
			return false
		}
	}
	// 右肩下がり
	for i := 0; i < 15; i++ {
		num := 0
		for _, c := range down[i] {
			if F[c.y][c.x] == 'Q' {
				num++
			}
		}
		if num >= 2 {
			return false
		}
	}

	return true
}

type Coord struct {
	y, x int
}

func initUpDown() {
	// up
	up = make([][]Coord, 15)
	idx := 0
	for i := 0; i < 8; i++ {
		cy, cx := i, 0
		for d := 0; ; d++ {
			ny, nx := cy-d, cx+d
			if ny < 0 || nx >= 8 {
				break
			} else {
				up[idx] = append(up[idx], Coord{y: ny, x: nx})
			}
		}
		idx++
	}
	for j := 1; j < 8; j++ {
		cy, cx := 7, j
		for d := 0; ; d++ {
			ny, nx := cy-d, cx+d
			if ny < 0 || nx >= 8 {
				break
			} else {
				up[idx] = append(up[idx], Coord{y: ny, x: nx})
			}
		}
		idx++
	}

	// down
	down = make([][]Coord, 15)
	idx = 0
	for i := 7; i >= 0; i-- {
		cy, cx := i, 0
		for d := 0; ; d++ {
			ny, nx := cy+d, cx+d
			if ny >= 8 || nx >= 8 {
				break
			} else {
				down[idx] = append(down[idx], Coord{y: ny, x: nx})
			}
		}
		idx++
	}
	for j := 1; j < 8; j++ {
		cy, cx := 0, j
		for d := 0; ; d++ {
			ny, nx := cy+d, cx+d
			if ny >= 8 || nx >= 8 {
				break
			} else {
				down[idx] = append(down[idx], Coord{y: ny, x: nx})
			}
		}
		idx++
	}
}

func buildField(p []int) {
	// 初期化
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			F[i][j] = C[i][j]
		}
	}

	// パターンに従ってQを配置
	idx := 0
	for i := 0; i < 8; i++ {
		if haveQween(F[i]) {
			continue
		} else {
			F[i][p[idx]] = 'Q'
			idx++
		}
	}
}

func haveQween(row []rune) bool {
	for _, r := range row {
		if r == 'Q' {
			return true
		}
	}
	return false
}

// PermutationPatterns returns all patterns of nPk of elems([]int).
func PermutationPatterns(elems []int, k int) [][]int {
	newResi := make([]int, len(elems))
	copy(newResi, elems)

	return permRec([]int{}, newResi, k)
}

// DFS function for PermutationPatterns.
func permRec(pattern, residual []int, k int) [][]int {
	if len(pattern) == k {
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

		res = append(res, permRec(newPattern, newResi, k)...)
	}

	return res
}

// DuplicatePatterns returns all patterns of n^k of elems([]int).
func DuplicatePatterns(elems []int, k int) [][]int {
	return dupliRec([]int{}, elems, k)
}

// DFS function for DuplicatePatterns.
func dupliRec(pattern, elems []int, k int) [][]int {
	if len(pattern) == k {
		return [][]int{pattern}
	}

	res := [][]int{}
	for _, e := range elems {
		newPattern := make([]int, len(pattern))
		copy(newPattern, pattern)
		newPattern = append(newPattern, e)

		res = append(res, dupliRec(newPattern, elems, k)...)
	}

	return res
}

/*
- まずは全探索を検討しましょう
- MODは最後にとりましたか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
- いきなりオーバーフローはしていませんか？
- MOD取る系はint64必須ですよ？
- 後ろ・逆・ゴールから考えましたか？
- 3者のうち真ん中に着目しましたか？
*/

/*******************************************************************/

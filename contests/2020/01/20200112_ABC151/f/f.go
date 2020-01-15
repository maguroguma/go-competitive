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

var n int
var X, Y []int

func main() {
	n = ReadInt()
	X, Y = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x, y := ReadInt2()
		X[i], Y[i] = x, y
	}

	D := float64(INF_BIT60)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sx, sy := float64(X[i]), float64(Y[i])
				tx, ty := float64(X[j]), float64(Y[j])
				ux, uy := float64(X[k]), float64(Y[k])

				var cx, cy, r *float64
				stat := calcCircleOf2Point(sx, sy, tx, ty, ux, uy, cx, cy, r)
				if stat >= 0.0 {
					if D < *r {
						D = *r
					}
				}
			}
		}
	}
	fmt.Println(D)
}

func calcCircleOf2Point(x1, y1, x2, y2, x3, y3 float64, cx, cy, r *float64) float64 {
	var ox, oy, a, b, c, d float64
	var r1, r2, r3 float64
	var stat float64

	a = x2 - x1
	b = y2 - y1
	c = x3 - x1
	d = y3 - y1
	stat = -1.0

	if (a > 0.0 && d > 0.0) || (b > 0.0 && c > 0.0) {
		ox = x1 + (d*(a*a+b*b)-b*(c*c+d*d))/(a*d-b*c)/2
		if b > 0.0 {
			oy = (a*(x1+x2-ox-ox) + b*(y1+y2)) / b / 2
		} else {
			oy = (c*(x1+x3-ox-ox) + d*(y1+y3)) / d / 2
		}
		r1 = math.Sqrt((ox-x1)*(ox-x1) + (oy-y1)*(oy-y1))
		r2 = math.Sqrt((ox-x2)*(ox-x2) + (oy-y2)*(oy-y2))
		r3 = math.Sqrt((ox-x3)*(ox-x3) + (oy-y3)*(oy-y3))
		*cx = ox
		*cy = oy
		*r = (r1 + r2 + r3) / 3
		stat = 1.0
	}

	return stat
}

// int calcCircleOf2Point(int x1, int y1, int x2, int y2, int x3, int y3, int *cx, int *cy, int *r)
// {
// 	long 	ox, oy, a, b, c, d ;
// 	long 	r1, r2, r3 ;
// 	int		stat ;

// 	a = x2 - x1 ;
// 	b = y2 - y1 ;
// 	c = x3 - x1 ;
// 	d = y3 - y1 ;

// 	stat = -1 ;

// 	if  ((a && d) || (b && c)) {
// 		ox = x1 + (d * (a * a + b * b) - b * (c * c + d * d)) / (a * d - b * c) / 2 ;
// 		if  (b) {
// 			oy = (a * (x1 + x2 - ox - ox) + b * (y1 + y2)) / b / 2 ;
// 		} else {
// 			oy = (c * (x1 + x3 - ox - ox) + d * (y1 + y3)) / d / 2 ;
// 		}
// 		r1   = sqrt((ox - x1) * (ox - x1) + (oy - y1) * (oy - y1)) ;
// 		r2   = sqrt((ox - x2) * (ox - x2) + (oy - y2) * (oy - y2)) ;
// 		r3   = sqrt((ox - x3) * (ox - x3) + (oy - y3) * (oy - y3)) ;
// 		*cx = ox ;
// 		*cy = oy ;
// 		*r  = (r1 + r2 + r3) / 3 ;
// 		stat = 0 ;
// 	}

// 	return stat ;
// }

/*
- まずは全探索を検討しましょう
- MODは最後にとりましたか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
- いきなりオーバーフローはしていませんか？
	- MOD取る系はint64必須ですよ？
*/

/*******************************************************************/

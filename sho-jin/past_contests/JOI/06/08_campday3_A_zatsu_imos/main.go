/*
URL:
https://atcoder.jp/contests/joisc2008/tasks/joisc2008_origami
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	n          int
	a, b       int
	P, Q, R, S []int32

	X, Y []int32

	A    [][]int
	ok   [][]bool
	maxS int
)

func main() {
	defer stdout.Flush()

	n = readi()
	a, b = readi2()
	for i := 0; i < n; i++ {
		p, q, r, s := readi4()
		p--
		q--
		r--
		s--
		P = append(P, int32(p))
		Q = append(Q, int32(q))
		R = append(R, int32(r))
		S = append(S, int32(s))
	}

	X = append(X, 0)
	X = append(X, int32(a-1))
	Y = append(Y, 0)
	Y = append(Y, int32(b-1))
	for i := 0; i < n; i++ {
		x1, x2 := P[i], R[i]
		// X = append(X, x1, x2, x1-1, x1+1, x2-1, x2+1)
		X = append(X, x1, x2, x1+1, x2+1)
		y1, y2 := Q[i], S[i]
		// Y = append(Y, y1, y2, y1-1, y1+1, y2-1, y2+1)
		Y = append(Y, y1, y2, y1+1, y2+1)
	}

	_, xotp, xpto := ZaAtsu1Dim(X, 0)
	_, yotp, ypto := ZaAtsu1Dim(Y, 0)
	// debugf("xp: %v\n", xp)
	// debugf("yp: %v\n", yp)
	// debugf("xotp: %v\n", xotp)
	// debugf("yotp: %v\n", yotp)
	// debugf("xpto: %v\n", xpto)
	// debugf("ypto: %v\n", ypto)

	xm := Max(Max(P...), Max(R...))
	ym := Max(Max(Q...), Max(S...))
	xx := xotp[xm]
	yy := yotp[ym]

	// A = make([][]int, Max(yp...)+5)
	A = make([][]int, yy+5)
	for i := 0; i < len(A); i++ {
		A[i] = make([]int, xx+5)
	}

	for i := 0; i < n; i++ {
		x, y := P[i], Q[i]
		left := xotp[x]
		up := yotp[y]

		x, y = R[i], S[i]
		right := xotp[x]
		down := yotp[y]

		A[up][left]++
		A[up][right+1]--
		A[down+1][left]--
		A[down+1][right+1]++
	}

	// 横方向に累積
	for i := 0; i < len(A); i++ {
		for j := 1; j < len(A[i]); j++ {
			A[i][j] += A[i][j-1]
		}
	}
	// 縦方向に累積
	for j := 0; j < len(A[0]); j++ {
		for i := 1; i < len(A); i++ {
			A[i][j] += A[i-1][j]
		}
	}
	// for i := 0; i < len(A); i++ {
	// 	debugf("%v\n", A[i])
	// }

	maxS = 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			ChMax(&maxS, A[i][j])
		}
	}

	ok = make([][]bool, len(A))
	for i := 0; i < len(A); i++ {
		ok[i] = make([]bool, len(A[0]))
	}
	menseki := int32(0)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if !ok[i][j] && A[i][j] == maxS {
				up, left := int32(i), int32(j)
				down, right := rd(up, left)
				ll := xpto[left]
				rr := xpto[right]
				uu := ypto[up]
				dd := ypto[down]

				menseki += (rr - ll + 1) * (dd - uu + 1)

				bomb(left, right, up, down)
			} else {
				ok[i][j] = true
			}
		}
	}

	fmt.Println(maxS)
	fmt.Println(menseki)
}

func rd(cy, cx int32) (int32, int32) {
	for A[cy+1][cx] == maxS {
		cy++
	}
	for A[cy][cx+1] == maxS {
		cx++
	}
	return cy, cx
}

func bomb(l, r, u, d int32) {
	for i := u; i <= d; i++ {
		for j := l; j <= r; j++ {
			ok[i][j] = true
		}
	}
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

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int32) int32 {
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

// Max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Max(integers ...int32) int32 {
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

// ZaAtsu1Dim returns 3 values.
// pressed: pressed slice of the original slice
// orgToPress: map for translating original value to pressed value
// pressToOrg: reverse resolution of orgToPress
// O(nlogn)
func ZaAtsu1Dim(org []int32, initVal int32) (pressed []int32, orgToPress, pressToOrg map[int32]int32) {
	pressed = make([]int32, len(org))
	copy(pressed, org)
	// sort.Sort(sort.IntSlice(pressed))
	sort.Slice(pressed, func(i, j int) bool {
		return pressed[i] < pressed[j]
	})

	orgToPress = make(map[int32]int32)
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

	pressToOrg = make(map[int32]int32)
	for k, v := range orgToPress {
		pressToOrg[v] = k
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

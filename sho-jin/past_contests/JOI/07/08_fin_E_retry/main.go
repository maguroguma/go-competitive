/*
URL:
https://atcoder.jp/contests/joi2008ho/tasks/joi2008ho_e
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

// max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func max(integers ...int32) int32 {
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

// min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func min(integers ...int32) int32 {
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

var (
	w, h       int32
	n          int
	A, B, C, D []int32

	// M [][]int16
	M     [2000 + 50][2000 + 50]int16
	steps [][2]int32
	W, H  int32
)

func main() {
	defer stdout.Flush()

	ww, hh := readi2()
	w, h = int32(ww), int32(hh)
	n = readi()
	A, B, C, D = make([]int32, n), make([]int32, n), make([]int32, n), make([]int32, n)
	for i := 0; i < n; i++ {
		x1, y1, x2, y2 := readi4()
		A[i], B[i], C[i], D[i] = int32(x1), int32(y1), int32(x2), int32(y2)
	}

	compX, compY := NewCompress(), NewCompress()

	for i := 0; i < n; i++ {
		a, c := int32(A[i]), int32(C[i])
		compX.Add(a, c)
		b, d := int32(B[i]), int32(D[i])
		compY.Add(b, d)
	}
	compX.Add(0, w)
	compY.Add(0, h)

	compX.Build()
	compY.Build()
	maxX := int32(len(compX.cs) - 1)
	maxY := int32(len(compY.cs) - 1)

	H = maxY + 1
	W = maxX + 1

	var i, j int32

	// M = make([][]int16, H)
	// for i = 0; i < H; i++ {
	// 	M[i] = make([]int16, W)
	// }

	for i := 0; i < n; i++ {
		a, b, c, d := A[i], B[i], C[i], D[i]
		t, s, v, u := compX.Get(a), compY.Get(b), compX.Get(c), compY.Get(d)

		M[s][t]++
		if u+1 < H {
			M[u][t]--
		}
		if v+1 < W {
			M[s][v]--
		}
		if u+1 < H && v+1 < W {
			M[u][v]++
		}
	}

	for i = 0; i < H; i++ {
		for j = 0; j < W-1; j++ {
			M[i][j+1] += M[i][j]
		}
	}
	for j = 0; j < W; j++ {
		for i = 0; i < H-1; i++ {
			M[i+1][j] += M[i][j]
		}
	}

	// for i := H - 1; i >= 0; i-- {
	// 	debugf("M[i]: %v\n", M[i])
	// }

	steps = [][2]int32{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	ans := 0
	for i = 0; i < H; i++ {
		for j = 0; j < W; j++ {
			if M[i][j] > 0 {
				continue
			}
			ans++
			// dfs(i, j)
			bfs(i, j)
		}
	}

	fmt.Println(ans)
}

func bfs(i, j int32) {
	M[i][j]++

	Q := [][2]int32{}
	Q = append(Q, [2]int32{i, j})
	for len(Q) > 0 {
		pop := Q[0]
		Q = Q[1:]
		y, x := pop[0], pop[1]
		for _, step := range steps {
			dy, dx := step[0], step[1]
			ny, nx := y+dy, x+dx
			if 0 <= ny && ny < H && 0 <= nx && nx < W && M[ny][nx] == 0 {
				M[ny][nx]++
				Q = append(Q, [2]int32{ny, nx})
			}
		}
	}
}

func dfs(i, j int32) {
	M[i][j]++
	for _, step := range steps {
		dy, dx := step[0], step[1]
		ny, nx := i+dy, j+dx
		if 0 <= ny && ny < H && 0 <= nx && nx < W && M[ny][nx] == 0 {
			dfs(ny, nx)
		}
	}
}

// NewCompress returns a compress algorithm.
func NewCompress() *Compress {
	c := new(Compress)
	c.xs = []int32{}
	c.cs = []int32{}

	return c
}

// Add can add any number of elements.
// Time complexity: O(1)
func (c *Compress) Add(X ...int32) {
	c.xs = append(c.xs, X...)
}

// Build compresses input elements by sorting.
// Time complexity: O(NlogN)
func (c *Compress) Build() {
	sort.Slice(c.xs, func(i, j int) bool {
		return c.xs[i] < c.xs[j]
	})

	if len(c.xs) == 0 {
		panic("Compress doesn't have any elements")
	}

	c.cs = append(c.cs, c.xs[0])
	for i := 1; i < len(c.xs); i++ {
		if c.xs[i-1] == c.xs[i] {
			continue
		}
		c.cs = append(c.cs, c.xs[i])
	}
}

// Get returns index that is equal to by binary search.
// Time complexity: O(logN)
func (c *Compress) Get(x int32) int32 {
	_abs := func(a int32) int32 {
		if a < 0 {
			return -a
		}
		return a
	}

	var ng, ok = int32(-1), int32(len(c.cs))
	for _abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if c.cs[mid] >= x {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

type Compress struct {
	xs []int32 // sorted original values
	cs []int32 // sorted compressed values
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

// // min returns the min integer among input set.
// // This function needs at least 1 argument (no argument causes panic).
// func min(integers ...int) int {
// 	m := integers[0]
// 	for i, integer := range integers {
// 		if i == 0 {
// 			continue
// 		}
// 		if m > integer {
// 			m = integer
// 		}
// 	}
// 	return m
// }

// // max returns the max integer among input set.
// // This function needs at least 1 argument (no argument causes panic).
// func max(integers ...int) int {
// 	m := integers[0]
// 	for i, integer := range integers {
// 		if i == 0 {
// 			continue
// 		}
// 		if m < integer {
// 			m = integer
// 		}
// 	}
// 	return m
// }

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

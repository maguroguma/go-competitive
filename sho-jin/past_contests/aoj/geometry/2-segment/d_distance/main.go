/*
URL:
https://onlinejudge.u-aizu.ac.jp/courses/library/4/CGL/2/CGL_2_D
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

	q int
)

func main() {
	defer stdout.Flush()

	q = readi()
	for i := 0; i < q; i++ {
		x0, y0, x1, y1 := readf(), readf(), readf(), readf()
		x2, y2, x3, y3 := readf(), readf(), readf(), readf()

		a, b := NewCoord(x0, y0), NewCoord(x1, y1)
		c, d := NewCoord(x2, y2), NewCoord(x3, y3)

		c1 := distanceLSwithP(a, b, c)
		c2 := distanceLSwithP(a, b, d)
		c3 := distanceLSwithP(c, d, a)
		c4 := distanceLSwithP(c, d, b)

		ans := math.MaxFloat64
		for _, v := range []float64{c1, c2, c3, c4} {
			ans = math.Min(ans, v)
		}
		// printf("%v\n", ans)

		// s := math.Min(dist(a, c), dist(a, d))
		// t := math.Min(dist(b, c), dist(b, d))
		// ans := math.Min(s, t)
		// ans = math.Sqrt(ans)

		ba := a.minus(b)
		dc := c.minus(d)
		if _equal(ba.det(dc), 0.0) {
			// 平行な場合
			ok := onSegment(a, b, c) || onSegment(a, b, d) || onSegment(c, d, a) || onSegment(c, d, b)
			if ok {
				printf("%v\n", 0.0)
			} else {
				printf("%v\n", ans)
			}
		} else {
			// 平行でない場合
			r := intersection(a, b, c, d)
			ok := onSegment(a, b, r) && onSegment(c, d, r)
			if ok {
				printf("%v\n", 0.0)
			} else {
				printf("%v\n", ans)
			}
		}
	}
}

// 二次元ベクトル構造体
type Coord struct {
	x, y float64
}

func NewCoord(x, y float64) Coord {
	return Coord{x: x, y: y}
}

func _add_eps(a, b float64) float64 {
	if math.Abs(a+b) < EPS*(math.Abs(a)+math.Abs(b)) {
		return 0
	}
	return a + b
}

func _equal(a, b float64) bool {
	// return math.Abs(a-b) < 1e-6
	return math.Abs(a-b) < 1e-15
}

// 内積
func (p Coord) dot(q Coord) float64 {
	return _add_eps(p.x*q.x, p.y*q.y)
}

// 外積
func (p Coord) det(q Coord) float64 {
	return _add_eps(p.x*q.y, -p.y*q.x)
}

// 距離の二乗
func dist(p, q Coord) float64 {
	tx := _add_eps(p.x, -q.x)
	ty := _add_eps(p.y, -q.y)
	tp := Coord{tx, ty}

	return tp.dot(tp)
}

// ノルム
func (p Coord) norm() float64 {
	return math.Sqrt(p.dot(p))
}

func (p Coord) minus(q Coord) Coord {
	x := _add_eps(p.x, -q.x)
	y := _add_eps(p.y, -q.y)
	return Coord{x, y}
}

func (p Coord) add(q Coord) Coord {
	x := _add_eps(p.x, q.x)
	y := _add_eps(p.y, q.y)
	return Coord{x, y}
}

func (p Coord) mul(a float64) Coord {
	return Coord{p.x * a, p.y * a}
}

func (p Coord) unitVector() Coord {
	d := math.Sqrt(p.dot(p))

	return p.mul(1.0 / d)
}

func (p Coord) rot(r float64) Coord {
	x := p.x*math.Cos(r) - p.y*math.Sin(r)
	y := p.x*math.Sin(r) + p.y*math.Cos(r)
	return Coord{x, y}
}

// https://www.hiramine.com/programming/graphics/2d_vectorangle.html
// Acosは非推奨？
func radian(p, q Coord) float64 {
	var rad float64

	cos := p.dot(q) / (p.norm() * q.norm())
	if _equal(math.Abs(cos), 1.0) {
		if cos < 0.0 {
			rad = math.Pi
		} else {
			rad = 0.0
		}
	} else {
		rad = math.Acos(cos)
	}

	if p.det(q) < 0 {
		rad *= -1
	}
	return rad
}

func (p Coord) equal(q Coord) bool {
	xok := _equal(p.x, q.x)
	yok := _equal(p.y, q.y)
	return xok && yok
}

// 線分p1-p2上に点qがあるか判定
func onSegment(p1, p2, q Coord) bool {
	a := p1.minus(q)
	b := p2.minus(q)
	return _equal(a.det(b), 0.0) && a.dot(b) <= 0.0
}

func intersection(p1, p2, q1, q2 Coord) Coord {
	a := p2.minus(p1)

	b := q2.minus(q1)
	c := q1.minus(p1)
	d := q2.minus(q1)
	e := p2.minus(p1)
	val := b.det(c) / d.det(e)

	a = a.mul(val)

	return p1.add(a)
}

// 点a, bを通る直線と点cとの距離
func distanceLwithP(a, b, c Coord) float64 {
	ab, ac := b.minus(a), c.minus(a)

	nume := math.Abs(ab.det(ac))
	deno := math.Abs(ab.norm())
	return nume / deno
}

// 点a, bを端点とする線分と点cとの距離
// http://www.deqnotes.net/acmicpc/2d_geometry/lines
func distanceLSwithP(a, b, c Coord) float64 {
	ab, ac := b.minus(a), c.minus(a)
	ba, bc := a.minus(b), c.minus(b)

	// if _equal(ab.dot(ac), 0.0) {
	if ab.dot(ac) < 0.0 {
		return ac.norm()
	}
	// if _equal(ba.dot(bc), 0.0) {
	if ba.dot(bc) < 0.0 {
		return bc.norm()
	}

	return distanceLwithP(a, b, c)
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

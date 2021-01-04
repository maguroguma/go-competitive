/*
URL:
https://atcoder.jp/contests/abc022/tasks/abc022_d
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

var (
	println = fmt.Println

	n    int
	A, B []Coord
)

func main() {
	defer stdout.Flush()

	n = readi()
	for i := 0; i < n; i++ {
		x, y := readi2()
		xf, yf := float64(x), float64(y)
		A = append(A, Coord{xf, yf})
	}
	for i := 0; i < n; i++ {
		x, y := readi2()
		xf, yf := float64(x), float64(y)
		B = append(B, Coord{xf, yf})
	}

	bef := MaxDistEasy(A)
	aft := MaxDistEasy(B)

	// bef := MaxDistHard(A)
	// aft := MaxDistHard(B)

	println(math.Sqrt(aft / bef))
}

// 二次元ベクトル構造体
type Coord struct {
	x, y float64
}

// 凸包を求める
func ConvexHull(ps []Coord) []Coord {
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].x != ps[j].x {
			return ps[i].x < ps[j].x
		}
		return ps[i].y < ps[j].y
	})

	n := len(ps)

	k := 0                   // 凸包の頂点数
	qs := make([]Coord, 2*n) // 構築中の凸包

	// 下側凸包の作成
	for i := 0; i < n; i++ {
		for k > 1 {
			lp, rp := qs[k-1].minus(qs[k-2]), ps[i].minus(qs[k-1])

			if lp.det(rp) > 0.0 {
				break
			}
			k--
		}

		qs[k] = ps[i]
		k++
	}

	// 上側凸包の作成
	for i, t := n-2, k; i >= 0; i-- {
		for k > t {
			lp, rp := qs[k-1].minus(qs[k-2]), ps[i].minus(qs[k-1])

			if lp.det(rp) > 0.0 {
				break
			}
			k--
		}

		qs[k] = ps[i]
		k++
	}

	return qs[:k-1]
}

// ※座標範囲が小さく、凸包を構成する点の数が場合しか解けない
// Time complexity: O(最大座標幅)
func MaxDistEasy(ps []Coord) float64 {
	qs := ConvexHull(ps)
	res := 0.0
	for i := 0; i < len(qs); i++ {
		for j := 0; j < i; j++ {
			res = math.Max(res, dist(qs[i], qs[j]))
		}
	}
	return res
}

func MaxDistHard(ps []Coord) float64 {
	_cmp_x := func(p, q Coord) bool {
		if p.x != q.x {
			return p.x < q.x
		}
		return p.y < q.y
	}

	qs := ConvexHull(ps)

	n := len(qs)
	if n == 2 {
		// 凸包が潰れている場合は特別処理
		return dist(qs[0], qs[1])
	}

	i, j := 0, 0 // ある方向に最も遠い点対
	// x軸方向に最も遠い点対を求める
	for k := 0; k < n; k++ {
		if !_cmp_x(qs[i], qs[k]) {
			i = k
		}
		if _cmp_x(qs[j], qs[k]) {
			j = k
		}
	}

	res := 0.0
	si, sj := i, j
	for i != sj || j != si {
		// 方向を180度変化させる
		res = math.Max(res, dist(qs[i], qs[j]))

		// 辺i-(i+1)の法線方向と辺j-(j+1)の法線の方向のどちらを先に向くか判定
		im, jm := (i+1)%n, (j+1)%n

		lp, rp := qs[im], qs[i]
		left := lp.minus(rp)

		lp, rp = qs[jm], qs[j]
		right := lp.minus(rp)

		if left.det(right) < 0 {
			i = im // 辺i-(i+1)の法線方向を先に向く
		} else {
			j = jm // 辺j-(j+1)の法線方向を先に向く
		}
	}

	return res
}

func _add_eps(a, b float64) float64 {
	if math.Abs(a+b) < EPS*(math.Abs(a)+math.Abs(b)) {
		return 0
	}
	return a + b
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

func (p Coord) minus(q Coord) Coord {
	x := _add_eps(p.x, -q.x)
	y := _add_eps(p.y, -q.y)
	return Coord{x, y}
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

/*
URL:
https://onlinejudge.u-aizu.ac.jp/courses/library/4/CGL/all/CGL_2_A
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
)

func main() {
	defer stdout.Flush()

	q := readi()
	for i := 0; i < q; i++ {
		x0, y0, x1, y1, x2, y2, x3, y3 := readf(), readf(), readf(), readf(), readf(), readf(), readf(), readf()

		a, b, c, d := NewPoint(x0, y0), NewPoint(x1, y1), NewPoint(x2, y2), NewPoint(x3, y3)
		l1, l2 := NewLine(a, b), NewLine(c, d)

		if IsOrthogonal(l1, l2) {
			printf("1\n")
		} else if IsParallel(l1, l2) {
			printf("2\n")
		} else {
			printf("0\n")
		}
	}
}

// originated from:
// https://ei1333.github.io/luzhiled/snippets/geometry/template.html

type Point struct {
	x, y float64
}

type Line struct {
	a, b *Point
}

type Segment struct {
	a, b *Point
}

type Circle struct {
	p *Point
	r float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Add(q *Point) *Point {
	x, y := p.x+q.x, p.y+q.y
	return NewPoint(x, y)
}

func (p *Point) Minus(q *Point) *Point {
	x, y := p.x-q.x, p.y-q.y
	return NewPoint(x, y)
}

func (p *Point) Mul(a float64) *Point {
	return NewPoint(p.x*a, p.y*a)
}

func (p *Point) Dot(q *Point) float64 {
	return p.x*q.x + p.y*q.y
}

func (p *Point) Cross(q *Point) float64 {
	return p.x*q.y - p.y*q.x
}

func Dot(p, q *Point) float64 {
	return p.Dot(q)
}

func Cross(p, q *Point) float64 {
	return p.Cross(q)
}

func (p *Point) Norm2() float64 {
	return p.x*p.x + p.y*p.y
}

func (p *Point) Norm() float64 {
	return gsqrt(p.Norm2())
}

const (
	G_EPS = 1e-10
	G_PI  = math.Pi

	G_ONLINE_FRONT      = -2
	G_CLOCKWISE         = -1
	G_ON_SEGMENT        = 0
	G_COUNTER_CLOCKWISE = 1
	G_ONLINE_BACK       = 2

	G_OUT = 0
	G_ON  = 1
	G_IN  = 2
)

var (
	gabs   = math.Abs
	gcos   = math.Cos
	gsin   = math.Sin
	gatan2 = math.Atan2
	gmin   = math.Min
	gsqrt  = math.Sqrt
)

func fEq(v, w float64) bool {
	return gabs(v-w) < G_EPS
}

func pEq(p, q *Point) bool {
	dx, dy := p.x-q.x, p.y-q.y
	return fEq(dx, 0.0) && fEq(dy, 0.0)
}

func RotateTheta(t float64, p *Point) *Point {
	x := gcos(t)*p.x - gsin(t)*p.y
	y := gsin(t)*p.x + gcos(t)*p.y
	return NewPoint(x, y)
}

func RadianToDegree(r float64) float64 {
	return (r * 180.0) / G_PI
}

func DegreeToRadian(d float64) float64 {
	return (d * G_PI) / 180.0
}

// a-b-cの角度のうち小さい方を返す
func Angle(a, b, c *Point) float64 {
	v := b.Minus(a)
	w := c.Minus(b)
	alpha := gatan2(v.y, v.x)
	beta := gatan2(w.y, w.x)

	if alpha > beta {
		alpha, beta = beta, alpha
	}

	theta := beta - alpha
	return gmin(theta, 2.0*G_PI-theta)
}

func pLess(a, b *Point) bool {
	if a.x < b.x {
		return true
	} else if a.x > b.x {
		return false
	} else {
		return a.y < b.y
	}
}

func NewLine(a, b *Point) *Line {
	na, nb := NewPoint(a.x, a.y), NewPoint(b.x, b.y)
	return &Line{a: na, b: nb}
}

func NewSegment(a, b *Point) *Segment {
	na, nb := NewPoint(a.x, a.y), NewPoint(b.x, b.y)
	return &Segment{a: na, b: nb}
}

func NewCircle(p *Point, r float64) *Circle {
	np := NewPoint(p.x, p.y)
	return &Circle{p: np, r: r}
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_1_C
// 点の回転方向
func Ccw(a, b, c *Point) int {
	d, e := b.Minus(a), c.Minus(a)

	cross := Cross(d, e)
	if !fEq(cross, 0.0) {
		if cross > 0.0 {
			return G_COUNTER_CLOCKWISE
		}
		return G_CLOCKWISE
	}

	dot := Dot(d, e)
	if !fEq(dot, 0.0) && dot < 0.0 {
		return G_ONLINE_BACK
	}
	if d.Norm2() < e.Norm2() {
		return G_ONLINE_FRONT
	}

	return G_ON_SEGMENT
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_2_A
// 平行判定
func IsParallel(a, b *Line) bool {
	AB := a.b.Minus(a.a)
	CD := b.b.Minus(b.a)
	cross := Cross(AB, CD)
	return fEq(cross, 0.0)
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_2_A
// 垂直判定
func IsOrthogonal(a, b *Line) bool {
	BA := a.a.Minus(a.b)
	DC := b.a.Minus(b.b)
	dot := Dot(BA, DC)
	return fEq(dot, 0.0)
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_1_A
// 射影
// 直線 l に p から垂線を引いた交点を求める
func ProjectionToLine(l *Line, p *Point) *Point {
	AP := p.Minus(l.a)
	BA := l.a.Minus(l.b)
	t := Dot(AP, BA) / BA.Norm2()
	return l.a.Add(BA.Mul(t))
}
func ProjectionToSegment(l *Segment, p *Point) *Point {
	line := NewLine(l.a, l.b)
	return ProjectionToLine(line, p)
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_1_B
// 反射
// 直線 l を対称軸として点 p  と線対称にある点を求める
func Reflection(l *Line, p *Point) *Point {
	plus := ProjectionToLine(l, p).Minus(p).Mul(2.0)
	return p.Add(plus)
}

// 交差判定
func IsIntersectLinePoint(l *Line, p *Point) bool {
	ccw := Ccw(l.a, l.b, p)
	return ccw != G_CLOCKWISE && ccw != G_COUNTER_CLOCKWISE
}
func IsIntersectLineLine(l, m *Line) bool {
	AB := l.b.Minus(l.a)
	CD := m.b.Minus(m.a)
	cross := Cross(AB, CD)
	return !fEq(cross, 0.0)
}
func IsIntersectSegmentPoint(s *Segment, p *Point) bool {
	return Ccw(s.a, s.b, p) == G_ON_SEGMENT
}
func IsIntersectLineSegment(l *Line, s *Segment) bool {
	AB := l.b.Minus(l.a)
	AC := s.a.Minus(l.a)
	AD := s.b.Minus(l.a)
	return Cross(AB, AC)*Cross(AB, AD) < G_EPS
}

// func IsIntersectCircleLine(c *Circle, l *Line) bool {

// }
// func IsIntersectCirclePoint(c *Circle, p *Point) bool {

// }

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_2_B
func IsIntersectSegmentSegment(s, t *Segment) bool {
	lb := Ccw(s.a, s.b, t.a)*Ccw(s.a, s.b, t.b) <= 0
	rb := Ccw(t.a, t.b, s.a)*Ccw(t.a, t.b, s.b) <= 0
	return lb && rb
}

// func IsIntersectCircleSegment(c *Circle, s *Segment) bool {

// }

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_7_A&lang=jp
// func IsIntersectCircleCircle(c1, c2 *Circle) bool {

// }

func DistancePointPoint(a, b *Point) float64 {
	AB := a.Minus(b)
	return AB.Norm()
}
func DistanceLinePoint(l *Line, p *Point) float64 {
	q := ProjectionToLine(l, p)
	QP := p.Minus(q)
	return QP.Norm()
}
func DistanceLineLine(l, m *Line) float64 {
	if IsIntersectLineLine(l, m) {
		return 0.0
	}
	return DistanceLinePoint(l, m.a)
}
func DistanceSegmentPoint(s *Segment, p *Point) float64 {
	r := ProjectionToSegment(s, p)

	if IsIntersectSegmentPoint(s, r) {
		RP := p.Minus(r)
		return RP.Norm()
	}

	PA := s.a.Minus(p)
	PB := s.b.Minus(p)
	return gmin(PA.Norm(), PB.Norm())
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_2_D
func DistanceSegmentSegment(a, b *Segment) float64 {
	if IsIntersectSegmentSegment(a, b) {
		return 0.0
	}

	d1 := DistanceSegmentPoint(a, b.a)
	d2 := DistanceSegmentPoint(a, b.b)
	d3 := DistanceSegmentPoint(b, a.a)
	d4 := DistanceSegmentPoint(b, a.b)

	return gmin(d1, gmin(d2, gmin(d3, d4)))
}
func DistanceLineSegment(l *Line, s *Segment) float64 {
	if IsIntersectLineSegment(l, s) {
		return 0.0
	}

	d1 := DistanceLinePoint(l, s.a)
	d2 := DistanceLinePoint(l, s.b)
	return gmin(d1, d2)
}

func CrossPointLineLine(l, m *Line) *Point {
	AB := l.b.Minus(l.a)
	CD := m.b.Minus(m.a)
	CB := l.b.Minus(m.a)
	A := Cross(AB, CD)
	B := Cross(AB, CB)

	if fEq(gabs(A), 0.0) && fEq(gabs(B), 0.0) {
		return NewPoint(m.a.x, m.a.y)
	}

	tmp := CD.Mul(B).Mul(1.0 / A)
	return m.a.Add(tmp)
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_2_C
func CrossPointSegmentSegment(l, m *Segment) *Point {
	a := NewLine(l.a, l.b)
	b := NewLine(m.a, m.b)
	return CrossPointLineLine(a, b)
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_7_D
// func CrossPointsCircleLine(c *Circle, l *Line) []*Point {

// }

// func CrossPointsCircleSegment(c *Circle, l *Segment) []*Point {

// }

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_7_E
// func CrossPointsCircleCircle(c1, c2 *Circle) []*Point {

// }

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_7_F
// 点 p を通る円 c の接線
// func Tangent(c1 *Circle, p2 *Point) []*Point {

// }

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_7_G
// 円 c1, c2 の共通接線
// func Tangent(c1, c2 *Circle) []*Line {

// }

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_3_B
// 凸性判定
func IsConvex(P []*Point) bool {
	n := len(P)
	for i := 0; i < n; i++ {
		a, b, c := P[(i+n-1)%n], P[i], P[(i+1)%n]
		if Ccw(a, b, c) == G_CLOCKWISE {
			return false
		}
	}
	return true
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_4_A
// 凸包
func ConvexHull(P []*Point) []*Point {
	n := len(P)
	k := 0

	if n <= 2 {
		return P
	}

	sort.Slice(P, func(i, j int) bool {
		return pLess(P[i], P[j])
	})

	ch := make([]*Point, 2*n)
	for i := 0; i < n; i++ {
		for k >= 2 {
			CB := ch[k-1].Minus(ch[k-2])
			BA := P[i].Minus(ch[k-1])
			if Cross(CB, BA) > 0.0 {
				break
			}
			k--
		}

		ch[k] = P[i]
		k++
	}
	for i, t := n-2, k; i >= 0; i-- {
		for k >= t {
			CB := ch[k-1].Minus(ch[k-2])
			BA := P[i].Minus(ch[k-1])
			if Cross(CB, BA) > 0.0 {
				break
			}
			k--
		}

		ch[k] = P[i]
		k++
	}

	return ch[:k-1]
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_3_C
// 多角形と点の包含判定
func Contains(P []*Point, q *Point) int {
	n := len(P)
	isIn := false

	for i := 0; i < n; i++ {
		a := P[i].Minus(q)
		b := P[(i+1)%n].Minus(q)

		if a.y > b.y {
			a, b = b, a
		}

		cross := Cross(a, b)
		dot := Dot(a, b)
		if a.y <= 0.0 && 0.0 < b.y && !fEq(cross, 0.0) && cross < 0.0 {
			isIn = !isIn
		}
		if fEq(cross, 0.0) && dot <= 0.0 {
			return G_ON
		}
	}

	if isIn {
		return G_IN
	}
	return G_OUT
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=1033
// 線分の重複除去
// func MergeSegments(segs []*Segment) {

// }

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=1033
// 線分アレンジメント
// 任意の2線分の交点を頂点としたグラフを構築する
// func SegmentArrangement(segs []*Segment, ps []*Point) [][]int {

// }

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_4_C
// 凸多角形の切断
// 直線 l.a-l.b で切断しその左側にできる凸多角形を返す
func ConvexCut(U []*Point, l *Line) []*Point {
	n := len(U)
	ret := []*Point{}

	for i := 0; i < n; i++ {
		a, b := U[i], U[(i+1)%n]
		now := NewPoint(a.x, a.y)
		nxt := NewPoint(b.x, b.y)

		if Ccw(l.a, l.b, now) != G_CLOCKWISE {
			ret = append(ret, now)
		}
		if Ccw(l.a, l.b, now)*Ccw(l.a, l.b, nxt) < 0 {
			cp := CrossPointLineLine(NewLine(now, nxt), l)
			ret = append(ret, cp)
		}
	}

	return ret
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_3_A
// 多角形の面積
func AreaPolygon(P []*Point) float64 {
	n := len(P)
	A := 0.0
	for i := 0; i < n; i++ {
		A += Cross(P[i], P[(i+1)%n])
	}
	return A
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_7_H
// 円と多角形の共通部分の面積
// func AreaPolygonCircle(P []*Point) float64 {

// }

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_4_B
// 凸多角形の直径(最遠頂点対間距離)
func ConvexDiameter(P []*Point) (maxDist float64, mi, mj int) {
	n := len(P)

	is, js := 0, 0
	for i := 1; i < n; i++ {
		if P[i].y > P[is].y {
			is = i
		}
		if P[i].y < P[js].y {
			js = i
		}
	}
	maxdis := P[is].Minus(P[js]).Norm2()

	i, maxi := is, is
	j, maxj := js, js
	for {
		if Cross(P[(i+1)%n].Minus(P[i]), P[(j+1)%n].Minus(P[j])) >= 0.0 {
			j = (j + 1) % n
		} else {
			i = (i + 1) % n
		}

		if P[i].Minus(P[j]).Norm2() > maxdis {
			maxdis = P[i].Minus(P[j]).Norm2()
			maxi, maxj = i, j
		}

		if !(i != is || j != js) {
			break
		}
	}

	return gsqrt(maxdis), maxi, maxj
}

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=CGL_5_A
// 最近点対
func ClosestPair(P []*Point) float64 {
	var _rec func(P []*Point, l, r int) float64
	_rec = func(P []*Point, l, r int) float64 {
		if r-l <= 1 {
			return 1e60
		}

		mid := (l + r) / 2
		x := P[mid].x
		d := gmin(_rec(P, l, mid), _rec(P, mid, r))

		// merge by order of y Pointinate.
		L, R := []*Point{}, []*Point{}
		for i := l; i < r; i++ {
			if i < mid {
				L = append(L, P[i])
			} else {
				R = append(R, P[i])
			}
		}
		cur, j := l, 0
		for i := 0; i < len(L); i++ {
			for j < len(R) && L[i].y > R[j].y {
				P[cur] = R[j]
				cur, j = cur+1, j+1
			}

			P[cur] = L[i]
			cur++
		}
		for ; j < len(R); j++ {
			P[cur] = R[j]
			cur++
		}

		nearLine := []*Point{}
		for i := l; i < r; i++ {
			if gabs(P[i].x-x) >= d {
				continue
			}

			sz := len(nearLine)
			for j := sz - 1; j >= 0; j-- {
				dx := P[i].x - nearLine[j].x
				dy := P[i].y - nearLine[j].y
				if dy >= d {
					break
				}
				d = gmin(d, gsqrt(dx*dx+dy*dy))
			}
			nearLine = append(nearLine, P[i])
		}

		return d
	}

	sort.Slice(P, func(i, j int) bool {
		return pLess(P[i], P[j])
	})

	return _rec(P, 0, len(P))
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

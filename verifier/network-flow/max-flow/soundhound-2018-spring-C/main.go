/*
URL:
https://atcoder.jp/contests/soundhound2018/tasks/soundhound2018_c
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

	h, w int
	C    [][]rune
)

func toid(y, x int) int { return y*w + x }
func toy(id int) int    { return id / w }
func tox(id int) int    { return id % w }

func main() {
	defer stdout.Flush()

	h, w = readi2()
	for i := 0; i < h; i++ {
		row := readrs()
		C = append(C, row)
	}

	left, right := solve(0), solve(1)
	debugf("0: %d, 1: %d\n", left, right)
	// ans := max(left, right)
	// ans := min(left, right)
	// ans := left
	ans := right

	// if h == 1 && w == 1 && C[0][0] == '.' {
	// 	chmax(&ans, 1)
	// }

	total := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if C[i][j] == '.' {
				total++
			}
		}
	}

	println(total - ans)
}

func solve(parity int) int {
	mf := NewMaxFlow(h*w + 2)

	sid, tid := h*w, h*w+1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if C[i][j] == '*' {
				continue
			}

			cid := toid(i, j)
			if (i+j)%2 == parity {
				mf.AddEdge(sid, cid, 1)
			} else {
				mf.AddEdge(cid, tid, 1)
			}
		}
	}

	steps := [][]int{
		{-1, 0}, {1, 0}, {0, 1}, {0, -1},
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if C[i][j] == '*' {
				continue
			}

			cid := toid(i, j)
			if (i+j)%2 == parity {
				for _, st := range steps {
					dy, dx := st[0], st[1]
					ny, nx := i+dy, j+dx
					// if 0 <= ny && ny < h && 0 <= nx && nx < w && C[ny][nx] == '.' {
					if 0 <= ny && ny < h && 0 <= nx && nx < w {
						nid := toid(ny, nx)
						mf.AddEdge(cid, nid, 1)
					}
				}
			}
		}
	}

	ans := mf.Flow(sid, tid)

	return ans
}

type Edge struct {
	from int
	to   int
	capa int
	flow int
}

func NewMaxFlow(n int) *MaxFlow {
	return &MaxFlow{
		n: n,
		g: make([][]_Edge, n),
	}
}

func (mf *MaxFlow) AddEdge(from, to, capa int) int {
	m := len(mf.pos)
	mf.pos = append(mf.pos, [2]int{from, len(mf.g[from])})
	mf.g[from] = append(mf.g[from], _Edge{to, len(mf.g[to]), capa})
	mf.g[to] = append(mf.g[to], _Edge{from, len(mf.g[from]) - 1, 0})
	return m
}

func (mf *MaxFlow) GetEdge(i int) Edge {
	_e := mf.g[mf.pos[i][0]][mf.pos[i][1]]
	_re := mf.g[_e.to][_e.rev]
	return Edge{mf.pos[i][0], _e.to, _e.capa + _re.capa, _re.capa}
}

func (mf *MaxFlow) EdgesList() []Edge {
	m := len(mf.pos)
	result := make([]Edge, 0, m)
	for i := 0; i < m; i++ {
		result = append(result, mf.GetEdge(i))
	}
	return result
}

func (mf *MaxFlow) ChangeEdge(i, newCapa, newFlow int) {
	_e := &mf.g[mf.pos[i][0]][mf.pos[i][1]]
	_re := &mf.g[_e.to][_e.rev]
	_e.capa = newCapa - newFlow
	_re.capa = newFlow
}

func (mf *MaxFlow) Flow(s, t int) int {
	return mf.FlowL(s, t, int(1e+18))
}

func (mf *MaxFlow) FlowL(s, t, flowLim int) int {
	level := make([]int, mf.n)
	iter := make([]int, mf.n)
	bfs := func() {
		for i := range level {
			level[i] = -1
		}
		level[s] = 0
		q := make([]int, 0, mf.n)
		q = append(q, s)
		for len(q) != 0 {
			v := q[0]
			q = q[1:]
			for _, e := range mf.g[v] {
				if e.capa == 0 || level[e.to] >= 0 {
					continue
				}
				level[e.to] = level[v] + 1
				if e.to == t {
					return
				}
				q = append(q, e.to)
			}
		}
	}
	var dfs func(v, up int) int
	dfs = func(v, up int) int {
		if v == s {
			return up
		}
		res := 0
		lv := level[v]
		for ; iter[v] < len(mf.g[v]); iter[v]++ {
			e := &mf.g[v][iter[v]]
			if lv <= level[e.to] || mf.g[e.to][e.rev].capa == 0 {
				continue
			}
			d := dfs(e.to, mf._smaller(up-res, mf.g[e.to][e.rev].capa))
			if d <= 0 {
				continue
			}
			mf.g[v][iter[v]].capa += d
			mf.g[e.to][e.rev].capa -= d
			res += d
			if res == up {
				break
			}
		}
		return res
	}
	flow := 0
	for flow < flowLim {
		bfs()
		if level[t] == -1 {
			break
		}
		for i := range iter {
			iter[i] = 0
		}
		for flow < flowLim {
			f := dfs(t, flowLim-flow)
			if f == 0 {
				break
			}
			flow += f
		}
	}
	return flow
}

func (mf *MaxFlow) MinCut(s int) []bool {
	visited := make([]bool, mf.n)
	q := make([]int, 0, mf.n)
	q = append(q, s)
	for len(q) != 0 {
		p := q[0]
		q = q[1:]
		visited[p] = true
		for _, e := range mf.g[p] {
			if e.capa > 0 && !visited[e.to] {
				visited[e.to] = true
				q = append(q, e.to)
			}
		}
	}
	return visited
}

type _Edge struct {
	to   int
	rev  int
	capa int
}

type MaxFlow struct {
	n   int
	pos [][2]int
	g   [][]_Edge
}

func (mf *MaxFlow) _smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
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

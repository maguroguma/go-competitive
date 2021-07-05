/*
URL:
https://atcoder.jp/contests/abc149/tasks/abc149_f
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

	n int

	G   [][]int
	M   [200000 + 100]int
	sub [200000 + 100]int
	E   [200000 + 100]int

	deg [200000 + 100]int
)

func main() {
	defer stdout.Flush()

	n = readi()
	G = make([][]int, n)
	for i := 0; i < n-1; i++ {
		a, b := readi2()
		a, b = a-1, b-1

		G[a] = append(G[a], b)
		G[b] = append(G[b], a)

		deg[a]++
		deg[b]++
	}

	for i := 0; i <= n; i++ {
		tmp := modpow(2, i, MOD)
		M[i] = ModInv(tmp, MOD)
	}

	ti := T(0)
	merge := func(l, r T) T {
		res := l + r
		return res
	}
	addNode := func(t T, idx int) T {
		res := T(int(t) + 1)
		return res
	}
	solver := NewReRooting(n, G, ti, merge, addNode)
	solver.Solve()

	for i := 0; i < n; i++ {
		// debugf("solver.dp[i]: %v\n", solver.dp[i])
		sub[i] = int(solver.dp[i])
	}

	rec(0, -1)

	ans := 0
	for i := 0; i < n; i++ {
		ans = mod(ans+E[i], MOD)
	}
	println(ans)
}

func rec(cid, pid int) {
	res := M[1]

	children := 0
	C := []int{} // 部分木に少なくとも1つ以上の黒がある確率
	D := []int{} // 部分木がすべて白の確率

	for _, nid := range G[cid] {
		if nid == pid {
			continue
		}
		num := sub[nid]
		children += num
		tmp := mod(1-M[num], MOD)

		C = append(C, tmp)
		D = append(D, mod(1-tmp, MOD))

		rec(nid, cid)
	}

	if pid != -1 {
		num := n - 1 - children
		tmp := mod(1-M[num], MOD)

		C = append(C, tmp)
		D = append(D, mod(1-tmp, MOD))
	}

	if deg[cid] > 1 {
		P := ProdOthers(D, func(x, y int) int {
			return (x * y) % MOD
		})
		ss := 0
		for i := 0; i < len(D); i++ {
			tmp := mod(P[i]*C[i], MOD)
			// tmp = mod(1-tmp, MOD)
			ss += tmp
			ss %= MOD
		}
		ss += mod(P[0]*D[0], MOD)
		ss %= MOD

		ss = mod(1-ss, MOD)
		res = mod(res*ss, MOD)

		E[cid] = res
	} else {
		E[cid] = 0
	}
}

// ProdOthers returns B that B[i] denotes f(A[:i]..., A[i+1:]...).
// Time complexity: O(n)
func ProdOthers(A []int, f func(x, y int) int) (B []int) {
	if len(A) < 2 {
		panic("A must be have more than one element")
	}

	n := len(A)
	L, R := make([]int, n), make([]int, n)

	L[0] = A[0]
	for i := 1; i < n; i++ {
		L[i] = f(L[i-1], A[i])
	}
	R[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		R[i] = f(R[i+1], A[i])
	}

	B = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			B[0] = R[1]
			continue
		}
		if i == n-1 {
			B[n-1] = L[n-2]
			continue
		}

		B[i] = f(L[i-1], R[i+1])
	}

	return B
}

// type T int
type T int

type ReRooting struct {
	n int
	G [][]int

	ti      T
	dp, res []T
	merge   func(l, r T) T
	addNode func(t T, idx int) T
}

func NewReRooting(
	n int, AG [][]int, ti T, merge func(l, r T) T, addNode func(t T, idx int) T,
) *ReRooting {
	s := new(ReRooting)
	s.n, s.G, s.ti, s.merge, s.addNode = n, AG, ti, merge, addNode
	s.dp, s.res = make([]T, n), make([]T, n)

	s.Solve()

	return s
}

func (s *ReRooting) Solve() {
	s.inOrder(0, -1)
	s.reroot(0, -1, s.ti)
}

func (s *ReRooting) Query(idx int) T {
	return s.res[idx]
}

func (s *ReRooting) inOrder(cid, pid int) T {
	res := s.ti

	for _, nid := range G[cid] {
		if nid == pid {
			continue
		}

		res = s.merge(res, s.inOrder(nid, cid))
	}
	res = s.addNode(res, cid)
	s.dp[cid] = res

	return s.dp[cid]
}

func (s *ReRooting) reroot(cid, pid int, parentValue T) {
	childValues := []T{}
	nexts := []int{}
	for _, nid := range G[cid] {
		if nid == pid {
			continue
		}
		childValues = append(childValues, s.dp[nid])
		nexts = append(nexts, nid)
	}

	// result of cid
	rootValue := s.ti
	for _, v := range childValues {
		rootValue = s.merge(rootValue, v)
	}
	rootValue = s.merge(rootValue, parentValue)
	rootValue = s.addNode(rootValue, cid)
	s.res[cid] = rootValue

	// for children
	accum := s.merge(s.ti, parentValue)
	length := len(childValues)
	if length == 0 {
		return
	}
	if length == 1 {
		s.reroot(nexts[0], cid, s.addNode(accum, cid))
		return
	}

	// cid has more than one child
	R, L := make([]T, length), make([]T, length)
	L[0] = s.merge(s.ti, childValues[0])
	for i := 1; i < length; i++ {
		L[i] = s.merge(L[i-1], childValues[i])
	}
	R[length-1] = s.merge(s.ti, childValues[length-1])
	for i := length - 2; i >= 0; i-- {
		R[i] = s.merge(R[i+1], childValues[i])
	}

	for i, nid := range nexts {
		if i == 0 {
			s.reroot(nid, cid, s.addNode(s.merge(accum, R[1]), cid))
		} else if i == length-1 {
			s.reroot(nid, cid, s.addNode(s.merge(accum, L[length-2]), cid))
		} else {
			s.reroot(nid, cid, s.addNode(s.merge(accum, s.merge(L[i-1], R[i+1])), cid))
		}
	}
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

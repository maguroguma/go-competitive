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
	n int
	G [100000 + 5][]Edge
	T [][]int
	q int

	prods []int
)

type Edge struct {
	to int
	c  int
}

func main() {
	n = ReadInt()
	T = make([][]int, n)
	for i := 0; i < n-1; i++ {
		u, v, c := ReadInt3()
		u--
		v--
		G[u] = append(G[u], Edge{to: v, c: c})
		G[v] = append(G[v], Edge{to: u, c: c})
		T[u] = append(T[u], v)
		T[v] = append(T[v], u)
	}

	prods = make([]int, n)
	prods[0] = 1
	dfs(0, -1, -1, 1)

	s := NewLCASolver(T, 0, n)

	q := ReadInt()
	for i := 0; i < q; i++ {
		m, p, x := ReadInt3()
		m--
		p--
		// res = make([]int, n)
		// res[m] = 1
		// dfs(m, -1, p, 1)
		// x *= res[p]
		// x %= MOD
		// fmt.Println(x)

		l := s.LCA(m, p)
		tmp := prods[m] * prods[p]
		tmp %= MOD
		tmp *= ModInv(prods[l], MOD)
		tmp %= MOD
		tmp *= ModInv(prods[l], MOD)
		tmp %= MOD
		x *= tmp
		x %= MOD
		fmt.Println(x)
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

func dfs(cid, pid, tid, cprod int) {
	for _, e := range G[cid] {
		if e.to == pid {
			continue
		}
		tmp := cprod * e.c
		tmp %= MOD
		prods[e.to] = tmp
		dfs(e.to, cid, tid, tmp)
	}
}

const (
	MAX_V     = 100000 + 5 // maximum node number of a tree
	MAX_LOG_V = 100 + 1    // maximum log(n)
)

type LCASolver struct {
	// Graph info
	G    [][]int // graph as adjacent list
	root int     // root node ID
	n    int     // node number

	// data structure for answer LCA
	parent [MAX_LOG_V][MAX_V]int
	depth  [MAX_V]int
}

func NewLCASolver(G [][]int, root, n int) *LCASolver {
	s := new(LCASolver)
	s.G, s.root, s.n = G, root, n
	s.initialize()
	return s
}

func (s *LCASolver) initialize() {
	s.dfs(s.root, -1, 0)
	for k := 0; k+1 < MAX_LOG_V; k++ {
		for v := 0; v < s.n; v++ {
			if s.parent[k][v] < 0 {
				s.parent[k+1][v] = -1
			} else {
				s.parent[k+1][v] = s.parent[k][s.parent[k][v]]
			}
		}
	}
}

func (s *LCASolver) dfs(v, p, d int) {
	s.parent[0][v] = p
	s.depth[v] = d
	for _, to := range s.G[v] {
		if to != p {
			s.dfs(to, v, d+1)
		}
	}
}

func (s *LCASolver) LCA(u, v int) int {
	if s.depth[u] > s.depth[v] {
		u, v = v, u
	}
	for k := 0; k < MAX_LOG_V; k++ {
		if ((s.depth[v] - s.depth[u]) >> uint(k) & 1) == 1 {
			v = s.parent[k][v]
		}
	}

	if u == v {
		return u
	}

	for k := MAX_LOG_V - 1; k >= 0; k-- {
		if s.parent[k][u] != s.parent[k][v] {
			u, v = s.parent[k][u], s.parent[k][v]
		}
	}

	return s.parent[0][u]
}

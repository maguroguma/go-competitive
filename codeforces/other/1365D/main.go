/*
URL:
https://codeforces.com/problemset/problem/1365/D
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

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

var (
	t int

	n, m int
	R    [][]rune

	N int
	G [][]int

	steps [4][2]int
)

func main() {
	steps = [4][2]int{
		[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0},
	}

	t = ReadInt()
	for i := 0; i < t; i++ {
		n, m = ReadInt2()
		R = [][]rune{}
		for j := 0; j < n; j++ {
			row := ReadRuneSlice()
			R = append(R, row)
		}

		solve()
	}
}

func solve() {
	// Gが一人もいなかったら "Yes"
	gcnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if R[i][j] == 'G' {
				gcnt++
			}
		}
	}
	if gcnt == 0 {
		fmt.Println("Yes")
		return
	}

	// Bの周りを埋められるなら埋める
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if R[i][j] != 'B' {
				continue
			}

			for _, s := range steps {
				dy, dx := s[0], s[1]
				ny, nx := i+dy, j+dx

				// 埋め立て
				if 0 <= ny && ny < n && 0 <= nx && nx < m && R[ny][nx] == '.' {
					R[ny][nx] = '#'
				}
			}
		}
	}

	// グラフ作成
	N = n * m
	G = make([][]int, N)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 壁は無視
			if R[i][j] == '#' {
				continue
			}

			cid := m*i + j

			for _, s := range steps {
				dy, dx := s[0], s[1]
				ny, nx := i+dy, j+dx
				// 壁は駄目
				if 0 <= ny && ny < n && 0 <= nx && nx < m && R[ny][nx] != '#' {
					nid := m*ny + nx

					G[cid] = append(G[cid], nid)
				}
			}
		}
	}

	if R[n-1][m-1] == '#' {
		fmt.Println("No")
		return
	}

	// ゴールからの最短路
	// _, visited := bfs()
	_, visited := SSSPByBFS(n*m-1, N, G[:N])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			id := m*i + j
			// Gに到達できないときは NO
			if R[i][j] == 'G' && !visited[id] {
				fmt.Println("No")
				return
			}
			// Bに到達できるときは NO
			if R[i][j] == 'B' && visited[id] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

func SSSPByBFS(sid, n int, AG [][]int) (dp []int, visited []bool) {
	dp = make([]int, n)
	visited = make([]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = INF_BIT30
		visited[i] = false
	}

	Q := []Node{}
	dp[sid] = 0
	visited[sid] = true
	Q = append(Q, Node{id: sid, cost: dp[sid]})

	for len(Q) > 0 {
		cnode := Q[0]
		Q = Q[1:]

		for _, nid := range G[cnode.id] {
			// 訪問済みならパス
			if visited[nid] {
				continue
			}

			dp[nid] = cnode.cost + 1
			visited[nid] = true
			Q = append(Q, Node{id: nid, cost: dp[nid]})
		}
	}

	return dp, visited
}

type Node struct {
	id, cost int
}

/*******************************************************************/

/*********** I/O ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
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

// PrintfDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintfDebug(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// PrintfBufStdout is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func PrintfBufStdout(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}

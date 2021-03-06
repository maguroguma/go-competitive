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

var (
	n    int
	A, B []int

	T [][]int
	G [][]int
)

func main() {
	n = ReadInt()

	// 2, 3頂点の場合はどんな順列でもいい
	if n == 2 {
		fmt.Println("1 2")
		return
	}
	if n == 3 {
		fmt.Println("1 2 3")
		return
	}

	A, B = make([]int, n-1), make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a, b := ReadInt2()
		a--
		b--
		A[i], B[i] = a, b
	}

	T = make([][]int, n)
	edges = make([][]Edge, n)
	for i := 0; i < n-1; i++ {
		a, b := A[i], B[i]
		T[a] = append(T[a], b)
		T[b] = append(T[b], a)

		edges[a] = append(edges[a], Edge{nid: b, weight: 1})
		edges[b] = append(edges[b], Edge{nid: a, weight: 1})
	}

	// 直径を調べる
	dia := Diameter()
	// PrintDebug("diameter: %d\n", dia)
	if dia < 3 {
		// どんな順列でもいい
		answers := []int{}
		for i := 1; i <= n; i++ {
			answers = append(answers, i)
			fmt.Println(PrintIntsLine(answers...))
			return
		}
	}

	// 距離3のノード間にエッジを張ったグラフを作る
	G = make([][]int, n)
	for i := 0; i < n; i++ {
		dfsThree(i, i, -1, 0)
	}
	// debug
	// for i := 0; i < n; i++ {
	// 	PrintDebug("node %d: %v\n", i, G[i])
	// }

	// Gに対して二部グラフ彩色を行う
	_, colors := BipartiteJudge(n)
	// PrintDebug("colors: %v\n", colors)

	// 割り振りを行う
	white, black := 0, 0
	for i := 0; i < len(colors); i++ {
		if colors[i] == 1 {
			white++
		} else {
			black++
		}
	}
	larger, smaller := white, black
	bef, after := 1, -1
	if larger < smaller {
		larger, smaller = smaller, larger
		bef, after = after, bef
	}
	// 先にbefの色からあまり1を割り振る
	answers := make([]int, n)
	curVal := 1
	for i := 0; i < len(colors); i++ {
		if colors[i] == bef {
			answers[i] = curVal
			curVal += 3
			if curVal > n {
				break
			}
		}
	}
	if curVal <= n {
		fmt.Println(-1)
		return
	}
	// 次にafterの色にあまり2を割り振る
	curVal = 2
	for i := 0; i < len(colors); i++ {
		if colors[i] == after {
			answers[i] = curVal
			curVal += 3
			if curVal > n {
				break
			}
		}
	}
	if curVal <= n {
		fmt.Println(-1)
		return
	}
	// 最後に埋まっていない部分に3の倍数を割り振る
	curVal = 3
	for i := 0; i < n; i++ {
		if answers[i] == 0 {
			answers[i] = curVal
			curVal += 3
		}
	}
	fmt.Println(PrintIntsLine(answers...))
}

func dfsThree(sid, cid, pid, depth int) {
	if depth == 3 {
		// 探索開始ノードに追加して終了
		G[sid] = append(G[sid], cid)
		return
	}

	for _, nid := range T[cid] {
		if nid == pid {
			continue
		}
		dfsThree(sid, nid, cid, depth+1)
	}
}

/* 二部グラフ彩色 */

// v: ノード数
// 隣接リストedgesをグローバルに設定する必要がある
func BipartiteJudge(v int) (bool, []int) {
	colors := make([]int, v)

	for i := 0; i < v; i++ {
		if colors[i] == 0 {
			// まだ頂点iが塗られていなければ1で塗る
			if !dfs(i, 1, colors) {
				return false, colors
			}
		}
	}

	return true, colors
}

// 頂点を1, -1で塗っていく
func dfs(v, c int, colors []int) bool {
	colors[v] = c
	// for _, nid := range edges[v] {
	for _, nid := range G[v] {
		// 隣接している頂点が同じ色ならfalse
		if colors[nid] == c {
			return false
		}
		// 隣接している頂点がまだ塗られていないなら-cで塗る
		if colors[nid] == 0 && !dfs(nid, -c, colors) {
			return false
		}
	}

	// すべての頂点を塗れたらtrue
	return true
}

/* 直径調査用 */

var edges [][]Edge

type Edge struct {
	// nid: 向き先ノードID, weight: 重み
	nid, weight int
}

type Result struct {
	// dist: 距離, nid: 終点ノードID
	dist, nid int
}

// 木の直径を返す
// O(|E|)
func Diameter() int {
	r := visit(-1, 0)     // nodeID: 0からの最遠ノード(とその距離)を計算
	t := visit(-1, r.nid) // 0からの最遠ノードからの最遠ノードとその距離を計算
	return t.dist         // 最遠距離のみを返す
}

// pidからcidに遷移したときの、cidからの最遠ノードを返す
// pid: 直前の遷移元ノードID, cid: 現在観ているノードID
func visit(pid, cid int) Result {
	r := Result{dist: 0, nid: cid}
	// DFS
	for _, e := range edges[cid] {
		if e.nid != pid {
			t := visit(cid, e.nid) // 次の遷移先へ
			t.dist += e.weight
			if r.dist < t.dist {
				r = t
			}
		}
	}
	return r
}

/*
- まずは全探索を検討しましょう
- MODは最後にとりましたか？
- 負のMODはちゃんと関数を使って処理していますか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
- いきなりオーバーフローはしていませんか？
- MOD取る系はint64必須ですよ？
- 後ろ・逆・ゴールから考えましたか？
- 3者のうち真ん中に着目しましたか？
*/

/*******************************************************************/

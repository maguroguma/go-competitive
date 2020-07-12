/*
URL:
https://atcoder.jp/contests/code-festival-2015-qualb/tasks/codefestival_2015_qualB_d
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
	MOD = 1000000000 + 7
	// MOD          = 998244353
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
	n    int
	S, C []int

	R map[int]int
)

func main() {
	n = ReadInt()
	for i := 0; i < n; i++ {
		s, c := ReadInt2()
		S = append(S, s)
		C = append(C, c)
	}

	R = make(map[int]int)
	tr := NewTreap()
	for i := 0; i < n; i++ {
		s, c := S[i], C[i]

	}
}

// Treap usage
// tr := NewTreap()
// tr.Insert(k)
// node := tr.Find(k)
// min := tr.FindMinimum()
// max := tr.FindMaximum()
// tr.Delete(k)
// node := tr.MinGeq(x)
// node := tr.MinGreater(x)
// node := tr.MaxLeq(x)
// node := tr.MaxLess(x)
// fmt.Println(PrintIntsLine(tr.Inorder()...))
// fmt.Println(PrintIntsLine(tr.Preorder()...))
// tr.InsertBySettingPri(k, p)

type Node struct {
	key, priority int
	right, left   *Node
}

type Treap struct {
	root *Node
}

/*************************************/
// Public method
/*************************************/

// NewTreap returns a pointer of a Treap instance.
func NewTreap() *Treap {
	tr := new(Treap)
	tr.root = nil
	return tr
}

// InsertBySettingPri method inserts a new node consisting of new key and priority.
// A duplicate key is ignored and nothing happens.
func (tr *Treap) InsertBySettingPri(key, priority int) {
	tr.root = tr.insert(tr.root, key, priority)
}

// for XorShift
var _gtx, _gty, _gtz, _gtw = 123456789, 362436069, 521288629, 88675123

// Insert method inserts a new node consisting o new key.
// The priority is automatically set by random value.
// A duplicate key is ignored and nothing happens.
func (tr *Treap) Insert(key int) {
	// XorShiftによる乱数生成
	// 下記URLを参考
	// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
	randInt := func() int {
		tt := (_gtx ^ (_gtx << 11))
		_gtx = _gty
		_gty = _gtz
		_gtz = _gtw
		_gtw = (_gtw ^ (_gtw >> 19)) ^ (tt ^ (tt >> 8))
		return _gtw
	}

	tr.root = tr.insert(tr.root, key, randInt())
}

// Find returns a node that has an argument key value.
// Find returns nil when there is no node that has an argument key value.
func (tr *Treap) Find(k int) *Node {
	u := tr.root
	for u != nil && k != u.key {
		if k < u.key {
			u = u.left
		} else {
			u = u.right
		}
	}
	return u
}

// FindMinimum returns a node that has the minimum key in the treap.
// FindMinimum returns nil when there is no nodes.
func (tr *Treap) FindMinimum() *Node {
	u := tr.root
	for u != nil && u.left != nil {
		u = u.left
	}
	return u
}

// FindMaximum returns a node that has the maximum key in the treap.
// FindMaximum returns nil when there is no nodes.
func (tr *Treap) FindMaximum() *Node {
	u := tr.root
	for u != nil && u.right != nil {
		u = u.right
	}
	return u
}

// Delete method deletes a node that has an argument key value.
// A duplicate key is ignored and nothing happens.
func (tr *Treap) Delete(key int) {
	tr.root = tr.delete(tr.root, key)
}

// Inorder returns a slice consisting of treap nodes in order of INORDER.
// The nodes are sorted by key values.
func (tr *Treap) Inorder() []int {
	res := make([]int, 0, 200000+5)
	tr.inorder(tr.root, &res)
	return res
}

// Preorder returns a slice consisting of treap nodes in order of PREORDER.
func (tr *Treap) Preorder() []int {
	res := make([]int, 0, 200000+5)
	tr.preorder(tr.root, &res)
	return res
}

// MinGeq returns a node that has MINIMUM KEY MEETING key >= x.
// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
func (tr *Treap) MinGeq(x int) *Node {
	return tr.biggerLowerBound(tr.root, x)
}

// MinGreater returns a node that has MINIMUM KEY MEETING key > x.
// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
func (tr *Treap) MinGreater(x int) *Node {
	return tr.biggerUpperBound(tr.root, x)
}

// MaxLeq returns a node that has MAXIMUM KEY MEETING key <= x.
// for AGC005-B
func (tr *Treap) MaxLeq(x int) *Node {
	return tr.smallerUpperBound(tr.root, x)
}

// MaxLess returns a node that has MAXIMUM KEY MEETING key < x.
// for AGC005-B
func (tr *Treap) MaxLess(x int) *Node {
	return tr.smallerLowerBound(tr.root, x)
}

/*************************************/
// Private method
/*************************************/

func (tr *Treap) insert(t *Node, key, priority int) *Node {
	// 葉に到達したら新しい節点を生成して返す
	if t == nil {
		node := new(Node)
		node.key, node.priority = key, priority
		return node
	}

	// 重複したkeyは無視
	if key == t.key {
		return t
	}

	if key < t.key {
		// 左の子へ移動
		t.left = tr.insert(t.left, key, priority) // 左の子へのポインタを更新
		// 左の子の方が優先度が高い場合右回転
		if t.priority < t.left.priority {
			t = tr.rightRotate(t)
		}
	} else {
		// 右の子へ移動
		t.right = tr.insert(t.right, key, priority) // 右の子へのポインタを更新
		if t.priority < t.right.priority {
			// 右の子の方が優先度が高い場合左回転
			t = tr.leftRotate(t)
		}
	}

	return t
}

// 削除対象の節点を回転によって葉まで移動させた後に削除する
func (tr *Treap) delete(t *Node, key int) *Node {
	if t == nil {
		return nil
	}

	// 削除対象を検索
	if key < t.key {
		t.left = tr.delete(t.left, key)
	} else if key > t.key {
		t.right = tr.delete(t.right, key)
	} else {
		// 削除対象を発見、葉ノードとなるように回転を繰り返す
		return tr._delete(t, key)
	}

	return t
}

// 削除対象の節点の場合
func (tr *Treap) _delete(t *Node, key int) *Node {
	if t.left == nil && t.right == nil {
		// 葉の場合
		return nil
	} else if t.left == nil {
		// 右の子のみを持つ場合は左回転
		t = tr.leftRotate(t)
	} else if t.right == nil {
		// 左の子のみを持つ場合は右回転
		t = tr.rightRotate(t)
	} else {
		// 優先度が高い方を持ち上げる
		if t.left.priority > t.right.priority {
			t = tr.rightRotate(t)
		} else {
			t = tr.leftRotate(t)
		}
	}

	return tr.delete(t, key)
}

func (tr *Treap) rightRotate(t *Node) *Node {
	s := t.left
	t.left = s.right
	s.right = t
	return s
}

func (tr *Treap) leftRotate(t *Node) *Node {
	s := t.right
	t.right = s.left
	s.left = t
	return s
}

// rootからスタートする
func (tr *Treap) biggerLowerBound(t *Node, x int) *Node {
	if t == nil {
		return nil
	} else if t.key >= x {
		// 探索キーxが現在のノードキー以下の場合、左を探索する
		node := tr.biggerLowerBound(t.left, x)
		if node != nil {
			return node
		} else {
			return t
		}
	} else {
		// 探索キーxが現在のノードキーより大きい場合、右を探索する
		return tr.biggerLowerBound(t.right, x)
	}
}

// rootからスタートする
func (tr *Treap) biggerUpperBound(t *Node, x int) *Node {
	if t == nil {
		return nil
	} else if t.key > x {
		// 探索キーxが現在のノードキーより小さい場合、左を探索する
		node := tr.biggerUpperBound(t.left, x)
		if node != nil {
			return node
		} else {
			return t
		}
	} else {
		// 探索キーxが現在のノードキー以上の場合、右を探索する
		return tr.biggerUpperBound(t.right, x)
	}
}

// rootからスタートする
func (tr *Treap) smallerUpperBound(t *Node, x int) *Node {
	if t == nil {
		return nil
	} else if t.key <= x {
		node := tr.smallerUpperBound(t.right, x)
		if node != nil {
			return node
		} else {
			return t
		}
	} else {
		return tr.smallerUpperBound(t.left, x)
	}
}

// rootからスタートする
func (tr *Treap) smallerLowerBound(t *Node, x int) *Node {
	if t == nil {
		return nil
	} else if t.key < x {
		node := tr.smallerLowerBound(t.right, x)
		if node != nil {
			return node
		} else {
			return t
		}
	} else {
		return tr.smallerLowerBound(t.left, x)
	}
}

func (tr *Treap) inorder(u *Node, res *[]int) {
	if u == nil {
		return
	}
	tr.inorder(u.left, res)
	*res = append(*res, u.key)
	tr.inorder(u.right, res)
}

func (tr *Treap) preorder(u *Node, res *[]int) {
	if u == nil {
		return
	}
	*res = append(*res, u.key)
	tr.preorder(u.left, res)
	tr.preorder(u.right, res)
}

func solve() {
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

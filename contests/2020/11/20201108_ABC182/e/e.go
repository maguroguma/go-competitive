/*
URL:
https://atcoder.jp/contests/abc182/tasks/abc182_e
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

type Item struct {
	y, x int
	kind int
}

const (
	block, light = 0, 1
)

var (
	h, w, n, m int
	A, B, C, D []int

	H [1500 + 50][]Item
	W [1500 + 50][]Item
)

func main() {
	defer stdout.Flush()

	h, w, n, m = readi4()
	A, B = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a, b := readi2()
		a--
		b--
		A[i], B[i] = a, b
	}
	C, D = make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		c, d := readi2()
		c--
		d--
		C[i], D[i] = c, d
	}

	for i := 0; i < n; i++ {
		a, b := A[i], B[i]
		li := Item{a, b, light}
		H[a] = append(H[a], li)
		W[b] = append(W[b], li)
	}
	for i := 0; i < m; i++ {
		c, d := C[i], D[i]
		bl := Item{c, d, block}
		H[c] = append(H[c], bl)
		W[d] = append(W[d], bl)
	}
	// 番兵
	for i := 0; i < h; i++ {
		left := Item{i, -1, block}
		right := Item{i, w, block}
		H[i] = append(H[i], left, right)
	}
	for i := 0; i < w; i++ {
		up := Item{-1, i, block}
		down := Item{h, i, block}
		W[i] = append(W[i], up, down)
	}
	// 行のソート
	for idx := 0; idx < h; idx++ {
		sort.Slice(H[idx], func(i, j int) bool {
			return H[idx][i].x < H[idx][j].x
		})
	}
	// 列のソート
	for idx := 0; idx < w; idx++ {
		sort.Slice(W[idx], func(i, j int) bool {
			return W[idx][i].y < W[idx][j].y
		})
	}

	ans := 0
	// 判定しながら各マスについて数える
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			y, x := i, j

			lidx := BinarySearch(0, len(H[y]), func(mid int) bool {
				return H[y][mid].x <= x
			})
			ridx := BinarySearch(len(H[y]), 0, func(mid int) bool {
				return H[y][mid].x >= x
			})
			uidx := BinarySearch(0, len(W[x]), func(mid int) bool {
				return W[x][mid].y <= y
			})
			didx := BinarySearch(len(W[x]), 0, func(mid int) bool {
				return W[x][mid].y >= y
			})

			L, R, U, D := H[y][lidx], H[y][ridx], W[x][uidx], W[x][didx]
			if L.kind == light || R.kind == light || U.kind == light || D.kind == light {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

func BinarySearch(initOK, initNG int, isOK func(mid int) bool) (ok int) {
	ng := initNG
	ok = initOK
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

// Treap usage
// tr := NewTreap() 				// constructor
// tr.Insert(key) 					// insert one key node
// cnt := tr.Count(key) 		// return a number of key nodes
// node := tr.Find(key) 		// return a pointer
// min := tr.FindMinimum() 	// return a pointer
// max := tr.FindMaximum() 	// return a pointer
// tr.Delete(key) 					// delete one key node
// node := tr.MinGeq(x) 		// return a pointer
// node := tr.MinGreater(x) // return a pointer
// node := tr.MaxLeq(x) 		// return a pointer
// node := tr.MaxLess(x) 		// return a pointer

// fmt.Println(PrintIntsLine(tr.Inorder()...))
// fmt.Println(PrintIntsLine(tr.Preorder()...))
// tr.InsertBySettingPri(key, p)

// type of key
type T struct {
	s, t int
}

type Node struct {
	key         T
	priority    int
	right, left *Node
}

type Treap struct {
	root    *Node
	cnts    map[T]int
	randInt func() int
	less    func(l, r T) bool // *strictly less*
}

/*************************************/
// Public method
/*************************************/

// NewTreap returns a pointer of a Treap instance.
func NewTreap(less func(l, r T) bool) *Treap {
	tr := new(Treap)

	tr.root = nil
	tr.cnts = make(map[T]int)
	tr.less = less

	// XorShiftによる乱数生成
	// 下記URLを参考
	// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
	_gtx, _gty, _gtz, _gtw := 123456789, 362436069, 521288629, 88675123
	tr.randInt = func() int {
		tt := (_gtx ^ (_gtx << 11))
		_gtx = _gty
		_gty = _gtz
		_gtz = _gtw
		_gtw = (_gtw ^ (_gtw >> 19)) ^ (tt ^ (tt >> 8))
		return _gtw
	}

	return tr
}

// Count method returns the number of the key.
// If there hasn't been the key in the treap, this returns 0.
func (tr *Treap) Count(key T) int {
	return tr.cnts[key]
}

// InsertBySettingPri method inserts a new node consisting of new key and priority.
// A duplicate key is ignored and nothing happens.
// func (tr *Treap) InsertBySettingPri(key, priority int) {
// 	tr.root = tr.insert(tr.root, key, priority)
// }

// Insert method inserts a new node consisting o new key.
// The priority is automatically set by random value.
// A duplicate key is ignored and nothing happens.
func (tr *Treap) Insert(key T) {
	preCnt := tr.Count(key)
	tr.increase(key, 1)
	if preCnt > 0 {
		return
	}

	tr.root = tr.insert(tr.root, key, tr.randInt())
}

// Find returns a node that has an argument key value.
// Find returns nil when there is no node that has an argument key value.
func (tr *Treap) Find(key T) *Node {
	cnt := tr.cnts[key]
	if cnt == 0 {
		return nil
	}

	u := tr.root
	for u != nil && key != u.key {
		if tr.less(key, u.key) {
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
func (tr *Treap) Delete(key T) {
	tr.decrease(key, 1)
	curCnt := tr.Count(key)
	if curCnt > 0 {
		return
	}

	tr.root = tr.delete(tr.root, key)
}

// Inorder returns a slice consisting of treap nodes in order of INORDER.
// The nodes are sorted by key values.
func (tr *Treap) Inorder() []T {
	res := make([]T, 0, 200000+5)
	tr.inorder(tr.root, &res)
	return res
}

// Preorder returns a slice consisting of treap nodes in order of PREORDER.
func (tr *Treap) Preorder() []T {
	res := make([]T, 0, 200000+5)
	tr.preorder(tr.root, &res)
	return res
}

// MinGeq returns a node that has MINIMUM KEY MEETING key >= x.
// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
func (tr *Treap) MinGeq(x T) *Node {
	return tr.biggerLowerBound(tr.root, x)
}

// MinGreater returns a node that has MINIMUM KEY MEETING key > x.
// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
func (tr *Treap) MinGreater(x T) *Node {
	return tr.biggerUpperBound(tr.root, x)
}

// MaxLeq returns a node that has MAXIMUM KEY MEETING key <= x.
// for AGC005-B
func (tr *Treap) MaxLeq(x T) *Node {
	return tr.smallerUpperBound(tr.root, x)
}

// MaxLess returns a node that has MAXIMUM KEY MEETING key < x.
// for AGC005-B
func (tr *Treap) MaxLess(x T) *Node {
	return tr.smallerLowerBound(tr.root, x)
}

/*************************************/
// Private method
/*************************************/

func (tr *Treap) increase(key T, num int) {
	tr.cnts[key] += num
}

func (tr *Treap) decrease(key T, num int) {
	curCnt := tr.cnts[key]
	if curCnt-num < 0 {
		panic("too many elements is deleted!")
	}

	tr.cnts[key] -= num
}

func (tr *Treap) insert(t *Node, key T, priority int) *Node {
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

	if tr.less(key, t.key) {
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
func (tr *Treap) delete(t *Node, key T) *Node {
	if t == nil {
		return nil
	}

	// 削除対象を検索
	if key == t.key {
		// 削除対象を発見、葉ノードとなるように回転を繰り返す
		return tr._delete(t, key)
	} else if tr.less(key, t.key) {
		t.left = tr.delete(t.left, key)
	} else {
		t.right = tr.delete(t.right, key)
	}

	return t
}

// 削除対象の節点の場合
func (tr *Treap) _delete(t *Node, key T) *Node {
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
func (tr *Treap) biggerLowerBound(t *Node, x T) *Node {
	if t == nil {
		return nil
	} else if tr.less(t.key, x) {
		// 探索キーxが現在のノードキーより大きい場合、右を探索する
		return tr.biggerLowerBound(t.right, x)
	} else {
		// 探索キーxが現在のノードキー以下の場合、左を探索する
		node := tr.biggerLowerBound(t.left, x)
		if node != nil {
			return node
		} else {
			return t
		}
	}
}

// rootからスタートする
func (tr *Treap) biggerUpperBound(t *Node, x T) *Node {
	if t == nil {
		return nil
	} else if tr.less(t.key, x) || t.key == x {
		// 探索キーxが現在のノードキー以上の場合、右を探索する
		return tr.biggerUpperBound(t.right, x)
	} else {
		// 探索キーxが現在のノードキーより小さい場合、左を探索する
		node := tr.biggerUpperBound(t.left, x)
		if node != nil {
			return node
		} else {
			return t
		}
	}
}

// rootからスタートする
func (tr *Treap) smallerUpperBound(t *Node, x T) *Node {
	if t == nil {
		return nil
	} else if tr.less(t.key, x) || t.key == x {
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
func (tr *Treap) smallerLowerBound(t *Node, x T) *Node {
	if t == nil {
		return nil
	} else if tr.less(t.key, x) {
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

func (tr *Treap) inorder(u *Node, res *[]T) {
	if u == nil {
		return
	}
	tr.inorder(u.left, res)
	*res = append(*res, u.key)
	tr.inorder(u.right, res)
}

func (tr *Treap) preorder(u *Node, res *[]T) {
	if u == nil {
		return
	}
	*res = append(*res, u.key)
	tr.preorder(u.left, res)
	tr.preorder(u.right, res)
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

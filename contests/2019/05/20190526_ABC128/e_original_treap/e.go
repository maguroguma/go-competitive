package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
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
	r.Buffer(make([]byte, 1024), int(1e+11))
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

/*********** DP sub-functions ***********/

// ChMin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

// ChMax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
}

// NthBit returns nth bit value of an argument.
// n starts from 0.
func NthBit(num, nth int) int {
	return num >> uint(nth) & 1
}

/*********** Arithmetic ***********/

// Max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Max(integers ...int) int {
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

// Min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func Min(integers ...int) int {
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

// DigitSum returns digit sum of a decimal number.
// DigitSum only accept a positive integer.
func DigitSum(n int) int {
	if n < 0 {
		return -1
	}

	res := 0

	for n > 0 {
		res += n % 10
		n /= 10
	}

	return res
}

// Sum returns multiple integers sum.
func Sum(integers ...int) int {
	s := 0

	for _, i := range integers {
		s += i
	}

	return s
}

// CeilInt returns the minimum integer larger than or equal to float(a/b).
func CeilInt(a, b int) int {
	res := a / b
	if a%b > 0 {
		res++
	}
	return res
}

// FloorInt returns the maximum integer smaller than or equal to float(a/b)
func FloorInt(a, b int) int {
	res := a / b
	return res
}

// PowInt is integer version of math.Pow
// PowInt calculate a power by Binary Power (二分累乗法(O(log e))).
func PowInt(a, e int) int {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}

	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := PowInt(a, halfE)
		return half * half
	}

	return a * PowInt(a, e-1)
}

// AbsInt is integer version of math.Abs
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Gcd returns the Greatest Common Divisor of two natural numbers.
// Gcd only accepts two natural numbers (a, b >= 1).
// 0 or negative number causes panic.
// Gcd uses the Euclidean Algorithm.
func Gcd(a, b int) int {
	if a <= 0 || b <= 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}
	if a < b {
		a, b = b, a
	}

	// Euclidean Algorithm
	for b > 0 {
		div := a % b
		a, b = b, div
	}

	return a
}

// Lcm returns the Least Common Multiple of two natural numbers.
// Lcd only accepts two natural numbers (a, b >= 1).
// 0 or negative number causes panic.
// Lcd uses the Euclidean Algorithm indirectly.
func Lcm(a, b int) int {
	if a <= 0 || b <= 0 {
		panic(errors.New("[argument error]: Gcd only accepts two NATURAL numbers"))
	}

	// a = a'*gcd, b = b'*gcd, a*b = a'*b'*gcd^2
	// a' and b' are relatively prime numbers
	// gcd consists of prime numbers, that are included in a and b
	gcd := Gcd(a, b)

	// not (a * b / gcd), because of reducing a probability of overflow
	return (a / gcd) * b
}

// Strtoi is a wrapper of `strconv.Atoi()`.
// If `strconv.Atoi()` returns an error, Strtoi calls panic.
func Strtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(errors.New("[argument error]: Strtoi only accepts integer string"))
	} else {
		return i
	}
}

/********** I/O usage **********/

//str := ReadString()
//i := ReadInt()
//X := ReadIntSlice(n)
//S := ReadRuneSlice()

/*******************************************************************/

const MOD = 1000000000 + 7
const ALPHABET_NUM = 26

var n, q int
var S, T, X []int
var D []int

type Item struct {
	key            int
	time, query, x int
}
type ItemList []*Item
type byQuery struct {
	ItemList
}

func (l ItemList) Len() int {
	return len(l)
}
func (l ItemList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l ItemList) Less(i, j int) bool {
	return l[i].key < l[j].key
}
func (b byQuery) Less(i, j int) bool {
	return b.ItemList[i].query < b.ItemList[j].query
}

// how to use
// L := make(ItemList, 0, 100000+1)
// L = append(L, &Item{key: intValue})
// sort.Sort(L)               // ASC
// sort.Sort(sort.Reverse(L)) // DESC

var L ItemList

func main() {
	n, q = ReadInt(), ReadInt()
	S, T, X = make([]int, n), make([]int, n), make([]int, n)
	L = make(ItemList, 0, 200005)
	for i := 0; i < n; i++ {
		s, t, x := ReadInt(), ReadInt(), ReadInt()
		S[i], T[i], X[i] = s, t, x

		L = append(L, &Item{key: s - x, time: s - x, query: 1, x: x})
		L = append(L, &Item{key: t - x, time: t - x, query: -1, x: x})
	}
	D = ReadIntSlice(q)

	sort.Stable(L)

	// sset = New()
	tr := NewTreap()
	curIdx := 0
	for _, item := range L {

		if D[curIdx] >= item.time {
			if item.query == 1 {
				// sset.AddOrUpdate(item.x, SCORE(item.x), 0)
				tr.Insert(item.x)
			} else {
				// sset.Remove(item.x)
				tr.Delete(item.x)
			}
			continue
		}

		for curIdx < q && D[curIdx] < item.time {
			// lowest := sset.GetByRank(1, false)
			lowest := tr.FindMinimum()
			if lowest != nil {
				// fmt.Println(lowest.score)
				fmt.Println(lowest.key)
			} else {
				fmt.Println(-1)
			}

			curIdx++
		}

		if curIdx == q {
			break
		} else {
			if item.query == 1 {
				// sset.AddOrUpdate(item.x, SCORE(item.x), 0)
				tr.Insert(item.x)
			} else {
				// sset.Remove(item.x)
				tr.Delete(item.x)
			}
		}
	}

	for i := curIdx; i < q; i++ {
		fmt.Println(-1)
	}
}

// Treap usage
// tr := NewTreap()
// tr.Insert(k)
// node := tr.Find(k)
// min := tr.FindMinimum()
// max := tr.FindMaximum()
// tr.Delete(k)
// node := tr.BigggerLowerBound(x)
// node := tr.BiggerUpperBound(x)
// node := tr.SmallerUpperBound(x)
// node := tr.SmallerLowerBound(x)
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

// BiggerLowerBound returns a node that has MINIMUM KEY MEETING key >= x.
// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
func (tr *Treap) BiggerLowerBound(x int) *Node {
	return tr.biggerLowerBound(tr.root, x)
}

// BiggerUpperBound returns a node that has MINIMUM KEY MEETING key > x.
// https://qiita.com/tubo28/items/f058582e457f6870a800#lower_bound-upper_bound
func (tr *Treap) BiggerUpperBound(x int) *Node {
	return tr.biggerUpperBound(tr.root, x)
}

// SmallerUpperBound returns a node that has MAXIMUM KEY MEETING key <= x.
// for AGC005-B
func (tr *Treap) SmallerUpperBound(x int) *Node {
	return tr.smallerUpperBound(tr.root, x)
}

// SmallerLowerBound returns a node that has MAXIMUM KEY MEETING key < x.
// for AGC005-B
func (tr *Treap) SmallerLowerBound(x int) *Node {
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

// MODはとったか？
// 遷移だけじゃなくて最後の最後でちゃんと取れよ？

/*******************************************************************/

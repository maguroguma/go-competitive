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

// OnBit returns the integer that has nth ON bit.
// If an argument has nth ON bit, OnBit returns the argument.
func OnBit(num, nth int) int {
	return num | (1 << uint(nth))
}

// OffBit returns the integer that has nth OFF bit.
// If an argument has nth OFF bit, OffBit returns the argument.
func OffBit(num, nth int) int {
	return num & ^(1 << uint(nth))
}

// PopCount returns the number of ON bit of an argument.
func PopCount(num int) int {
	res := 0

	for i := 0; i < 70; i++ {
		if ((num >> uint(i)) & 1) == 1 {
			res++
		}
	}

	return res
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

// Sum returns multiple integers sum.
func Sum(integers ...int) int {
	s := 0

	for _, i := range integers {
		s += i
	}

	return s
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
		// str := strconv.FormatInt(A[i], 10)  // 64bit int version
		res = append(res, []rune(str)...)

		if i != len(A)-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
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

/*******************************************************************/

const MOD = 1000000000 + 7
const ALPHABET_NUM = 26
const INF_INT64 = math.MaxInt64
const INF_BIT60 = 1 << 60

var n int
var A, B []int

type Pair struct {
	key  int
	a, b int
}
type PairList []*Pair
type byKey struct {
	PairList
}

func (l PairList) Len() int {
	return len(l)
}
func (l PairList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byKey) Less(i, j int) bool {
	// return l.PairList[i].key < l.PairList[j].key
	if l.PairList[i].b < l.PairList[j].b {
		return true
	} else if l.PairList[i].b > l.PairList[j].b {
		return false
	} else {
		return l.PairList[i].a < l.PairList[j].a
	}
}

// how to use
// L := make(PairList, 0, 200000+5)
// L = append(L, &Pair{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

func main() {
	n = ReadInt()
	A = ReadIntSlice(n)
	B = ReadIntSlice(n)

	L := make(PairList, 0)
	for i := 0; i < n; i++ {
		a, b := A[i], B[i]
		L = append(L, &Pair{a: a, b: b})
	}
	sort.Stable(byKey{L})
	// for i := 0; i < n; i++ {
	// 	fmt.Printf("a: %d, b: %d\n", L[i].a, L[i].b)
	// }
	AA := make([]int, n)
	BB := make([]int, n)
	for i := 0; i < n; i++ {
		AA[i] = L[i].a
		BB[i] = L[i].b
	}

	st := NewSegTree(AA)

	passed := 0
	exchanged := 0
	for i := 0; i < n; i++ {
		if AA[i] <= BB[i] {
			passed++
			continue
		}

		mini, minIdx := st.Query(i, n)
		minIdx = minIdx - (st.num - 1) + 1
		// fmt.Printf("l, r: %d, %d\n", i, minIdx)
		if mini > BB[i] {
			fmt.Println("No")
			return
		}

		exchanged++
		st.Exchange(i, minIdx)

		// 		mini, minIdx = st.Query(minIdx, n)
		// 		minIdx = minIdx - (st.num - 1) + 1
		// 		// fmt.Printf("l, r: %d, %d\n", i, minIdx)
		// 		if mini > BB[minIdx] {
		// 			fmt.Println("No")
		// 			return
		// 		}
	}

	// fmt.Println(exchanged)
	if exchanged <= n-2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

type SegTree struct {
	num    int
	minVal [2*200000 - 1]int
	index  [2*200000 - 1]int
}

func NewSegTree(A []int) *SegTree {
	st := new(SegTree)

	st.num = 1
	for st.num < len(A) {
		st.num *= 2
	}

	for i := 0; i < len(A); i++ {
		st.minVal[i+(st.num-1)] = A[i]
		st.index[i+(st.num-1)] = i + (st.num - 1)
	}

	for i := st.num - 2; i >= 0; i-- {
		// 中間ノードの初期化
		if st.minVal[2*i+1] < st.minVal[2*i+2] {
			st.minVal[i] = st.minVal[2*i+1]
			st.index[i] = st.index[2*i+1]
		} else {
			st.minVal[i] = st.minVal[2*i+2]
			st.index[i] = st.index[2*i+2]
		}
	}

	return st
}

func (st *SegTree) Exchange(l, r int) {
	ll, rr := l+(st.num-1), r+(st.num-1)
	vl, vr := st.minVal[ll], st.minVal[rr]

	st.minVal[ll] = vr
	for ll > 0 {
		ll = (ll - 1) / 2
		if st.minVal[2*ll+1] < st.minVal[2*ll+2] {
			st.minVal[ll] = st.minVal[2*ll+1]
			st.index[ll] = st.index[2*ll+1]
		} else {
			st.minVal[ll] = st.minVal[2*ll+2]
			st.index[ll] = st.index[2*ll+2]
		}
	}
	st.minVal[rr] = vl
	for rr > 0 {
		rr = (rr - 1) / 2
		if st.minVal[2*rr+1] < st.minVal[2*rr+2] {
			st.minVal[rr] = st.minVal[2*rr+1]
			st.index[rr] = st.index[2*rr+1]
		} else {
			st.minVal[rr] = st.minVal[2*rr+2]
			st.index[rr] = st.index[2*rr+2]
		}
	}
}

func (st *SegTree) Query(a, b int) (int, int) {
	return st.subQuery(a, b, 0, 0, -1)
}

func (st *SegTree) subQuery(a, b, k, l, r int) (int, int) {
	if r < 0 {
		r = st.num
	}

	if b <= l || r <= a {
		return INF_BIT60, -1
	}
	if a <= l && r <= b {
		return st.minVal[k], st.index[k]
	}

	vl, idxl := st.subQuery(a, b, 2*k+1, l, (l+r)/2)
	vr, idxr := st.subQuery(a, b, 2*k+2, (l+r)/2, r)

	// return Min(vl, vr)
	if vl < vr {
		// return st.index[2*k+1]
		return vl, idxl
	} else {
		// return st.index[2*k+2]
		return vr, idxr
	}
}

// // RMQを高速で処理するセグメントツリー構造体
// type SegTreeRMQ struct {
// 	ElemNum    int               // 配列の要素数
// 	Dat        [2*200005 - 1]int // セグメント木のノード集合
// 	INIT_VALUE int               // 初期化用の値
// }

// // 配列の要素数と初期化用の値を渡す
// func NewSegTreeRMQ(elemNum, initValue int) *SegTreeRMQ {
// 	st := new(SegTreeRMQ)
// 	st.INIT_VALUE = initValue

// 	// 要素数を2べきの数に補正する
// 	st.ElemNum = 1
// 	for st.ElemNum < elemNum {
// 		st.ElemNum *= 2
// 	}

// 	// すべての節点を初期化する
// 	for i := 0; i < 2*st.ElemNum-1; i++ {
// 		st.Dat[i] = st.INIT_VALUE
// 	}

// 	return st
// }

// // k番目の値（0-based）をaに変更する
// func (st *SegTreeRMQ) Update(k, a int) {
// 	// 葉の節点を変更
// 	k += st.ElemNum - 1
// 	st.Dat[k] = a

// 	// 登りながら各節点を変更
// 	for k > 0 {
// 		k = (k - 1) / 2
// 		st.Dat[k] = int(math.Min(float64(st.Dat[2*k+1]), float64(st.Dat[2*k+2])))
// 	}
// }

// // [a, b)の区間の最小値を返す
// // 再帰関数のラッパー関数
// func (st *SegTreeRMQ) Query(a, b int) int {
// 	return st.subQuery(a, b, 0, 0, st.ElemNum)
// }

// // [a, b)の区間の最小値を返す
// // k: 検索先のセグメントツリーのノード番号(0-based)
// // l, r: k番目のセグメントツリーノードの半開区間の左端と右端, [l, r)
// func (st *SegTreeRMQ) subQuery(a, b, k, l, r int) int {
// 	if r <= a || b <= l {
// 		return st.INIT_VALUE
// 	}

// 	if a <= l && r <= b {
// 		return st.Dat[k]
// 	}

// 	vl, vr := st.subQuery(a, b, 2*k+1, l, (l+r)/2), st.subQuery(a, b, 2*k+2, (l+r)/2, r)
// 	return int(math.Min(float64(vl), float64(vr)))
// }

// MODはとったか？
// 遷移だけじゃなくて最後の最後でちゃんと取れよ？

/*******************************************************************/

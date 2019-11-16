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

type Hero struct {
	key  int
	p, s int
}
type HeroList []*Hero
type byKey struct {
	HeroList
}

func (l HeroList) Len() int {
	return len(l)
}
func (l HeroList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byKey) Less(i, j int) bool {
	return l.HeroList[i].key < l.HeroList[j].key
}

// how to use
// L := make(HeroList, 0, 200000+5)
// L = append(L, &Hero{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

var t int
var n, m int
var A, P, S []int
var st *SegTreeRMQ

func main() {
	t = ReadInt()

	for tc := 0; tc < t; tc++ {
		n = ReadInt()
		A = ReadIntSlice(n)
		m = ReadInt()
		P, S = make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			p, s := ReadInt2()
			P[i], S[i] = p, s
		}

		solve()
	}
}

func solve() {
	L := make(HeroList, 0)
	for i := 0; i < m; i++ {
		p, s := P[i], S[i]
		L = append(L, &Hero{key: s, p: p, s: s})
	}
	sort.Stable(byKey{L})

	// enduranceでソートされたヒーロー列
	SS := make([]int, m)
	for i := 0; i < m; i++ {
		SS[i] = L[i].s
	}
	// enduranceのソート列に従ってpowerをセグメントツリーに放り込む
	st = NewSegTreeRMQ(m, 0)
	for i := 0; i < m; i++ {
		st.Update(i, L[i].p)
	}

	l := 0
	ans := 0
	for l < n {
		length := sub2(l, SS)
		if length == -1 {
			fmt.Println(-1)
			return
		} else {
			l += length
			ans++
		}
	}

	fmt.Println(ans)
}

func sub(l int, H []int) int {
	// m は中央を意味する何らかの値
	isOK := func(m int) bool {
		if H[m] >= l {
			return true
		}
		return false
	}

	ng, ok := -1, len(H)
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

func sub2(s int, H []int) int {
	maxMP := A[s]
	length := -1
	for l := 0; s+l < n; l++ {
		maxMP = Max(maxMP, A[s+l])
		idx := sub(l+1, H)
		if idx == len(H) {
			break
		}

		maxHP := st.Query(idx, len(H))
		if maxHP >= maxMP {
			length = l + 1
		} else {
			break
		}
	}

	return length
}

// RMQを高速で処理するセグメントツリー構造体
type SegTreeRMQ struct {
	ElemNum int // 配列の要素数
	// Dat     [2*200005 - 1]int // セグメント木のノード集合
	Dat        []int // セグメント木のノード集合
	INIT_VALUE int   // 初期化用の値
}

// 配列の要素数と初期化用の値を渡す
func NewSegTreeRMQ(elemNum, initValue int) *SegTreeRMQ {
	st := new(SegTreeRMQ)
	st.INIT_VALUE = initValue
	// st.Dat = make([]int, 4*elemNum+10)

	// 要素数を2べきの数に補正する
	st.ElemNum = 1
	for st.ElemNum < elemNum {
		st.ElemNum *= 2
	}
	st.Dat = make([]int, 2*st.ElemNum)

	// すべての節点を初期化する
	for i := 0; i < 2*st.ElemNum-1; i++ {
		st.Dat[i] = st.INIT_VALUE
	}

	return st
}

// k番目の値（0-based）をaに変更する
func (st *SegTreeRMQ) Update(k, a int) {
	// 葉の節点を変更
	k += st.ElemNum - 1
	st.Dat[k] = a

	// 登りながら各節点を変更
	for k > 0 {
		k = (k - 1) / 2
		st.Dat[k] = int(math.Max(float64(st.Dat[2*k+1]), float64(st.Dat[2*k+2])))
	}
}

// [a, b)の区間の最小値を返す
// 再帰関数のラッパー関数
func (st *SegTreeRMQ) Query(a, b int) int {
	return st.subQuery(a, b, 0, 0, st.ElemNum)
}

// [a, b)の区間の最小値を返す
// k: 検索先のセグメントツリーのノード番号(0-based)
// l, r: k番目のセグメントツリーノードの半開区間の左端と右端, [l, r)
func (st *SegTreeRMQ) subQuery(a, b, k, l, r int) int {
	if r <= a || b <= l {
		return st.INIT_VALUE
	}

	if a <= l && r <= b {
		return st.Dat[k]
	}

	vl, vr := st.subQuery(a, b, 2*k+1, l, (l+r)/2), st.subQuery(a, b, 2*k+2, (l+r)/2, r)
	return int(math.Max(float64(vl), float64(vr)))
}

// MODはとったか？
// 遷移だけじゃなくて最後の最後でちゃんと取れよ？

/*******************************************************************/

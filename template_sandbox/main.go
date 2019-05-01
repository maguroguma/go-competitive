package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
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

// GetZeroPaddingRuneSlice returns binary expressions of integer n with zero padding.
// For debugging use.
func GetZeroPaddingRuneSlice(n, digitsNum int) []rune {
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

// GetNthBit returns nth bit value of an argument.
// n starts from 0.
func GetNthBit(num, nth int) int {
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

// GetDigitSum returns digit sum of a decimal number.
// GetDigitSum only accept a positive integer.
func GetDigitSum(n int) int {
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

/*********** Utilities ***********/

// Strtoi is a wrapper of `strconv.Atoi()`.
// If `strconv.Atoi()` returns an error, Strtoi calls panic.
func Strtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(errors.New("[argument error]: Strtoi only accepts integer string"))
	} else {
		return i
	}
}

/*********** Permutation ***********/

// memo: 10! == 3628800 > 3M
func CalcFactorialPatterns(elements []rune) [][]rune {
	copiedResidual := make([]rune, len(elements))
	copy(copiedResidual, elements)
	return factorialRecursion([]rune{}, copiedResidual)
}
func factorialRecursion(interim, residual []rune) [][]rune {
	if len(residual) == 0 {
		return [][]rune{interim}
	}

	res := [][]rune{}
	for idx, elem := range residual {
		copiedInterim := make([]rune, len(interim))
		copy(copiedInterim, interim)
		copiedInterim = append(copiedInterim, elem)
		copiedResidual := genDeletedSlice(idx, residual)
		res = append(res, factorialRecursion(copiedInterim, copiedResidual)...)
	}

	return res
}
func genDeletedSlice(delId int, S []rune) []rune {
	res := []rune{}
	res = append(res, S[:delId]...)
	res = append(res, S[delId+1:]...)
	return res
}

// memo: 3**10 == 59049
func CalcDuplicatePatterns(elements []rune, digit int) [][]rune {
	return duplicateRecursion([]rune{}, elements, digit)
}
func duplicateRecursion(interim, elements []rune, digit int) [][]rune {
	if len(interim) == digit {
		return [][]rune{interim}
	}

	res := [][]rune{}
	for i := 0; i < len(elements); i++ {
		copiedInterim := make([]rune, len(interim))
		copy(copiedInterim, interim)
		copiedInterim = append(copiedInterim, elements[i])
		res = append(res, duplicateRecursion(copiedInterim, elements, digit)...)
	}

	return res
}

// usage
//tmp := CalcFactorialPatterns([]rune{'a', 'b', 'c'})
//expected := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
//tmp := CalcDuplicatePatterns([]rune{'a', 'b', 'c'}, 3)
//expected := []string{"aaa", "aab", "aac", "aba", "abb", "abc", ...}

/*********** Union Find ***********/

// UnionFind provides disjoint set algorithm.
// It accepts both 0-based and 1-based setting.
type UnionFind struct {
	parents []int
}

// NewUnionFind returns a pointer of a new instance of UnionFind.
func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.parents = make([]int, n+1)

	for i := 0; i <= n; i++ {
		uf.parents[i] = -1
	}

	return uf
}

// Root method returns root node of an argument node.
// Root method is a recursive function.
func (uf *UnionFind) Root(x int) int {
	if uf.parents[x] < 0 {
		return x
	}

	// route compression
	uf.parents[x] = uf.Root(uf.parents[x])
	return uf.parents[x]
}

// Unite method merges a set including x and a set including y.
func (uf *UnionFind) Unite(x, y int) bool {
	xp := uf.Root(x)
	yp := uf.Root(y)

	if xp == yp {
		return false
	}

	// merge: xp -> yp
	// merge larger set to smaller set
	if uf.CcSize(xp) > uf.CcSize(yp) {
		xp, yp = yp, xp
	}
	// update set size
	uf.parents[yp] += uf.parents[xp]
	// finally, merge
	uf.parents[xp] = yp

	return true
}

// Same method returns whether x is in the set including y or not.
func (uf *UnionFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

// CcSize method returns the size of a set including an argument node.
func (uf *UnionFind) CcSize(x int) int {
	return -uf.parents[uf.Root(x)]
}

/*********** Factorization, Prime Number ***********/

// TrialDivision returns the result of prime factorization of integer N.
// Complicity: O(n)
func TrialDivision(n int) map[int]int {
	if n <= 1 {
		panic(errors.New("[argument error]: TrialDivision only accepts a NATURAL number"))
	}

	p := map[int]int{}
	for i := 2; i*i <= n; i++ {
		exp := 0
		for n%i == 0 {
			exp++
			n /= i
		}

		if exp == 0 {
			continue
		}
		p[i] = exp
	}
	if n > 1 {
		p[n] = 1
	}

	return p
}

// IsPrime judges whether an argument integer is a prime number or not.
func IsPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

/*********** Inverse Element ***********/

// CalcNegativeMod can calculate a right residual whether value is positive or negative.
func CalcNegativeMod(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
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

// CalcModInv returns $a^{-1} mod m$ by Fermat's little theorem.
// O(1), but C is nearly equal to 30 (when m is 1000000000+7).
func CalcModInv(a, m int) int {
	return modpow(a, m-2, m)
}

/********** I/O usage **********/

//str := ReadString()
//i := ReadInt()
//X := ReadIntSlice(n)
//S := ReadRuneSlice()

/********** String Split **********/

//strs := strings.Split(string(runeSlice), "+")

/*******************************************************************/

const MOD = 1000000000 + 7
const ALPHABET_NUM = 26

func main() {
	fmt.Println(len(strings.Split("Hello Hello World", " ")))
	fmt.Println(math.Abs(1))
}

// MODはとったか？
// 遷移だけじゃなくて最後の最後でちゃんと取れよ？

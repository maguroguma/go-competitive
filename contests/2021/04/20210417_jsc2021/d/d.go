/*
URL:
https://atcoder.jp/contests/jsc2021/tasks/jsc2021_d
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
	println          = fmt.Println
	yes, no, invalid = "Yes", "No", -1

	n, p int
)

func main() {
	defer stdout.Flush()

	n, p = readi2()

	SetMod(Mod1000000007)

	a := Mint(p - 1)
	b := Mint(p - 2)
	c := b.Pow(Mint(n - 1))
	ans := a.Mul(c)

	println(ans)
}

// Originated from @ccppjsrb
// e.g.: https://atcoder.jp/contests/arc106/submissions/17669427

// Mod constants.
const (
	Mod1000000007 = 1000000007
	Mod998244353  = 998244353
)

var _mod Mint
var _fmod func(Mint) Mint

// Mint treats the modular arithmetic
type Mint int64

// SetMod sets the mod. It must be called first.
func SetMod(newmod Mint) {
	switch newmod {
	case Mod1000000007:
		_fmod = staticMod1000000007
	case Mod998244353:
		_fmod = staticMod998244353
	default:
		_mod = newmod
		_fmod = dynamicMod
	}
}
func dynamicMod(m Mint) Mint {
	m %= _mod
	if m < 0 {
		return m + _mod
	}
	return m
}
func staticMod1000000007(m Mint) Mint {
	m %= Mod1000000007
	if m < 0 {
		return m + Mod1000000007
	}
	return m
}
func staticMod998244353(m Mint) Mint {
	m %= Mod998244353
	if m < 0 {
		return m + Mod998244353
	}
	return m
}

// Mod returns m % mod.
func (m Mint) Mod() Mint {
	return _fmod(m)
}

// Inv returns modular multiplicative inverse
func (m Mint) Inv() Mint {
	return m.Pow(Mint(0).Sub(2))
}

// Pow returns m^n
func (m Mint) Pow(n Mint) Mint {
	p := Mint(1)
	for n > 0 {
		if n&1 == 1 {
			p.MulAs(m)
		}
		m.MulAs(m)
		n >>= 1
	}
	return p
}

// Add returns m+x
func (m Mint) Add(x Mint) Mint {
	return (m + x).Mod()
}

// Sub returns m-x
func (m Mint) Sub(x Mint) Mint {
	return (m - x).Mod()
}

// Mul returns m*x
func (m Mint) Mul(x Mint) Mint {
	return (m * x).Mod()
}

// Div returns m/x
func (m Mint) Div(x Mint) Mint {
	return m.Mul(x.Inv())
}

// AddAs assigns *m + x to *m and returns m
func (m *Mint) AddAs(x Mint) *Mint {
	*m = m.Add(x)
	return m
}

// SubAs assigns *m - x to *m and returns m
func (m *Mint) SubAs(x Mint) *Mint {
	*m = m.Sub(x)
	return m
}

// MulAs assigns *m * x to *m and returns m
func (m *Mint) MulAs(x Mint) *Mint {
	*m = m.Mul(x)
	return m
}

// DivAs assigns *m / x to *m and returns m
func (m *Mint) DivAs(x Mint) *Mint {
	*m = m.Div(x)
	return m
}

// cf := NewCombFactorial(2000000)
// maxNum == "maximum n" * 2 (for H(n,r))
// res := cf.C(n, r) 	// 組み合わせ
// res := cf.H(n, r) 	// 重複組合せ
// res := cf.P(n, r) 	// 順列

type CombFactorial struct {
	maxNum Mint
	fact   func(x Mint) Mint
	invf   func(x Mint) Mint
}

func NewCombFactorial(maxNum Mint) *CombFactorial {
	cf := new(CombFactorial)
	cf.maxNum = maxNum
	cf.initCF()

	return cf
}

func (c *CombFactorial) initCF() {
	var i Mint

	factTable := make([]Mint, c.maxNum+50)
	invfTable := make([]Mint, c.maxNum+50)

	factTable[0] = 1
	invfTable[0] = factTable[0].Inv()
	for i = 1; i <= c.maxNum; i++ {
		val := factTable[i-1].Mul(Mint(i))
		factTable[i] = val
		invfTable[i] = factTable[i].Inv()
	}

	c.fact = func(x Mint) Mint { return factTable[x] }
	c.invf = func(x Mint) Mint { return invfTable[x] }
}
func (c *CombFactorial) C(n, r Mint) Mint {
	var res Mint

	res = Mint(1).
		Mul(c.fact(n)).
		Mul(c.invf(r)).
		Mul(c.invf(n - r))

	return res
}
func (c *CombFactorial) P(n, r Mint) Mint {
	var res Mint

	res = 1
	res.MulAs(c.fact(n)).MulAs(c.invf(n - r))

	return res
}
func (c *CombFactorial) H(n, r Mint) Mint {
	return c.C(n-1+r, r)
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

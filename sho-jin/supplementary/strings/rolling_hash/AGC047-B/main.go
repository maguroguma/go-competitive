/*
URL:
https://atcoder.jp/contests/agc047/tasks/agc047_b
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"strconv"
	"time"
)

var (
	n int
	S []string

	gm map[V]int
)

type V struct {
	hv uint64
	c  byte
}

func main() {
	defer stdout.Flush()

	n = readi()
	for i := 0; i < n; i++ {
		str := readrs()
		ReverseMyself(str)
		S = append(S, string(str))
	}

	sort.Slice(S, func(i, j int) bool {
		return len(S[i]) < len(S[j])
	})

	gm = make(map[V]int)

	ans := 0
	for i := 0; i < n; i++ {
		memo := [ALPHABET_NUM]int{}
		for _, r := range S[i] {
			memo[r-'a']++
		}

		rh := NewRHash(S[i])
		for j := 0; j < len(S[i]); j++ {
			h := rh.SliceHash(0, j)
			for k := 0; k < ALPHABET_NUM; k++ {
				if memo[k] > 0 {
					v := V{h, byte(k) + 'a'}
					ans += gm[v]
				}
			}

			memo[S[i][j]-'a']--
		}

		v := V{rh.SliceHash(0, rh.Len()-1), S[i][len(S[i])-1]}
		gm[v]++
	}

	printf("%d\n", ans)
}

// rolling hash (by keymoon@atcoder)
// originated from: https://qiita.com/keymoon/items/11fac5627672a6d6a9f6
// reference: https://atcoder.jp/contests/abc141/submissions/7717102

// NewRHash returns rolling hashs of the string.
func NewRHash(s string) *RHash {
	if !_isInitialized {
		initRHashConfing()
		_isInitialized = true
	}

	rh := new(RHash)

	rh.hash = make([]uint64, len(s)+1)
	for i := 0; i < len(s); i++ {
		rh.hash[i+1] = rhCalcMod(rhMul(rh.hash[i], _rhBase) + uint64(s[i]))
	}

	return rh
}

// SliceHash returns a rolling hash of a slice of the string.
// The slice is expressed like [l, r).
// This function can be used like Golang slice(S[l:r]).
func (rh *RHash) SliceHash(l, r int) uint64 {
	begin, length := l, r-l
	return rhCalcMod(
		rh.hash[begin+length] + _RH_POSITIVIZER - rhMul(rh.hash[begin], _rhPowMemo[length]),
	)
}

// OffsetHash returns a rolling hash of a slice of the string.
// The slice is expressed like [begin, begin+length).
func (rh *RHash) OffsetHash(begin, length int) uint64 {
	return rhCalcMod(
		rh.hash[begin+length] + _RH_POSITIVIZER - rhMul(rh.hash[begin], _rhPowMemo[length]),
	)
}

// Len returns a length of an original string.
func (rh *RHash) Len() int {
	return len(rh.hash) - 1
}

type RHash struct {
	hash []uint64
}

const (
	_RH_MASK30       uint64 = (1 << 30) - 1
	_RH_MASK31       uint64 = (1 << 31) - 1
	_RH_MOD          uint64 = (1 << 61) - 1
	_RH_POSITIVIZER  uint64 = _RH_MOD * ((1 << 3) - 1)
	_RH_MAX_S_LENGTH        = 2000000 + 50
)

var (
	_rhBase        uint64
	_rhPowMemo     []uint64
	_isInitialized = false
)

func initRHashConfing() {
	rand.Seed(time.Now().UnixNano())

	_rhBase = uint64(rand.Int31n(math.MaxInt32-129)) + uint64(129)
	_rhPowMemo = make([]uint64, _RH_MAX_S_LENGTH)
	_rhPowMemo[0] = 1
	for i := 1; i < len(_rhPowMemo); i++ {
		_rhPowMemo[i] = rhCalcMod(rhMul(_rhPowMemo[i-1], _rhBase))
	}
}

func rhMul(l, r uint64) uint64 {
	var lu uint64 = l >> 31
	var ld uint64 = l & _RH_MASK31
	var ru uint64 = r >> 31
	var rd uint64 = r & _RH_MASK31
	var middleBit uint64 = ld*ru + lu*rd

	return ((lu * ru) << 1) + ld*rd + ((middleBit & _RH_MASK30) << 31) + (middleBit >> 30)
}

func rhCalcMod(val uint64) uint64 {
	val = (val & _RH_MOD) + (val >> 61)
	if val > _RH_MOD {
		val -= _RH_MOD
	}
	return val
}

func ReverseMyself(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

/*******************************************************************/

/********** common constants **********/

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

// modi can calculate a right residual whether value is positive or negative.
func modi(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// modll can calculate a right residual whether value is positive or negative.
func modll(val, m int64) int64 {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	reads = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
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

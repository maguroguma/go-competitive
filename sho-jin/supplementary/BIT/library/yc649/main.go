/*
URL:
https://yukicoder.me/problems/no/649
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

var (
	q, k    int
	queries [][]int
)

func main() {
	q, k = ReadInt2()

	for i := 0; i < q; i++ {
		c := ReadInt()
		if c == 1 {
			v := ReadInt()
			queries = append(queries, []int{c, v})
		} else {
			queries = append(queries, []int{c})
		}
	}

	elems := []int{}
	for i := 0; i < q; i++ {
		if len(queries[i]) == 1 {
			continue
		}

		elems = append(elems, queries[i][1])
	}

	_, otop, ptoo := ZaAtsu1Dim(elems, 1)

	ft := NewFenwickTreeSet(200000 + 5)
	for i := 0; i < q; i++ {
		que := queries[i]
		if que[0] == 1 {
			v := que[1]
			vv := otop[v]
			ft.Insert(vv, 1)
		} else {
			if ft.Count(ft.n) < k {
				fmt.Println(-1)
			} else {
				v := ft.Kth(k)
				vv := ptoo[v]
				fmt.Println(vv)
				ft.Delete(v, 1)
			}
		}
	}
}

// Public methods
// ft := NewFenwickTreeSet(200000 + 5)
// c := ft.Count(b.n)
// ft.Insert(val, 1)
// ft.Delete(val, 1)
// ans := ft.Kth(k)

type FenwickTreeSet struct {
	dat     []int
	n       int
	minPow2 int
}

// n(>=1) is maximum integer for the set.
func NewFenwickTreeSet(n int) *FenwickTreeSet {
	newBit := new(FenwickTreeSet)

	newBit.dat = make([]int, n+1)
	newBit.n = n

	newBit.minPow2 = 1
	for {
		if (newBit.minPow2 << 1) > n {
			break
		}
		newBit.minPow2 <<= 1
	}

	return newBit
}

// Count returns number of elements less or equal than e in the set.
// b.Count(b.n) returns total number of elements in the set.
// O(logN)
func (ft *FenwickTreeSet) Count(e int) int {
	s := 0

	for e > 0 {
		s += ft.dat[e]
		e -= e & (-e)
	}

	return s
}

// Insert e(1<=e<=n) for num(>= 1) times.
func (ft *FenwickTreeSet) Insert(e, num int) {
	for e <= ft.n {
		ft.dat[e] += num
		e += e & (-e)
	}
}

// Delete e(1<=e<=n) for num(>= 1) times.
func (ft *FenwickTreeSet) Delete(e, num int) {
	num *= -1
	for e <= ft.n {
		ft.dat[e] += num
		e += e & (-e)
	}
}

// Kth returns kth(>=0) element in the set
func (ft *FenwickTreeSet) Kth(kth int) int {
	if kth <= 0 {
		return 0
	}

	x := 0
	for k := ft.minPow2; k > 0; k /= 2 {
		if x+k <= ft.n && ft.dat[x+k] < kth {
			kth -= ft.dat[x+k]
			x += k
		}
	}

	return x + 1
}

// ZaAtsu1Dim returns 3 values.
// pressed: pressed slice of the original slice
// orgToPress: map for translating original value to pressed value
// pressToOrg: reverse resolution of orgToPress
// O(nlogn)
func ZaAtsu1Dim(org []int, initVal int) (pressed []int, orgToPress, pressToOrg map[int]int) {
	pressed = make([]int, len(org))
	copy(pressed, org)
	sort.Sort(sort.IntSlice(pressed))

	orgToPress = make(map[int]int)
	for i := 0; i < len(org); i++ {
		if i == 0 {
			orgToPress[pressed[0]] = initVal
			continue
		}

		if pressed[i-1] != pressed[i] {
			initVal++
			orgToPress[pressed[i]] = initVal
		}
	}

	for i := 0; i < len(org); i++ {
		pressed[i] = orgToPress[org[i]]
	}

	pressToOrg = make(map[int]int)
	for k, v := range orgToPress {
		pressToOrg[v] = k
	}

	return
}

/*******************************************************************/

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
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

/*
URL:
https://yukicoder.me/problems/no/430
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var (
	S []rune
	m int
	C []string
)

func main() {
	defer stdout.Flush()

	S = readrs()
	m = readi()
	for i := 0; i < m; i++ {
		str := reads()
		C = append(C, str)
	}

	trie := NewTrie('A')
	for i := 0; i < m; i++ {
		trie.Insert(C[i])
	}

	memo := map[string]int{}
	for i := 0; i < len(S); i++ {
		for j := 1; j <= 10; j++ {
			if i+j > len(S) {
				break
			}

			sub := string(S[i : i+j])
			if trie.Find(sub) {
				memo[sub]++
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		ans += memo[C[i]]
	}
	printf("%d\n", ans)
}

// Nodes are managed by slice(vector).

const (
	// assume words consisting of only lower case or upper case
	_TRIE_CHAR_SIZE = 26
)

// NewTrie returns a trie managing words starting from base character.
// e.g.: NewTrie('a')
func NewTrie(base rune) *Trie {
	t := new(Trie)

	t.root = 0
	t.nodes = append(t.nodes, newTrieNode(t.root)) // 始めは根だけ
	t.base = base

	return t
}

// Insert a word.
func (t *Trie) Insert(word string) {
	t._insert(word, t.nodes[0].common)
}

// Find returns whether the trie has the word or not.
func (t *Trie) Find(word string) bool {
	return t._search(word, false)
}

// FindStartWith returns whether the trie has the word having the prefix or not.
func (t *Trie) FindStartWith(prefix string) bool {
	return t._search(prefix, true)
}

// CountWord returns the number of the words that the trie has.
// CountWord cannot count the number of UNIQUE words.
func (t *Trie) CountWord() int {
	return t.nodes[0].common
}

// SizeTrie returns the number of the nodes that the trie has.
func (t *Trie) SizeTrie() int {
	return len(t.nodes)
}

type Trie struct {
	nodes []*trieNode // nodes managed by the trie
	root  int         // root node id
	base  rune        // base character
}

type trieNode struct {
	next   [_TRIE_CHAR_SIZE]int // child id that a character ith has (NIL == -1)
	accept []int                // accept has string id(s) whose last character is equal to this node
	c      int                  // offset of this node's character from the base character
	common int                  // number of strings that share this node
}

func newTrieNode(c int) *trieNode {
	tn := new(trieNode)

	tn.c = c
	tn.common = 0
	for i := 0; i < _TRIE_CHAR_SIZE; i++ {
		tn.next[i] = -1
	}

	return tn
}

func (t *Trie) _insert(word string, wordID int) {
	nodeID := t.root

	for _, r := range word {
		c := int(r - t.base)

		nextID := &t.nodes[nodeID].next[c]
		if *nextID == -1 {
			// add nodes when there is not the next node
			*nextID = len(t.nodes)
			t.nodes = append(t.nodes, newTrieNode(c))
		}
		t.nodes[nodeID].common++
		nodeID = *nextID
	}
	t.nodes[nodeID].common++
	t.nodes[nodeID].accept = append(t.nodes[nodeID].accept, wordID)
}

func (t *Trie) _search(word string, isPrefix bool) bool {
	nodeID := t.root

	for _, r := range word {
		c := int(r - t.base)

		nextID := t.nodes[nodeID].next[c]
		if nextID == -1 {
			return false
		}
		nodeID = nextID
	}

	if isPrefix {
		return true
	}

	// check whether the word is accepted or not
	return len(t.nodes[nodeID].accept) > 0
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

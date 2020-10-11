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
	"os"
	"reflect"
	"sort"
	"strconv"
)

var (
	n int
	S []string
)

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

	trie := NewTrie('a')

	ans := 0
	for i := 0; i < n; i++ {
		memo := [ALPHABET_NUM]int{}
		for _, r := range S[i] {
			memo[r-'a']++
		}

		op := func(cid int, c int) {
			for i := 0; i < _TRIE_CHAR_SIZE; i++ {
				nid := trie.nodes[cid].next[i]
				if nid != -1 && trie.IsAccept(nid) && memo[i] > 0 {
					ans++
				}
			}

			memo[c]--
		}
		trie.Traverse(S[i][:len(S[i])-1], op)

		trie.Insert(S[i])
	}

	printf("%d\n", ans)
}

func ReverseMyself(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// Nodes are managed by slice(vector).

const (
	// assume words consisting of only lower case or upper case
	_TRIE_CHAR_SIZE = 26
)

// Operate do something in a trie node while traversing automaton.
type Operate func(curNodeID int, c int)

// NewTrie returns a trie managing words starting from base character.
// e.g.: NewTrie('a')
func NewTrie(base rune) *Trie {
	t := new(Trie)

	t.root = 0
	t.nodes = append(t.nodes, newTrieNode(t.root))
	t.base = base
	t.noop = func(curNodeID int, c int) {}

	return t
}

// Insert a word.
func (t *Trie) Insert(word string) {
	t._insert(word, t.nodes[0].common)
}

// Find returns whether the trie has the word or not.
func (t *Trie) Find(word string) bool {
	return t._search(word, false, t.noop)
}

// FindStartWith returns whether the trie has the word having the prefix or not.
func (t *Trie) FindStartWith(prefix string) bool {
	return t._search(prefix, true, t.noop)
}

// Traverse walk on the tree while operating something.
// Operation function receives current node id and current node offset from a base.
func (t *Trie) Traverse(word string, op Operate) {
	t._search(word, false, op)
}

// CountWord returns the number of the words that the trie has.
// CountWord can NOT count the number of UNIQUE words.
func (t *Trie) CountWord() int {
	return t.nodes[0].common
}

// SizeTrie returns the number of the nodes that the trie has.
func (t *Trie) SizeTrie() int {
	return len(t.nodes)
}

// IsAccept returns whether a trie node says accept or not.
func (t *Trie) IsAccept(curNodeID int) bool {
	return len(t.nodes[curNodeID].accept) > 0
}

type Trie struct {
	nodes []*trieNode // nodes managed by the trie
	root  int         // root node id
	base  rune        // base character
	noop  Operate     // do nothing
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
	curNodeID := t.root

	for _, r := range word {
		c := int(r - t.base)

		nextID := &t.nodes[curNodeID].next[c]
		if *nextID == -1 {
			// add nodes when there is not the next node
			*nextID = len(t.nodes)
			t.nodes = append(t.nodes, newTrieNode(c))
		}
		t.nodes[curNodeID].common++
		curNodeID = *nextID
	}
	t.nodes[curNodeID].common++
	t.nodes[curNodeID].accept = append(t.nodes[curNodeID].accept, wordID)
}

func (t *Trie) _search(word string, isPrefix bool, op Operate) bool {
	curNodeID := t.root

	for _, r := range word {
		c := int(r - t.base)

		// operate something (do nothing if op is noop)
		op(curNodeID, c)

		nextNodeID := t.nodes[curNodeID].next[c]
		if nextNodeID == -1 {
			return false
		}
		curNodeID = nextNodeID
	}

	if isPrefix {
		return true
	}

	// check whether the word is accepted or not
	return t.IsAccept(curNodeID)
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

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
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

var n int
var A []int

var sset *SortedSet

type Item struct {
	key int
	idx int
}
type ItemList []*Item

func (l ItemList) Len() int {
	return len(l)
}
func (l ItemList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l ItemList) Less(i, j int) bool {
	return l[i].key < l[j].key
}

// how to use
// L := make(ItemList, 0, 100000+1)
// L = append(L, &Item{key: intValue})
// sort.Sort(L)               // ASC
// sort.Sort(sort.Reverse(L)) // DESC

func main() {
	n = ReadInt()
	A = ReadIntSlice(n)
	sset = New()

	L := make(ItemList, n)
	for i := 0; i < n; i++ {
		L[i] = &Item{key: A[i], idx: i}
	}
	sort.Stable(L)

	ans := 0
	for i := 0; i < n; i++ {
		item := L[i]
		sset.AddOrUpdate(item.idx, SCORE(item.idx), 0)
		// fmt.Printf("set size: %d\n", sset.GetCount())

		// a のランクxと、ランクx-1, x+1のキー（スコア）が知りたい
		arank := sset.FindRank(item.idx)

		left := 0
		if arank > 1 {
			lnode := sset.GetByRank(arank-1, false)
			left = lnode.key + 1
		}
		right := n - 1
		if arank < sset.GetCount() {
			rnode := sset.GetByRank(arank+1, false)
			right = rnode.key - 1
		}

		// fmt.Printf("left: %d, mid: %d, right: %d\n", left, item.idx, right)
		ans += item.key * (item.idx - left + 1) * (right - item.idx + 1)
	}

	fmt.Println(ans)
}

// MODはとったか？
// 遷移だけじゃなくて最後の最後でちゃんと取れよ？

/*******************************************************************/

// Copyright (c) 2016, Jerry.Wang
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//  list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//  this list of conditions and the following disclaimer in the documentation
//  and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// package sortedset

type SortedSetLevel struct {
	forward *SortedSetNode
	span    int64
}

// Node in skip list
type SortedSetNode struct {
	key      int         // unique key of this node
	Value    interface{} // associated data
	score    SCORE       // score to determine the order of this node in the set
	backward *SortedSetNode
	level    []SortedSetLevel
}

// Get the key of the node
func (this *SortedSetNode) Key() int {
	return this.key
}

// Get the node of the node
func (this *SortedSetNode) Score() SCORE {
	return this.score
}

// Copyright (c) 2016, Jerry.Wang
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//  list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//  this list of conditions and the following disclaimer in the documentation
//  and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// package sortedset

// import (
// 	"math/rand"
// )

type SCORE int64 // the type of score

const SKIPLIST_MAXLEVEL = 32 /* Should be enough for 2^32 elements */
const SKIPLIST_P = 0.25      /* Skiplist P = 1/4 */

type SortedSet struct {
	header *SortedSetNode
	tail   *SortedSetNode
	length int64
	level  int
	dict   map[int]*SortedSetNode
}

func createNode(level int, score SCORE, key int, value interface{}) *SortedSetNode {
	node := SortedSetNode{
		score: score,
		key:   key,
		Value: value,
		level: make([]SortedSetLevel, level),
	}
	return &node
}

// Returns a random level for the new skiplist node we are going to create.
// The return value of this function is between 1 and SKIPLIST_MAXLEVEL
// (both inclusive), with a powerlaw-alike distribution where higher
// levels are less likely to be returned.
func randomLevel() int {
	level := 1
	for float64(rand.Int31()&0xFFFF) < float64(SKIPLIST_P*0xFFFF) {
		level += 1
	}
	if level < SKIPLIST_MAXLEVEL {
		return level
	}

	return SKIPLIST_MAXLEVEL
}

func (this *SortedSet) insertNode(score SCORE, key int, value interface{}) *SortedSetNode {
	var update [SKIPLIST_MAXLEVEL]*SortedSetNode
	var rank [SKIPLIST_MAXLEVEL]int64

	x := this.header
	for i := this.level - 1; i >= 0; i-- {
		/* store rank that is crossed to reach the insert position */
		if this.level-1 == i {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && // score is the same but the key is different
					x.level[i].forward.key < key)) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	/* we assume the key is not already inside, since we allow duplicated
	 * scores, and the re-insertion of score and redis object should never
	 * happen since the caller of Insert() should test in the hash table
	 * if the element is already inside or not. */
	level := randomLevel()

	if level > this.level { // add a new level
		for i := this.level; i < level; i++ {
			rank[i] = 0
			update[i] = this.header
			update[i].level[i].span = this.length
		}
		this.level = level
	}

	x = createNode(level, score, key, value)
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		/* update span covered by update[i] as x is inserted here */
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	/* increment span for untouched levels */
	for i := level; i < this.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == this.header {
		x.backward = nil
	} else {
		x.backward = update[0]
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		this.tail = x
	}
	this.length++
	return x
}

/* Internal function used by delete, DeleteByScore and DeleteByRank */
func (this *SortedSet) deleteNode(x *SortedSetNode, update [SKIPLIST_MAXLEVEL]*SortedSetNode) {
	for i := 0; i < this.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span -= 1
		}
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		this.tail = x.backward
	}
	for this.level > 1 && this.header.level[this.level-1].forward == nil {
		this.level--
	}
	this.length--
	delete(this.dict, x.key)
}

/* Delete an element with matching score/key from the skiplist. */
func (this *SortedSet) delete(score SCORE, key int) bool {
	var update [SKIPLIST_MAXLEVEL]*SortedSetNode

	x := this.header
	for i := this.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score &&
					x.level[i].forward.key < key)) {
			x = x.level[i].forward
		}
		update[i] = x
	}
	/* We may have multiple elements with the same score, what we need
	 * is to find the element with both the right score and object. */
	x = x.level[0].forward
	if x != nil && score == x.score && x.key == key {
		this.deleteNode(x, update)
		// free x
		return true
	}
	return false /* not found */
}

// Create a new SortedSet
func New() *SortedSet {
	sortedSet := SortedSet{
		level: 1,
		dict:  make(map[int]*SortedSetNode),
	}
	sortedSet.header = createNode(SKIPLIST_MAXLEVEL, 0, -1000000000000, nil)
	return &sortedSet
}

// Get the number of elements
func (this *SortedSet) GetCount() int {
	return int(this.length)
}

// get the element with minimum score, nil if the set is empty
//
// Time complexity of this method is : O(log(N))
func (this *SortedSet) PeekMin() *SortedSetNode {
	return this.header.level[0].forward
}

// get and remove the element with minimal score, nil if the set is empty
//
// // Time complexity of this method is : O(log(N))
func (this *SortedSet) PopMin() *SortedSetNode {
	x := this.header.level[0].forward
	if x != nil {
		this.Remove(x.key)
	}
	return x
}

// get the element with maximum score, nil if the set is empty
// Time Complexity : O(1)
func (this *SortedSet) PeekMax() *SortedSetNode {
	return this.tail
}

// get and remove the element with maximum score, nil if the set is empty
//
// Time complexity of this method is : O(log(N))
func (this *SortedSet) PopMax() *SortedSetNode {
	x := this.tail
	if x != nil {
		this.Remove(x.key)
	}
	return x
}

// Add an element into the sorted set with specific key / value / score.
// if the element is added, this method returns true; otherwise false means updated
//
// Time complexity of this method is : O(log(N))
func (this *SortedSet) AddOrUpdate(key int, score SCORE, value interface{}) bool {
	var newNode *SortedSetNode = nil

	found := this.dict[key]
	if found != nil {
		// score does not change, only update value
		if found.score == score {
			found.Value = value
		} else { // score changes, delete and re-insert
			this.delete(found.score, found.key)
			newNode = this.insertNode(score, key, value)
		}
	} else {
		newNode = this.insertNode(score, key, value)
	}

	if newNode != nil {
		this.dict[key] = newNode
	}
	return found == nil
}

// Delete element specified by key
//
// Time complexity of this method is : O(log(N))
func (this *SortedSet) Remove(key int) *SortedSetNode {
	found := this.dict[key]
	if found != nil {
		this.delete(found.score, found.key)
		return found
	}
	return nil
}

type GetByScoreRangeOptions struct {
	Limit        int  // limit the max nodes to return
	ExcludeStart bool // exclude start value, so it search in interval (start, end] or (start, end)
	ExcludeEnd   bool // exclude end value, so it search in interval [start, end) or (start, end)
}

// Get the nodes whose score within the specific range
//
// If options is nil, it searchs in interval [start, end] without any limit by default
//
// Time complexity of this method is : O(log(N))
func (this *SortedSet) GetByScoreRange(start SCORE, end SCORE, options *GetByScoreRangeOptions) []*SortedSetNode {

	// prepare parameters
	var limit int = 2147483648
	if options != nil && options.Limit > 0 {
		limit = options.Limit
	}

	excludeStart := options != nil && options.ExcludeStart
	excludeEnd := options != nil && options.ExcludeEnd
	reverse := start > end
	if reverse {
		start, end = end, start
		excludeStart, excludeEnd = excludeEnd, excludeStart
	}

	//////////////////////////
	var nodes []*SortedSetNode

	//determine if out of range
	if this.length == 0 {
		return nodes
	}
	//////////////////////////

	if reverse { // search from end to start
		x := this.header

		if excludeEnd {
			for i := this.level - 1; i >= 0; i-- {
				for x.level[i].forward != nil &&
					x.level[i].forward.score < end {
					x = x.level[i].forward
				}
			}
		} else {
			for i := this.level - 1; i >= 0; i-- {
				for x.level[i].forward != nil &&
					x.level[i].forward.score <= end {
					x = x.level[i].forward
				}
			}
		}

		for x != nil && limit > 0 {
			if excludeStart {
				if x.score <= start {
					break
				}
			} else {
				if x.score < start {
					break
				}
			}

			next := x.backward

			nodes = append(nodes, x)
			limit--

			x = next
		}
	} else {
		// search from start to end
		x := this.header
		if excludeStart {
			for i := this.level - 1; i >= 0; i-- {
				for x.level[i].forward != nil &&
					x.level[i].forward.score <= start {
					x = x.level[i].forward
				}
			}
		} else {
			for i := this.level - 1; i >= 0; i-- {
				for x.level[i].forward != nil &&
					x.level[i].forward.score < start {
					x = x.level[i].forward
				}
			}
		}

		/* Current node is the last with score < or <= start. */
		x = x.level[0].forward

		for x != nil && limit > 0 {
			if excludeEnd {
				if x.score >= end {
					break
				}
			} else {
				if x.score > end {
					break
				}
			}

			next := x.level[0].forward

			nodes = append(nodes, x)
			limit--

			x = next
		}
	}

	return nodes
}

// Get nodes within specific rank range [start, end]
// Note that the rank is 1-based integer. Rank 1 means the first node; Rank -1 means the last node;
//
// If start is greater than end, the returned array is in reserved order
// If remove is true, the returned nodes are removed
//
// Time complexity of this method is : O(log(N))
func (this *SortedSet) GetByRankRange(start int, end int, remove bool) []*SortedSetNode {

	/* Sanitize indexes. */
	if start < 0 {
		start = int(this.length) + start + 1
	}
	if end < 0 {
		end = int(this.length) + end + 1
	}
	if start <= 0 {
		start = 1
	}
	if end <= 0 {
		end = 1
	}

	reverse := start > end
	if reverse { // swap start and end
		start, end = end, start
	}

	var update [SKIPLIST_MAXLEVEL]*SortedSetNode
	var nodes []*SortedSetNode
	var traversed int = 0

	x := this.header
	for i := this.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			traversed+int(x.level[i].span) < start {
			traversed += int(x.level[i].span)
			x = x.level[i].forward
		}
		if remove {
			update[i] = x
		} else {
			if traversed+1 == start {
				break
			}
		}
	}

	traversed++
	x = x.level[0].forward
	for x != nil && traversed <= end {
		next := x.level[0].forward

		nodes = append(nodes, x)

		if remove {
			this.deleteNode(x, update)
		}

		traversed++
		x = next
	}

	if reverse {
		for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
			nodes[i], nodes[j] = nodes[j], nodes[i]
		}
	}
	return nodes
}

// Get node by rank.
// Note that the rank is 1-based integer. Rank 1 means the first node; Rank -1 means the last node;
//
// If remove is true, the returned nodes are removed
// If node is not found at specific rank, nil is returned
//
// Time complexity of this method is : O(log(N))
func (this *SortedSet) GetByRank(rank int, remove bool) *SortedSetNode {
	nodes := this.GetByRankRange(rank, rank, remove)
	if len(nodes) == 1 {
		return nodes[0]
	}
	return nil
}

// Get node by key
//
// If node is not found, nil is returned
// Time complexity : O(1)
func (this *SortedSet) GetByKey(key int) *SortedSetNode {
	return this.dict[key]
}

// Find the rank of the node specified by key
// Note that the rank is 1-based integer. Rank 1 means the first node
//
// If the node is not found, 0 is returned. Otherwise rank(> 0) is returned
//
// Time complexity of this method is : O(log(N))
func (this *SortedSet) FindRank(key int) int {
	var rank int = 0
	node := this.dict[key]
	if node != nil {
		x := this.header
		for i := this.level - 1; i >= 0; i-- {
			for x.level[i].forward != nil &&
				(x.level[i].forward.score < node.score ||
					(x.level[i].forward.score == node.score &&
						x.level[i].forward.key <= node.key)) {
				rank += int(x.level[i].span)
				x = x.level[i].forward
			}

			if x.key == key {
				return rank
			}
		}
	}
	return 0
}

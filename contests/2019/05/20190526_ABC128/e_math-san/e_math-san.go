package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Io struct {
	reader    *bufio.Reader
	writer    *bufio.Writer
	tokens    []string
	nextToken int
}

func NewIo() *Io {
	return &Io{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (io *Io) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *Io) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

func (io *Io) Next() string {
	for io.nextToken >= len(io.tokens) {
		line := io.NextLine()
		io.tokens = strings.Fields(line)
		io.nextToken = 0
	}
	r := io.tokens[io.nextToken]
	io.nextToken++
	return r
}

func (io *Io) NextInt() int {
	i, err := strconv.Atoi(io.Next())
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) NextInt64() int64 {
	i, err := strconv.ParseInt(io.Next(), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) NextFloat() float64 {
	i, err := strconv.ParseFloat(io.Next(), 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) PrintLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *Io) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (io *Io) PrintBool(b bool, trueS, falseS string) {
	if b {
		io.PrintLn(trueS)
	} else {
		io.PrintLn(falseS)
	}
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absi(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

type Int64s []int64

func (f Int64s) Len() int {
	return len(f)
}
func (f Int64s) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f Int64s) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func abcInt(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func gcdInt(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

type CompareResult int

const (
	LESS CompareResult = iota - 1
	EQUAL
	GREATER
)

func ToCompareResult(diff int) CompareResult {
	switch {
	case diff < 0:
		return LESS
	case diff > 0:
		return GREATER
	default:
		return EQUAL
	}
}

func Max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

type ValType int

func (l ValType) Compare(r ValType) CompareResult {
	return ToCompareResult(int(l) - int(r))
}

type avlNode struct {
	data                ValType
	parent, left, right *avlNode
	height              int
}

type AvlTree struct {
	root *avlNode
	size int
}

func getHeight(node *avlNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func getBalance(node *avlNode) int {
	ret := 0
	if node.left != nil {
		ret += node.left.height
	}
	if node.right != nil {
		ret -= node.right.height
	}
	return ret
}

func recalcHeight(node *avlNode) {
	node.height = 1 + Max(getHeight(node.left), getHeight(node.right))
}

func rotateRight(node *avlNode) *avlNode {
	newRoot := node.left
	oldRoot := node

	if newRoot.right != nil {
		newRoot.right.parent = oldRoot
	}
	oldRoot.left = newRoot.right
	newRoot.right = oldRoot
	newRoot.parent = oldRoot.parent
	oldRoot.parent = newRoot

	recalcHeight(oldRoot)
	recalcHeight(newRoot)

	return newRoot
}

func rotateLeft(node *avlNode) *avlNode {
	newRoot := node.right
	oldRoot := node

	if newRoot.left != nil {
		newRoot.left.parent = oldRoot
	}
	oldRoot.right = newRoot.left
	newRoot.left = oldRoot
	newRoot.parent = oldRoot.parent
	oldRoot.parent = newRoot

	recalcHeight(oldRoot)
	recalcHeight(newRoot)

	return newRoot
}

func insert(node *avlNode, val ValType) (*avlNode, bool) {
	if node == nil {
		return &avlNode{val, nil, nil, nil, 1}, true
	}

	switch val.Compare(node.data) {
	case EQUAL:
		//duplication is not allowed
		return node, false
	case LESS:
		var inserted bool
		node.left, inserted = insert(node.left, val)
		node.left.parent = node
		if !inserted {
			return node, false
		}
	case GREATER:
		var inserted bool
		node.right, inserted = insert(node.right, val)
		node.right.parent = node
		if !inserted {
			return node, false
		}
	}
	recalcHeight(node)

	balance := getBalance(node)

	if balance > 1 {
		switch val.Compare(node.left.data) {
		case EQUAL:
		case LESS:
			// LL
			return rotateRight(node), true
		case GREATER:
			// LR
			node.left = rotateLeft(node.left)
			return rotateRight(node), true
		}
	} else if balance < -1 {
		switch val.Compare(node.right.data) {
		case EQUAL:
		case LESS:
			// RL
			node.right = rotateRight(node.right)
			return rotateLeft(node), true
		case GREATER:
			// RR
			return rotateLeft(node), true
		}
	} else {
		// balanced
		return node, true
	}

	panic("unreachable")
}

func delete(node *avlNode, val ValType) (*avlNode, bool) {
	if node == nil {
		return node, false
	}

	switch val.Compare(node.data) {
	case EQUAL:
		if node.right == nil {
			return node.left, true
		} else if node.left == nil {
			return node.right, true
		} else {
			//find min
			swapped := node.right
			for swapped.left != nil {
				swapped = swapped.left
			}
			node.data = swapped.data
			node.right, _ = delete(node.right, node.data)
		}
	case LESS:
		var deleted bool
		node.left, deleted = delete(node.left, val)
		if !deleted {
			return node, deleted
		}
	case GREATER:
		var deleted bool
		node.right, deleted = delete(node.right, val)
		if !deleted {
			return node, deleted
		}
	}

	recalcHeight(node)
	balance := getBalance(node)

	if balance > 1 {
		childBalance := getBalance(node.left)
		if childBalance >= 0 {
			return rotateRight(node), true
		} else {
			node.left = rotateLeft(node.left)
			return rotateRight(node), true
		}
	} else if balance < -1 {
		childBalance := getBalance(node.right)
		if childBalance <= 0 {
			return rotateLeft(node), true
		} else {
			node.right = rotateRight(node.right)
			return rotateLeft(node), true
		}
	} else {
		return node, true
	}
}

func walkForToList(node *avlNode, ret *[]ValType) {
	if node == nil {
		return
	}
	walkForToList(node.left, ret)
	*ret = append(*ret, node.data)
	walkForToList(node.right, ret)
}

func NewAvlTree() *AvlTree {
	return &AvlTree{nil, 0}
}

func (avlTree *AvlTree) toList() []ValType {
	ret := make([]ValType, 0, avlTree.size)
	walkForToList(avlTree.root, &ret)
	return ret
}

func (avlTree *AvlTree) Insert(val ValType) {
	var inserted bool
	avlTree.root, inserted = insert(avlTree.root, val)
	if inserted {
		avlTree.size += 1
	}
}

func (avlTree *AvlTree) Delete(val ValType) {
	var deleted bool
	avlTree.root, deleted = delete(avlTree.root, val)
	if deleted {
		avlTree.size -= 1
	}
}

func (avlTree *AvlTree) Smallest() ValType {
	ret := avlTree.root
	for ret.left != nil {
		ret = ret.left
	}
	return ret.data
}

var (
	io *Io
)

type Pii struct {
	f, s int
}

type Piis []Pii

func (f Piis) Len() int {
	return len(f)
}
func (f Piis) Less(i, j int) bool {
	if f[i].f != f[j].f {
		return f[i].f < f[j].f
	}
	return f[i].s < f[j].s
}

func (f Piis) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func solve() {
	n, q := io.NextInt(), io.NextInt()
	var vp []Pii
	for i := 0; i < n; i++ {
		l, r, d := io.NextInt(), io.NextInt(), io.NextInt()
		vp = append(vp, Pii{l - d, d})
		vp = append(vp, Pii{r - d, -d})
	}
	sort.Sort(Piis(vp))

	mp := NewAvlTree()

	for i := 0; i < q; i++ {
		d := io.NextInt()
		for len(vp) > 0 && vp[0].f <= d {
			if vp[0].s > 0 {
				mp.Insert(ValType(vp[0].s))
			} else {
				mp.Delete(ValType(-vp[0].s))
			}
			vp = vp[1:]
		}

		ans := -1
		if mp.size > 0 {
			cur := mp.Smallest()
			ans = int(cur)
		}

		io.PrintLn(ans)
	}
}

func main() {
	io = NewIo()
	f, err := os.Open("src/input")
	if err == nil {
		io.reader = bufio.NewReader(f)
	}
	defer io.Flush()

	solve()
}

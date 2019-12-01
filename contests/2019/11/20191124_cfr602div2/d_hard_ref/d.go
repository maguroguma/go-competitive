// https://codeforces.com/contest/1262/submission/65729700

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"sort"
)

type pstNode struct {
	l, r   int // ノードの担当範囲
	lo, ro *pstNode
	sum    int
}
type pSegmentTree struct {
	nodes        []pstNode
	versionRoots []*pstNode
}

// 初期化は行わない
func newPST(n int) *pSegmentTree {
	maxNodeSize := n * (3 + bits.Len(uint(n)))
	return &pSegmentTree{
		make([]pstNode, 0, maxNodeSize), // 最初は長さ0、十分に容量を確保しておく
		make([]*pstNode, 1, n+1),        // 最初は長さ1、n+1だけ容量を確保
	}
}

// initメソッドから呼ばれる
// t._build(1, n)
func (t *pSegmentTree) _build(l, r int) *pstNode {
	t.nodes = append(t.nodes, pstNode{l: l, r: r}) // 実オブジェクトスライスの方を拡張していく
	o := &t.nodes[len(t.nodes)-1]                  // 最後尾のポインタ
	if l == r {
		return o
	}
	mid := (l + r) >> 1 // 2で割っている
	o.lo = t._build(l, mid)
	o.ro = t._build(mid+1, r)
	return o
}

// 再帰的に呼ばれる
// TreeNodeのポインタを返す
func (t *pSegmentTree) _update(o *pstNode, idx int) *pstNode {
	t.nodes = append(t.nodes, *o)
	o = &t.nodes[len(t.nodes)-1]
	if o.l == o.r {
		o.sum++
		return o
	}
	if idx <= o.lo.r {
		o.lo = t._update(o.lo, idx)
	} else {
		o.ro = t._update(o.ro, idx)
	}
	o.sum = o.lo.sum + o.ro.sum
	return o
}

func (t *pSegmentTree) _queryKth(o1, o2 *pstNode, k int) (idx int) {
	if o1.l == o1.r {
		return o1.l
	}
	if d := o2.lo.sum - o1.lo.sum; d >= k {
		return t._queryKth(o1.lo, o2.lo, k)
	} else {
		return t._queryKth(o1.ro, o2.ro, k-d)
	}
}

func (t *pSegmentTree) init(n int) {
	t.versionRoots[0] = t._build(1, n)
}

// baseVersionは0-basedなAの降順ソート済み配列のインデックスで呼ばれる
// idxは1-basedなAの配列のインデックスで呼ばれる
// len(A)回呼ばれる
func (t *pSegmentTree) update(baseVersion int, idx int) {
	t.versionRoots = append(t.versionRoots, t._update(t.versionRoots[baseVersion], idx))
}

func (t *pSegmentTree) queryKth(l, r int, k int) (idx int) {
	return t._queryKth(t.versionRoots[l-1], t.versionRoots[r], k)
}

// github.com/EndlessCheng/codeforces-go
func Sol1262D2(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n := read()
	a := make([]int, n)
	type pair struct{ val, i int }
	ps := make([]pair, n)
	for i := range ps {
		a[i] = read()
		ps[i] = pair{a[i], i}
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].val > ps[j].val || ps[i].val == ps[j].val && ps[i].i < ps[j].i
	})
	st := newPST(n)
	st.init(n)
	for i, p := range ps {
		st.update(i, p.i+1)
	}

	for m := read(); m > 0; m-- {
		idx := st.queryKth(1, read(), read()) - 1
		Fprintln(out, a[idx])
	}
}

func main() {
	Sol1262D2(os.Stdin, os.Stdout)
}

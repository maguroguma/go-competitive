package main

import "sort"

// ※隣接リストのような形で辺を持たない
// ベルマンフォード法のような辺の管理を行う

// 頂点fromから頂点toへのコストcostの辺
// ソートする必要があるためsorterインタフェースを実装する
// type Edge struct {
// 	from, to, cost int
// }

var es []Edge // 辺
var v, e int  // vは頂点数, eは辺数

func kruskal() int {
	L := EdgeList{}
	sort.Stable(byKey{L})

	// 閉路判定用のUF木
	uf := NewUnionFind(v)

	res := 0
	// すべての辺についてチェックする
	for _, e := range L {
		// 辺の両端が同じ連結成分に属していない場合のみ全域木のコストに加算する
		if !uf.Same(e.from, e.to) {
			uf.Unite(e.from, e.to)
			res += e.cost
		}
	}

	return res
}

type Edge struct {
	key            int
	from, to, cost int
}
type EdgeList []*Edge
type byKey struct {
	EdgeList
}

func (l EdgeList) Len() int {
	return len(l)
}
func (l EdgeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byKey) Less(i, j int) bool {
	return l.EdgeList[i].key < l.EdgeList[j].key
}

// how to use
// L := make(EdgeList, 0, 200000+5)
// L = append(L, &Edge{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

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

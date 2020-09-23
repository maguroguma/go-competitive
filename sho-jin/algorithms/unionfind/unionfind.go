package unionfind

import "errors"

// 各ノード番号は0-based index
type UnionFindTree struct {
	parents [100005]int

	// 連結成分のサイズも計算できる
	ccsize [100005]int
}

// 最大ノード数は100005
func NewUnionFindTree(n int) *UnionFindTree {
	if n > 100005 {
		panic(errors.New("argument error: max node number is 100005"))
	}

	uft := new(UnionFindTree)
	// 最初はすべてのノードが根ノード
	for i := 0; i < n; i++ {
		uft.parents[i] = i
		uft.ccsize[i] = 1
	}

	return uft
}

// xとyの属する集合を併合
func (uft *UnionFindTree) Unite(x, y int) {
	xp := uft.root(x)
	yp := uft.root(y)
	// もともと併合済の場合は経路圧縮のみがなされる
	if xp == yp {
		return
	}

	// ypが2つの集合の新しい根ノードになる
	uft.parents[xp] = yp

	// 連結成分のサイズを更新
	ss := uft.ccsize[xp] + uft.ccsize[yp]
	uft.ccsize[yp] = ss
}

// xとyが同じ集合に属するか否かを判定
func (uft *UnionFindTree) Same(x, y int) bool {
	return uft.root(x) == uft.root(y)
}

// 木の根を求める、Union Findは根ノードを求める関数を中心に動作する
// **経路圧縮: 上向きにたどって再帰的に根を調べる際に、調べたら辺を根に直接つなぎ直す（xの親を根に変える）**
// 再帰関数
func (uft *UnionFindTree) root(x int) int {
	if uft.parents[x] == x { // 根に到着
		return x
	} else { // 根に向かって進行
		// 引数xの親ノードに対して、さらに根ノードを求める
		// **returnする前にuft.parents[x]を更新することで経路圧縮を行う
		// 再帰的に実行されるためxの親であるparents[x]の親も更新される（＝巡回するすべての親ノードの親が根ノードに更新される）**
		uft.parents[x] = uft.root(uft.parents[x])
		return uft.parents[x]
	}
}

func (uft *UnionFindTree) ccSize(x int) int {
	return uft.ccsize[uft.root(x)]
}

package main

const MAX_V = 100000

var G [MAX_V][]int // グラフの隣接リスト表現
var root int       // 根ノードの番号

var parent [MAX_V]int // 親ノード（根ノードの親は-1とする）
var depth [MAX_V]int

func dfs(v, p, d int) {
	parent[v] = p
	depth[v] = d
	for _, to := range G[v] {
		if to != p {
			dfs(to, v, d+1)
		}
	}
}

// 初期化
func initialize() {
	// parentとdepthを初期化する
	dfs(root, -1, 0)
}

// uとvのLCAを求める
func lca(u, v int) int {
	//u, vの深さがおなじになるまで親をたどる
	for depth[u] > depth[v] {
		u = parent[u]
	}
	for depth[v] > depth[u] {
		v = parent[v]
	}
	// 同じ頂点に到達するまで親をたどる
	for u != v {
		u = parent[u]
		v = parent[v]
	}
	return u
}

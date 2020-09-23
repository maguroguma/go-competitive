package main

const MAX_V = 100000 + 5
const MAX_LOG_V = 60 + 1

var G [MAX_V][]int // グラフの隣接リスト表現
var root int       // 根ノードの番号
var V int          // ノードの数

// 親を2^k回たどって到達する頂点（根を通り過ぎる場合は-1とする）
var parent [MAX_LOG_V][MAX_V]int

// 根からの深さ
var depth [MAX_V]int

func dfs(v, p, d int) {
	parent[0][v] = p
	depth[v] = d
	for _, to := range G[v] {
		if to != p {
			dfs(to, v, d+1)
		}
	}
}

// 初期化
func initialize() {
	// parent[0]とdepthを初期化する
	dfs(root, -1, 0)
	// parentを初期化する
	for k := 0; k+1 < MAX_LOG_V; k++ {
		for v := 0; v < V; v++ {
			if parent[k][v] < 0 {
				parent[k+1][v] = -1
			} else {
				parent[k+1][v] = parent[k][parent[k][v]]
			}
		}
	}
}

// uとvのLCAを求める
func lca(u, v int) int {
	// u, vの深さがおなじになるまで親をたどる
	if depth[u] > depth[v] {
		u, v = v, u
	}
	for k := 0; k < MAX_LOG_V; k++ {
		if ((depth[v] - depth[u]) >> uint(k) & 1) == 1 {
			v = parent[k][v]
		}
	}

	if u == v {
		return u
	}

	// 二分探索でLCAを求める
	for k := MAX_LOG_V - 1; k >= 0; k-- {
		if parent[k][u] != parent[k][v] {
			u = parent[k][u]
			v = parent[k][v]
		}
	}
	return parent[0][u]
}

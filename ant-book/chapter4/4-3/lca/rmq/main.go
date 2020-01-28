package main

const MAX_V = 100000 + 5

var G [MAX_V][]int
var root int

var vs [MAX_V*2 - 1]int    // DFSでの訪問順
var depth [MAX_V*2 - 1]int // 根からの深さ
var id [MAX_V]int          // 各頂点がvsにはじめて登場するインデックス

func dfs(v, p, d int, k *int) {
	id[v] = *k
	vs[*k] = v
	depth[*k] = d
	(*k)++

	for _, to := range G[v] {
		if to != p {
			dfs(to, v, d+1, k)
			vs[*k] = v
			depth[*k] = d
			(*k)++
		}
	}
}

// 初期化
func initialize(V int) {
	// vs, depth, idを初期化する
	k := 0
	dfs(root, -1, 0, &k)
	// RMQを初期化する（最小値ではなく、最小値のインデックスを返すようにする）
	// rmq_init(depth, V*2-1)
}

// u, vのLCAを求める
func lca(u, v int) int {
	// return vs[query(Min(id[u], id[v]), Max(id[u], id[v]) + 1)]
	return -1
}

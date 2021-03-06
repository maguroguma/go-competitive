package main

// ※隣接リストのような形で辺を持たない

// 頂点fromから頂点toへのコストcostの辺
type Edge struct {
	from, to, cost int
}

const INF = 1 << 60

var es []Edge  // 辺
var dist []int // 最短距離
var v, e int   // vは頂点数, eは辺数

// s番目の頂点から各頂点への最短距離を求める
func shortestPath(s int) {
	// 初期化
	for i := 0; i < v; i++ {
		dist[i] = INF
	}
	dist[s] = 0

	for {
		isUpdate := false

		for i := 0; i < e; i++ {
			e := es[i]
			if dist[e.from] != INF && dist[e.to] > dist[e.from]+e.cost {
				dist[e.to] = dist[e.from] + e.cost
				isUpdate = true
			}
		}

		// 更新がなかったらループを抜ける
		if !isUpdate {
			break
		}
	}
}

// trueなら負の閉路が存在する
func findNegativeLoop() bool {
	// すべてを始点とみなすことで、すべての負の閉路を検出する（？）
	for i := 0; i < v; i++ {
		dist[i] = 0
	}

	for i := 0; i < v; i++ {
		for j := 0; j < e; j++ {
			e := es[j]
			if dist[e.to] > dist[e.from]+e.cost {
				dist[e.to] = dist[e.from] + e.cost

				// n回目にも更新があるなら負の閉路が存在する
				if i == v-1 {
					return true
				}
			}
		}
	}

	return false
}

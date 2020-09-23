package main

// 隣接行列でグラフを表す

// dist[u][v]は辺e=(u, v)のコスト
// 存在しない場合はINF
// ただしd[i][i] = 0とする

// 今回は配列の再利用をせずに3次元のDPテーブルを保持したままとする
// dp[k+1][i][j] := 頂点0~kとi, jのみを使う場合の、iからjへの最短路
// 知りたい答えは dp[MAX_V][i][j]
var dp [][][]int
var v int

func warshallFloyd() {
	for k := 0; k < v; k++ {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				// 1. 頂点kをちょうど一度通る場合
				// 2. 頂点kを全く通らない場合
				// の排反な2ケースを加味したもの
				dp[k+1][i][j] = Min(dp[k][i][j], dp[k][i][k]+dp[k][k][j])
			}
		}
	}
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

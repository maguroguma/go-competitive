# ダイクストラ法

辺の重みが **非負数** の場合の **単一始点** 最短経路問題を解くための **最良優先探索** によるアルゴリズム。

## 計算量

priority queueはかならず使う。
計算量的に使わないと無理なケースが多い上に、実装もだいぶシンプルになる。

優先度付きキュー（二分ヒープ）: `O((E+V)log(V))`

Golangのheapパッケージの実装は調査できていないが、最悪を考慮してこの計算量を認識しておく。

だいぶ雑だが、 `N=(エッジ数 + 頂点数)` とした `O(NlogN)` だと思っておく。

## 擬似コード（スニペット）

グラフを拡張する場合はもっと複雑になるが、基本的にはこのような形になるはず。

```golang
// 全頂点について、始点からの距離は無限大としておく
for (すべての頂点) {
  dp[i] = INF
}
// 始点は当然0
dp[0] = 0

// queueの初期化
heap.Push(pq, &Node{priority: dp[0], id: 0})
for pq.Len() > 0 {
  // この時点でノードの最短距離が確定（フラグ立てができる）
  node := heap.Pop(pq).(*Node)

  // 選択中のノードのエッジをすべて見る
  for _, next := range edges[node.id] {
    // 選択中のノードを経由した場合のコストが暫定よりも小さかった場合、更新してqueueに突っ込む
    if dp[next.tid] > node.priority + next.cost {
      dp[next.tid] = node.priority + next.cost
      heap.Push(pq, &Node{priority: dp[next.tid], id: next.tid})
    }
  }
}

// queueが空になるまで続けるので、（到達できるノードであれば）最短距離が求まる
fmt.Println(dp[n-1])
```

## 最短経路の復元

2019/07/07時点でまだ解いたことがないが、最短距離を持つ `dp` テーブルのほか、
1つ前のノードを持つ `prev` テーブルを各ノードに対して用意してやれば良い。
**最短路が更新されたときと同時に、 `dp` に加えて `prev` も更新してやれば良い。**

## 練習問題など

- [いろはちゃんコンテストDay2 - G:通学路](https://atcoder.jp/contests/iroha2019-day2/tasks/iroha2019_day2_g)
  - twitterでみかけたもの。これはダイクストラ？


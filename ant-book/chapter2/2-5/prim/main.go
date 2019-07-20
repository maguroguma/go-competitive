package main

import "container/heap"

// 隣接リストで辺を管理
// ダイクストラ法に非常に似ている

type Edge struct {
	to, cost int
}

const INF = 1 << 60

var v int       // 頂点数
var G [][]Edge  // 隣接リスト（グラフそのものを表すデータと言えるため、Gと命名）
var used []bool // 頂点iがXに含まれているか（暫定の全域木Tに含まれているか）

func prim() int {
	// 昇順で取り出せるpriority queue
	temp := make(NodePQ, 0, 100000+1)
	pq := &temp
	heap.Init(pq)

	// 初期化
	// 0から暫定の全域木Tを作っていく
	for _, e := range G[0] {
		heap.Push(pq, &Node{pri: e.cost, cost: e.cost, id: e.to})
	}

	res := 0
	for pq.Len() > 0 {
		node := heap.Pop(pq).(*Node)
		v := node.id

		// すでに全域木Tに組み込まれている場合は無視する
		if used[v] {
			continue
		}

		// ノードへ到達するための辺のコストを加算する
		res += node.cost

		for _, e := range G[v] {
			// vから伸びる先のノードがすでに全域木Tに組み込まれている場合は、その辺は無視する
			if used[e.to] {
				continue
			}
			heap.Push(pq, &Node{pri: e.cost, cost: e.cost, id: e.to})
		}
	}

	return res
}

type Node struct {
	pri      int
	cost, id int
}
type NodePQ []*Node

func (pq NodePQ) Len() int           { return len(pq) }
func (pq NodePQ) Less(i, j int) bool { return pq[i].pri < pq[j].pri } // <: ASC, >: DESC
func (pq NodePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *NodePQ) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}
func (pq *NodePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// how to use
// temp := make(NodePQ, 0, 100000+1)
// pq := &temp
// heap.Init(pq)
// heap.Push(pq, &Node{pri: intValue})
// popped := heap.Pop(pq).(*Node)

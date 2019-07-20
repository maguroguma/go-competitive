package main

import "container/heap"

// 隣接リストで辺を管理

type Edge struct {
	to, cost int
}

const INF = 1 << 60

var v int      // 頂点数
var G [][]Edge // 隣接リスト（グラフそのものを表すデータと言えるため、Gと命名）
var dist []int // 最短距離

func dijkstra(s int) {
	// 昇順で取り出せるpriority queue
	temp := make(NodePQ, 0, 100000+1)
	pq := &temp
	heap.Init(pq)

	// 初期化
	for i := 0; i < v; i++ {
		dist[i] = INF
	}
	dist[s] = 0
	heap.Push(pq, &Node{pri: dist[s], dist: dist[s], id: s})

	for pq.Len() > 0 {
		node := heap.Pop(pq).(*Node)
		v := node.id

		// 最短距離でなければ無視する
		if dist[v] < node.dist {
			continue
		}

		for _, e := range G[v] {
			if dist[e.to] > dist[v]+e.cost {
				dist[e.to] = dist[v] + e.cost
				heap.Push(pq, &Node{pri: dist[e.to], dist: dist[e.to], id: e.to})
			}
		}
	}
}

type Node struct {
	pri      int
	dist, id int
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

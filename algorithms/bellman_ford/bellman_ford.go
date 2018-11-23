package bellman_ford

import (
	"github.com/myokoyama0712/go-competitive/algorithms/graph"
)

type BellmanFord struct {
	g           *graph.Graph
	start, goal int
	nodeMap     NodeMap
}

type Node struct {
	cost     int
	previous int
}
type NodeMap map[int]*Node

func NewBellmanFord(graph *graph.Graph, start, goal int) *BellmanFord {
	bf := new(BellmanFord)
	bf.start, bf.goal = start, goal
	bf.g = graph
	nodeMap := make(NodeMap)
	for i := 0; i < graph.GetNodeNumber(); i++ {
		node := new(Node)
		if i == start {
			node.cost = 0
		} else {
			node.cost = 1000
		}
		node.previous = -1
		nodeMap[i] = node
	}
	bf.nodeMap = nodeMap

	return bf
}

func (bf *BellmanFord) GetShortestPath() int {
	// 0. 初期化

	// 1. 辺の緩和（relaxing）を反復
	// 更新がなくなる判定は行わない
	for counter := 0; counter < bf.g.GetNodeNumber()-1; counter++ {
		// 全エッジに対してループ
		for i := 0; i < bf.g.GetNodeNumber(); i++ {
			for j := 0; j < bf.g.GetNodeNumber(); j++ {
				if w, e := bf.g.GetEdgeWeight(i, j); w >= 0 && e == nil { // FIXME: ここでは e == nil だけでは正しく動かない
					source, destination, cost := bf.nodeMap[i], bf.nodeMap[j], w
					// 行き先へのもともとのコストよりも、注目エッジを介したコストのほうが小さい場合、更新する
					oldCost := destination.cost
					newCost := source.cost + cost
					if oldCost > newCost {
						destination.cost = newCost
						destination.previous = i
					}
				}
			}
		}
	}

	// 負の重みの閉路がないかチェック（収束後も更新が行われるかどうか）
	// 全エッジに対してループ
	for i := 0; i < bf.g.GetNodeNumber(); i++ {
		for j := 0; j < bf.g.GetNodeNumber(); j++ {
			if w, e := bf.g.GetEdgeWeight(i, j); w >= 0 && e == nil { // FIXME: ここでは e == nil だけでは正しく動かない
				source, destination, cost := bf.nodeMap[i], bf.nodeMap[j], w
				// 収束後、あるエッジを介して目的地へ向かったときに、得られているはずの目的地への最短距離よりもさらに小さくできる場合、閉路が存在する
				if source.cost+cost < destination.cost {
					return -1000000007
				}
			}
		}
	}

	return bf.nodeMap[bf.goal].cost
}

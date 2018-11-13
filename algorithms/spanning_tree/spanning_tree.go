package spanning_tree

import (
	"sort"

	"github.com/myokoyama0712/algorithm/graph"
)

type Kruskal struct {
	g           *graph.Graph
	scanedEdges EdgeSlice
	trees       map[int]*Tree
	maxTreeId   int
}

type Edge struct {
	startNodeId int
	endNodeId   int
	weight      int
}
type EdgeSlice []*Edge

func (es EdgeSlice) Len() int {
	return len(es)
}
func (es EdgeSlice) Less(i, j int) bool {
	return es[i].weight < es[j].weight
}
func (es EdgeSlice) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

type Tree struct {
	nodeList map[int]int
	edgeList []*Edge
}

func (t *Tree) HaveNode(nodeId int) bool {
	for _, id := range t.nodeList {
		if id == nodeId {
			return true
		}
	}
	return false
}

func NewKruskal(g *graph.Graph) *Kruskal {
	kruskal := new(Kruskal)
	kruskal.g = g
	kruskal.trees = map[int]*Tree{}
	kruskal.scanedEdges = []*Edge{}

	// setting tree slice (forest)
	for i := 0; i < g.GetNodeNumber(); i++ {
		kruskal.trees[i] = &Tree{nodeList: map[int]int{i: i}, edgeList: []*Edge{}}
		kruskal.maxTreeId = i
	}

	// setting scanned edge slice
	for i := 0; i < g.GetNodeNumber(); i++ {
		for j := i + 1; j < g.GetNodeNumber(); j++ {
			if weight, err := g.GetEdgeWeight(i, j); weight >= 0 && err == nil {
				e := &Edge{startNodeId: i, endNodeId: j, weight: weight}
				kruskal.scanedEdges = append(kruskal.scanedEdges, e)
			}
		}
	}
	sort.Sort(kruskal.scanedEdges)

	return kruskal
}

func (kruskal *Kruskal) IsConnectionEdge(edge *Edge) bool {
	sid := edge.startNodeId
	eid := edge.endNodeId

	for _, tree := range kruskal.trees {
		sbool := tree.HaveNode(sid)
		ebool := tree.HaveNode(eid)
		if (sbool && !ebool) || (!sbool && ebool) {
			return true
		}
	}
	return false
}

func (kruskal *Kruskal) ConnectTree(edge *Edge) {
	// get two tree including edge's start point and end point
	//  and then, delete these two trees from Kruskal structure
	var tree1, tree2 *Tree
	sid := edge.startNodeId
	eid := edge.endNodeId
	for key, tree := range kruskal.trees {
		if tree.HaveNode(sid) {
			tree1 = tree
			delete(kruskal.trees, key)
		} else if tree.HaveNode(eid) {
			tree2 = tree
			delete(kruskal.trees, key)
		}
	}

	// merge the two trees and add it to Kruskal structure
	tree1.edgeList = append(tree1.edgeList, edge)
	for _, edge := range tree2.edgeList {
		tree1.edgeList = append(tree1.edgeList, edge)
	}
	for key, node := range tree2.nodeList {
		tree1.nodeList[key] = node
	}
	kruskal.maxTreeId++
	kruskal.trees[kruskal.maxTreeId] = tree1
}

func (kruskal *Kruskal) ComputeMinimalSpanningTree() *Tree {
	for _, edge := range kruskal.scanedEdges {
		if kruskal.IsConnectionEdge(edge) {
			kruskal.ConnectTree(edge)
		}
	}

	return kruskal.trees[kruskal.maxTreeId]
}

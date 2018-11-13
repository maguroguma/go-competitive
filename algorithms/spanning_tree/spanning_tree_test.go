package spanning_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/myokoyama0712/algorithm/graph"
)

func setupExampleGraph() *graph.Graph {
	g := graph.NewGraph(7)
	edges := [11][3]int{
		[3]int{0, 1, 7},
		[3]int{1, 2, 8},
		[3]int{0, 3, 5},
		[3]int{1, 3, 9},
		[3]int{1, 4, 7},
		[3]int{2, 4, 5},
		[3]int{3, 4, 15},
		[3]int{3, 5, 6},
		[3]int{4, 5, 8},
		[3]int{4, 6, 9},
		[3]int{5, 6, 11},
	}
	for _, edge := range edges {
		g.SetEdgeWeight(edge[0], edge[1], edge[2])
	}
	return g
}

func TestIsExampleGraphCorrect(t *testing.T) {
	g := setupExampleGraph()
	actual, _ := g.GetEdgeWeight(1, 0)
	assert.Equal(t, 7, actual)
	actual, _ = g.GetEdgeWeight(1, 5)
	assert.Equal(t, -1, actual)
}

func TestInitKruskalOfExampleGraph(t *testing.T) {
	g := setupExampleGraph()
	kruskal := NewKruskal(g)
	sortedEdgeWeights := []int{5, 5, 6, 7, 7, 8, 8, 9, 9, 11, 15}
	actual := []int{}
	for _, edge := range kruskal.scanedEdges {
		actual = append(actual, edge.weight)
	}
	assert.Equal(t, 11, len(kruskal.scanedEdges))
	assert.Equal(t, map[int]int{0: 0}, kruskal.trees[0].nodeList)
	assert.Equal(t, []*Edge{}, kruskal.trees[0].edgeList)
	assert.Equal(t, sortedEdgeWeights, actual)
}

func TestConnectingTreesByEdge(t *testing.T) {
	g := setupExampleGraph()
	kruskal := NewKruskal(g)
	isConnectionEdge := kruskal.IsConnectionEdge(kruskal.scanedEdges[0])
	assert.True(t, isConnectionEdge)
	kruskal.ConnectTree(kruskal.scanedEdges[0])
	assert.Equal(t,
		&Tree{nodeList: map[int]int{0: 0, 3: 3}, edgeList: []*Edge{&Edge{startNodeId: 0, endNodeId: 3, weight: 5}}},
		kruskal.trees[7])
}

func TestKruskalAlgorithm(t *testing.T) {
	g := setupExampleGraph()
	kruskal := NewKruskal(g)
	answer := [6][3]int{
		[3]int{0, 1, 7},
		[3]int{0, 3, 5},
		[3]int{1, 4, 7},
		[3]int{2, 4, 5},
		[3]int{3, 5, 6},
		[3]int{4, 6, 9},
	}
	minimalSpanningTree := []*Edge{}
	for _, e := range answer {
		minimalSpanningTree = append(minimalSpanningTree,
			&Edge{startNodeId: e[0], endNodeId: e[1], weight: e[2]})
	}
	actual := kruskal.ComputeMinimalSpanningTree()
	assert.ElementsMatch(t, minimalSpanningTree, actual.edgeList)
	assert.Equal(t, 1, len(kruskal.trees))
}

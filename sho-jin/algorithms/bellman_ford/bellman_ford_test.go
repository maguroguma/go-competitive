package bellman_ford

import (
	"testing"

	"github.com/myokoyama0712/go-competitive/algorithms/graph"
	"github.com/stretchr/testify/assert"
)

func setupGraphExample() *graph.Graph {
	g := graph.NewGraph(6)
	g.SetEdgeWeight(0, 1, 7)
	g.SetEdgeWeight(0, 2, 9)
	g.SetEdgeWeight(0, 5, 14)
	g.SetEdgeWeight(1, 2, 10)
	g.SetEdgeWeight(1, 3, 15)
	g.SetEdgeWeight(2, 3, 11)
	g.SetEdgeWeight(2, 5, 2)
	g.SetEdgeWeight(3, 4, 6)
	g.SetEdgeWeight(4, 5, 9)
	return g
}

func TestShortestPathByDijkstra(t *testing.T) {
	g := setupGraphExample()
	dijkstra := graph.NewDijkstra(g, 0, 4)
	actual := dijkstra.GetShortestPath()
	assert.Equal(t, "0, 2, 5, 4, : cost: 20, step: 6", actual)
}

func TestShortestPathByBellmanFord(t *testing.T) {
	g := setupGraphExample()
	bfAlgorithm := NewBellmanFord(g, 0, 4)
	shortestPath := bfAlgorithm.GetShortestPath()
	assert.Equal(t, 20, shortestPath)

	bfAlgorithm = NewBellmanFord(g, 0, 2)
	shortestPath = bfAlgorithm.GetShortestPath()
	assert.Equal(t, 9, shortestPath)

	bfAlgorithm = NewBellmanFord(g, 5, 1)
	shortestPath = bfAlgorithm.GetShortestPath()
	assert.Equal(t, 12, shortestPath)
}

package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDijkstraInitialStateCorrect(t *testing.T) {
	graph := setupGraphExample()
	dijkstra := NewDijkstra(graph, 0, 4)
	assert.Equal(t, 0, dijkstra.startNodeId)
	assert.Equal(t, 4, dijkstra.goalNodeId)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, dijkstra.candidates)
	expected := make(NodeMap)
	for i := 0; i < graph.GetNodeNumber(); i++ {
		expected[i] = &Node{cost: 1000, previousNodeId: -1}
	}
	expected[0].cost = 0
	assert.Equal(t, expected, dijkstra.nodeMap)
}

func TestPopMinimalCostNode(t *testing.T) {
	graph := setupGraphExample()
	dijkstra := NewDijkstra(graph, 0, 4)
	assert.Equal(t, 0, dijkstra.PopMinimalCostNode())
	assert.Equal(t, []int{1, 2, 3, 4, 5}, dijkstra.candidates)
	dijkstra = NewDijkstra(graph, 3, 4)
	assert.Equal(t, 3, dijkstra.PopMinimalCostNode())
	assert.Equal(t, []int{0, 1, 2, 4, 5}, dijkstra.candidates)
	assert.Equal(t, 0, dijkstra.PopMinimalCostNode())
	assert.Equal(t, []int{1, 2, 4, 5}, dijkstra.candidates)
	assert.Equal(t, 1, dijkstra.PopMinimalCostNode())
	assert.Equal(t, []int{2, 4, 5}, dijkstra.candidates)
}

func TestShortestPath(t *testing.T) {
	graph := setupGraphExample()
	dijkstra := NewDijkstra(graph, 0, 4)
	actual := dijkstra.GetShortestPath()
	assert.Equal(t, "0, 2, 5, 4, : 20", actual)
	dijkstra = NewDijkstra(graph, 2, 1)
	assert.Equal(t, "2, 1, : 10", dijkstra.GetShortestPath())
	dijkstra = NewDijkstra(graph, 5, 3)
	assert.Equal(t, "5, 2, 3, : 13", dijkstra.GetShortestPath())
}

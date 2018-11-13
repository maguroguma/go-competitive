package warshall_floyd

import (
	"testing"

	"github.com/myokoyama0712/algorithm/graph"

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

func TestIsWarshallFloydStructCorrect(t *testing.T) {
	g := setupGraphExample()
	warflo := NewWarshallFloyd(g)
	assert.Equal(t, &ShortestPath{cost: 9, path: []int{2, 0}}, warflo.matrix[2][0])
	assert.Equal(t, &ShortestPath{cost: 9, path: []int{0, 2}}, warflo.matrix[0][2])
	assert.Equal(t, &ShortestPath{cost: 1000000, path: []int{3}}, warflo.matrix[3][0])
	assert.Equal(t, &ShortestPath{cost: 1000000, path: []int{0}}, warflo.matrix[0][3])
	assert.Equal(t, &ShortestPath{cost: 1000000, path: []int{0}}, warflo.matrix[0][0])
	assert.Equal(t, &ShortestPath{cost: 1000000, path: []int{1}}, warflo.matrix[1][1])
}

func TestShortestPathByWarshallFloyd(t *testing.T) {
	g := setupGraphExample()
	warflo := NewWarshallFloyd(g)
	shortestPath, _ := warflo.GetShortestPath(0, 4)
	assert.Equal(t, &ShortestPath{cost: 20, path: []int{0, 2, 5, 4}}, shortestPath)
	shortestPath, _ = warflo.GetShortestPath(0, 2)
	assert.Equal(t, &ShortestPath{cost: 9, path: []int{0, 2}}, shortestPath)
	shortestPath, _ = warflo.GetShortestPath(5, 1)
	assert.Equal(t, &ShortestPath{cost: 12, path: []int{5, 2, 1}}, shortestPath)
}

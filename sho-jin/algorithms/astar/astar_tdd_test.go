package astar

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/myokoyama0712/algorithm/graph"
)

func setupGraphExample() *graph.Graph {
	g := graph.NewGraph(29)
	edgeList := [][2]int{
		{0, 1}, {1, 2},
		{0, 3}, {3, 5}, {5, 11}, {11, 14}, {14, 19}, {19, 22},
		{22, 23}, {23, 24},
		{2, 4}, {4, 6}, {6, 12}, {12, 15}, {15, 20}, {20, 24},
		{6, 7}, {7, 8}, {8, 9}, {9, 10}, {10, 13}, {13, 18}, {18, 17}, {17, 16}, {16, 21}, {21, 26},
		{24, 25}, {25, 26}, {26, 27}, {27, 28},
	}
	for _, e := range edgeList {
		g.SetEdgeWeight(e[0], e[1], 1)
	}
	return g
}

func setupAstarExample(startNodeId, goalNodeId int) *Astar {
	g := setupGraphExample()
	coordinatesMap := make(map[int][2]int)
	coordinatesList := [29][2]int{
		{0, 0}, {1, 0}, {2, 0},
		{0, 1}, {2, 1},
		{0, 2}, {2, 2}, {3, 2}, {4, 2}, {5, 2}, {6, 2},
		{0, 3}, {2, 3}, {6, 3},
		{0, 4}, {2, 4}, {4, 4}, {5, 4}, {6, 4},
		{0, 5}, {2, 5}, {4, 5},
		{0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6},
	}
	for i, c := range coordinatesList {
		coordinatesMap[i] = c
	}
	astar := NewAstar(g, startNodeId, goalNodeId, coordinatesMap)
	return astar
}

func TestIsExampleGraphCorrect(t *testing.T) {
	g := setupGraphExample()
	assert.Equal(t, 30, g.GetEdgeNumber())
}

func TestGetShortestPathByDijkstra(t *testing.T) {
	g := setupGraphExample()
	dijkstra := graph.NewDijkstra(g, 6, 28)
	assert.Equal(t, "6, 12, 15, 20, 24, 25, 26, 27, 28, : 8", dijkstra.GetShortestPath())
}

func TestHeulisticCost(t *testing.T) {
	astar := setupAstarExample(6, 28)
	actual := []int{}
	actual = append(actual, astar.nodeMap[6].heulisticCostToGoal)
	actual = append(actual, astar.nodeMap[24].heulisticCostToGoal)
	actual = append(actual, astar.nodeMap[8].heulisticCostToGoal)
	actual = append(actual, astar.nodeMap[0].heulisticCostToGoal)
	actual = append(actual, astar.nodeMap[28].heulisticCostToGoal)
	assert.Equal(t, []int{8, 4, 6, 12, 0}, actual)
}

func TestShortestPath(t *testing.T) {
	astar := setupAstarExample(6, 28)
	assert.Equal(t, "6, 12, 15, 20, 24, 25, 26, 27, 28, : 8", astar.GetShortestPath())
	astar = setupAstarExample(9, 27)
	assert.Equal(t, "9, 10, 13, 18, 17, 16, 21, 26, 27, : 8", astar.GetShortestPath())
}

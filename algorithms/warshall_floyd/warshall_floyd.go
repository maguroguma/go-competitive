package warshall_floyd

import (
	"errors"

	"github.com/myokoyama0712/algorithm/graph"
)

type WarshallFloyd struct {
	g      *graph.Graph
	matrix [][]*ShortestPath
}

type ShortestPath struct {
	cost int
	path []int
}

func NewWarshallFloyd(g *graph.Graph) *WarshallFloyd {
	pseudoInfinite := 1000000
	warflo := new(WarshallFloyd)
	warflo.g = g
	warflo.matrix = make([][]*ShortestPath, g.GetNodeNumber())
	for i := 0; i < g.GetNodeNumber(); i++ {
		row := make([]*ShortestPath, g.GetNodeNumber())
		for j := 0; j < g.GetNodeNumber(); j++ {
			shortestPath := new(ShortestPath)
			shortestPath.path = []int{i}
			if weight, err := g.GetEdgeWeight(i, j); weight >= 0 && err == nil {
				shortestPath.cost = weight
				shortestPath.path = append(shortestPath.path, j)
			} else if weight == -1 {
				shortestPath.cost = pseudoInfinite
			}
			row[j] = shortestPath
		}
		warflo.matrix[i] = row
	}

	return warflo
}

func (w *WarshallFloyd) GetShortestPath(startNodeId, goalNodeId int) (*ShortestPath, error) {
	if startNodeId < 0 || w.g.GetNodeNumber() <= startNodeId ||
		goalNodeId < 0 || w.g.GetNodeNumber() <= goalNodeId {
		return nil, errors.New("out of range error")
	} else if startNodeId == goalNodeId {
		return nil, errors.New("same argument error")
	}

	for k := 0; k < w.g.GetNodeNumber(); k++ {
		for i := 0; i < w.g.GetNodeNumber(); i++ {
			for j := 0; j < w.g.GetNodeNumber(); j++ {
				if w.matrix[i][j].cost > w.matrix[i][k].cost+w.matrix[k][j].cost {
					w.matrix[i][j].cost = w.matrix[i][k].cost + w.matrix[k][j].cost
					w.matrix[i][j].path = append(w.matrix[i][k].path, w.matrix[k][j].path[1:]...)
				}
			}
		}
	}

	return w.matrix[startNodeId][goalNodeId], nil
}

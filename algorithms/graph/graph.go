package graph

import "errors"

type Graph struct {
	adjMatrix [][]int
}

//type NodeNumberError struct {
//	msg string
//}
//func (e *NodeNumberError) Error() string {
//	return e.msg
//}
// usage:
//return &NodeNumberError{"[Argument Error]: too small, or too large"}
//return &NodeNumberError{"[Argument Error]: cannot make edge between same node"}

func NewGraph(nodeNumber int) *Graph {
	graph := new(Graph)
	matrix := make([][]int, nodeNumber)
	for i := 0; i < nodeNumber; i++ {
		row := make([]int, nodeNumber)
		for j := 0; j < nodeNumber; j++ {
			row[j] = -1
		}
		matrix[i] = row
	}
	graph.adjMatrix = matrix

	return graph
}

func (g *Graph) GetNodeNumber() int {
	return len(g.adjMatrix)
}

func (g *Graph) GetEdgeWeight(i, j int) (int, error) {
	if i < 0 || len(g.adjMatrix) <= i || j < 0 || len(g.adjMatrix) <= j {
		return -1, errors.New("out of range error")
	}
	if i == j {
		return -1, errors.New("same argument error")
	}
	return g.adjMatrix[i][j], nil
}

func (g *Graph) SetEdgeWeight(i, j, weight int) error {
	if i < 0 || len(g.adjMatrix) <= i || j < 0 || len(g.adjMatrix) <= j {
		return errors.New("out of range error")
	}
	if i == j {
		return errors.New("same argument error")
	}
	g.adjMatrix[i][j] = weight
	g.adjMatrix[j][i] = weight
	return nil
}

func (g *Graph) GetEdgeNumber() int {
	edgeNumber := 0
	for i := 0; i < g.GetNodeNumber(); i++ {
		for j := i + 1; j < g.GetNodeNumber(); j++ {
			if weight, err := g.GetEdgeWeight(i, j); weight >= 0 && err == nil {
				edgeNumber++
			}
		}
	}
	return edgeNumber
}

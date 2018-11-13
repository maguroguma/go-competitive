package graph

import "fmt"

type Dijkstra struct {
	startNodeId, goalNodeId int
	graph                   *Graph
	candidates              []int
	nodeMap                 NodeMap
}
type Node struct {
	cost           int
	previousNodeId int
}
type NodeMap map[int]*Node

func NewDijkstra(graph *Graph, startNodeId, goalNodeId int) *Dijkstra {
	dijkstra := new(Dijkstra)
	dijkstra.startNodeId = startNodeId
	dijkstra.goalNodeId = goalNodeId
	dijkstra.graph = graph
	candidates := []int{}
	nodeMap := make(NodeMap)
	for i := 0; i < graph.GetNodeNumber(); i++ {
		node := new(Node)
		if i == startNodeId {
			node.cost = 0
		} else {
			node.cost = 1000
		}
		node.previousNodeId = -1
		nodeMap[i] = node
		candidates = append(candidates, i)
	}
	dijkstra.candidates = candidates
	dijkstra.nodeMap = nodeMap

	return dijkstra
}

// PopMinimalCostNode returns node that has minimal cost
//  within candidates nodes.
// And the returned node is deleted from candidates nodes.
func (d *Dijkstra) PopMinimalCostNode() int {
	minimalCost := 1000000
	nodeId := -1
	for _, id := range d.candidates {
		if d.nodeMap[id].cost < minimalCost {
			minimalCost = d.nodeMap[id].cost
			nodeId = id
		}
	}
	d.deleteCandidate(nodeId)
	return nodeId
}

func (d *Dijkstra) deleteCandidate(deletedNodeId int) {
	newCandidates := []int{}
	for _, id := range d.candidates {
		if id != deletedNodeId {
			newCandidates = append(newCandidates, id)
		}
	}
	d.candidates = newCandidates
}

func (d *Dijkstra) GetShortestPath() string {
	shortestPath := ""
	costOfShortestPath := 0
	step := 0
	for {
		step++
		// choose a node that have minimal cost within residual candidate nodes
		checkNodeId := d.PopMinimalCostNode()
		if checkNodeId == d.goalNodeId {
			costOfShortestPath = d.nodeMap[d.goalNodeId].cost
			break
		}

		for _, id := range d.candidates {
			if cost, err := d.graph.GetEdgeWeight(checkNodeId, id); cost >= 0 && err == nil {
				if d.nodeMap[checkNodeId].cost+cost < d.nodeMap[id].cost {
					d.nodeMap[id].cost = d.nodeMap[checkNodeId].cost + cost
					d.nodeMap[id].previousNodeId = checkNodeId
				}
			}
		}
	}

	order := []int{}
	prevNodeId := d.nodeMap[d.goalNodeId].previousNodeId
	order = append(order, d.goalNodeId)
	for {
		if prevNodeId == -1 {
			break
		}
		order = append(order, prevNodeId)
		prevNodeId = d.nodeMap[prevNodeId].previousNodeId
	}
	for i := len(order) - 1; i >= 0; i-- {
		shortestPath += fmt.Sprintf("%d, ", order[i])
	}

	return fmt.Sprintf("%s: cost: %d, step: %d", shortestPath, costOfShortestPath, step)
}

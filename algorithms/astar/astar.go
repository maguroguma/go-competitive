package astar

import (
	"fmt"
	"math"

	"github.com/myokoyama0712/algorithm/graph"
)

type Astar struct {
	startNodeId, goalNodeId int
	g                       *graph.Graph
	candidates              []int
	nodeMap                 NodeMap
}

type Node struct {
	realCost            int
	previousNodeId      int
	heulisticCostToGoal int
}
type NodeMap map[int]*Node

func NewAstar(g *graph.Graph, startNodeId, goalNodeId int, coordinateMap map[int][2]int) *Astar {
	astar := new(Astar)
	astar.startNodeId = startNodeId
	astar.goalNodeId = goalNodeId
	astar.g = g
	candidates := []int{}
	nodeMap := make(NodeMap)
	for i := 0; i < g.GetNodeNumber(); i++ {
		node := new(Node)
		if i == startNodeId {
			node.realCost = 0
		} else {
			node.realCost = 1000
		}
		node.previousNodeId = -1
		f1 := math.Abs(float64(coordinateMap[i][0] - coordinateMap[goalNodeId][0]))
		f2 := math.Abs(float64(coordinateMap[i][1] - coordinateMap[goalNodeId][1]))
		node.heulisticCostToGoal = int(f1 + f2)
		nodeMap[i] = node
		candidates = append(candidates, i)
	}
	astar.candidates = candidates
	astar.nodeMap = nodeMap

	return astar
}

// PopMinHeulisticTotalCostNode returns node that has minimal predicted cost
//  to the goal using heulistic cost within candidates nodes.
// And the returned node is deleted from candidates nodes.
func (astar *Astar) PopMinHeulisticTotalCostNode() int {
	minimalCost := 1000000
	nodeId := -1
	for _, id := range astar.candidates {
		predictedCost := astar.nodeMap[id].realCost + astar.nodeMap[id].heulisticCostToGoal
		if predictedCost < minimalCost {
			minimalCost = predictedCost
			nodeId = id
		}
	}
	astar.deleteCandidate(nodeId)
	return nodeId
}
func (astar *Astar) deleteCandidate(deletedNodeId int) {
	newCandidates := []int{}
	for _, id := range astar.candidates {
		if id != deletedNodeId {
			newCandidates = append(newCandidates, id)
		}
	}
	astar.candidates = newCandidates
}

func (astar *Astar) GetShortestPath() string {
	shortestPath := ""
	costOfShortestPath := 0
	step := 0
	for {
		step++
		// choose a node that have minimal **total** cost within residual candidate nodes
		checkNodeId := astar.PopMinHeulisticTotalCostNode()
		if checkNodeId == astar.goalNodeId {
			costOfShortestPath = astar.nodeMap[astar.goalNodeId].realCost
			break
		}

		for _, id := range astar.candidates {
			if edgeCost, err := astar.g.GetEdgeWeight(checkNodeId, id); edgeCost >= 0 && err == nil {
				checkingPathCost := astar.nodeMap[checkNodeId].realCost + edgeCost + astar.nodeMap[id].heulisticCostToGoal
				interimBestCost := astar.nodeMap[id].realCost + astar.nodeMap[id].heulisticCostToGoal
				if checkingPathCost < interimBestCost {
					astar.nodeMap[id].realCost = astar.nodeMap[checkNodeId].realCost + edgeCost
					astar.nodeMap[id].previousNodeId = checkNodeId
				}
			}
		}
	}

	order := []int{}
	prevNodeId := astar.nodeMap[astar.goalNodeId].previousNodeId
	order = append(order, astar.goalNodeId)
	for {
		if prevNodeId == -1 {
			break
		}
		order = append(order, prevNodeId)
		prevNodeId = astar.nodeMap[prevNodeId].previousNodeId
	}
	for i := len(order) - 1; i >= 0; i-- {
		shortestPath += fmt.Sprintf("%d, ", order[i])
	}

	return fmt.Sprintf("%s: cost: %d, step: %d", shortestPath, costOfShortestPath, step)
}

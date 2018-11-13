package tree

import (
	"fmt"
)

/*
Node is struct denoting tree node.
*/
type Node struct {
	id       int
	children []*Node
}

/*
NewNode returns pointer of Node struct.
*/
func NewNode(id int) *Node {
	node := new(Node)
	node.id = id
	node.children = []*Node{}
	return node
}

/*
ToString returns string formatted like
`
id: self, children: [child1, child2, ...]
	id: self, children: [child1, child2, ...]
		id: self, children: [child1, child2, ...]
		id: self, children: [child1, child2, ...]
	id: self, children: [child1, child2, ...]
		id: self, children: [child1, child2, ...]
	...
`
*/
func (n *Node) ToString() string {
	return fmt.Sprintf("id: %d, children: %v", n.id, n.children)
}

/*
AddChild append new child node to its childlen.
*/
func (n *Node) AddChild(id int) {
	newChild := new(Node)
	newChild.id = id
	newChild.children = []*Node{}

	n.children = append(n.children, newChild)
}

/*
FullSearchByBfs returns string denoting order of passed nodes by BFS
use queue structure
*/
func (n *Node) FullSearchByBfs() string {
	nodeQueue := make([]*Node, 0)
	nodeQueue = append(nodeQueue, n)
	order := ""

	for {
		// pop
		currentNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		// check
		order += fmt.Sprintf("%d ", currentNode.id)

		// push
		for _, child := range currentNode.children {
			nodeQueue = append(nodeQueue, child)
		}

		if len(nodeQueue) == 0 {
			break
		}
	}

	return order
}

/*
FullSearchByDfs returns string denoting order of passed nodes by DFS
(pre-order: use recursion)
*/
func (n *Node) FullSearchByDfs() string {
	order := ""
	recursion(n, &order)
	return order
}
func recursion(currentNode *Node, order *string) {
	*order += fmt.Sprintf("%d ", currentNode.id)

	for _, child := range currentNode.children {
		recursion(child, order)
	}
}

/*
FullSearchByDfsStackVersion is the version of DFS using stack
(another pre-order: use stack structure)
*/
func (n *Node) FullSearchByDfsStackVersion() string {
	order := ""
	nodeStack := make([]*Node, 0)
	nodeStack = append(nodeStack, n)

	for {
		currentNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		order += fmt.Sprintf("%d ", currentNode.id)

		for _, child := range currentNode.children {
			nodeStack = append(nodeStack, child)
		}

		if len(nodeStack) == 0 {
			break
		}
	}

	return order
}

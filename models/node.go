package models

type Node interface {
	GetID() int
}

type NodeWithScore struct {
	Node  Node
	Score float32
}

func NewNodeWithScore(node Node, score float32) *NodeWithScore {
	return &NodeWithScore{
		Node:  node,
		Score: score,
	}
}

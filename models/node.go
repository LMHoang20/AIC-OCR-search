package models

type Node struct {
	Frames   map[*Frame]bool
	Children map[rune]*Node
}

func NewNode() *Node {
	return &Node{
		Frames:   make(map[*Frame]bool),
		Children: make(map[rune]*Node),
	}
}

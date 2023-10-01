package models

type RAMNode struct {
	frames   map[string]int
	children map[rune]Node
}

func NewRAMNode() *RAMNode {
	return &RAMNode{
		frames:   map[string]int{},
		children: map[rune]Node{},
	}
}

func (n *RAMNode) GetFrames() *map[string]int {
	return &n.frames
}

func (n *RAMNode) GetChildren() *map[rune]Node {
	return &n.children
}

func (n *RAMNode) GetChild(character rune) Node {
	val, ok := n.children[character]
	if !ok {
		return nil
	}
	return val
}

func (n *RAMNode) AddChild(character rune, node Node) {
	n.children[character] = node
}

func (n *RAMNode) AddFrame(frame *RAMFrame) {
	n.frames[frame.String()] += 1
}

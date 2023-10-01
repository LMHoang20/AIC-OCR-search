package models

type RAMNode struct {
	frames    map[*Frame]bool
	childrens map[rune]*Node
}

func NewRAMNode() *RAMNode {
	return &RAMNode{
		frames:    map[*Frame]bool{},
		childrens: map[rune]*Node{},
	}
}

func (n *RAMNode) GetFrames() *map[*Frame]bool {
	return &n.frames
}

func (n *RAMNode) GetChildrens() *map[rune]*Node {
	return &n.childrens
}

func (n *RAMNode) GetChild(character rune) *Node {
	val, ok := n.childrens[character]
	if !ok {
		return nil
	}
	return val
}

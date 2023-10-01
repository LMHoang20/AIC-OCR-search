package models

type Node interface {
	GetFrames() *map[*Frame]bool
	GetChildrens() *map[rune]*Node
	GetChild(rune) *Node
}

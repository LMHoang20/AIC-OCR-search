package repositories

import (
	"OCRsearch/database"
	"OCRsearch/models"
)

type RAMGetter struct {
	db *database.RAM
}

func NewRAMGetter() *RAMGetter {
	return &RAMGetter{
		db: database.RAMInstance(),
	}
}

func (r *RAMGetter) GetRoot() models.Node {
	return r.db.GetRoot()
}

func (r *RAMGetter) GetFramesOfNode(node models.Node) map[string]int {
	return *node.(*models.RAMNode).GetFrames()
}

func (r *RAMGetter) GetChildrensOfNode(node models.Node) map[rune]models.Node {
	return *node.(*models.RAMNode).GetChildren()
}

func (r *RAMGetter) GetChildOfNode(node models.Node, character rune) models.Node {
	return node.(*models.RAMNode).GetChild(character)
}

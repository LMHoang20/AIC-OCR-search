package repositories

import "OCRsearch/models"

type Getter interface {
	GetRoot() models.Node
	GetFramesOfNode(node models.Node) map[string]int
	GetChildrensOfNode(node models.Node) map[rune]models.Node
	GetChildOfNode(node models.Node, character rune) models.Node
}

type Finder interface {
	Find(characters []rune, limit int) []models.NodeWithScore
}

type Interface interface {
	Getter
	Finder
}

func NewGetter(DBType string) Getter {
	switch DBType {
	case "RAM":
		return NewRAMGetter()
	default:
		return NewRAMGetter()
	}
}

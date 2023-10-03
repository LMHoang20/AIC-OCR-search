package repositories

import (
	"OCRsearch/models"
)

type Exact struct {
	Finder
	g Getter
}

func NewExact(DBType string) *Exact {
	return &Exact{
		g: NewGetter(DBType),
	}
}

func (r *Exact) Find(characters []rune, limit int) []models.NodeWithScore {
	result := make([]models.NodeWithScore, 0)

	currentNode := r.GetRoot()

	for _, character := range characters {
		currentNode = r.GetChildOfNode(currentNode, character)
		if currentNode == nil {
			return result
		}
	}

	result = append(result, *models.NewNodeWithScore(currentNode, 1))

	return result
}

func (r *Exact) GetRoot() models.Node {
	return r.g.GetRoot()
}

func (r *Exact) GetFramesOfNode(node models.Node) map[string]int {
	return r.g.GetFramesOfNode(node)
}

func (r *Exact) GetChildrensOfNode(node models.Node) map[rune]models.Node {
	return r.g.GetChildrensOfNode(node)
}

func (r *Exact) GetChildOfNode(node models.Node, character rune) models.Node {
	return r.g.GetChildOfNode(node, character)
}

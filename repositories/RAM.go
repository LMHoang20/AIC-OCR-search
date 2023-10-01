package repositories

import (
	"OCRsearch/database"
	"OCRsearch/models"
)

type RAM struct {
	db *database.RAM
}

var ramInstance *RAM

func RAMInstance() *RAM {
	if ramInstance == nil {
		ramInstance = &RAM{
			db: database.RAMInstance(),
		}
	}
	return ramInstance
}

func (r *RAM) FindExact(word string) (models.Node, error) {
	currentNode := r.db.GetRoot()

	for _, character := range word {
		currentNode = r.GetChildOfNode(currentNode, character)
		if currentNode == nil {
			return nil, nil
		}
	}

	return currentNode, nil
}

func (r *RAM) GetFramesOfNode(node models.Node) map[string]bool {
	return *node.(*models.RAMNode).GetFrames()
}

func (r *RAM) GetChildrensOfNode(node models.Node) map[rune]models.Node {
	return *node.(*models.RAMNode).GetChildren()
}

func (r *RAM) GetChildOfNode(node models.Node, character rune) models.Node {
	return node.(*models.RAMNode).GetChild(character)
}

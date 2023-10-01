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

func (r *RAM) FindExact(word string) (*models.Node, error) {
	currentNode := r.db.GetRoot()

	for _, character := range word {
		currentNode = (*currentNode).GetChild(character)
		if currentNode == nil {
			return nil, nil
		}
	}

	return currentNode, nil
}

package database

import "OCRsearch/models"

type RAM struct {
	root models.Node
}

var ramInstance *RAM

func RAMInstance() *RAM {
	if ramInstance == nil {
		ramInstance = &RAM{
			root: models.NewRAMNode(),
		}
	}
	return ramInstance
}

func (r *RAM) GetRoot() *models.Node {
	return &r.root
}

func (r *RAM) Initialize() error {
	return nil
}

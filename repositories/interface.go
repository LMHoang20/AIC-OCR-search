package repositories

import "OCRsearch/models"

type Interface interface {
	FindExact(word string) (models.Node, error)
	GetFramesOfNode(node models.Node) map[string]bool
	GetChildrensOfNode(node models.Node) map[rune]models.Node
	GetChildOfNode(node models.Node, character rune) models.Node
}

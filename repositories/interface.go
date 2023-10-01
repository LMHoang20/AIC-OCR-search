package repositories

import "OCRsearch/models"

type Interface interface {
	FindExact(word string) (*models.Node, error)
}

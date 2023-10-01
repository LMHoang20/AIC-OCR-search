package searchers

import "OCRsearch/models"

type Searcher interface {
	Search(query string) ([]models.Frame, error)
}

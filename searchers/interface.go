package searchers

import "OCRsearch/models"

type Interface interface {
	Search(query string, limit int) ([]models.Candidate, error)
}

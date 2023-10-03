package searchers

import "OCRsearch/models"

type Interface interface {
	Search(query string, limit int) ([]models.Candidate, error)
}

func NewSearcher(searcherType string, dbType string) Interface {
	switch searcherType {
	case "exact":
		return NewExact(dbType)
	case "fuzzy":
		return NewFuzzy(dbType)
	default:
		return NewExact(dbType)
	}
}

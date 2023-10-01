package searchers

import (
	"OCRsearch/models"
	"OCRsearch/repositories"
	"sort"
	"strings"
)

type Exact struct {
	r repositories.Interface
}

func NewExact(repoType string) Interface {
	switch repoType {
	default:
		return &Exact{r: repositories.RAMInstance()}
	}
}

func (e *Exact) Search(query string, limit int) ([]models.Candidate, error) {
	words := strings.Split(query, " ")

	var cnt map[*models.Frame]int

	for _, word := range words {
		node, err := e.r.FindExact(word)
		if err != nil {
			return nil, err
		} else if node == nil {
			return make([]models.Candidate, 0), nil
		}
		for frame := range *(*node).GetFrames() {
			cnt[frame]++
		}
	}

	var results []models.Candidate
	for frame, score := range cnt {
		results = append(results, *models.NewCandidate(frame, score))
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results[:limit], nil
}

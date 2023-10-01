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

	cnt := make(map[string]int)

	for _, word := range words {
		node, err := e.r.FindExact(word)
		if err != nil {
			return nil, err
		} else if node == nil {
			return make([]models.Candidate, 0), nil
		}
		for frame := range e.r.GetFramesOfNode(node) {
			cnt[frame]++
		}
	}

	results := make([]models.Candidate, 0)
	for frame, score := range cnt {
		results = append(results, *models.NewCandidate(models.NewRAMFrameFromString(frame), score))
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if len(results) < limit {
		return results, nil
	}
	return results[:limit], nil
}

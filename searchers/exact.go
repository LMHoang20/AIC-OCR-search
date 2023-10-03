package searchers

import (
	"OCRsearch/helpers"
	"OCRsearch/models"
	"OCRsearch/repositories"
	"sort"
	"strings"
)

type Exact struct {
	r repositories.Interface
}

func NewExact(DBType string) Interface {
	return &Exact{r: repositories.NewExact(DBType)}
}

func (e *Exact) Search(query string, limit int) ([]models.Candidate, error) {
	words := strings.Split(query, " ")

	scores := make(map[string]float32)

	for _, word := range words {
		characters := helpers.GetCharacters(word)
		nodes := e.r.Find(characters, 1)
		for _, nodeWithScore := range nodes {
			for frame, occurences := range e.r.GetFramesOfNode(nodeWithScore.Node) {
				scores[frame] += float32(occurences) * nodeWithScore.Score
			}
		}
	}

	results := make([]models.Candidate, 0)
	for frame, score := range scores {
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

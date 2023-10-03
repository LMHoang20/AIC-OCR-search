package searchers

import (
	"OCRsearch/helpers"
	"OCRsearch/models"
	"OCRsearch/repositories"
	"sort"
	"strings"
)

type Fuzzy struct {
	r repositories.Interface
}

func NewFuzzy(repoType string) *Fuzzy {
	return &Fuzzy{
		r: repositories.NewFuzzy(repoType),
	}
}

func (f *Fuzzy) Search(query string, limit int) ([]models.Candidate, error) {
	words := strings.Fields(query)

	scores := make(map[string]float32)

	for _, word := range words {
		characters := helpers.GetCharacters(word)
		nodes := f.r.Find(characters, 1)
		bestScoreOfFrame := make(map[string]float32)
		for _, nodeWithScore := range nodes {
			for frame, occurences := range f.r.GetFramesOfNode(nodeWithScore.Node) {
				score := float32(occurences) * nodeWithScore.Score
				val, ok := bestScoreOfFrame[frame]
				if !ok || val < score {
					bestScoreOfFrame[frame] = score
				}
			}
		}
		for frame, score := range bestScoreOfFrame {
			scores[frame] += score
		}
	}

	results := make([]models.Candidate, 0)
	for frame, score := range scores {
		results = append(results, *models.NewCandidate(models.NewRAMFrameFromString(frame), float32(score)))
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if len(results) < limit {
		return results, nil
	}

	return results[:limit], nil
}

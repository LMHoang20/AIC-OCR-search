package searchers

import (
	"OCRsearch/helpers"
	"OCRsearch/models"
	"OCRsearch/repositories"
	"fmt"
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
	words := strings.Split(query, " ")

	scores := make(map[string]float32)

	for _, word := range words {
		characters := helpers.GetCharacters(word)
		nodes := f.r.Find(characters, 1)
		for _, nodeWithScore := range nodes {
			for frame, occurences := range f.r.GetFramesOfNode(nodeWithScore.Node) {
				fmt.Println(frame, occurences, nodeWithScore)
				scores[frame] += float32(occurences) * nodeWithScore.Score
			}
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

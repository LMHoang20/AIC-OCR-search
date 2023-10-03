package repositories

import (
	"OCRsearch/constants"
	"OCRsearch/models"
	"container/heap"
	"fmt"
)

type Fuzzy struct {
	Finder
	g Getter
}

func NewFuzzy(DBType string) *Fuzzy {
	return &Fuzzy{
		g: NewGetter(DBType),
	}
}

func calculateScore(len int, penalty float32) float32 {
	return 1 - penalty/float32(len)
}

func stateString(node models.Node, match int) string {
	return fmt.Sprintf("%d %d", node.GetID(), match)
}

func pushMemoize(pq *models.PriorityQueue, dp *map[string]bool, node models.Node, penalty float32, match int) {
	if !(*dp)[stateString(node, match)] {
		(*dp)[stateString(node, match)] = true
		heap.Push(pq, models.NewItem(node, []float32{penalty, -float32(match)}))
	}
}

func (r *Fuzzy) Find(characters []rune, limit int) []models.NodeWithScore {
	result := make([]models.NodeWithScore, 0)

	n := len(characters)
	maxPenalty := float32(n) * constants.FuzzyTolerate

	pq := models.NewPriorityQueue()
	pq.Push(models.NewItem(r.GetRoot(), []float32{0, 0}))

	visited := map[string]bool{}
	visited[stateString(r.GetRoot(), 0)] = true

	added := map[int]bool{}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*models.Item)
		node := item.Value.(models.Node)
		penalty := item.Priority[0]
		match := -int(item.Priority[1])
		if match == n {
			if !added[node.GetID()] {
				added[node.GetID()] = true
				result = append(result, *models.NewNodeWithScore(node, calculateScore(n, penalty)))
			}
			continue
		}
		for character, child := range r.g.GetChildrensOfNode(node) {
			if character == rune(characters[match]) {
				// match
				pushMemoize(&pq, &visited, child, penalty, match+1)
			} else if penalty+1 < maxPenalty {
				// replace
				pushMemoize(&pq, &visited, child, penalty+1, match+1)
				// insert
				pushMemoize(&pq, &visited, child, penalty+1, match)
				// delete
				pushMemoize(&pq, &visited, node, penalty+1, match+1)
			}
		}
	}

	return result
}

func (r *Fuzzy) GetRoot() models.Node {
	return r.g.GetRoot()
}

func (r *Fuzzy) GetFramesOfNode(node models.Node) map[string]int {
	return r.g.GetFramesOfNode(node)
}

func (r *Fuzzy) GetChildrensOfNode(node models.Node) map[rune]models.Node {
	return r.g.GetChildrensOfNode(node)
}

func (r *Fuzzy) GetChildOfNode(node models.Node, character rune) models.Node {
	return r.g.GetChildOfNode(node, character)
}

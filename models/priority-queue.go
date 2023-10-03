package models

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value    interface{}
	Priority []float32
	Index    int
}

type PriorityQueue []*Item

func NewItem(Value interface{}, Priority []float32) *Item {
	return &Item{
		Value:    Value,
		Priority: Priority,
	}
}

func NewPriorityQueue() PriorityQueue {
	return make(PriorityQueue, 0)
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// need a little fix but it's ok for now
	for k := 0; k < len(pq[i].Priority); k++ {
		if pq[i].Priority[k] != pq[j].Priority[k] {
			return pq[i].Priority[k] < pq[j].Priority[k]
		}
	}
	return false
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Clear() {
	*pq = make(PriorityQueue, 0)
}

func (pq *PriorityQueue) update(item *Item, Value interface{}, Priority []float32) {
	item.Value = Value
	item.Priority = Priority
	heap.Fix(pq, item.Index)
}

func (pq *PriorityQueue) String() string {
	result := ""
	for _, item := range *pq {
		result += fmt.Sprintf("%v ", item.Priority)
	}
	return result
}

func test() {
	// Some items and their priorities.
	items := map[string][]float32{
		"banana": {3, 4}, "apple": {2, 2}, "pear": {2, 2},
	}

	// Create a Priority queue, put the items in it, and
	// establish the Priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for Value, Priority := range items {
		pq[i] = &Item{
			Value:    Value,
			Priority: Priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its Priority.
	item := &Item{
		Value:    "orange",
		Priority: []float32{2},
	}
	heap.Push(&pq, item)
	pq.update(item, item.Value, []float32{1})

	// Take the items out; they arrive in decreasing Priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2f:%s ", item.Priority, item.Value)
	}
}

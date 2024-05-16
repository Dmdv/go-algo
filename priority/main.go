package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value    string
	Priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) UpdatePriority(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	pq := make(PriorityQueue, 0)

	for i := 0; i < 10; i++ {
		item := &Item{
			Value:    "item",
			Priority: i,
			index:    i,
		}
		pq = append(pq, item)
	}

	heap.Init(&pq)

	item := &Item{
		Value:    "item",
		Priority: 100,
	}

	heap.Push(&pq, item)
	pq.UpdatePriority(item, "item", 1000)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Value: %s, Priority: %d\n", item.Value, item.Priority)
	}
}

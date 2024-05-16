# Priority Heap

```go
package main

import (
    "fmt"
    "github.com/dmdv/ds/heap"
)

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

```
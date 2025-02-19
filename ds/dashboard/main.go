package main

import (
	"container/heap"
	"sort"
)

type CurrencyCount struct {
	Currency string
	Count    int
}

type CurrencyCounter struct {
	counts map[string]int  // Global count tracker
	heap   []string        // Min-heap of currencies (sorted by count)
	inHeap map[string]bool // Tracks currencies in the heap
}

func NewCurrencyCounter() *CurrencyCounter {
	return &CurrencyCounter{
		counts: make(map[string]int),
		heap:   make([]string, 0, 10), // Preallocate heap capacity
		inHeap: make(map[string]bool),
	}
}

// Heap interface implementation (min-heap by count)

func (cc *CurrencyCounter) Len() int           { return len(cc.heap) }
func (cc *CurrencyCounter) Less(i, j int) bool { return cc.counts[cc.heap[i]] < cc.counts[cc.heap[j]] }
func (cc *CurrencyCounter) Swap(i, j int)      { cc.heap[i], cc.heap[j] = cc.heap[j], cc.heap[i] }
func (cc *CurrencyCounter) Push(x interface{}) {
	currency := x.(string)
	cc.heap = append(cc.heap, currency)
	cc.inHeap[currency] = true
}
func (cc *CurrencyCounter) Pop() interface{} {
	currency := cc.heap[len(cc.heap)-1]
	cc.heap = cc.heap[:len(cc.heap)-1]
	delete(cc.inHeap, currency)
	return currency
}

func (cc *CurrencyCounter) RecordRequest(currency string) {
	cc.counts[currency]++ // Update global count

	if cc.inHeap[currency] {
		// Rebalance, if currency is already in the heap
		for i, c := range cc.heap {
			if c == currency {
				heap.Fix(cc, i)
				break
			}
		}
	} else if len(cc.heap) < 10 || cc.counts[currency] > cc.counts[cc.heap[0]] {
		// Add new currency if it qualifies for the top 10
		if len(cc.heap) >= 10 {
			heap.Pop(cc) // Remove the smallest (root)
		}
		heap.Push(cc, currency)
	}
}

func (cc *CurrencyCounter) GetTop10() []CurrencyCount {
	// Convert heap currencies to CurrencyCount structs
	result := make([]CurrencyCount, len(cc.heap))
	for i, currency := range cc.heap {
		result[i] = CurrencyCount{
			Currency: currency,
			Count:    cc.counts[currency],
		}
	}

	// Sort by count descending
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	return result
}

func main() {
	counter := NewCurrencyCounter()
	counter.RecordRequest("USD")
	counter.RecordRequest("USD")
	counter.RecordRequest("EUR")
	top10 := counter.GetTop10() // Returns [{USD 2}, {EUR 1}]
	_ = top10
}

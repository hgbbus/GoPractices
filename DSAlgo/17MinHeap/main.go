package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

// First implement the sort.Interface methods as defined in heap.Interface for IntHeap
func (h IntHeap) Len() int		   		{ return len(h) }
func (h IntHeap) Less(i, j int) bool	{ return h[i] < h[j] }	// strictly less
func (h IntHeap) Swap(i, j int)			{ h[i], h[j] = h[j], h[i] }

// Now implement the last two methods of heap.Interface for IntHeap
func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))	// type assertion to int
}
func (h *IntHeap) Pop() any {
	old := *h		// old slice
	n := len(old)
	x := old[n-1]	// last element
	*h = old[:n-1]	// shrink slice
	return x
}

// Now, let's implement a Priority Queue in general.

// An item is something we manage in a priority queue.
type Item[T any] struct {
	Value	 T      // The actual value of the item; arbitrary.
	Priority int    // The priority of the item in the queue (smaller value means higher priority).
	index    int	// The index (internal tracker) of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }
func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest priority (lowest value) so we use less than here.
	return pq[i].Priority < pq[j].Priority
}
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i		// update index of item at i
	pq[j].index = j		// update index of item at j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])	// type assertion to *Item[T]
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1] // last item
	old[n-1] = nil   // avoid memory leak
	item.index = -1  // for safety
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue[T]) update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)	// re-heapify after updating the priority
}

func main() {
	h := IntHeap{6, 18, 29, 20, 28, 39, 66, 37, 26, 76, 32, 74}
	heap.Init(&h)
	printIntHeap(h)

	heap.Push(&h, 89)	// don't call h.Push(89) directly
	printIntHeap(h)

	heap.Push(&h, 8)
	printIntHeap(h)

	heap.Pop(&h)	// don't call h.Pop() directly
	printIntHeap(h)

	heap.Push(&h, 6)
	printIntHeap(h)
	
	items := map[string]int{
		"John Smith": 10, "Jane Fonda": 51, "Vincent Foxx": 2,
		"Johnny Cash": 18, "Jennifer Lopez": 4, "Jon Ossoff": 31,
		"Donald Trump": 13, "Ryan Paul": 5, "Nancy Pelosi": 23,
		"Bill Me": 64, "Joy Hash": 29,
	}

	pq := make(PriorityQueue[string], len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item[string]{
			Value:    value,
			Priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	printPriorityQueue(pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[string])
		fmt.Printf("%d: %s\n", item.Priority, item.Value)
	}
	fmt.Println()
}

func printIntHeap(h IntHeap) {
	for _, v := range h {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func printPriorityQueue[T any](pq PriorityQueue[T]) {
	for _, item := range pq {
		fmt.Printf("%d:%v ", item.Priority, item.Value)
	}
	fmt.Println()
}
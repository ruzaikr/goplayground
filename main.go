// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"sort"

	"github.com/davecgh/go-spew/spew"
)

// An Item is something we manage in a priority queue.
type Item struct {
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use lesser than here.
	return pq[i].priority < pq[j].priority
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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	} else if len(intervals) == 1 {
		return 1
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return false
	})

	var pq = make(PriorityQueue, 1)
	pq[0] = &Item{
		priority: intervals[0][1],
		index:    0,
	}

	heap.Init(&pq)

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= pq[0].priority {
			pq.Pop()
		}
		var item = &Item{
			priority: intervals[i][1],
			index:    0,
		}
		pq.Push(item)

		spew.Dump(i, pq)
	}

	return len(pq)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {

}

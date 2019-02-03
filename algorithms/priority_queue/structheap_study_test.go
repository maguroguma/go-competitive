package priority_queue

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://golang.org/pkg/container/heap/
// https://golang.org/src/container/heap/heap.go

type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
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
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// `update` method modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

/******************************* test cases ********************************/

func Test構造体のヒープ(t *testing.T) {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))

	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	actual := []*Item{}
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		actual = append(actual, item)
	}

	expected := []*Item{
		&Item{value: "apple", priority: 2, index: -1},
		&Item{value: "banana", priority: 3, index: -1},
		&Item{value: "pear", priority: 4, index: -1},
		&Item{value: "orange", priority: 5, index: -1},
	}
	assert.Equal(t, expected, actual)
}

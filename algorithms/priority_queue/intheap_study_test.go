package priority_queue

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://golang.org/pkg/container/heap/
// https://golang.org/src/container/heap/heap.go

type IntHeap []int

// 5つのメソッドを実装する
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's LENGTH, NOT JUST ITS CONTENTS.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/******************************* test cases ********************************/

func Test標準パッケージのheapのテスト(t *testing.T) {
	h := &IntHeap{3, 6, 1, 2}

	// IntHeap型はheap.Interfaceを実装しているため、Init関数に渡せる
	heap.Init(h)
	// Pushはheapの関数にInterfaceを実装した型と追加したいオブジェクトを渡すことで実行する
	heap.Push(h, 3)

	actual := []int{}
	for h.Len() > 0 {
		// Popはheapの関数にInterfaceを実装した型を渡すことで実行する
		actual = append(actual, heap.Pop(h).(int))
	}

	assert.Equal(t, []int{1, 2, 3, 3, 6}, actual)
}

func TestPopせずに中身をチェック(t *testing.T) {
	h := &IntHeap{3, 6, 1, 2}
	// IntHeap型はheap.Interfaceを実装しているため、Init関数に渡せる
	heap.Init(h)

	actual := []int{}
	for i := 0; i < h.Len(); i++ {
		actual = append(actual, (*h)[i])
	}
	assert.Equal(t, []int{1, 2, 3, 6}, actual)
}

func Test空の状態からヒープを作成(t *testing.T) {
	h := &IntHeap{}
	heap.Push(h, 3)
	heap.Push(h, 6)
	heap.Push(h, 1)
	heap.Push(h, 2)

	actual := []int{}
	for i := 0; i < h.Len(); i++ {
		actual = append(actual, (*h)[i])
	}
	assert.Equal(t, []int{1, 2, 3, 6}, actual)
}

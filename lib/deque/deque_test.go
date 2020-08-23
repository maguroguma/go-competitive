package deque

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	pushFronts, pushBacks   []int
	popFrontNum, popBackNum int
	popFronts, popBacks     []int
	queue                   []int
}

func TestLen(t *testing.T) {
	d := NewDeque()
	l := d.Len()
	if l != 0 {
		t.Errorf("got %v, want %v", l, 0)
	}
}

func TestDeque(t *testing.T) {
	testcases := []TestCase{
		{
			pushFronts:  []int{1, 2, 3},
			pushBacks:   []int{-1, -2, -3},
			popFrontNum: 2, popBackNum: 2,
			popFronts: []int{3, 2}, popBacks: []int{-3, -2},
			queue: []int{-1, 1},
		},
		{
			pushFronts:  []int{1, 2, 3},
			pushBacks:   []int{},
			popFrontNum: 1, popBackNum: 1,
			popFronts: []int{3}, popBacks: []int{1},
			queue: []int{2},
		},
		{
			pushFronts:  []int{},
			pushBacks:   []int{-1, -2, -3},
			popFrontNum: 1, popBackNum: 1,
			popFronts: []int{-1}, popBacks: []int{-3},
			queue: []int{-2},
		},
	}

	for i, tc := range testcases {
		subTest := fmt.Sprintf("Deque test %d", i)
		t.Run(subTest, func(t *testing.T) {
			d := new(Deque)

			for _, e := range tc.pushFronts {
				d.PushFront(e)
			}
			for _, e := range tc.pushBacks {
				d.PushBack(e)
			}

			popF, popB := []int{}, []int{}
			for i := 0; i < tc.popFrontNum; i++ {
				popF = append(popF, d.PopFront())
			}
			for i := 0; i < tc.popBackNum; i++ {
				popB = append(popB, d.PopBack())
			}

			queue := d.List()

			if !reflect.DeepEqual(popF, tc.popFronts) {
				t.Errorf("popFront is wrong, got %v, want %v", popF, tc.popFronts)
			}
			if !reflect.DeepEqual(popB, tc.popBacks) {
				t.Errorf("popBack is wrong, got %v, want %v", popB, tc.popBacks)
			}
			if !reflect.DeepEqual(queue, tc.queue) {
				t.Errorf("queue is wrong, got %v, want %v", queue, tc.queue)
			}
		})
	}
}

package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueの基本操作(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	assert.Equal(t, 1, q.GetSize())
	assert.Equal(t, 1, q.Dequeue())
}

func Test異なるデータ型を格納するQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(3.0)
	q.Enqueue("string")
	assert.Equal(t, 3, q.GetSize())
	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 3.0, q.Dequeue())
	assert.Equal(t, "string", q.Dequeue())
}

package queue

// Queue data structure
type Queue struct {
	data []interface{}
	size int
}

// NewQueue returns an empty queue.
func NewQueue() *Queue {
	q := new(Queue)
	q.data = []interface{}{}
	q.size = 0
	return q
}

// Enqueue accept an object in the last.
func (q *Queue) Enqueue(obj interface{}) {
	q.data = append(q.data, obj)
	q.size++
}

// Dequeue pop an object from the head.
func (q *Queue) Dequeue() interface{} {
	obj := q.data[0]
	q.data = q.data[1:]
	q.size--
	return obj
}

// GetSize returns the size of itself.
func (q *Queue) GetSize() int {
	return q.size
}

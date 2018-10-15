package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test整数型のキューの基本操作(t *testing.T) {
	q := NewQueue([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	assert.Equal(t, 10, q.Length())

}

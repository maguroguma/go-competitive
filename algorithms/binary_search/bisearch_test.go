package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test異なる要素からなるスライスの2分探索(t *testing.T) {
	test := []int{1, 3, 5, 11, 12, 13, 17, 22, 25, 28}
	assert.Equal(t, 8, BinarySearch(test, 25))
	assert.Equal(t, -1, BinarySearch(test, 7))
}

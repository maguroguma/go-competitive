package full_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test部分和問題(t *testing.T) {
	assert.Equal(t, "Yes", partialSum(4, []int{1, 2, 4, 7}, 13))
	assert.Equal(t, "No", partialSum(4, []int{1, 2, 4, 7}, 15))
}

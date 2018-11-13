package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowerBound(t *testing.T) {
	test := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}

	assert.Equal(t, 3, LowerBound(test, 51))
	assert.Equal(t, 0, LowerBound(test, 1))
	assert.Equal(t, 9, LowerBound(test, 910))

	assert.Equal(t, 6, LowerBound(test, 52))
	assert.Equal(t, 0, LowerBound(test, 0))
	assert.Equal(t, 10, LowerBound(test, 911))
}

func TestUpperBound(t *testing.T) {
	test := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}

	assert.Equal(t, 5, UpperBound(test, 51))
	assert.Equal(t, 0, UpperBound(test, 1))
	assert.Equal(t, 9, UpperBound(test, 910))

	assert.Equal(t, 5, UpperBound(test, 52))
	assert.Equal(t, -1, UpperBound(test, 0))
	assert.Equal(t, 9, UpperBound(test, 911))
}

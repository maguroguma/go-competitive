package cumulative_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2次元累積和(t *testing.T) {
	sampleRectangle := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	recSum := NewRectangleSum(sampleRectangle)
	assert.Equal(t, 34, recSum.GetSum(1, 1, 2, 2))
	assert.Equal(t, 38, recSum.GetSum(1, 2, 2, 3))
	assert.Equal(t, 6, recSum.GetSum(1, 1, 1, 1))
	assert.Equal(t, 10, recSum.GetSum(2, 1, 2, 1))
	assert.Equal(t, 9, recSum.GetSum(2, 0, 2, 0))
}

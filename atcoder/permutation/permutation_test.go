package permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test順列(t *testing.T) {
	expected := [][]int{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{2, 1, 3},
		[]int{2, 3, 1},
		[]int{3, 1, 2},
		[]int{3, 2, 1},
	}
	actual := GeneratePermutation(3)
	assert.Equal(t, expected, actual)
}

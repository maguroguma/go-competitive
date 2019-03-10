package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test10進数の桁和(t *testing.T) {
	assert.Equal(t, 45, GetDigitSum(123456789))
	assert.Equal(t, 45, GetDigitSum(102030405060708090))
	assert.Equal(t, -1, GetDigitSum(-123))
	assert.Equal(t, 45, GetDigitSum(908070605040302010))
}

func Test整数スライスの総和(t *testing.T) {
	assert.Equal(t, 45, Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	assert.Equal(t, 45, Sum(0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7, 0, 8, 0, 9))
	assert.Equal(t, 45, Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
}

func Test整数スライスの累積和スライス(t *testing.T) {
	assert.Equal(t, []int{1, 3, 6, 10, 15, 21, 28, 36, 45},
		GetCumulativeSums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

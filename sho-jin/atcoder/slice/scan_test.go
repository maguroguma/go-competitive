package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test両端から中心に向かってのスキャン(t *testing.T) {
	zero := []int{}
	one := []int{0}
	two := []int{0, 1}
	three := []int{0, 2, 1}
	ten := []int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}
	eleven := []int{0, 10, 1, 9, 2, 8, 3, 7, 4, 6, 5}
	assert.Equal(t, zero, PincerScan(0))
	assert.Equal(t, one, PincerScan(1))
	assert.Equal(t, two, PincerScan(2))
	assert.Equal(t, three, PincerScan(3))
	assert.Equal(t, ten, PincerScan(10))
	assert.Equal(t, eleven, PincerScan(11))
}

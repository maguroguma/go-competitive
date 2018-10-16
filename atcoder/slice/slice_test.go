package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test誤ったデリート(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := append(a[:3], a[4:]...)
	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 7, 8, 9}, b)
	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 7, 8, 9, 9}, a)
}

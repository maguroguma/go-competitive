package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test任意個の整数引数のうち最大のものを取得(t *testing.T) {
	assert.Equal(t, 100, Max(0, 100, -100, 1))
	assert.Equal(t, 1, Max(1))
	assert.Equal(t, 0, Max(0))
	assert.Equal(t, 100, Max([]int{100, -100}...))
}

func Test任意個の整数引数のうち最小のものを取得(t *testing.T) {
	assert.Equal(t, -100, Min(0, 100, -100, 1))
	assert.Equal(t, 1, Min(1))
	assert.Equal(t, 0, Min(0))
	assert.Equal(t, -100, Min([]int{100, -100}...))
}

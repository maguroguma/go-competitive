package pascal_triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testパスカルの三角形による二項係数の計算(t *testing.T) {
	assert.Equal(t, 1, Combination(0, 0)) // (0, 0)は1
	assert.Equal(t, 10, Combination(5, 3))
	assert.Equal(t, 126, Combination(9, 4))
}

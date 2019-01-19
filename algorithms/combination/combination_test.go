package combination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testオーソドックスな組み合わせ計算(t *testing.T) {
	assert.Equal(t, 1, CalcComb(0, 0)) // (0, 0)は1
	assert.Equal(t, 10, CalcComb(5, 3))
	assert.Equal(t, 126, CalcComb(9, 4))
}

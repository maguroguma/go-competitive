package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test整数のべき乗(t *testing.T) {
	assert.Equal(t, 4, PowInt(2, 2))
	assert.Equal(t, 1, PowInt(10000, 0))
	assert.Equal(t, 1024, PowInt(2, 10))
}

func Test整数の絶対値(t *testing.T) {
	assert.Equal(t, 0, AbsInt(0))
	assert.Equal(t, 0, AbsInt(-0))
	assert.Equal(t, 1, AbsInt(1))
	assert.Equal(t, 1, AbsInt(-1))
	assert.Equal(t, 1000, AbsInt(1000))
	assert.Equal(t, 1000, AbsInt(-1000))
}

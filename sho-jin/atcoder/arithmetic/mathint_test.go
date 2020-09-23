package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test整数のべき乗(t *testing.T) {
	assert.Equal(t, 4, PowInt(2, 2))
	assert.Equal(t, 1, PowInt(10000, 0))
	assert.Equal(t, 1024, PowInt(2, 10))
	assert.Equal(t, 1073741824, PowInt(2, 30))
	assert.Equal(t, 25937424601, PowInt(11, 10))
}

func Test整数の絶対値(t *testing.T) {
	assert.Equal(t, 0, AbsInt(0))
	assert.Equal(t, 0, AbsInt(-0))
	assert.Equal(t, 1, AbsInt(1))
	assert.Equal(t, 1, AbsInt(-1))
	assert.Equal(t, 1000, AbsInt(1000))
	assert.Equal(t, 1000, AbsInt(-1000))
}

func Test整数で閉じた天井関数(t *testing.T) {
	assert.Equal(t, 1, CeilInt(2, 3))
	assert.Equal(t, 0, CeilInt(0, 3))
	assert.Equal(t, 1, CeilInt(1, 1000000))
	assert.Equal(t, 100, CeilInt(200, 2))
	assert.Equal(t, 5, CeilInt(14, 3))
}

func Test整数で閉じた床関数(t *testing.T) {
	assert.Equal(t, 0, FloorInt(2, 3))
	assert.Equal(t, 0, FloorInt(0, 3))
	assert.Equal(t, 0, FloorInt(1, 1000000))
	assert.Equal(t, 100, FloorInt(200, 2))
	assert.Equal(t, 4, FloorInt(14, 3))
}

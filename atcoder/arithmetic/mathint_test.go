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

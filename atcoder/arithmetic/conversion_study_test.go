package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test浮動小数点数と整数の変換(t *testing.T) {
	a := 1.99
	b := 1.01
	// intでのキャストにより小数点以下切り捨てになる
	assert.Equal(t, 1, int(a))
	assert.Equal(t, 1, int(b))
}

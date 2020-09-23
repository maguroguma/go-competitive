package strprocess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test文字列から整数への変換(t *testing.T) {
	assert.Equal(t, 0, Strtoi("0"))
	assert.Equal(t, 1, Strtoi("1"))
	assert.Equal(t, -1, Strtoi("-1"))
	assert.Equal(t, -11, Strtoi("-11"))
}

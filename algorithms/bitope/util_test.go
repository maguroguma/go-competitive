package bitope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0埋め済みのruneスライス取得(t *testing.T) {
	// 45 = 0b101101
	// 25 = 0b011001
	//  9 = 0b001001

	a := 25
	S := GetZeroPaddingRuneSlice(a, 10)
	assert.Equal(t, S, []rune("0000011001"))
	assert.Equal(t, GetZeroPaddingRuneSlice(45, 10), []rune("0000101101"))
	assert.Equal(t, GetZeroPaddingRuneSlice(9, 10), []rune("0000001001"))

	// 桁が足りないケース
	assert.Equal(t, GetZeroPaddingRuneSlice(25, 5), []rune("11001"))
	assert.Equal(t, GetZeroPaddingRuneSlice(45, 4), []rune("101101"))
	assert.Equal(t, GetZeroPaddingRuneSlice(9, -1), []rune("1001"))
}

package xor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test単位元(t *testing.T) {
	// 同じモノ同士のXORは0となる
	// 各ビットが打ち消し合うので0となる
	assert.Equal(t, 0, 1^1)
	assert.Equal(t, 0, 10^10)
	assert.Equal(t, 0, 100^100)
	assert.Equal(t, 0, 999^999)
	assert.Equal(t, 0, 11111^11111)
}

func Test偶数とそれより1大きい奇数のXORは1(t *testing.T) {
	// ビットの差分が最下位桁の0->1のみなので成り立つ
	assert.Equal(t, 1, 2^3)
	assert.Equal(t, 1, 100^101)
	assert.Equal(t, 1, 10000^10001)
	assert.Equal(t, 1, 123456^123457)
	assert.Equal(t, 1, 9876^9877)
}

package modpow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test負数に対する剰余(t *testing.T) {
	actual := -17 % 5
	assert.Equal(t, -2, actual) // 期待される値とは異なる
}

func Test負数に対する剰余を計算する独自関数(t *testing.T) {
	actual := CalcNegativeMod(-17, 5)
	assert.Equal(t, actual, 3) // 期待される値

	const mod = 1000000000 + 7
	a, b := 2000000020, 20
	actual = (a - b) % mod
	assert.Equal(t, actual, 999999993)
	actual = (a%mod - b%mod) % mod
	assert.Equal(t, actual, -14)
	actual = CalcNegativeMod(2000000020%mod-20%mod, mod)
	assert.Equal(t, actual, 999999993)
}

func Testフェルマーの小定理に基づいた法13の逆元の計算(t *testing.T) {
	assert.Equal(t, CalcModInv(1, 13), 1)
	assert.Equal(t, CalcModInv(2, 13), 7)
	assert.Equal(t, CalcModInv(3, 13), 9)
	assert.Equal(t, CalcModInv(4, 13), 10)
	assert.Equal(t, CalcModInv(5, 13), 8)
	assert.Equal(t, CalcModInv(6, 13), 11)
	assert.Equal(t, CalcModInv(7, 13), 2)
	assert.Equal(t, CalcModInv(8, 13), 5)
	assert.Equal(t, CalcModInv(9, 13), 3)
	assert.Equal(t, CalcModInv(10, 13), 4)
	assert.Equal(t, CalcModInv(11, 13), 6)
	assert.Equal(t, CalcModInv(12, 13), 12)
}

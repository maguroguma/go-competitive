package arithmetic

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
基本的に入力・出力ともに浮動小数点数のみ許容する
*/

func Test最大値最小値を取得する(t *testing.T) {
	// 浮動小数点数
	assert.Equal(t, 1.0, math.Max(1.0, -1.0))
	assert.Equal(t, -1.0, math.Min(1.0, -1.0))
	// 整数（返り値は浮動小数点数）
	assert.Equal(t, 1.0, math.Max(1, -1))
	assert.Equal(t, -1.0, math.Min(1, -1))
}

func Test天井関数と床関数(t *testing.T) {
	// 天井関数
	assert.Equal(t, 100.0, math.Ceil(99.01))
	assert.Equal(t, 99.0, math.Ceil(99.00))
	// 床関数
	assert.Equal(t, 99.0, math.Floor(99.99))
	assert.Equal(t, 99.0, math.Floor(99.00))
	// 整数に変換するならキャストでもよい
	a := 99.01
	assert.Equal(t, 99, int(a))
}

func Test絶対値(t *testing.T) {
	// 浮動小数点数
	assert.Equal(t, 1.0, math.Abs(-1.0))
	// 整数（返り値は浮動小数点数）
	assert.Equal(t, 1.0, math.Abs(-1))
}

func Testべき乗(t *testing.T) {
	// 浮動小数点数
	assert.Equal(t, 8.0, math.Pow(2.0, 3.0))
	// 整数（返り値は浮動小数点数）
	assert.Equal(t, 8.0, math.Pow(2, 3))
}

package random

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test整数の乱数生成(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	max := 100

	ok := true
	observedMax := -1
	for i := 0; i < 100000; i++ {
		// 標準関数は[0, n)の整数を返す
		r := rand.Intn(max + 1)
		// 範囲チェック
		if !(0 <= r && r <= max) {
			ok = false
		}
		// 最大値チェック
		if observedMax < r {
			observedMax = r
		}
	}
	assert.True(t, ok)
	assert.Equal(t, max, observedMax)
}

func Test範囲を指定した整数の乱数生成(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	max := 6
	min := 1

	ok := true
	observedMax := -1
	observedMin := 1000000
	for i := 0; i < 100000; i++ {
		br := BoundedInt(min, max)
		// 範囲チェック
		if !(min <= br && br <= max) {
			ok = false
		}
		// 最大値・最小値チェック
		if observedMax < br {
			observedMax = br
		}
		if observedMin > br {
			observedMin = br
		}
	}
	assert.True(t, ok)
	assert.Equal(t, max, observedMax)
	assert.Equal(t, min, observedMin)
}

func Test浮動小数点数の乱数生成(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	ok := true
	for i := 0; i < 100000; i++ {
		r := rand.Float64()
		if !(0.0 <= r && r <= 1.0) {
			ok = false
		}
	}
	assert.True(t, ok)
}

func Test範囲を指定した浮動小数点数の乱数生成(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	max, min := 100.0, 10.0
	requiredError := 0.01

	ok := true
	observedMax := -1.0
	observedMin := 1000000.0
	for i := 0; i < 100000; i++ {
		bfr := BoundedFloat64(min, max)
		// 範囲チェック
		if !(min <= bfr && bfr <= max) {
			ok = false
		}
		// 最大値・最小値チェック
		if observedMax < bfr {
			observedMax = bfr
		}
		if observedMin > bfr {
			observedMin = bfr
		}
	}
	assert.True(t, ok)
	maxError := math.Abs(max - observedMax)
	minError := math.Abs(min - observedMin)
	assert.True(t, maxError < requiredError)
	assert.True(t, minError < requiredError)
}

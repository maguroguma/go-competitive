package bit

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2進数変換(t *testing.T) {
	assert.Equal(t, "0", strconv.FormatInt(0, 2))
	assert.Equal(t, "100", strconv.FormatInt(4, 2))

	// カウントアップ
	a := int64(0)
	assert.Equal(t, "0", strconv.FormatInt(a, 2))
	a++
	assert.Equal(t, "1", strconv.FormatInt(a, 2))
	a++
	assert.Equal(t, "10", strconv.FormatInt(a, 2))
	a++
	assert.Equal(t, "11", strconv.FormatInt(a, 2))
	a++
	assert.Equal(t, "100", strconv.FormatInt(a, 2))
}

func Test2進数を1ビットずつ取り出す(t *testing.T) {
	a := byte(100)
	assert.Equal(t, "1100100", strconv.FormatInt(int64(a), 2))
	actual := [8]bool{}
	assert.Equal(t, [8]bool{false, false, false, false, false, false, false, false}, actual)
	for i := byte(0); i < 8; i++ {
		b := (a >> i) & 1 // 右にiビットシフトして最下位ビットを取得する
		if b == 1 {
			actual[i] = true
		} else {
			actual[i] = false
		}
	}
	assert.Equal(t, [8]bool{false, false, true, false, false, true, true, false}, actual)
}

func Test2進数の任意ビットの1での編集(t *testing.T) {
	a := byte(0)
	b := a
	actual := [8]byte{}
	// 1での上書きはOR演算
	b |= a | (1 << 1)
	b |= a | (1 << 3)
	b |= a | (1 << 5)
	b |= a | (1 << 7)
	for i := 0; i < 8; i++ {
		actual[i] = 1 & (b >> byte(i))
	}
	assert.Equal(t, [8]byte{0, 1, 0, 1, 0, 1, 0, 1}, actual)
}

func TestNbitの2進数の全列挙(t *testing.T) {
	actual := []string{}
	// 2の3乗通りの全列挙のforループ
	for i := 0; i < (1 << 3); i++ {
		actual = append(actual, strconv.FormatInt(int64(i), 2))
	}
	assert.Equal(t, []string{"0", "1", "10", "11", "100", "101", "110", "111"}, actual)
}

package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test整数スライスをスペース区切り改行なし文字列に変換(t *testing.T) {
	assert.Equal(t, "1 10 100 1000", PrintIntsLine([]int{1, 10, 100, 1000}...))
}

package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowerBound(t *testing.T) {
	// 昇順ソート済み
	test := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}

	// keyが（複数）存在する場合、一番左のidxを返す
	assert.Equal(t, 3, LowerBound(test, 51))
	assert.Equal(t, 0, LowerBound(test, 1))
	assert.Equal(t, 9, LowerBound(test, 910))

	// 存在しないkeyを指定する場合、どこに挿入すべきかを表すidxを返す
	assert.Equal(t, 6, LowerBound(test, 52))
	assert.Equal(t, 0, LowerBound(test, 0))
	assert.Equal(t, 10, LowerBound(test, 911))
}

func TestUpperBound(t *testing.T) {
	// 昇順ソート済み
	test := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}

	// keyより大きい要素のうちの一番左のidxを返す（keyの存在の真偽を問わない）
	assert.Equal(t, 6, UpperBound(test, 51))
	assert.Equal(t, 1, UpperBound(test, 1))
	assert.Equal(t, 10, UpperBound(test, 910))

	// keyより大きい要素のうちの一番左のidxを返す
	assert.Equal(t, 6, UpperBound(test, 52))
	assert.Equal(t, 0, UpperBound(test, 0))
	assert.Equal(t, 10, UpperBound(test, 911))
}

func Test一致するキーの個数の計算(t *testing.T) {
	// 昇順ソート済み
	test := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 10, 10, 10, 20, 20, 20, 30, 30, 30}

	assert.Equal(t, 1, UpperBound(test, 1)-LowerBound(test, 1))
	assert.Equal(t, 2, UpperBound(test, 2)-LowerBound(test, 2))
	assert.Equal(t, 3, UpperBound(test, 3)-LowerBound(test, 3))
	assert.Equal(t, 4, UpperBound(test, 4)-LowerBound(test, 4))
	assert.Equal(t, 5, UpperBound(test, 5)-LowerBound(test, 5))

	// keyが存在しない場合にも対応
	assert.Equal(t, 0, UpperBound(test, 15)-LowerBound(test, 15))
	assert.Equal(t, 0, UpperBound(test, 100)-LowerBound(test, 100))
}

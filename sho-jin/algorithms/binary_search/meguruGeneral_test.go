package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneralLowerBound(t *testing.T) {
	// 昇順ソート済み
	test := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}

	// keyが（複数）存在する場合、一番左のidxを返す
	assert.Equal(t, 3, GeneralLowerBound(test, 51))
	assert.Equal(t, 0, GeneralLowerBound(test, 1))
	assert.Equal(t, 9, GeneralLowerBound(test, 910))

	// 存在しないkeyを指定する場合、どこに挿入すべきかを表すidxを返す
	assert.Equal(t, 6, GeneralLowerBound(test, 52))
	assert.Equal(t, 0, GeneralLowerBound(test, 0))
	assert.Equal(t, 10, GeneralLowerBound(test, 911))
}

func TestGeneralUpperBound(t *testing.T) {
	// 昇順ソート済み
	test := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}

	// keyより大きい要素のうちの一番左のidxを返す（keyの存在の真偽を問わない）
	assert.Equal(t, 6, GeneralUpperBound(test, 51))
	assert.Equal(t, 1, GeneralUpperBound(test, 1))
	assert.Equal(t, 10, GeneralUpperBound(test, 910))

	// keyより大きい要素のうちの一番左のidxを返す
	assert.Equal(t, 6, GeneralUpperBound(test, 52))
	assert.Equal(t, 0, GeneralUpperBound(test, 0))
	assert.Equal(t, 10, GeneralUpperBound(test, 911))
}

func Test真めぐる式二分探索実装による一致するキーの個数の計算(t *testing.T) {
	// 昇順ソート済み
	test := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 10, 10, 10, 20, 20, 20, 30, 30, 30}

	assert.Equal(t, 1, GeneralUpperBound(test, 1)-GeneralLowerBound(test, 1))
	assert.Equal(t, 2, GeneralUpperBound(test, 2)-GeneralLowerBound(test, 2))
	assert.Equal(t, 3, GeneralUpperBound(test, 3)-GeneralLowerBound(test, 3))
	assert.Equal(t, 4, GeneralUpperBound(test, 4)-GeneralLowerBound(test, 4))
	assert.Equal(t, 5, GeneralUpperBound(test, 5)-GeneralLowerBound(test, 5))

	// keyが存在しない場合にも対応
	assert.Equal(t, 0, GeneralUpperBound(test, 15)-GeneralLowerBound(test, 15))
	assert.Equal(t, 0, GeneralUpperBound(test, 100)-GeneralLowerBound(test, 100))
}

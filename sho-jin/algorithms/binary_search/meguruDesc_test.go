package binary_search

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneralDescLowerBound(t *testing.T) {
	// 降順ソート済み
	test := []int{910, 750, 419, 243, 51, 51, 51, 32, 14, 1}

	assert.Equal(t, 4, GeneralDescLowerBound(test, 51))
	assert.Equal(t, 9, GeneralDescLowerBound(test, 1))
	assert.Equal(t, 0, GeneralDescLowerBound(test, 910))

	assert.Equal(t, 4, GeneralDescLowerBound(test, 52))
	assert.Equal(t, 10, GeneralDescLowerBound(test, 0))
	assert.Equal(t, 0, GeneralDescLowerBound(test, 911))
}

func TestGeneralDescUpperBound(t *testing.T) {
	// 降順ソート済み
	test := []int{910, 750, 419, 243, 51, 51, 51, 32, 14, 1}

	assert.Equal(t, 7, GeneralDescUpperBound(test, 51))
	assert.Equal(t, 10, GeneralDescUpperBound(test, 1))
	assert.Equal(t, 1, GeneralDescUpperBound(test, 910))

	assert.Equal(t, 4, GeneralDescUpperBound(test, 52))
	assert.Equal(t, 10, GeneralDescUpperBound(test, 0))
	assert.Equal(t, 0, GeneralDescUpperBound(test, 911))
}

func Test真めぐる式二分探索実装による降順ソート済みスライスに対する一致するキーの個数の計算(t *testing.T) {
	// 降順ソート済み
	test := []int{30, 30, 30, 20, 20, 20, 10, 10, 10, 5, 5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 2, 2, 1}
	sort.Sort(sort.Reverse(sort.IntSlice(test)))

	assert.Equal(t, 1, GeneralDescUpperBound(test, 1)-GeneralDescLowerBound(test, 1))
	assert.Equal(t, 2, GeneralDescUpperBound(test, 2)-GeneralDescLowerBound(test, 2))
	assert.Equal(t, 3, GeneralDescUpperBound(test, 3)-GeneralDescLowerBound(test, 3))
	assert.Equal(t, 4, GeneralDescUpperBound(test, 4)-GeneralDescLowerBound(test, 4))
	assert.Equal(t, 5, GeneralDescUpperBound(test, 5)-GeneralDescLowerBound(test, 5))

	assert.Equal(t, 0, GeneralDescUpperBound(test, 15)-GeneralDescLowerBound(test, 15))
	assert.Equal(t, 0, GeneralDescUpperBound(test, 100)-GeneralDescLowerBound(test, 100))
}

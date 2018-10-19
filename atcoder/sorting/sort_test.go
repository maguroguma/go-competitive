package sorting

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test整数の昇順ソート(t *testing.T) {
	s := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	// 昇順ソート（Lessメソッドの自然な実装）がデフォルト
	sort.Sort(sort.IntSlice(s))
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s)
}

func Test整数の降順ソート(t *testing.T) {
	s := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	// Lessメソッドをデフォルトのものを反転するように再実装した型を用いる
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, s)
}

func Test浮動小数点数の昇順ソート(t *testing.T) {
	s := []float64{5.0, 6.0, 4.0, 7.0, 3.0, 8.0, 2.0, 9.0, 1.0, 0.0}
	sort.Sort(sort.Float64Slice(s))
	assert.Equal(t, []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}, s)
}

func Test文字列の辞書順ソート(t *testing.T) {
	s := []string{"php", "c", "go", "rust", "haskell"}
	sort.Sort(sort.StringSlice(s))
	assert.Equal(t, []string{"c", "go", "haskell", "php", "rust"}, s)
}

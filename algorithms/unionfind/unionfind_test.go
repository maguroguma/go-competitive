package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type query struct {
	q    int // 0: 併合, 1: 判定
	a, b int // 頂点
}

// https://atc001.contest.atcoder.jp/tasks/unionfind_aのサンプルテストケース
func TestUnionFind木(t *testing.T) {
	uniteQueries := []query{
		{0, 1, 2},
		{0, 3, 2},
		{1, 1, 3},
		{1, 1, 4},
		{0, 2, 4},
		{1, 4, 1},
		{0, 4, 2},
		{0, 0, 0},
		{1, 0, 0},
	}
	results := []string{}
	ccSizes := []int{}

	uft := NewUnionFindTree(8)

	for _, qu := range uniteQueries {
		if qu.q == 0 {
			uft.Unite(qu.a, qu.b)
		} else {
			isSame := uft.Same(qu.a, qu.b)
			if isSame {
				results = append(results, "Yes")
			} else {
				results = append(results, "No")
			}
			ccSizes = append(ccSizes, uft.ccSize(qu.b))
		}
	}

	assert.Equal(t, []string{"Yes", "No", "Yes", "Yes"}, results)
	assert.Equal(t, []int{3, 1, 4, 1}, ccSizes)
}

func TestUnionFind木改(t *testing.T) {
	uniteQueries := []query{
		{0, 1, 2},
		{0, 3, 2},
		{1, 1, 3},
		{1, 1, 4},
		{0, 2, 4},
		{1, 4, 1},
		{0, 4, 2},
		{0, 0, 0},
		{1, 0, 0},
	}
	results := []string{}
	ccSizes := []int{}

	uf := NewUnionFind(8)

	for _, qu := range uniteQueries {
		if qu.q == 0 {
			uf.Unite(qu.a, qu.b)
		} else {
			isSame := uf.Same(qu.a, qu.b)
			if isSame {
				results = append(results, "Yes")
			} else {
				results = append(results, "No")
			}
			ccSizes = append(ccSizes, uf.CcSize(qu.b))
		}
	}

	assert.Equal(t, []string{"Yes", "No", "Yes", "Yes"}, results)
	assert.Equal(t, []int{3, 1, 4, 1}, ccSizes)
}

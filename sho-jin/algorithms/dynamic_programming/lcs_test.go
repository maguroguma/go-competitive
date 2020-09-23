package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test最長部分列を1つ求める(t *testing.T) {
	// https://qiita.com/drken/items/a5e6fe22863b7992efdb
	assert.Equal(t, "abe", string(LCS([]rune("abcde"), []rune("acbef"))))
	assert.Equal(t, "priapr", string(LCS([]rune("pirikapirirara"), []rune("poporinapeperuto"))))

	// https://atcoder.jp/contests/dp/tasks/dp_f
	assert.Equal(t, "axb", string(LCS([]rune("axyb"), []rune("abyxb"))))
	assert.Equal(t, "aa", string(LCS([]rune("aa"), []rune("xayaz"))))
	assert.Equal(t, "", string(LCS([]rune("a"), []rune("z"))))
	assert.Equal(t, "aaadara", string(LCS([]rune("abracadabra"), []rune("avadakedavra"))))
}

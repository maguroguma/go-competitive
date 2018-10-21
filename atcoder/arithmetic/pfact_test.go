package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test試し割りによる素因数分解(t *testing.T) {
	assert.Equal(t, map[int]int{2: 3, 3: 2, 5: 1}, TrialDivision(360))
	assert.Equal(t, map[int]int{1: 1}, TrialDivision(1))
	assert.Equal(t, map[int]int{13: 1}, TrialDivision(13))
	assert.Equal(t, map[int]int{2: 1, 13: 1}, TrialDivision(26))
}

func Testエラトステネスの篩による素数集合の取得(t *testing.T) {
	actual := SieveOfEratosthenes(120)
	assert.Equal(t, []int{}, actual)
}

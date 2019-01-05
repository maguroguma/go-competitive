package prime

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

func Test試し割りによる素数判定(t *testing.T) {
	assert.True(t, IsPrime(2))
	assert.True(t, IsPrime(3))
	assert.True(t, IsPrime(5))
	assert.True(t, IsPrime(3571))
	assert.True(t, IsPrime(3559))
	assert.True(t, IsPrime(3557))

	assert.False(t, IsPrime(1))
	assert.False(t, IsPrime(4))
	assert.False(t, IsPrime(6))
	assert.False(t, IsPrime(3570))
	assert.False(t, IsPrime(3549))
	assert.False(t, IsPrime(3509))
}

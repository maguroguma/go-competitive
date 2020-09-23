package divisor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test試し割りによる約数列挙(t *testing.T) {
	assert.Equal(t, map[int]int{1: 1, 2: 1, 7: 1, 14: 1}, Divisors(14))
	assert.Equal(t, map[int]int{1: 1, 11: 1, 121: 1}, Divisors(121))
}

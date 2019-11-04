package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test等差数列の和(t *testing.T) {
	assert.Equal(t, 69, ArithmeticSequenceSum(4, 3, 6))
	assert.Equal(t, 2550, ArithmeticSequenceSum(2, 2, 50))
}

func Test等比数列の和(t *testing.T) {
	assert.Equal(t, 121, GeometricSequenceSum(1, 3, 5))
}

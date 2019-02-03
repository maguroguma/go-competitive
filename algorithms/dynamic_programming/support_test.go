package dynamic_programming

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateValue(t *testing.T) {
	actual := math.MaxInt32
	ChMin(&actual, 500)
	assert.Equal(t, 500, actual)
	ChMin(&actual, 501)
	assert.Equal(t, 500, actual)
	ChMin(&actual, -100)
	assert.Equal(t, -100, actual)

	actual = -1
	ChMax(&actual, 100)
	assert.Equal(t, 100, actual)
	ChMax(&actual, 99)
	assert.Equal(t, 100, actual)
	ChMax(&actual, 500)
	assert.Equal(t, 500, actual)
}

func TestDPUsage(t *testing.T) {
	dp := [2][2]int{
		[2]int{0, 0},
		[2]int{0, 0},
	}

	ChMax(&dp[0][0], 100)
	assert.Equal(t, 100, dp[0][0])
	ChMin(&dp[0][1], -100)
	assert.Equal(t, -100, dp[0][1])
}

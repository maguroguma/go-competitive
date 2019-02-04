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

func TestGetNthBit(t *testing.T) {
	i := 1 + 2*2 + 2*2*2*2 + 2*2*2*2*2*2*2*2*2
	assert.Equal(t, GetNthBit(i, 0), 1)
	assert.Equal(t, GetNthBit(i, 1), 0)
	assert.Equal(t, GetNthBit(i, 2), 1)
	assert.Equal(t, GetNthBit(i, 3), 0)
	assert.Equal(t, GetNthBit(i, 4), 1)
	assert.Equal(t, GetNthBit(i, 5), 0)
	assert.Equal(t, GetNthBit(i, 6), 0)
	assert.Equal(t, GetNthBit(i, 7), 0)
	assert.Equal(t, GetNthBit(i, 8), 0)
	assert.Equal(t, GetNthBit(i, 9), 1)
	assert.Equal(t, GetNthBit(i, 10), 0)
	assert.Equal(t, GetNthBit(i, 11), 0)
}

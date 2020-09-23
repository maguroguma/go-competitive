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

func TestOnBit(t *testing.T) {
	// i == 1000010101
	i := 1 + 2*2 + 2*2*2*2 + 2*2*2*2*2*2*2*2*2
	assert.Equal(t, OnBit(i, 0), i)
	assert.Equal(t, OnBit(i, 1), i+2)
	assert.Equal(t, OnBit(i, 2), i)
	assert.Equal(t, OnBit(i, 3), i+2*2*2)
	assert.Equal(t, OnBit(i, 4), i)
	assert.Equal(t, OnBit(i, 5), i+2*2*2*2*2)
	assert.Equal(t, OnBit(i, 6), i+2*2*2*2*2*2)
	assert.Equal(t, OnBit(i, 7), i+2*2*2*2*2*2*2)
	assert.Equal(t, OnBit(i, 8), i+2*2*2*2*2*2*2*2)
	assert.Equal(t, OnBit(i, 9), i)
	assert.Equal(t, OnBit(i, 10), i+2*2*2*2*2*2*2*2*2*2)
	assert.Equal(t, OnBit(i, 11), i+2*2*2*2*2*2*2*2*2*2*2)
}

func TestOffBit(t *testing.T) {
	// i == 1000010101
	i := 1 + 2*2 + 2*2*2*2 + 2*2*2*2*2*2*2*2*2
	assert.Equal(t, OffBit(i, 0), i-1)
	assert.Equal(t, OffBit(i, 1), i)
	assert.Equal(t, OffBit(i, 2), i-2*2)
	assert.Equal(t, OffBit(i, 3), i)
	assert.Equal(t, OffBit(i, 4), i-2*2*2*2)
	assert.Equal(t, OffBit(i, 5), i)
	assert.Equal(t, OffBit(i, 6), i)
	assert.Equal(t, OffBit(i, 7), i)
	assert.Equal(t, OffBit(i, 8), i)
	assert.Equal(t, OffBit(i, 9), i-2*2*2*2*2*2*2*2*2)
	assert.Equal(t, OffBit(i, 10), i)
	assert.Equal(t, OffBit(i, 11), i)
}

func TestPopCount(t *testing.T) {
	// i == 1000010101
	i := 1 + 2*2 + 2*2*2*2 + 2*2*2*2*2*2*2*2*2
	assert.Equal(t, PopCount(i), 4)
	assert.Equal(t, PopCount(1<<5), 1)
	assert.Equal(t, PopCount(OnBit(0, 1)+OnBit(0, 55)), 2)
	assert.Equal(t, PopCount((1<<63)-1), 63)
	assert.Equal(t, PopCount(OnBit(0, 1)+OnBit(0, 66)), 1)
	assert.Equal(t, PopCount(OnBit(0, 70)), 0)
}

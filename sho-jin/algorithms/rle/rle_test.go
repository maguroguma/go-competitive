package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var S []rune
var A []int
var L []int

func TestRLE(t *testing.T) {
	S, L = RunLengthEncoding([]rune("aaabbbaaa"))
	assert.Equal(t, []rune("aba"), S)
	assert.Equal(t, []int{3, 3, 3}, L)

	S, L = RunLengthEncoding([]rune("abcdefg"))
	assert.Equal(t, []rune("abcdefg"), S)
	assert.Equal(t, []int{1, 1, 1, 1, 1, 1, 1}, L)

	S, L = RunLengthEncoding([]rune("aaabbbc"))
	assert.Equal(t, []rune("abc"), S)
	assert.Equal(t, []int{3, 3, 1}, L)

	S, L = RunLengthEncoding([]rune("abccccc"))
	assert.Equal(t, []rune("abc"), S)
	assert.Equal(t, []int{1, 1, 5}, L)

	S, L = RunLengthEncoding([]rune("aaaaaaa"))
	assert.Equal(t, []rune("a"), S)
	assert.Equal(t, []int{7}, L)

	S, L = RunLengthEncoding([]rune("a"))
	assert.Equal(t, []rune("a"), S)
	assert.Equal(t, []int{1}, L)

	S, L = RunLengthEncoding([]rune("ab"))
	assert.Equal(t, []rune("ab"), S)
	assert.Equal(t, []int{1, 1}, L)
}

func TestRLEIntVer(t *testing.T) {
	A, L = RunLengthEncodingIntVer([]int{1, 1, 1, 2, 2, 2, 3, 3, 3})
	assert.Equal(t, []int{1, 2, 3}, A)
	assert.Equal(t, []int{3, 3, 3}, L)

	A, L = RunLengthEncodingIntVer([]int{1, 2, 3, 4, 5, 6, 7})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, A)
	assert.Equal(t, []int{1, 1, 1, 1, 1, 1, 1}, L)

	A, L = RunLengthEncodingIntVer([]int{1, 1, 1, 2, 2, 2, 3})
	assert.Equal(t, []int{1, 2, 3}, A)
	assert.Equal(t, []int{3, 3, 1}, L)

	A, L = RunLengthEncodingIntVer([]int{1, 2, 3, 3, 3, 3, 3})
	assert.Equal(t, []int{1, 2, 3}, A)
	assert.Equal(t, []int{1, 1, 5}, L)

	A, L = RunLengthEncodingIntVer([]int{1, 1, 1, 1, 1, 1, 1})
	assert.Equal(t, []int{1}, A)
	assert.Equal(t, []int{7}, L)

	A, L = RunLengthEncodingIntVer([]int{1})
	assert.Equal(t, []int{1}, A)
	assert.Equal(t, []int{1}, L)

	A, L = RunLengthEncodingIntVer([]int{1, 2})
	assert.Equal(t, []int{1, 2}, A)
	assert.Equal(t, []int{1, 1}, L)
}

func TestRLD(t *testing.T) {
	S = RunLengthDecoding(RunLengthEncoding([]rune("aaabbbaaa")))
	assert.Equal(t, []rune("aaabbbaaa"), S)
}

func TestRLDIntVer(t *testing.T) {
	A = RunLengthDecodingIntVer(RunLengthEncodingIntVer([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}))
	assert.Equal(t, []int{1, 1, 1, 2, 2, 2, 3, 3, 3}, A)
}

package permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test順列(t *testing.T) {
	expected := []string{
		"abc",
		"acb",
		"bac",
		"bca",
		"cab",
		"cba",
	}
	tmp := CalcFactorialPatterns([]rune{'a', 'b', 'c'})
	actual := []string{}
	for _, t := range tmp {
		actual = append(actual, string(t))
	}
	assert.Equal(t, expected, actual)
}

func Test重複順列(t *testing.T) {
	expected := []string{
		"aaa", "aab", "aac",
		"aba", "abb", "abc",
		"aca", "acb", "acc",
		"baa", "bab", "bac",
		"bba", "bbb", "bbc",
		"bca", "bcb", "bcc",
		"caa", "cab", "cac",
		"cba", "cbb", "cbc",
		"cca", "ccb", "ccc",
	}
	tmp := CalcDuplicatePatterns([]rune{'a', 'b', 'c'}, 3)
	actual := []string{}
	for _, t := range tmp {
		actual = append(actual, string(t))
	}
	assert.Equal(t, expected, actual)
}

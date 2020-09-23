package permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test階乗(t *testing.T) {
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

func TestN個からK個とる順列(t *testing.T) {
	A := []rune{'a', 'b', 'c', 'd', 'e'}
	expected := []string{
		"ab", "ac", "ad", "ae",
		"ba", "bc", "bd", "be",
		"ca", "cb", "cd", "ce",
		"da", "db", "dc", "de",
		"ea", "eb", "ec", "ed",
	}
	tmp := CalcPermutationPatterns(A, 2)
	actual := []string{}
	for _, t := range tmp {
		actual = append(actual, string(t))
	}

	assert.Equal(t, expected, actual)

	// nP0 と nPn
	assert.Equal(t, [][]rune{[]rune{}}, CalcPermutationPatterns(A, 0))
	assert.Equal(t, CalcFactorialPatterns(A), CalcPermutationPatterns(A, len(A)))
}

func TestN個からK個取る組み合わせ(t *testing.T) {
	A := []rune{'a', 'b', 'c', 'd', 'e'}
	expected := []string{
		"abc", "abd", "abe",
		"acd", "ace",
		"ade",
		"bcd", "bce",
		"bde",
		"cde",
	}
	tmp := CalcCombinationPatterns(A, 3)
	actual := []string{}
	for _, t := range tmp {
		actual = append(actual, string(t))
	}

	assert.Equal(t, expected, actual)

	// nC0 と nCn
	assert.Equal(t, [][]rune{[]rune{}}, CalcCombinationPatterns(A, 0))
	assert.Equal(t, [][]rune{[]rune{'a', 'b', 'c', 'd', 'e'}}, CalcCombinationPatterns(A, len(A)))
}

func Test重複順列(t *testing.T) {
	A := []rune{'a', 'b', 'c'}
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
	tmp := CalcDuplicatePatterns(A, 3)
	actual := []string{}
	for _, t := range tmp {
		actual = append(actual, string(t))
	}
	assert.Equal(t, expected, actual)

	// n^0
	assert.Equal(t, [][]rune{[]rune{}}, CalcDuplicatePatterns(A, 0))
}

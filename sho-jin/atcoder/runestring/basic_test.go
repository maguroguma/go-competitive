package runestring

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var lowerAlphabetsRunes = []rune{'a', 'b', 'c', 'd', 'e',
	'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y',
	'z'}
var upperAlphabetsRunes = []rune{'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y',
	'Z'}
var otherRunes = []rune{'1', '(', 'あ', '漢'}

const lowerAlphabetsString string = "abcdefghijklmnopqrstuvwsyz"
const upperAlphabetsString string = "ABCDEFGHIJKLMNOPQRSTUVWSYZ"

func TestRuneのアルファベットを大文字もしくは小文字に変換しそれ以外は不変(t *testing.T) {
	for i := 0; i < len(lowerAlphabetsRunes); i++ {
		actual := UpperRune(lowerAlphabetsRunes[i])
		assert.Equal(t, upperAlphabetsRunes[i], actual)
		actual = LowerRune(lowerAlphabetsRunes[i])
		assert.Equal(t, lowerAlphabetsRunes[i], actual)
		actual = LowerRune(upperAlphabetsRunes[i])
		assert.Equal(t, lowerAlphabetsRunes[i], actual)
		actual = UpperRune(upperAlphabetsRunes[i])
		assert.Equal(t, upperAlphabetsRunes[i], actual)
	}
	for i := 0; i < len(otherRunes); i++ {
		actual := UpperRune(otherRunes[i])
		assert.Equal(t, otherRunes[i], actual)
		actual = LowerRune(otherRunes[i])
		assert.Equal(t, otherRunes[i], actual)
	}
}

func Testアルファベットのみを反転(t *testing.T) {
	for i := 0; i < len(lowerAlphabetsRunes); i++ {
		actual := ToggleRune(lowerAlphabetsRunes[i])
		assert.Equal(t, upperAlphabetsRunes[i], actual)
		actual = ToggleRune(upperAlphabetsRunes[i])
		assert.Equal(t, lowerAlphabetsRunes[i], actual)
	}
	for i := 0; i < len(otherRunes); i++ {
		actual := ToggleRune(otherRunes[i])
		assert.Equal(t, otherRunes[i], actual)
		actual = ToggleRune(otherRunes[i])
		assert.Equal(t, otherRunes[i], actual)
	}
}

func Test文字列中のアルファベットを反転(t *testing.T) {
	testStr := "AbCdEfGhIjKlMnOpQrStUvWxYz(){}+-*/あ漢"
	assert.Equal(t, "aBcDeFgHiJkLmNoPqRsTuVwXyZ(){}+-*/あ漢", ToggleString(testStr))
}

/* STUDY TEST */

func TestRune型の文字を文字列に変換(t *testing.T) {
	r := 'A' // 型名はintとなることにも注意
	rr := []rune{'A', 'B', 'C'}
	assert.Equal(t, "A", string(r))
	assert.Equal(t, "ABC", string(rr))
	// 以下は不可
	//	s := "A"
	//	assert.Equal(t, 'A', rune(s))
}

func Test文字列型の文字列の大文字小文字を変換(t *testing.T) {
	assert.Equal(t, upperAlphabetsString, strings.ToUpper(lowerAlphabetsString))
	assert.Equal(t, lowerAlphabetsString, strings.ToLower(upperAlphabetsString))
	assert.Equal(t, upperAlphabetsString, strings.ToUpper(upperAlphabetsString))
	assert.Equal(t, lowerAlphabetsString, strings.ToLower(lowerAlphabetsString))
}

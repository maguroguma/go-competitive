package strprocess

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test文字列から数値の変換(t *testing.T) {
	intStr1 := "123456789"
	intStr2 := "01"
	intStr3 := "000010"
	floatStr := "123.456789"
	// 文字列から数値への変換はstrconv.Atox, strconv.ParseXxx（errorも返される）
	actual1, _ := strconv.Atoi(intStr1)
	actual2, _ := strconv.Atoi(intStr2)
	actual3, _ := strconv.Atoi(intStr3)
	actual4, _ := strconv.ParseFloat(floatStr, 64)
	assert.Equal(t, 123456789, actual1)
	assert.Equal(t, 1, actual2)
	assert.Equal(t, 10, actual3)
	assert.Equal(t, 123.456789, actual4)
}

func Test数値から文字列への変換(t *testing.T) {
	testInt := 111
	testFloat := 11.1
	// 数値から文字列はstrconv.Itoa, strconv.FormatXxx
	assert.Equal(t, "111", strconv.Itoa(testInt))
	assert.Equal(t, "11.1000", strconv.FormatFloat(testFloat, 'f', 4, 64))
}

func Test文字列とruneの変換(t *testing.T) {
	exampleStr := "競技programming"
	// 文字列からruneスライスへ変換（キャスト）
	exampleRunes := []rune(exampleStr)
	assert.Equal(t, []rune{'競', '技', 'p', 'r', 'o', 'g', 'r', 'a', 'm', 'm', 'i', 'n', 'g'}, exampleRunes)
	// rune or runeスライスから文字列へ変換（キャスト）
	assert.Equal(t, "競", string(exampleRunes[0]))
	assert.Equal(t, "o", string(exampleRunes[4]))
	assert.Equal(t, "技prog", string(exampleRunes[1:6]))
	// 長さ（文字列型の長さは注意）
	assert.Equal(t, 17, len(exampleStr))
	assert.Equal(t, 13, len(exampleRunes))
}

func Test数値をruneスライスとして扱う(t *testing.T) {
	runesInt := []rune{'1', '2', '3', '0'}
	assert.Equal(t, 51, int(runesInt[2]))      // intに変換するだけではダメ
	assert.Equal(t, int32(3), runesInt[2]-'0') // *runeの実体はint32のエイリアス* であるため、このようにrune同士の演算で正しく扱える
}

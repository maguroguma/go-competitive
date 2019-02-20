package bitope

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testビットの基本演算(t *testing.T) {
	// golangのAND/OR演算は、それぞれ&/|

	// 45 = 0b101101
	// 25 = 0b011001
	//  9 = 0b001001
	assert.Equal(t, 45&25, 9)
	// 61 = 0b111101
	assert.Equal(t, 45|25, 61)
}

func Testビット表記の文字列取得(t *testing.T) {
	// 主にデバッグ用に（2進数の桁DPなどでは使いやすくなるかもしれない）
	// 最上位が1のビットで打ち止めなので、場合によっては手動で0埋めするなどの処理が必要
	a, b := 45, 25
	sa, sb := fmt.Sprintf("%b", a), fmt.Sprintf("%b", b)
	assert.Equal(t, "101101", sa)
	assert.Equal(t, "11001", sb)
}

func Testビット操作の基本(t *testing.T) {
	// 45 = 0b101101
	// フラグ立て（"0...010...0"というマスクでOR演算を取る）
	assert.Equal(t, "111101", fmt.Sprintf("%b", 45|(1<<4)))
	assert.Equal(t, "101101", fmt.Sprintf("%b", 45|(1<<3)))

	// フラグ降ろし（"1...101...1"というマスクでAND演算を取る、ビット反転演算を利用している）
	assert.Equal(t, "101001", fmt.Sprintf("%b", 45&^(1<<2)))
	assert.Equal(t, "101101", fmt.Sprintf("%b", 45&^(1<<1)))
}

func Testマスクビット(t *testing.T) {
	// 複数のフラグをまとめて立てる
	// 複数のフラグをまとめて消す
	// 必要な情報だけを取り出すために、マスクした部分の情報のみを取り出す
	const BITFLAGDAMAGE = (1 << 0)
	const BITFLAGDOKU = (1 << 1)
	const BITFLAGMAHI = (1 << 2)
	const BITFLAGSENTOFUNO = (1 << 3)

	const MASKATTACK = BITFLAGDAMAGE
	const MASKPUNCH = BITFLAGDAMAGE | BITFLAGMAHI
	const MASKDEFEAT = BITFLAGDAMAGE | BITFLAGSENTOFUNO
	const MASKDOKUMAHI = BITFLAGDOKU | BITFLAGMAHI

	status := 0
	// 立てる系は|=
	assert.Equal(t, "0", fmt.Sprintf("%b", status))
	status |= MASKATTACK
	assert.Equal(t, "1", fmt.Sprintf("%b", status))
	status |= MASKPUNCH
	assert.Equal(t, "101", fmt.Sprintf("%b", status))
	assert.True(t, status&MASKDOKUMAHI > 0)
	// 降ろす系は反転マスクを&=
	status &= ^MASKDOKUMAHI
	assert.Equal(t, "1", fmt.Sprintf("%b", status))

	status |= MASKDEFEAT
	assert.Equal(t, "1001", fmt.Sprintf("%b", status))
	// 毒・麻痺を回復しても戦闘不能は治らずそのまま
	status &= ^MASKDOKUMAHI
	assert.Equal(t, "1001", fmt.Sprintf("%b", status))
}

func Test任意の部分集合の列挙(t *testing.T) {
	n := 10
	A := (1 << 2) | (1 << 3) | (1 << 5) | (1 << 7)

	expected := "[2 3 5 7][3 5 7][2 5 7][5 7][2 3 7][3 7][2 7][7][2 3 5][3 5][2 5][5][2 3][3][2][]"

	actual := ""
	for bit := A; ; bit = (bit - 1) & A {
		S := []int{}
		for i := 0; i < n; i++ {
			// iがbitに入るかどうか
			if bit&(1<<uint(i)) > 0 {
				S = append(S, i)
			}
		}

		actual += fmt.Sprintf("%v", S)

		// 最後の0でbreak
		if bit == 0 {
			break
		}
	}

	assert.Equal(t, expected, actual)
}

func TestBitDPで巡回セールマン問題(t *testing.T) {
}

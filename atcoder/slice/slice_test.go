package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteElement(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = DeleteElement(a, 4)
	assert.Equal(t, []int{0, 1, 2, 3, 5, 6, 7, 8, 9}, a)
	assert.Equal(t, 9, len(a))
	assert.Equal(t, 9, cap(a))
	a = DeleteElement(a, 0)
	assert.Equal(t, []int{1, 2, 3, 5, 6, 7, 8, 9}, a)
	assert.Equal(t, 8, len(a))
	assert.Equal(t, 8, cap(a))
	a = DeleteElement(a, len(a)-1)
	assert.Equal(t, []int{1, 2, 3, 5, 6, 7, 8}, a)
	assert.Equal(t, 7, len(a))
	assert.Equal(t, 7, cap(a))

	// 副作用がないことの確認
	b := []int{0, 1, 2, 3, 4, 5}
	c := SafeDeleteElement(b, 0)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, b)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, c)
}

func TestSafeDeleteElement(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5}
	a = SafeDeleteElement(a, 0)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, a)
	assert.Equal(t, 5, len(a))
	assert.Equal(t, 5, cap(a))
	a = SafeDeleteElement(a, len(a)-1)
	assert.Equal(t, []int{1, 2, 3, 4}, a)
	assert.Equal(t, 4, len(a))
	assert.Equal(t, 4, cap(a))
	a = SafeDeleteElement(a, 1)
	assert.Equal(t, []int{1, 3, 4}, a)
	assert.Equal(t, 3, len(a))
	assert.Equal(t, 3, cap(a))

	// 副作用がないことの確認
	b := []int{0, 1, 2, 3, 4, 5}
	c := SafeDeleteElement(b, 0)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, b)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, c)
}

func TestConcat(t *testing.T) {
	s := []rune{'b', 'e', 'f', 'o', 'r', 'e'}
	tt := []rune{'a', 'f', 't', 'e', 'r'}
	n := Concat(s, tt)
	assert.Equal(t, "beforeafter", string(n))
	assert.Equal(t, len(s)+len(tt), len(n))
	assert.Equal(t, cap(s)+cap(tt), cap(n))

	// 副作用がないことの確認
	assert.Equal(t, []rune{'b', 'e', 'f', 'o', 'r', 'e'}, s)
	assert.Equal(t, []rune{'a', 'f', 't', 'e', 'r'}, tt)
}

func TestSafeConcat(t *testing.T) {
	s := []rune{'b', 'e', 'f', 'o', 'r', 'e'}
	tt := []rune{'a', 'f', 't', 'e', 'r'}
	n := SafeConcat(s, tt)
	assert.Equal(t, "beforeafter", string(n))
	assert.Equal(t, len(s)+len(tt), len(n))
	assert.Equal(t, cap(s)+cap(tt), cap(n))

	// 副作用がないことの確認
	assert.Equal(t, []rune{'b', 'e', 'f', 'o', 'r', 'e'}, s)
	assert.Equal(t, []rune{'a', 'f', 't', 'e', 'r'}, tt)
}

/* STUDY TESTS */

func Testスライスのデリート(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := append(a[:3], a[4:]...)
	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 7, 8, 9}, b)
	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 7, 8, 9, 9}, a) // **もとのスライスは使えなくなる**
	assert.Equal(t, 10, len(a))
	assert.Equal(t, 10, cap(a))
	assert.Equal(t, 9, len(b))
	assert.Equal(t, 10, cap(b))
	// もとの変数を継続的に使いたい場合は、参照先を変更する
	a = b
	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 7, 8, 9}, a)
	assert.Equal(t, len(b), len(a))
	assert.Equal(t, cap(b), cap(a))
	// 同じ要領でデリートの繰り返し
	b = append(a[:1], a[2:]...)
	a = b
	assert.Equal(t, []int{0, 2, 4, 5, 6, 7, 8, 9}, a)
	assert.Equal(t, len(b), len(a))
	assert.Equal(t, cap(b), cap(a))
	b = a[1:] // 先頭のデリート
	a = b
	assert.Equal(t, []int{2, 4, 5, 6, 7, 8, 9}, a)
	assert.Equal(t, len(b), len(a))
	assert.Equal(t, cap(b), cap(a))
	b = a[:len(a)-1] // 末尾のデリート
	a = b
	assert.Equal(t, []int{2, 4, 5, 6, 7, 8}, a)
	assert.Equal(t, 6, len(a))
	assert.Equal(t, 9, cap(a))
}
func Test副作用を許容するスライスのデリート(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = append(a[:3], a[4:]...)
	assert.Equal(t, []int{0, 1, 2, 4, 5, 6, 7, 8, 9}, a)
	a = append(a[:3], a[4:]...)
	assert.Equal(t, []int{0, 1, 2, 5, 6, 7, 8, 9}, a)
}

func Testスライスのコピー(t *testing.T) {
	// スライスは参照型
	a := []int{0, 1, 2}
	b := a
	b[0] = 100
	assert.Equal(t, []int{100, 1, 2}, a)
	assert.Equal(t, []int{100, 1, 2}, b)
	assert.Equal(t, &b, &a)
	// 簡易スライス式を使っても、自動拡張が起こらない限り参照先は同じ
	a = []int{0, 1, 2}
	b = a[:]
	b[0] = 100
	assert.Equal(t, []int{100, 1, 2}, a)
	assert.Equal(t, []int{100, 1, 2}, b)
	assert.Equal(t, &b, &a)

	// 1. len=0, cap=len(copied) の新しいスライスを定義しappendする（自動拡張分のパフォーマンスを気にしないならcapの指定は不要）
	a = []int{0, 1, 2}
	b = make([]int, 0, len(a))
	b = append(b, a...)
	b[0] = 100
	assert.Equal(t, []int{0, 1, 2}, a)
	assert.Equal(t, []int{100, 1, 2}, b)
	assert.NotEqual(t, &b, &a)
	// 2. len=cap=len(copied) の新しいスライスを定義しcopy関数で塗りつぶす
	a = []int{0, 1, 2}
	b = make([]int, len(a))
	copy(b, a)
	b[0] = 100
	assert.Equal(t, []int{0, 1, 2}, a)
	assert.Equal(t, []int{100, 1, 2}, b)
	assert.NotEqual(t, &b, &a)
}

func Test組み込みのcopy関数(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11}
	copy(s1, s2)
	assert.Equal(t, []int{10, 11, 3, 4, 5}, s1)
	assert.Equal(t, []int{10, 11}, s2)
	assert.NotEqual(t, &s1, &s2)

	// len(src) > len(dst) でもpanicにならない
	s1 = []int{1, 2, 3, 4, 5}
	s2 = []int{10, 11, 12, 13, 14, 15, 16}
	copy(s1, s2)
	assert.Equal(t, []int{10, 11, 12, 13, 14}, s1)
	assert.Equal(t, []int{10, 11, 12, 13, 14, 15, 16}, s2)
	assert.NotEqual(t, &s1, &s2)
}

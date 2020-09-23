package mapuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mapKey struct {
	id   int
	name string
}

func Test構造体をキーとするマップ(t *testing.T) {
	a := mapKey{id: 1, name: "one"}
	b := mapKey{id: 2, name: "two"}
	aa := mapKey{id: 1, name: "one"}
	bb := mapKey{id: 2, name: "two"}
	m := map[mapKey]string{}
	m[a] = "a is set"
	m[b] = "b is set"
	_, aaIsSet := m[aa]
	_, bbIsSet := m[bb]
	assert.True(t, aaIsSet)
	assert.True(t, bbIsSet)
	assert.Equal(t, "a is set", m[aa])
	assert.Equal(t, "b is set", m[bb])
	assert.Equal(t, 2, len(m))
}

func Test構造体のポインタをキーとするマップ(t *testing.T) {
	a := mapKey{id: 1, name: "one"}
	b := mapKey{id: 2, name: "two"}
	aa := mapKey{id: 1, name: "one"}
	bb := mapKey{id: 2, name: "two"}
	m := map[*mapKey]string{}
	m[&a] = "a is set"
	m[&b] = "b is set"
	_, aaIsSet := m[&aa]
	_, bbIsSet := m[&bb]
	assert.False(t, aaIsSet)
	assert.False(t, bbIsSet)
	assert.Equal(t, "", m[&aa])
	assert.Equal(t, "", m[&bb])
	assert.Equal(t, 2, len(m))
}

func Test初期化されていないマップの操作(t *testing.T) {
	m := make(map[int]int)
	m[0]++
	assert.Equal(t, 1, m[0])
	assert.Equal(t, 1, len(m))
}

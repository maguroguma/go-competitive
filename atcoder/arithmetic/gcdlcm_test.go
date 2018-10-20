package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test最大公約数(t *testing.T) {
	expect := 5
	assert.Equal(t, expect, Gcd(110, 35))
	assert.Equal(t, 1, Gcd(110/expect, 35/expect))
	expect = 40
	assert.Equal(t, expect, Gcd(200, 80))
	assert.Equal(t, 1, Gcd(200/expect, 80/expect))
	assert.Equal(t, 1, Gcd(1, 1))
	assert.Equal(t, 1, Gcd(3, 2))
	assert.Equal(t, 1, Gcd(2, 3))
	assert.Equal(t, 21, Gcd(1071, 1029))
	assert.Equal(t, 21, Gcd(1029, 1071))
}

func Test最小公倍数(t *testing.T) {
	assert.Equal(t, 210, Lcm(30, 42))
	assert.Equal(t, 1, Lcm(1, 1))
	assert.Equal(t, 6, Lcm(2, 3))
	assert.Equal(t, 6, Lcm(3, 2))
}

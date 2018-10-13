package strprocess

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test文字列のデリミタによる分割(t *testing.T) {
	testStr := "123+456+789"
	actual := strings.Split(testStr, "+")
	assert.Equal(t, []string{"123", "456", "789"}, actual)
}

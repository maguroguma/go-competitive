package measuring_worm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testしゃくとり法の基本(t *testing.T) {
	n := 12
	A := []int{4, 6, 7, 8, 1, 2, 110, 2, 4, 12, 3, 9}
	x := 25

	ans := 0

	r := 0
	sum := 0
	for l := 0; l < n; l++ {
		// f(left)の計算（rightを前進させられるだけさせる）
		for r < n && sum+A[r] <= x {
			sum += A[r]
			r++
		}

		// f(left)が求まっている、半開区間なので+1はしないこと
		ans += r - l

		// leftを前進させる前の準備
		if r == l {
			r++ // rがlに重なったらrを1つ前進させておく
		} else {
			sum -= A[l] // leftのみが前進するので、sumからA[l]を引く
		}
	}

	assert.Equal(t, 32, ans)
}

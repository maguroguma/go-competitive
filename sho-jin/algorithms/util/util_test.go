package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test10進数の桁和(t *testing.T) {
	assert.Equal(t, 45, GetDigitSum(123456789))
	assert.Equal(t, 45, GetDigitSum(102030405060708090))
	assert.Equal(t, -1, GetDigitSum(-123))
	assert.Equal(t, 45, GetDigitSum(908070605040302010))
}

func Test整数スライスの総和(t *testing.T) {
	assert.Equal(t, 45, Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	assert.Equal(t, 45, Sum(0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7, 0, 8, 0, 9))
	assert.Equal(t, 45, Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
}

func Test整数スライスの累積和スライス(t *testing.T) {
	assert.Equal(t, []int{0, 1, 3, 6, 10, 15, 21, 28, 36, 45},
		GetCumulativeSums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func Test累積和スライスを用いた部分和の算出(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sums := GetCumulativeSums(A)

	// 半開区間として扱える
	// ex.) 0<=m<=n のとき A の [m, n) の和は sums[n] - sums[m]
	assert.Equal(t, A[0]+A[1]+A[2]+A[3], sums[4]-sums[0])
	assert.Equal(t, A[3]+A[4]+A[5]+A[6], sums[7]-sums[3])
	// コーナーケース（？）
	assert.Equal(t, 45, sums[len(A)])

	// 備忘録
	// ↓これはセーフ
	assert.Equal(t, []int{}, A[len(A):len(A)])
	// ↓これはアウト
	// assert.Equal(t, []int{}, A[len(A)+1:len(A)+1])
}

func Test整数型除算の切り上げ(t *testing.T) {
	assert.Equal(t, 2, Kiriage(5, 3))
	assert.Equal(t, 2, Kiriage(6, 3))
	assert.Equal(t, 0, Kiriage(0, 100))
	assert.Equal(t, 1, Kiriage(1<<60-1, 1<<60))
	assert.Equal(t, 1, Kiriage(1, 1<<60))
	assert.Equal(t, 0, Kiriage(0, 1<<60))
}

func Test10進数の桁数(t *testing.T) {
	assert.Equal(t, 0, DigitNumOfDecimal(0))
	assert.Equal(t, 1, DigitNumOfDecimal(1))
	assert.Equal(t, 1, DigitNumOfDecimal(9))
	assert.Equal(t, 2, DigitNumOfDecimal(10))
	assert.Equal(t, 3, DigitNumOfDecimal(100))
}

package util

// GetDigitSum returns digit sum of a decimal number.
// GetDigitSum only accept a positive integer.
func GetDigitSum(n int) int {
	if n < 0 {
		return -1
	}

	res := 0

	for n > 0 {
		res += n % 10
		n /= 10
	}

	return res
}

// Sum returns multiple integers sum.
func Sum(integers ...int) int {
	s := 0

	for _, i := range integers {
		s += i
	}

	return s
}

// GetCumulativeSums returns cumulative sums.
// Length of result slice is equal to that of an argument +1.
func GetCumulativeSums(integers []int) []int {
	res := make([]int, len(integers)+1)

	res[0] = 0
	for i, a := range integers {
		res[i+1] = res[i] + a
	}

	return res
}

// Kiriage returns Ceil(a/b)
// a >= 0, b > 0
func Kiriage(a, b int) int {
	return (a + (b - 1)) / b
}

// 任意のスライスを反転した、新しいスライスを返す

// unshift

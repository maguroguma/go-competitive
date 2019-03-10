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
// Length of result slice is equal to that of an argument.
func GetCumulativeSums(integers []int) []int {
	res := make([]int, len(integers))

	currentSum := 0
	for i, a := range integers {
		currentSum += a
		res[i] = currentSum
	}

	return res
}

package dynamic_programming

// ChMin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

// ChMax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
}

// GetNthBit returns nth bit value of an argument.
// n starts from 0.
func GetNthBit(num, nth int) int {
	return num >> uint(nth) & 1
}

// func OnBit(num, nth int) int {

// }

// func OffBit(num, nth int) int {

// }

// func PopCount(num int) int {

// }

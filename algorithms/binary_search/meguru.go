package binary_search

// LowerBound returns an index of a slice whose value is EQUAL TO AND LARGER THAN A KEY VALUE.
func LowerBound(s []int, key int) int {
	isLarger := func(index, key int) bool {
		if s[index] >= key {
			return true
		} else {
			return false
		}
	}

	left, right := -1, len(s)

	for right-left > 1 {
		mid := left + (right-left)/2
		if isLarger(mid, key) {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}

// UpperBound returns an index of a slice whose value is EQUAL TO AND SMALLER THAN A KEY VALUE.
func UpperBound(s []int, key int) int {
	isSmaller := func(index, key int) bool {
		if s[index] <= key {
			return true
		} else {
			return false
		}
	}

	left, right := -1, len(s)

	for right-left > 1 {
		mid := left + (right-left)/2
		if isSmaller(mid, key) {
			left = mid
		} else {
			right = mid
		}
	}

	return left
}

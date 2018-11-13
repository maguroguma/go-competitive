package binary_search

func LowerBound(s []int, key int) int {
	left, right := -1, len(s)

	for right-left > 1 {
		mid := left + (right-left)/2
		if isLarger(s, mid, key) {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}

func isLarger(s []int, index, key int) bool {
	if s[index] >= key {
		return true
	} else {
		return false
	}
}

func UpperBound(s []int, key int) int {
	left, right := -1, len(s)

	for right-left > 1 {
		mid := left + (right-left)/2
		if isSmaller(s, mid, key) {
			left = mid
		} else {
			right = mid
		}
	}

	return left
}

func isSmaller(s []int, index, key int) bool {
	if s[index] <= key {
		return true
	} else {
		return false
	}
}

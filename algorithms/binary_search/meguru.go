package binary_search

// LowerBound returns an index of a slice whose value(s[idx]) is EQUAL TO AND LARGER THAN A KEY.
// The idx is the most left one when there are many keys.
// In other words, the idx is the point where the argument key should be inserted.
func LowerBound(s []int, key int) int {
	isLargerAndEqual := func(index, key int) bool {
		if s[index] >= key {
			return true
		}
		return false
	}

	left, right := -1, len(s)

	for right-left > 1 {
		mid := left + (right-left)/2
		if isLargerAndEqual(mid, key) {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}

// UpperBound returns an index of a slice whose value(s[idx]) is LARGER THAN A KEY.
// The idx is the most right one when there are many keys.
// In other words, the idx is the point where the argument key should be inserted.
func UpperBound(s []int, key int) int {
	isLarger := func(index, key int) bool {
		if s[index] > key {
			return true
		}
		return false
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

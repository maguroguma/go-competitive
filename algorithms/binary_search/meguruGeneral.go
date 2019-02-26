package binary_search

import "math"

func GeneralLowerBound(s []int, key int) int {
	isOK := func(index, key int) bool {
		if s[index] >= key {
			return true
		}
		return false
	}

	ng, ok := -1, len(s)
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func GeneralUpperBound(s []int, key int) int {
	isOK := func(index, key int) bool {
		if s[index] > key {
			return true
		}
		return false
	}

	ng, ok := -1, len(s)
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func GeneralDescLowerBound(s []int, key int) int {
	isOK := func(index, key int) bool {
		if s[index] <= key {
			return true
		}
		return false
	}

	ng, ok := -1, len(s)
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func GeneralDescUpperBound(s []int, key int) int {
	isOK := func(index, key int) bool {
		if s[index] < key {
			return true
		}
		return false
	}

	ng, ok := -1, len(s)
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid, key) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

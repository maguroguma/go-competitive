package binary_search

func BinarySearch(s []int, goal int) int {
	return recursion(s, goal, 0, len(s)-1)
}

func recursion(s []int, goal, start, end int) int {
	if start > end {
		return -1
	}

	center := start + (end-start)/2
	if s[center] > goal {
		return recursion(s, goal, start, center-1)
	} else if s[center] < goal {
		return recursion(s, goal, center+1, end)
	} else {
		return center
	}
}

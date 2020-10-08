package strings

// KMPTable calculates "tagged borders".
// Time complexity: O(len(pattern))
func KMPTable(pattern []rune) (T []int) {
	T = make([]int, len(pattern)+1)

	j := -1
	T[0] = j
	for i := 0; i < len(pattern); i++ {
		for j >= 0 && pattern[i] != pattern[j] {
			j = T[j]
		}
		j++
		T[i+1] = j
	}

	return T
}

// KMPSearch find all indices of a text that match a pattern.
// Indices are 0-index.
// Time complexity: O(len(text))
func KMPSearch(text, pattern []rune, kmpTable []int) []int {
	res := []int{}

	if len(pattern) == 0 {
		return res
	}

	j := 0
	for i := 0; i < len(text); i++ {
		for j >= 0 && text[i] != pattern[j] {
			j = kmpTable[j]
		}
		j++
		if j == len(pattern) {
			res = append(res, i-j+1)
			j = kmpTable[j]
		}
	}

	return res
}

// KMPStringPeriod returns periods of strings P[:i-1] for all i < len(P).
// Time complexity: O(len(P))
func KMPStringPeriod(kmpTable []int) []int {
	res := make([]int, len(kmpTable)-1)

	for i := 1; i < len(kmpTable); i++ {
		res[i-1] = i - kmpTable[i]
	}

	return res
}

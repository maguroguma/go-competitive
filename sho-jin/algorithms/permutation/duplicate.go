package permutation

// CalcDuplicatePatterns returns all patterns of n^k of elems([]rune).
func CalcDuplicatePatterns(elems []rune, k int) [][]rune {
	return dupliRec([]rune{}, elems, k)
}

// DFS function for CalcDuplicatePatterns.
func dupliRec(pattern, elems []rune, k int) [][]rune {
	if len(pattern) == k {
		return [][]rune{pattern}
	}

	res := [][]rune{}
	for _, e := range elems {
		newPattern := make([]rune, len(pattern))
		copy(newPattern, pattern)
		newPattern = append(newPattern, e)

		res = append(res, dupliRec(newPattern, elems, k)...)
	}

	return res
}

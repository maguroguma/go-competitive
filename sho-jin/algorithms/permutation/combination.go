package permutation

// CalcCombinationPatterns returns all patterns of nCk of elems([]rune).
func CalcCombinationPatterns(elems []rune, k int) [][]rune {
	newResi := make([]rune, len(elems))
	copy(newResi, elems)

	return combRec([]rune{}, newResi, k)
}

// DFS function for CalcCombinationPatterns.
func combRec(pattern, residual []rune, k int) [][]rune {
	if len(pattern) == k {
		return [][]rune{pattern}
	}

	res := [][]rune{}
	for i, e := range residual {
		newPattern := make([]rune, len(pattern))
		copy(newPattern, pattern)
		newPattern = append(newPattern, e)

		newResi := []rune{}
		newResi = append(newResi, residual[i+1:]...)

		res = append(res, combRec(newPattern, newResi, k)...)
	}

	return res
}

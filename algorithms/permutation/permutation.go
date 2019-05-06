package permutation

// CalcPermutationPatterns returns all patterns of nPk of elems([]rune).
func CalcPermutationPatterns(elems []rune, k int) [][]rune {
	newResi := make([]rune, len(elems))
	copy(newResi, elems)

	return permRec([]rune{}, newResi, k)
}

// DFS function for CalcPermutationPatterns.
func permRec(pattern, residual []rune, k int) [][]rune {
	if len(pattern) == k {
		return [][]rune{pattern}
	}

	res := [][]rune{}
	for i, e := range residual {
		newPattern := make([]rune, len(pattern))
		copy(newPattern, pattern)
		newPattern = append(newPattern, e)

		newResi := []rune{}
		newResi = append(newResi, residual[:i]...)
		newResi = append(newResi, residual[i+1:]...)

		res = append(res, permRec(newPattern, newResi, k)...)
	}

	return res
}

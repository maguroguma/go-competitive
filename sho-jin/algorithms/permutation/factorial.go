package permutation

// CalcFactorialPatterns returns all patterns of n! of elems([]rune).
func CalcFactorialPatterns(elems []rune) [][]rune {
	newResi := make([]rune, len(elems))
	copy(newResi, elems)

	return factRec([]rune{}, newResi)
}

// DFS function for CalcFactorialPatterns.
func factRec(pattern, residual []rune) [][]rune {
	if len(residual) == 0 {
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

		res = append(res, factRec(newPattern, newResi)...)
	}

	return res
}

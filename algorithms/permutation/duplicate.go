package permutation

func CalcDuplicatePatterns(elements []rune, digit int) [][]rune {
	return duplicateRecursion([]rune{}, elements, digit)
}

func duplicateRecursion(interim, elements []rune, digit int) [][]rune {
	if len(interim) == digit {
		return [][]rune{interim}
	}

	res := [][]rune{}
	for i := 0; i < len(elements); i++ {
		copiedInterim := make([]rune, len(interim))
		copy(copiedInterim, interim)
		copiedInterim = append(copiedInterim, elements[i])
		res = append(res, duplicateRecursion(copiedInterim, elements, digit)...)
	}

	return res
}

package permutation

func CalcFactorialPatterns(elements []rune) [][]rune {
	copiedResidual := make([]rune, len(elements))
	copy(copiedResidual, elements)
	return factorialRecursion([]rune{}, copiedResidual)
}

func factorialRecursion(interim, residual []rune) [][]rune {
	if len(residual) == 0 {
		return [][]rune{interim}
	}

	res := [][]rune{}
	for idx, elem := range residual {
		copiedInterim := make([]rune, len(interim))
		copy(copiedInterim, interim)
		copiedInterim = append(copiedInterim, elem)
		copiedResidual := genDeletedSlice(idx, residual)
		res = append(res, factorialRecursion(copiedInterim, copiedResidual)...)
	}

	return res
}

func genDeletedSlice(delId int, S []rune) []rune {
	res := []rune{}
	res = append(res, S[:delId]...)
	res = append(res, S[delId+1:]...)
	return res
}

package permutation

// GeneratePermutation returns n! in a [][]int style.
// Each pattern consists of 1 to n integers.
func GeneratePermutation(n int) [][]int {
	interim, residual := []int{}, []int{}
	for i := 1; i < n+1; i++ {
		residual = append(residual, i)
	}

	return recursion(interim, residual)
}

// recursion finally returns only leaf node of a tree diagram
func recursion(interim, residual []int) [][]int {
	if len(residual) == 0 {
		return [][]int{interim}
	}

	permutation := [][]int{}
	for i, r := range residual {
		copiedInterim := make([]int, len(interim))
		copy(copiedInterim, interim)
		copiedResidual := deleteElement(residual, i)

		copiedInterim = append(copiedInterim, r)
		p := recursion(copiedInterim, copiedResidual)
		permutation = append(permutation, p...)
	}
	return permutation
}

func deleteElement(s []int, i int) []int {
	newS := []int{}
	for j, e := range s {
		if j == i {
			continue
		}
		newS = append(newS, e)
	}
	return newS
}

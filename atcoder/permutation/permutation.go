package permutation

func GeneratePermutation(n int) [][]int {
	interim, residual := []int{}, []int{}
	for i := 1; i < n+1; i++ {
		residual = append(residual, i)
	}

	return recursion(interim, residual)
}

// 樹形図の葉ノード列のみを取得する
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
	new_s := []int{}
	for j, e := range s {
		if j == i {
			continue
		}
		new_s = append(new_s, e)
	}
	return new_s
}

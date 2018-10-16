package full_search

var n, k int
var a []int

func partialSum(n_arg int, a_arg []int, k_arg int) string {
	n = n_arg
	a = a_arg
	k = k_arg

	if solve(0, 0) {
		return "Yes"
	} else {
		return "No"
	}
}

// 状態（和）を記憶しながら深さ優先探索する再帰関数
func solve(i, sum int) bool {
	// 木の最下層のチェックが終わった
	if i == n {
		return sum == k
	}

	// 木の右側（ノードa[i]を使う方）へ下る
	if solve(i+1, sum+a[i]) {
		return true
	}

	// 木の左側（ノードa[i]を使わない方）へ下る
	if solve(i+1, sum) {
		return true
	}

	// 最下層に達してもkに等しいものが現れなかったとき
	return false
}

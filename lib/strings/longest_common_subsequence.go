package strings

// LCS returns one of the Longest Common Subsequence of S and T.
// O(|S| * |T|)
func LCS(S, T []rune) []rune {
	_chmax := func(updatedValue *int, target int) bool {
		if *updatedValue < target {
			*updatedValue = target
			return true
		}
		return false
	}

	dp := [][]int{}
	for i := 0; i <= len(S); i++ {
		row := make([]int, len(T)+1)
		dp = append(dp, row)
	}

	for i := 0; i < len(S); i++ {
		for j := 0; j < len(T); j++ {
			if S[i] == T[j] {
				_chmax(&dp[i+1][j+1], dp[i][j]+1)
			}
			_chmax(&dp[i+1][j+1], dp[i+1][j])
			_chmax(&dp[i+1][j+1], dp[i][j+1])
		}
	}

	revRes := make([]rune, 0, dp[len(S)][len(T)])
	si, ti := len(S), len(T)
	for si > 0 && ti > 0 {
		if dp[si][ti] == dp[si-1][ti] {
			si--
		} else if dp[si][ti] == dp[si][ti-1] {
			ti--
		} else {
			revRes = append(revRes, S[si-1])
			si--
			ti--
		}
	}

	res := make([]rune, len(revRes))
	for i := len(revRes) - 1; i >= 0; i-- {
		res[len(revRes)-1-i] = revRes[i]
	}

	return res
}

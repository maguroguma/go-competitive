package strings

// Levenshtein distance
// 1. change a character of S.
// 2. delete a character of S.
// 3. insert an any character to S.
// Time complexity: O(|S| * |T|)
func LevenshteinDistance(S, T []rune) int {
	_chmin := func(updatedValue *int, target int) bool {
		if *updatedValue > target {
			*updatedValue = target
			return true
		}
		return false
	}

	const LD_INF = 1 << 30

	dp := [][]int{}
	for i := 0; i <= len(S); i++ {
		row := make([]int, len(T)+1)
		dp = append(dp, row)
	}

	for i := 0; i <= len(S); i++ {
		for j := 0; j <= len(T); j++ {
			dp[i][j] = LD_INF
		}
	}

	dp[0][0] = 0
	for i := 0; i <= len(S); i++ {
		for j := 0; j <= len(T); j++ {
			// change S
			if i > 0 && j > 0 {
				if S[i-1] == T[j-1] {
					_chmin(&dp[i][j], dp[i-1][j-1])
				} else {
					_chmin(&dp[i][j], dp[i-1][j-1]+1)
				}
			}

			// delete S
			if i > 0 {
				_chmin(&dp[i][j], dp[i-1][j]+1)
			}

			// insert T
			if j > 0 {
				_chmin(&dp[i][j], dp[i][j-1]+1)
			}
		}
	}

	return dp[len(S)][len(T)]
}

package strings

// LongestCommonSubstring returns length of it.
// Time complexity: O(|S|*|T|)
func LongestCommonSubstring(S, T []rune) int {
	_min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	_chmax := func(updatedValue *int, target int) bool {
		if *updatedValue < target {
			*updatedValue = target
			return true
		}
		return false
	}
	_rle := func(S []int) ([]int, []int) {
		runes := []int{}
		lengths := []int{}

		l := 0
		for i := 0; i < len(S); i++ {
			if i == 0 {
				l = 1
				continue
			}

			if S[i-1] == S[i] {
				l++
			} else {
				runes = append(runes, S[i-1])
				lengths = append(lengths, l)
				l = 1
			}
		}
		runes = append(runes, S[len(S)-1])
		lengths = append(lengths, l)

		return runes, lengths
	}

	_solve := func(S, T []rune) int {
		match := make([]int, _min(len(S), len(T)))
		for i := 0; i < len(match); i++ {
			if S[i] == T[i] {
				match[i] = 1
			}
		}

		comp, cnts := _rle(match)

		res := 0
		for i := 0; i < len(comp); i++ {
			if comp[i] == 1 {
				_chmax(&res, cnts[i])
			}
		}

		return res
	}

	res := 0

	for i := 0; i < len(S); i++ {
		_chmax(&res, _solve(S[i:], T))
	}
	for i := 0; i < len(T); i++ {
		_chmax(&res, _solve(S, T[i:]))
	}

	return res
}

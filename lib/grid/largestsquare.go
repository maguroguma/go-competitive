package grid

// LargestSquare calculates maximum width of the largest square in the grid.
// Time complexity: O(h*w)
func LargestSquare(G [][]int) (maxWidth int) {
	const _G_BLOCK = 1

	h, w := len(G), len(G[0])
	W := make([][]int, h)
	for i := 0; i < h; i++ {
		W[i] = make([]int, w)
	}
	_max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	_min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 0; i < h; i++ {
		if G[i][0] == _G_BLOCK {
			W[i][0] = 0
		} else {
			W[i][0] = 1
		}
	}
	for j := 0; j < w; j++ {
		if G[0][j] == _G_BLOCK {
			W[0][j] = 0
		} else {
			W[0][j] = 1
		}
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if G[i][j] == _G_BLOCK {
				W[i][j] = 0
			} else {
				W[i][j] = _min(W[i-1][j-1], _min(W[i-1][j], W[i][j-1])) + 1
			}
		}
	}

	maxWidth = 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			maxWidth = _max(maxWidth, W[i][j])
		}
	}

	return maxWidth
}

package arithmetic

// LargestRectangle calculates an area of largest rectangle in a histgram.
// Time complexity: O(len(height))
func LargestRectangle(height []int64) int64 {
	H := make([]int64, len(height))
	copy(H, height)

	_max := func(x, y int64) int64 {
		if x > y {
			return x
		}
		return y
	}
	_top := func(S []int64) int64 { return S[len(S)-1] }

	st := []int64{}
	H = append(H, 0)
	left := make([]int64, len(H))
	res := int64(0)

	for i := int64(0); i < int64(len(H)); i++ {
		for len(st) > 0 && H[_top(st)] >= H[i] {
			res = _max(res, (i-left[_top(st)]-1)*H[_top(st)])
			st = st[:len(st)-1]
		}

		if len(st) == 0 {
			left[i] = -1
		} else {
			left[i] = _top(st)
		}

		st = append(st, i)
	}

	return res
}

package csum

type RectangleSum struct {
	recSum [][]int64
}

func NewRectangleSum(m [][]int64) *RectangleSum {
	rs := new(RectangleSum)

	h, w := len(m), len(m[0])
	for y := 0; y < h; y++ {
		tmp := make([]int64, w)
		rs.recSum = append(rs.recSum, tmp)
	}

	// Build
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rs.recSum[y][x] = m[y][x]
			if y > 0 {
				rs.recSum[y][x] += rs.recSum[y-1][x]
			}
			if x > 0 {
				rs.recSum[y][x] += rs.recSum[y][x-1]
			}
			if y > 0 && x > 0 {
				rs.recSum[y][x] -= rs.recSum[y-1][x-1]
			}
		}
	}

	return rs
}

// RangeSum returns a result of \sum_{i=top to bottom, j=left to right}
// Time complexity: O(1)
func (rs *RectangleSum) RangeSum(top, left, bottom, right int) int64 {
	res := rs.recSum[bottom][right]
	if left > 0 {
		res -= rs.recSum[bottom][left-1]
	}
	if top > 0 {
		res -= rs.recSum[top-1][right]
	}
	if left > 0 && top > 0 {
		res += rs.recSum[top-1][left-1]
	}
	return res
}

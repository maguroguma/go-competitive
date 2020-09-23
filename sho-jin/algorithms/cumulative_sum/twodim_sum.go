package cumulative_sum

type RectangleSum struct {
	matrix [][]int
	recSum [][]int
}

// NewRectangleSum は2次元累積和を計算するための構造体のポインタを返す
func NewRectangleSum(m [][]int) *RectangleSum {
	rs := new(RectangleSum)
	rs.matrix = m

	h, w := len(m), len(m[0])
	for y := 0; y < h; y++ {
		tmp := make([]int, w)
		rs.recSum = append(rs.recSum, tmp)
	}

	// 1行ずつスキャンする
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rs.recSum[y][x] = rs.matrix[y][x] // 同じ座標の値を加算
			if y > 0 {
				rs.recSum[y][x] += rs.recSum[y-1][x] // 1マス上の座標と原点座標がなす長方形の和を加算
			}
			if x > 0 {
				rs.recSum[y][x] += rs.recSum[y][x-1] // 1マス左の座標と原点座標がなす長方形の和を加算
			}
			if y > 0 && x > 0 {
				rs.recSum[y][x] -= rs.recSum[y-1][x-1] // 過剰に加算した部分（左上のマスと原点座標がなす長方形の和）を減算
			}
		}
	}

	return rs
}

// GetSum は2次元累積和の初期化と逆の要領で、グリッド内の任意の長方形の和を計算し返す
func (rs *RectangleSum) GetSum(top, left, bottom, right int) int {
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

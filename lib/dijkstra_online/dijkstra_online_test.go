package dijkstra_online

import (
	"os"
	"testing"
)

// https://atcoder.jp/contests/abc170/tasks/abc170_f
// Sample No.1

var (
	h, w, k        int
	x1, y1, x2, y2 int
	C              [][]rune

	steps [4][2]int
	N     int
	G     [4000000 + 50][]Edge

	Expected int
)

func initTest() {
	h, w, k = 3, 5, 2
	y1, x1, y2, x2 = 3, 2, 3, 4
	x1--
	y1--
	x2--
	y2--

	C = [][]rune{
		[]rune("....."),
		[]rune(".@..@"),
		[]rune("..@.."),
	}

	Expected = 5
}

func TestMain(m *testing.M) {
	println("before all...")

	initTest()

	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func TestDijkstraGenericOnline(t *testing.T) {
	steps = [4][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	N = h * w * 4
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			cbid := toNodeId(i, j)
			for _, step := range steps {
				dy, dx := step[0], step[1]
				ny, nx := i+dy, j+dx
				if 0 <= ny && ny < h && 0 <= nx && nx < w && C[ny][nx] == '.' {
					nid := toNodeId(ny, nx)
					w := Weight{cost: 1}
					G[cbid] = append(G[cbid], Edge{to: nid, w: w})
				}
			}
		}
	}

	vinf := Value{num: INF_B60, nokori: -1}
	vinit := Value{num: 0, nokori: 0}
	less := func(l, r Value) bool {
		if l.num < r.num {
			return true
		} else if l.num > r.num {
			return false
		} else {
			return l.nokori > r.nokori
		}
	}
	transit := func(cv *Vertex, AG [][]Edge) []*Vertex {
		res := []*Vertex{}

		cbid := fromVidToBid(cv.vid)
		prevd := tod(cv.vid)
		for _, e := range AG[cbid] {
			nbid := e.to
			nextd := dir(cbid, nbid)
			nval := Value{num: cv.v.num, nokori: cv.v.nokori}

			if prevd != nextd || nval.nokori == 0 {
				// 方向転換 or 残り回数0から次に進む
				nval.num++
				nval.nokori = k - 1
			} else {
				nval.nokori--
			}

			nvid := nbid*4 + nextd
			nv := &Vertex{nvid, nval}
			res = append(res, nv)
		}

		return res
	}
	ds := NewDijkstraSolver(vinf, less, transit)
	S := []StartPoint{}
	for d := 0; d < 4; d++ {
		S = append(S, StartPoint{vid: toid(y1, x1, d), vzero: vinit})
	}
	dp := ds.Dijkstra(S, N, G[:N])

	ans := INF_B60
	for d := 0; d < 4; d++ {
		id := toid(y2, x2, d)
		chmin(&ans, dp[id].num)
	}

	if ans != Expected {
		t.Errorf("got %v, want %v", ans, Expected)
	}
}

const (
	L, R, U, D = 0, 1, 2, 3
	INF_B60    = 1 << 60
)

func chmin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

func toNodeId(y, x int) int {
	return y*w + x
}
func fromVidToBid(vid int) int {
	return vid / 4
}
func dir(cbid, nbid int) int {
	cy, cx := cbid/w, cbid%w
	ny, nx := nbid/w, nbid%w

	if ny-cy == -1 {
		return U
	} else if ny-cy == 1 {
		return D
	} else if nx-cx == -1 {
		return L
	}
	return R
}
func toid(y, x, d int) int {
	return (y*w+x)*4 + d
}
func toy(id int) int {
	return id / (w * 4)
}
func tox(id int) int {
	return id / 4 % w
}
func tod(id int) int {
	return id % 4
}

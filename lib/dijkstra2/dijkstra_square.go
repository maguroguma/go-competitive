package dijkstra2

// type Value and Weight should be modified according to problems.

// DP value type
type Value struct {
	gas, times int
}

// weight of edges
type Weight struct {
	gas int
}

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	id    int
	vzero Value
}

// Less returns l < r, and shared with pq.
type Less func(l, r Value) bool

// Estimate returns next value considered by transition.
type Estimate func(cv Value, w Weight) Value

func NewDijkstraSolver(vinf Value, winf Weight, less Less, estimate Estimate) *DijkstraSolver {
	ds := new(DijkstraSolver)

	ds.vinf, ds.winf = vinf, winf
	ds.less, ds.estimate = less, estimate

	return ds
}

// verified by [ABC143-E](https://atcoder.jp/contests/abc143/tasks/abc143_e)
func (ds *DijkstraSolver) Dijkstra(S []StartPoint, n int, AG [][]Weight) []Value {
	// initialize data
	dp, colors := ds._initAll(n)

	// configure about start points (some problems have multi start points)
	ds._initStartPoint(S, dp, colors)

	// body of dijkstra algorithm (O(n^2))
	for {
		minv, u := ds.vinf, -1

		// find next optimal node
		for i := 0; i < n; i++ {
			if ds.less(dp[i], minv) && colors[i] != BLACK {
				u = i
				minv = dp[i]
			}
		}
		if u == -1 {
			break
		}

		colors[u] = BLACK

		// update all nodes v from node u
		for v := 0; v < n; v++ {
			if colors[v] != BLACK && AG[u][v] != ds.winf {
				nv := ds.estimate(dp[u], AG[u][v])
				if ds.less(nv, dp[v]) {
					dp[v] = nv
					colors[v] = GRAY
				}
			}
		}
	}

	return dp
}

type DijkstraSolver struct {
	vinf     Value
	winf     Weight
	less     Less
	estimate Estimate
}

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

// _initAll returns initialized dp and colors slices.
func (ds *DijkstraSolver) _initAll(n int) (dp []Value, colors []int) {
	dp, colors = make([]Value, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = ds.vinf
		colors[i] = WHITE
	}

	return dp, colors
}

// _initStartPoint returns initialized priority queue, and update dp and colors slices.
// *This function update arguments (side effects).*
func (ds *DijkstraSolver) _initStartPoint(S []StartPoint, dp []Value, colors []int) {
	for _, sp := range S {
		dp[sp.id] = sp.vzero
		colors[sp.id] = GRAY
	}
}

package dijkstra2

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

// DP value type
type V struct {
	// {{
	gas, times int
	// }}
}

// weight of edges
type EdgeWeight struct {
	// {{
	gas int
	// }}
}

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	id    int
	vzero V
}

type DijkstraSolver struct {
	vinf  V
	ewinf EdgeWeight
	Less  func(l, r V) bool          // Less returns l < r.
	NextV func(cv V, e EdgeWeight) V // NextV returns next value considered by transition.
}

func NewDijkstraSolver(
	vinf V, ewinf EdgeWeight, Less func(l, r V) bool, NextV func(cv V, ew EdgeWeight) V,
) *DijkstraSolver {
	ds := new(DijkstraSolver)

	ds.vinf, ds.ewinf = vinf, ewinf
	ds.Less, ds.NextV = Less, NextV

	return ds
}

// verified by [ABC143-E](https://atcoder.jp/contests/abc143/tasks/abc143_e)
func (ds *DijkstraSolver) Dijkstra(S []StartPoint, n int, AG [][]EdgeWeight) []V {
	// initialize data
	dp, colors := ds.initAll(n)

	// configure about start points (some problems have multi start points)
	ds.initStartPoint(S, dp, colors)

	// body of dijkstra algorithm (O(n^2))
	for {
		minv, u := ds.vinf, -1

		// find next optimal node
		for i := 0; i < n; i++ {
			if ds.Less(dp[i], minv) && colors[i] != BLACK {
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
			if colors[v] != BLACK && AG[u][v] != ds.ewinf {
				nv := ds.NextV(dp[u], AG[u][v])
				if ds.Less(nv, dp[v]) {
					dp[v] = nv
					colors[v] = GRAY
				}
			}
		}
	}

	return dp
}

// initAll returns initialized dp and colors slices.
func (ds *DijkstraSolver) initAll(n int) (dp []V, colors []int) {
	dp, colors = make([]V, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = ds.vinf
		colors[i] = WHITE
	}

	return dp, colors
}

// initStartPoint returns initialized priority queue, and update dp and colors slices.
// *This function update arguments (side effects).*
func (ds *DijkstraSolver) initStartPoint(S []StartPoint, dp []V, colors []int) {
	for _, sp := range S {
		dp[sp.id] = sp.vzero
		colors[sp.id] = GRAY
	}
}

package main

const (
	MAX_V     = 200000 + 5 // maximum node number of a tree
	MAX_LOG_V = 100 + 1    // maximum log(n)
)

type LCASolver struct {
	// Graph info
	G    [][]int // graph as adjacent list
	root int     // root node ID
	n    int     // node number

	// data structure for answer LCA
	parent [MAX_LOG_V][MAX_V]int
	depth  [MAX_V]int
}

func NewLCASolver(G [][]int, root, n int) *LCASolver {
	s := new(LCASolver)
	s.G, s.root, s.n = G, root, n
	s.initialize()
	return s
}

func (s *LCASolver) initialize() {
	s.dfs(s.root, -1, 0)
	for k := 0; k+1 < MAX_LOG_V; k++ {
		for v := 0; v < s.n; v++ {
			if s.parent[k][v] < 0 {
				s.parent[k+1][v] = -1
			} else {
				s.parent[k+1][v] = s.parent[k][s.parent[k][v]]
			}
		}
	}
}

func (s *LCASolver) dfs(v, p, d int) {
	s.parent[0][v] = p
	s.depth[v] = d
	for _, to := range s.G[v] {
		if to != p {
			s.dfs(to, v, d+1)
		}
	}
}

func (s *LCASolver) LCA(u, v int) int {
	if s.depth[u] > s.depth[v] {
		u, v = v, u
	}
	for k := 0; k < MAX_LOG_V; k++ {
		if ((s.depth[v] - s.depth[u]) >> uint(k) & 1) == 1 {
			v = s.parent[k][v]
		}
	}

	if u == v {
		return u
	}

	for k := MAX_LOG_V - 1; k >= 0; k-- {
		if s.parent[k][u] != s.parent[k][v] {
			u, v = s.parent[k][u], s.parent[k][v]
		}
	}

	return s.parent[0][u]
}

package treediameter

type Edge struct {
	to, cost int
}

// TreeDiameter returns a diameter of a tree graph.
// Time complexity: O(|E|)
func TreeDiameter(TG [][]Edge) int {
	// temporal data type
	type result struct {
		dist int // distance
		tid  int // terminal node id
	}

	// recursive function (DFS)
	// visit returns the farthest node from cid, when pid -> cid.
	var visit func(pid, cid int) result
	visit = func(pid, cid int) result {
		r := result{dist: 0, tid: cid}
		// DFS
		for _, e := range TG[cid] {
			if e.to != pid {
				t := visit(cid, e.to) // next transition (neighbor node)
				t.dist += e.cost
				if r.dist < t.dist {
					r = t
				}
			}
		}
		return r
	}

	// main algorithm (0 -> u, u -> v, ans := dist(u, v))
	r := visit(-1, 0)
	t := visit(-1, r.tid)
	return t.dist
}

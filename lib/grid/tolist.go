package grid

// GridToAdjacencyList provides basic function
//  for converting 2-dimensional grid to graph as adjacency list format.
// Grid: size is H*W, and no wall.
func GridToAdjacencyList(h, w int) (G [][]int, N int) {
	steps := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	toid := func(i, j int) int { return w*i + j }
	isInGrid := func(y, x int) bool {
		return 0 <= y && y < h && 0 <= x && x < w
	}

	N = h * w
	G = make([][]int, N)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			cid := toid(i, j)

			for _, s := range steps {
				dy, dx := s[0], s[1]
				ny, nx := i+dy, j+dx
				if isInGrid(ny, nx) {
					nid := toid(ny, nx)

					G[cid] = append(G[cid], nid)
				}
			}
		}
	}

	return
}

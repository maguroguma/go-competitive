package sssp

// verified by https://codeforces.com/contest/1320/problem/B
func SSSPByBFS(sid int, n int, AG [][]int, edgeCost int) (dp []int, visited []bool) {
	const INF_SSSP = 1 << uint(30)

	dp = make([]int, n)
	visited = make([]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = INF_SSSP
		visited[i] = false
	}

	Q := []Node{}
	dp[sid] = 0
	visited[sid] = true
	Q = append(Q, Node{id: sid, cost: dp[sid]})

	for len(Q) > 0 {
		cnode := Q[0]
		Q = Q[1:]

		for _, nid := range AG[cnode.id] {
			if visited[nid] {
				continue
			}

			dp[nid] = cnode.cost + edgeCost
			visited[nid] = true
			Q = append(Q, Node{id: nid, cost: dp[nid]})
		}
	}

	return dp, visited
}

type Node struct {
	id, cost int
}

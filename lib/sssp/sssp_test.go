package sssp

import "testing"

// https://codeforces.com/contest/1320/problem/B
// Sample-3
func TestSSSPByBFS(t *testing.T) {
	// n, m := 8, 13
	n := 8
	tmp := [][2]int{
		{8, 7},
		{8, 6},
		{7, 5},
		{7, 4},
		{6, 5},
		{6, 4},
		{5, 3},
		{5, 2},
		{4, 3},
		{4, 2},
		{3, 1},
		{2, 1},
		{1, 8},
	}
	k := 5
	P := []int{8, 7, 5, 2, 1}
	for i := 0; i < k; i++ {
		P[i]--
	}
	expMini, expMaxi := 0, 3

	G := make([][]int, n)
	RG := make([][]int, n)
	for _, e := range tmp {
		u, v := e[0], e[1]
		u--
		v--

		G[v] = append(G[v], u)
		RG[u] = append(RG[u], v)
	}

	dp, _ := SSSPByBFS(P[k-1], n, G, 1)

	cnts := make([]int, 200000+50)
	for i := 0; i < n; i++ {
		cnts[dp[i]]++
	}

	mini, maxi := 0, 0
	for i := 0; i < k-1; i++ {
		cid, nid := P[i], P[i+1]

		if dp[cid]-1 < dp[nid] {
			mini++
			maxi++
			continue
		}

		ok := false
		for _, e := range RG[cid] {
			if dp[cid] == dp[e]+1 && e != nid {
				ok = true
				break
			}
		}
		if ok {
			maxi++
		}
	}

	if mini != expMini || maxi != expMaxi {
		t.Errorf("got (%v, %v), want (%v, %v)", mini, maxi, expMini, expMaxi)
	}
}

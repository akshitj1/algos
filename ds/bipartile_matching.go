package ds

// https://cp-algorithms.com/graph/kuhn_maximum_bipartite_matching.html
// https://www.geeksforgeeks.org/maximum-bipartite-matching/

func findAugPath(u int, m []int, g [][]int, vis []bool) bool {
	if vis[u] {
		return false
	}
	vis[u] = true
	for _, v := range g[u] {
		if m[v] == -1 || findAugPath(m[v], m, g, vis) {
			m[v] = u
			return true
		}
	}
	return false
}

func MaxBiPartileMatching(g [][]int) int {
	N := len(g)
	M := 0
	for u := 0; u < N; u++ {
		for _, v := range g[u] {
			M = MaxInt(M, v+1)
		}
	}
	m := make([]int, M)
	fillArray(m, -1)
	numMatched := 0
	for u := 0; u < N; u++ {
		vis := make([]bool, N)
		if findAugPath(u, m, g, vis) {
			numMatched++
		}
	}
	return numMatched
}

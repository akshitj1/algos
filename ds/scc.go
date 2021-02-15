package ds

// returns visit order by finish time decreasing for first pass
func dfsG(u int, g [][]int, vis []bool) []int {
	if vis[u] {
		return []int{}
	}
	vis[u] = true
	visOrder := []int{u}
	for _, v := range g[u] {
		visOrder = append(visOrder, dfsG(v, g, vis)...)
	}
	return visOrder
}

// StronglyConnectedComponents gives list of nodes in each scc
func StronglyConnectedComponents(g [][]int) [][]int {
	numV := len(g)
	gT := make([][]int, numV)
	for u := range g {
		for _, v := range g[u] {
			gT[v] = append(gT[v], u)
		}
	}
	vis := make([]bool, numV)
	latestVisit := []int{}
	for u := range g {
		latestVisit = append(latestVisit, dfsG(u, g, vis)...)
	}
	// reset vis for scc
	vis = make([]bool, numV)
	scc := make([][]int, 0)
	for _, u := range latestVisit {
		if !vis[u] {
			// order is not really important
			scc = append(scc, dfsG(u, gT, vis))
		}
	}
	return scc
}

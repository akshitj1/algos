package ds

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func lPath(u int, p int, g [][]int) []int {
	lP := []int{}
	for _, v := range g[u] {
		if v == p {
			continue
		}
		cPath := lPath(v, u, g)
		if len(cPath) > len(lP) {
			lP = cPath
		}
	}
	return append(lP, u)
}

// https://leetcode.com/problems/minimum-height-trees
func findMinHeightTrees(n int, edges [][]int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	lP := lPath(0, -1, g)
	eNode := lP[0]
	lP = lPath(eNode, -1, g)
	if len(lP)%2 == 1 {
		return []int{lP[(len(lP)-1)/2]}
	}
	return []int{lP[(len(lP)-1)/2], lP[len(lP)/2]}
}

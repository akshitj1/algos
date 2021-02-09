package ds

func postOrder(u int, vis []TopologicalState, g [][]int) ([]int, bool) {
	if vis[u] == C {
		return []int{}, true
	}
	if vis[u] == P {
		return []int{}, false
	}
	vis[u] = P
	cO := make([]int, 0)
	for _, v := range g[u] {
		cOt, valid := postOrder(v, vis, g)
		if !valid {
			return []int{}, false
		}
		cO = append(cO, cOt...)
	}
	cO = append(cO, u)
	vis[u] = C
	return cO, true
}

// https://leetcode.com/problems/course-schedule-ii/submissions/
func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make([][]int, numCourses)
	vis := make([]TopologicalState, numCourses)
	for _, e := range prerequisites {
		v, u := e[0], e[1]
		g[u] = append(g[u], v)
	}
	cO := make([]int, 0)
	for u := 0; u < numCourses; u++ {
		cOt, valid := postOrder(u, vis, g)
		if !valid {
			return []int{}
		}
		cO = append(cO, cOt...)
	}
	// reverse
	for l, r := 0, len(cO)-1; l < r; l, r = l+1, r-1 {
		cO[l], cO[r] = cO[r], cO[l]
	}
	return cO
}

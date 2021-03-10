package ds

import (
	"sort"
	"testing"
)

func addUnDirectedEdge(u, v int, g [][]int) {
	g[u] = append(g[u], v)
	g[v] = append(g[v], u)
}

//https://www.geeksforgeeks.org/articulation-points-or-cut-vertices-in-a-graph/
func TestFindCutPoints(t *testing.T) {
	g := make([][]int, 5)
	addUnDirectedEdge(0, 1, g)
	addUnDirectedEdge(0, 2, g)
	addUnDirectedEdge(0, 3, g)
	addUnDirectedEdge(1, 2, g)
	addUnDirectedEdge(3, 4, g)
	got := FindCutPoints(g)
	sort.Ints(got)
	exp := []int{0, 3}
	if !ArrayEquals(got, exp) {
		t.Errorf("exp: %v, got: %v\n", exp, got)
	}
}

package ds

import (
	"testing"
)

func addEdge(u, v, eCap int, g [][]int, cap [][]int) {
	g[u] = append(g[u], v)
	cap[u][v] = eCap
}

// https://cp-algorithms.com/graph/edmonds_karp.html
func TestMaxFlow(t *testing.T) {
	source := 0
	sink := 5
	numV := 6
	g := make([][]int, numV)
	var cap [][]int
	for i := 0; i < numV; i++ {
		cap = append(cap, make([]int, numV))
	}
	addEdge(0, 1, 7, g, cap)
	addEdge(0, 4, 4, g, cap)
	addEdge(1, 2, 5, g, cap)
	addEdge(1, 3, 3, g, cap)
	addEdge(4, 1, 3, g, cap)
	addEdge(4, 3, 2, g, cap)
	addEdge(3, 2, 3, g, cap)
	addEdge(3, 5, 5, g, cap)
	addEdge(2, 5, 8, g, cap)
	gotFlow := MaxFlow(source, sink, g, cap)
	expFlow := 10
	if expFlow != gotFlow {
		t.Errorf("exp: %v, got %v", expFlow, gotFlow)
	}
}

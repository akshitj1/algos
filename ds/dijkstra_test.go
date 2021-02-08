package ds

import (
	"testing"
)

func TestDijkstraShortestPath(t *testing.T) {
	source := 0
	dest := 20
	var g [][]DijkstraNode
	for i := source; i <= dest; i++ {
		g = append(g, make([]DijkstraNode, 0))
	}
	for v := 1; v <= dest; v++ {
		u := v - 1
		g[u] = append(g[u], DijkstraNode{v, 1})
	}
	g[source] = append(g[source], DijkstraNode{dest, 100})
	gotSourceDestDist := DijkstraShortestPath(g, source, dest)
	expSourceDestDist := 20
	if gotSourceDestDist != expSourceDestDist {
		t.Errorf("Dijkstra %v to %d shortest path exp: %v, got: %v", source, dest, expSourceDestDist, gotSourceDestDist)
	}
}

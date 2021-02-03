package ds

import "testing"

func TestBreadthFirstSearch(t *testing.T) {

	g := make([][]int, 5)
	g[0] = append(g[0], 1)
	g[1] = append(g[1], 2)
	g[2] = append(g[2], 3)
	g[0] = append(g[0], 4)
	g[4] = append(g[4], 3)
	// 0->3
	gotPath, gotPathLen := ShortestPath(g, 0, 3)
	expPathLen := 2
	expPath := []int{0, 4, 3}
	if !(gotPathLen == expPathLen && ArrayEquals(gotPath, expPath)) {
		t.Errorf("exp: path: %v\ngot: %v", expPath, gotPath)
	}
}

package ds

import "testing"

func TestMaxBiPartileMatching(t *testing.T) {
	N := 6
	g := make([][]int, N)
	g[0] = []int{1, 2}
	g[1] = []int{0, 3}
	g[2] = []int{2}
	g[3] = []int{2, 3}
	g[4] = []int{}
	g[5] = []int{5}
	exp := 5
	got := MaxBiPartileMatching(g)
	if exp != got {
		t.Errorf("exp: %v, got: %v\n", exp, got)
	}
}

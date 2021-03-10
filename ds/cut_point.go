package ds

// https://cp-algorithms.com/graph/cutpoints.html

func findCutPoints(u, p int, t_in []int, t_min []int, timer *int, g [][]int) []int {
	t_in[u] = *timer
	t_min[u] = *timer
	*timer++
	numChild := 0
	cutP := []int{}
	for _, v := range g[u] {
		if v == p {
			continue
		}
		if t_in[v] != -1 {
			t_min[u] = MinInt(t_min[u], t_in[v])
		} else {
			cutP = append(cutP, findCutPoints(v, u, t_in, t_min, timer, g)...)
			t_min[u] = MinInt(t_min[u], t_min[v])
			if p != -1 && t_min[v] >= t_in[u] {
				cutP = append(cutP, u)
			}
			numChild++
		}
	}
	if p == -1 && numChild > 1 {
		cutP = append(cutP, u)
	}
	return cutP
}

func FindCutPoints(g [][]int) []int {
	N := len(g)
	cutPoints := []int{}
	timer := 0
	t_in := make([]int, N)
	t_min := make([]int, N)
	fillArray(t_in, -1)
	fillArray(t_min, -1)

	for u := 0; u < N; u++ {
		if t_in[u] == -1 {
			cutPoints = append(cutPoints, findCutPoints(u, -1, t_in, t_min, &timer, g)...)
		}
	}
	return cutPoints
}

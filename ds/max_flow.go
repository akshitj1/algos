package ds

import "math"

func findAugmentingPath(source int, sink int, g [][]int, cap [][]int, path []int) int {
	fillArray(path, -1)
	flow := math.MaxInt32
	q := NewQueue()
	q.Enqueue(source)
	path[source] = -2
	for u, exists := q.Dequeue(); exists; u, exists = q.Dequeue() {
		for _, v := range g[u] {
			if path[v] == -1 && cap[u][v] > 0 {
				path[v] = u
				q.Enqueue(v)
				flow = MinInt(flow, cap[u][v])
				if v == sink {
					return flow
				}
			}
		}
	}
	return 0
}

func fillArray(x []int, fillValue int) {
	for i := 0; i < len(x); i++ {
		x[i] = fillValue
	}
}

// MaxFlow computes max flow where g encodes connection info. in adjecency list.
// capacity is encoded in dense fixed size matrix. We use v,u for reverse capacity
func MaxFlow(source int, sink int, g [][]int, cap [][]int) int {
	numV := len(g)
	gA := make([][]int, numV)
	for u := range g {
		for _, v := range g[u] {
			gA[u] = append(gA[u], v)
			gA[v] = append(gA[v], u)
		}
	}
	if len(cap) != numV {
		panic("len mismatch")
	}
	maxFlow := 0
	// path stores idx of parent node
	path := make([]int, numV)
	for iter := 0; iter < numV; iter++ {
		augFlow := findAugmentingPath(source, sink, gA, cap, path)
		if augFlow == 0 {
			return maxFlow
		}
		for u, v := path[sink], sink; u != -2; u, v = path[u], u {
			cap[u][v] -= augFlow
			cap[v][u] += augFlow
		}
		maxFlow += augFlow
	}
	return maxFlow
}

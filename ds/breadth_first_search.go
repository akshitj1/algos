package ds

import (
	"math"
)

// BreadthFirstSearch returns length of shortest path, and parent from source to all reachable nodes in uniform weighted graph
func BreadthFirstSearch(graph [][]int, source int) (dist []int, parent []int) {
	numNodes := len(graph)
	dist = make([]int, numNodes)
	parent = make([]int, numNodes)
	for i := range dist {
		dist[i] = math.MaxInt32
		parent[i] = -1
	}
	q := Queue{}
	dist[source] = 0
	q.Enqueue(source)
	for !q.Empty() {
		u, _ := q.Dequeue()
		for _, v := range graph[u] {
			if dist[v] == math.MaxInt32 {
				q.Enqueue(v)
				dist[v] = dist[u] + 1
				parent[v] = u
			}
		}
	}
	return dist, parent
}

func reverse(x []int) []int {
	for l, r := 0, len(x)-1; l <= r; l, r = l+1, r-1 {
		x[l], x[r] = x[r], x[l]
	}
	return x
}

// ShortestPath finds shortest path and its lenght
func ShortestPath(graph [][]int, source int, dest int) ([]int, int) {
	d, p := BreadthFirstSearch(graph, source)
	if p[dest] == -1 {
		return []int{}, d[dest]
	}
	rPath := []int{dest}
	for parent := p[dest]; parent != -1; parent = p[parent] {
		rPath = append(rPath, parent)
	}
	return reverse(rPath), d[dest]
}

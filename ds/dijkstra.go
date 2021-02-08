package ds

import (
	"container/heap"
	"math"
)

type DijkstraNode struct {
	Id, W int
}

type DijkstraHeap []DijkstraNode

func (h DijkstraHeap) Len() int            { return len(h) }
func (h DijkstraHeap) Less(i, j int) bool  { return h[i].W < h[j].W }
func (h DijkstraHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *DijkstraHeap) Push(x interface{}) { *h = append(*h, x.(DijkstraNode)) }
func (h *DijkstraHeap) Pop() interface{} {
	topEl := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return topEl
}

// DijkstraShortestPath returns lenght of shortest path in non-negative graph of edge weights
func DijkstraShortestPath(graph [][]DijkstraNode, source int, dest int) int {
	h := &DijkstraHeap{}
	heap.Init(h)
	heap.Push(h, DijkstraNode{source, 0})
	numNodes := len(graph)
	vis := make([]bool, numNodes)
	numVis := 0
	for numVis < numNodes {
		u := heap.Pop(h).(DijkstraNode)
		if vis[u.Id] {
			continue
		}
		if u.Id == dest {
			return u.W
		}
		for _, v := range graph[u.Id] {
			heap.Push(h, DijkstraNode{v.Id, u.W + v.W})
		}
		vis[u.Id] = true
		numVis++
	}
	return math.MaxInt32
}

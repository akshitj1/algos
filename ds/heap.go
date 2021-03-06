package ds

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	topEl := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return topEl
}

// HeapSTL demonstrates use of heap stl
func HeapSTL(els []int) {
	h := &IntHeap{}
	heap.Push(h, 0)
	heap.Push(h, 2)
	heap.Push(h, 1)
	minEl := heap.Pop(h)
	_ = minEl
}

// Min Heap
type Heap struct {
	els []int
}

func NewHeap() *Heap {
	// insert 0 indexed empty node
	return &Heap{make([]int, 1)}
}

func (h *Heap) root() int {
	return 1
}

func (h *Heap) left(parent int) int {
	return parent << 1
}
func (h *Heap) right(parent int) int {
	return h.left(parent) + 1
}

func (h *Heap) parent(child int) int {
	return child >> 1
}

func (h *Heap) valid(node int) bool {
	return node >= h.root() && node < len(h.els)
}

func (h *Heap) swap(x, y int) {
	h.els[x], h.els[y] = h.els[y], h.els[x]
}

func (h *Heap) Empty() bool {
	return len(h.els) <= 1
}

func (h *Heap) Push(val int) {
	h.els = append(h.els, val)
	node := len(h.els) - 1
	// heapify bottom up
	h.heapify(node)
	// for parent := h.parent(node); h.valid(parent) && h.els[parent] > h.els[node]; node, parent = parent, h.parent(parent) {
	// 	h.swap(node, parent)
	// }
}

func (h *Heap) Pop() int {
	minVal := h.els[h.root()]
	h.els[h.root()] = h.els[len(h.els)-1]
	h.els = h.els[:len(h.els)-1]
	h.heapify(h.root())
	return minVal
}

func (h *Heap) Top(val int) int {
	return h.els[h.root()]
}

func (h *Heap) heapify(node int) {
	if !h.valid(node) {
		return
	}
	if parent := h.parent(node); h.valid(parent) && h.els[node] < h.els[parent] {
		h.swap(node, parent)
		h.heapify(parent)
		// we assume that heapify is either up or down during master call
		return
	}

	smallest := node
	for _, evalNode := range []int{h.left(node), h.right(node)} {
		if h.valid(evalNode) && h.els[smallest] > h.els[evalNode] {
			smallest = evalNode
		}
	}
	if smallest == node {
		return
	}
	h.swap(node, smallest)
	h.heapify(smallest)
}

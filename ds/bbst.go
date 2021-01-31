package ds

import "fmt"

// Array implementation
type BalancedBinaryTree struct {
	// root idx -> 1
	nodes []int
}

// if travelled from left, right or parent
type PathTrace int

const (
	Parent PathTrace = iota
	Left
	Right
	Nil
)

func (t *BalancedBinaryTree) root() int {
	return 1
}

func (t *BalancedBinaryTree) isValid(nodeIdx int) bool {
	return nodeIdx <= len(t.nodes) && nodeIdx > 0
}

func (t *BalancedBinaryTree) Parent(nodeIdx int) int {
	return nodeIdx >> 1
}
func (t *BalancedBinaryTree) Left(nodeIdx int) int {
	return nodeIdx << 1
}
func (t *BalancedBinaryTree) Right(nodeIdx int) int {
	return t.Left(nodeIdx) + 1
}

func (t *BalancedBinaryTree) Data(nodeIdx int) *int {
	return &t.nodes[nodeIdx-1]
}

func isLeftSubTree(nodeIdx int) bool {
	return nodeIdx%2 == 0
}

func isRightSubTree(nodeIdx int) bool {
	return !isLeftSubTree(nodeIdx)
}

func isRoot(nodeIdx int) bool {
	return nodeIdx == 1
}

func (t *BalancedBinaryTree) Swap(x int, y int) {
	*t.Data(x), *t.Data(y) = *t.Data(y), *t.Data(x)
}

// swaps till correct place is found
func (t *BalancedBinaryTree) Adjust(node int, trace PathTrace) {
	if parent := t.Parent(node); t.isValid(parent) && trace != Parent {
		if isLeftSubTree(node) && *t.Data(node) > *t.Data(parent) {
			t.Swap(node, parent)
			t.Adjust(parent, Left)
			return
		}
		if isRightSubTree(node) && *t.Data(node) < *t.Data(parent) {
			t.Swap(node, parent)
			t.Adjust(parent, Right)
			return
		}
	}
	if left := t.Left(node); t.isValid(left) && trace != Left {
		if *t.Data(node) < *t.Data(left) {
			t.Swap(node, left)
			t.Adjust(left, Parent)
			return
		}
	}
	if right := t.Right(node); t.isValid(right) && trace != Right {
		if *t.Data(node) > *t.Data(right) {
			t.Swap(node, right)
			t.Adjust(right, Parent)
			return
		}
	}
}

func (t *BalancedBinaryTree) Insert(val int) {
	t.nodes = append(t.nodes, val)
	// array idx is nodeIdx-1
	node := len(t.nodes)
	t.Adjust(node, Nil)
}

func log2(x int) int {
	for pos := 0; ; pos++ {
		if x == 0 {
			return pos
		}
		x >>= 1
	}
}

func floorPow2(x int) int {
	log2 := log2(x)
	if log2 == 0 {
		return 0
	}
	return 1 << (log2 - 1)
}

// root -> height 1
func (t *BalancedBinaryTree) Height(nodeIdx int) int {
	// position of MSB
	return log2(nodeIdx) + 1
}

func (t *BalancedBinaryTree) preorder(node int) []int {
	if !t.isValid(node) {
		return []int{}
	}
	vals := []int{*t.Data(node)}
	vals = append(vals, t.preorder(t.Left(node))...)
	vals = append(vals, t.preorder(t.Right(node))...)
	return vals
}

func (t *BalancedBinaryTree) String() string {
	return fmt.Sprintf("#Nodes: %v\nHeight: %v\npreorder: %v\n", len(t.nodes), t.Height(len(t.nodes)), t.preorder(t.root()))
}

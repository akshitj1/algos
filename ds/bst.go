package ds

import "fmt"

// Node is node for BinarySearchTree
type node struct {
	Data  int
	Left  *node
	Right *node
}

// BinarySearchTree data structure
type BinarySearchTree struct {
	root *node
}

func Add(nd *node, x int) *node {
	if nd == nil {
		nd = &node{x, nil, nil}
	} else if x <= nd.Data {
		nd.Left = Add(nd.Left, x)
	} else {
		nd.Right = Add(nd.Right, x)
	}
	return nd
}

// Add inserts x to appropriately sorted location
func (t *BinarySearchTree) Add(x int) {
	t.root = Add(t.root, x)
}

func inorder(nd *node) []int {
	if nd == nil {
		return []int{}
	}
	vals := inorder(nd.Left)
	vals = append(vals, nd.Data)
	vals = append(vals, inorder(nd.Right)...)
	return vals
}

func preorder(nd *node) []int {
	if nd == nil {
		return []int{}
	}
	vals := []int{nd.Data}
	vals = append(vals, preorder(nd.Left)...)
	vals = append(vals, preorder(nd.Right)...)
	return vals
}

func (t BinarySearchTree) String() string {
	return fmt.Sprintf("%v", preorder(t.root))
}

func addSortedArray(nd *node, vals []int) *node {
	if len(vals) == 0 {
		return nil
	}
	rootIdx := len(vals) / 2
	nd = &node{vals[rootIdx], nil, nil}
	nd.Left = addSortedArray(nd.Left, vals[:rootIdx])
	nd.Right = addSortedArray(nd.Right, vals[rootIdx+1:])
	return nd
}

func (t *BinarySearchTree) AddSortedArray(vals []int) {
	t.root = addSortedArray(t.root, vals)
}

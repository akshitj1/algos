package ds

import "fmt"

// Node is node for BinarySearchTree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// BinarySearchTree data structure
type BinarySearchTree struct {
	root *Node
}

func Add(nd *Node, x int) *Node {
	if nd == nil {
		nd = &Node{x, nil, nil}
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

func inorder(nd *Node) []int {
	if nd == nil {
		return []int{}
	}
	vals := inorder(nd.Left)
	vals = append(vals, nd.Data)
	vals = append(vals, inorder(nd.Right)...)
	return vals
}

func preorder(nd *Node) []int {
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

func addSortedArray(nd *Node, vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	rootIdx := len(vals) / 2
	nd = &Node{vals[rootIdx], nil, nil}
	nd.Left = addSortedArray(nd.Left, vals[:rootIdx])
	nd.Right = addSortedArray(nd.Right, vals[rootIdx+1:])
	return nd
}

func (t *BinarySearchTree) AddSortedArray(vals []int) {
	t.root = addSortedArray(t.root, vals)
}

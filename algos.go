package main

import (
	"fmt"

	"github.com/akshitj1/algos/ds"
)

func main() {
	tree := ds.BinarySearchTree{}
	vals := []int{1, 4, 5, 9, 12, 15}
	tree.AddSortedArray(vals)
	fmt.Println(tree)
	foo := [3]int{1, 2, 3}
	fmt.Printf("%v$%v#%v\n", foo)
}

package main

import (
	"fmt"
	"math/rand"

	"github.com/akshitj1/algos/ds"
)

func main() {
	tree := ds.BalancedBinaryTree{}
	vals := []int{1, 4, 5, 9, 12, 15}
	rand.Shuffle(len(vals), func(i, j int) { vals[i], vals[j] = vals[j], vals[i] })
	fmt.Printf("Insert order: %v\n", vals)
	for _, v := range vals {
		fmt.Println("inserting ", v)
		tree.Insert(v)
		fmt.Println(&tree)
	}

}

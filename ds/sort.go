package ds

import "sort"

type Ints []int

func (x Ints) Less(i, j int) bool {
	return x[i] < x[j]
}

func (x Ints) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x Ints) Len() int {
	return len(x)
}

// SortInts sorts integers
func SortInts(els []int) []int {
	sortedEls := els
	sort.Sort(Ints(sortedEls))
	return sortedEls
}

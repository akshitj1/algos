package ds

import "fmt"

type BinaryIndexedTree struct {
	data []int
	size int
}

func (t *BinaryIndexedTree) init(size int) {
	t.data = make([]int, size)
	t.size = size
}

func (t *BinaryIndexedTree) Put(idx int, val int) {
	for idx <= t.size {
		t.data[idx] += val
		idx += idx & -idx
	}
}

// CumulativeSum till idx inclusive
func (t *BinaryIndexedTree) CumulativeSum(idx int) int {
	cumSum := 0
	for idx > 0 {
		cumSum += t.data[idx]
		idx -= idx & -idx
	}
	return cumSum
}

// RangeSum gives l and r inclusive
func (t *BinaryIndexedTree) RangeSum(l, r int) int {
	if r < l {
		panic(fmt.Sprintf("malformed range: (%v -> %v)", l, r))
	}
	rangeSum := t.CumulativeSum(r)
	if l > 0 {
		rangeSum -= t.CumulativeSum(l - 1)
	}
	return rangeSum
}

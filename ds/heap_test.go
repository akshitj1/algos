package ds

import "testing"

func TestHeap(t *testing.T) {
	els := []int{4, 2, 9, 0, 1}
	exp := []int{0, 1, 2, 4, 9}
	h := NewHeap()
	for _, el := range els {
		h.Push(el)
	}
	var got []int
	for !h.Empty() {
		got = append(got, h.Pop())
	}
	if !ArrayEquals(got, exp) {
		t.Errorf("heap malfunction. exp: %v, got: %v", exp, got)
	}
}

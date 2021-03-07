package ds

import "testing"

func TestSegTree(t *testing.T) {
	tr := NewSegTree(4)
	tr.Update(0, 1, 0)
	tr.Update(0, 2, 2)
	tr.Update(0, 3, 3)
	res := make([]int, 0)
	t.Log(tr.data)
	for i := 0; i < 4; i++ {
		res = append(res, tr.Query(i))
	}
	t.Log(res)
}

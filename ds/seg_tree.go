package ds

func CeilPow2(x int) int {
	for y := x & (-x); y != x; x += y {
	}
	return x
}

type SegNode struct {
	a, b, h int
	l, r    *SegNode
}

func NewSegTree(a, b int) *SegNode {
	if a > b {
		return nil
	}
	t := &SegNode{}
	if a == b {
		return t
	}
	m := (a + b) / 2
	t.l = NewSegTree(a, m)
	t.r = NewSegTree(m+1, b)
	return t
}

func (t *SegNode) Query(p int) int {
	if t.a == t.b {
		return t.h
	}
	m := (t.a + t.b) / 2
	if p <= m {
		return MaxInt(t.h, t.l.Query(p))
	}
	return MaxInt(t.h, t.r.Query(p))
}

func (t *SegNode) Update(a, b, h int) {
	if b < t.a || t.b < a {
		return
	}
	if a <= t.a && t.b <= b {
		t.h = MaxInt(t.h, h)
		return
	}
	t.l.Update(a, b, h)
	t.r.Update(a, b, h)
}

// type SegTree struct {
// 	data         []int
// 	rootL, rootR int
// }

// func (t *SegTree) Leaf2Node(p int) int {
// 	return CeilPow2(t.rootR) + p
// }

// func NewSegTree(numEls int) *SegTree {
// 	treeSize := 2 * CeilPow2(numEls)
// 	data := make([]int, treeSize)
// 	fillArray(data, -1)
// 	return &SegTree{data, 0, numEls}
// }

// func (t *SegTree) Update(l, r, val int) {
// 	t.update(t.rootL, t.rootR, 1, l, r, val)
// }

// func left(nd, l, r int) (int, int, int) {
// 	return nd << 1, l, (l + r) / 2
// }

// func right(nd, l, r int) (int, int, int) {
// 	return nd<<1 + 1, (l+r)/2 + 1, r
// }

// func (t *SegTree) update(l, r, node, L, R, val int) {
// 	if r < L || R < l {
// 		return
// 	}
// 	if L <= l && r <= R {
// 		if t.data[node] >= 0 {
// 			t.data[t.Leaf2Node(t.data[node])] = val
// 		} else {
// 			t.data[node] = val
// 		}
// 		return
// 	}
// 	nnode, nl, nr := left(node, l, r)
// 	t.update(nl, nr, nnode, L, R, val)
// 	nnode, nl, nr = right(node, l, r)
// 	t.update(nl, nr, nnode, L, R, val)
// }

// func (t *SegTree) Query(p int) int {
// 	if t.data[t.Leaf2Node(p)] >= 0 {
// 		return t.data[t.Leaf2Node(p)]
// 	}
// 	return t.query(t.rootL, t.rootR, 1, p)
// }

// func (t *SegTree) query(l, r, node, p int) int {
// 	if t.data[node] != -1 {
// 		return t.data[node]
// 	}
// 	nnode, nl, nr := left(node, l, r)
// 	if nl <= p && p <= nr {
// 		return t.query(nl, nr, nnode, p)
// 	}
// 	nnode, nl, nr = right(node, l, r)
// 	return t.query(nl, nr, nnode, p)
// }

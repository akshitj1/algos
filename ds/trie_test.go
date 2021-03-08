package ds

// https://practice.geeksforgeeks.org/problems/minimum-xor-value-pair/1/
// https://www.interviewbit.com/problems/min-xor-value/
import "testing"

func TestTrie(t *testing.T) {
	a := []int{0, 4, 7, 9}
	exp := 3
	tr := NewTrie()
	minXor := 1<<30 - 1
	for _, x := range a {
		if tr.root != nil {
			y := tr.Query(x)
			minXor = MinInt(minXor, x^y)
		}
		tr.Insert(x)
	}
	got := minXor
	if got != exp {
		t.Errorf("exp: %v, got: %v\n", exp, got)
	}
}

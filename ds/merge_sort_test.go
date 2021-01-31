package ds

import "testing"

func TestMergeSort(t *testing.T) {
	inp := []int{7, 5, 1, 3, 8, 10}
	exp := []int{1, 3, 5, 7, 8, 10}
	out := MergeSort(inp)
	if len(out) != len(exp) {
		t.Errorf("length mismatch after sorting. expected: %v, got: %v", len(exp), len(out))
	}
	for i := 0; i < len(out); i++ {
		if out[i] != exp[i] {
			t.Errorf("Wrongly sorted: exp: %v, got: %v", exp, out)
			break
		}
	}
}

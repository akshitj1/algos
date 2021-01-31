package ds

import "testing"

func TestTwoPointerSearch(t *testing.T) {
	els := []int{1, 2, 3, 4, 4}
	sum := 8
	sumExists := TwoPointerSearch(els, sum)
	if !sumExists {
		t.Errorf("Could not find sum %v in array %v", sum, els)
	}
}

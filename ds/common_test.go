package ds

import (
	"fmt"
	"testing"
)

func TestTwoPointerSearch(t *testing.T) {
	els := []int{1, 2, 3, 4, 4}
	sum := 8
	sumExists := TwoPointerSearch(els, sum)
	if !sumExists {
		t.Errorf("Could not find sum %v in array %v", sum, els)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	subLen := LengthOfLongestSubstring(s)
	if subLen != 3 {
		t.Errorf("exp: %v, got: %v", 3, subLen)
	}
}

func ArrayEquals(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestSortInts(t *testing.T) {
	els := []int{2, 3, 10, 9, 1}
	exp := []int{1, 2, 3, 9, 10}
	got := SortInts(els)
	if !ArrayEquals(exp, got) {
		fmt.Printf("exp: %v, got: %v", exp, got)
	}

}

package ds

import "testing"

func compareRunTime(n int) bool { return 100*n*n < (1 << n) }

func TestBinarySearch(t *testing.T) {
	// Exercise 1.2-3, CSLR
	// smallest value of n where 100*n^2 is faster than 2^n
	n := BinarySearch(1, 30, compareRunTime)
	if !(compareRunTime(n-1) == false && compareRunTime(n) == true) {
		t.Errorf("%v is not critical point. Got transistion from %v to %v", n, compareRunTime(n-1), compareRunTime(n))
	}
}

package ds

// BinarySearch takes check function which goes from false to true in real line.
// will return first value at which its true
// check(l) should be false and check(r) should be true
func BinarySearch(l, r int, check func(int) bool) int {
	if r-l <= 1 {
		return r
	}
	n := (l + r) / 2
	if check(n) {
		return BinarySearch(l, n, check)
	}
	return BinarySearch(n, r, check)
}

package ds

// TwoPointerSearch assumes that a is sorted and moves start and end till a[l]+a[r]=sum
// exercise 2.3-7 CSLR
func TwoPointerSearch(a []int, sum int) bool {
	for start, end := 0, len(a)-1; start < end; start++ {
		for start < end && a[start]+a[end] > sum {
			end--
		}
		if start < end && a[start]+a[end] == sum {
			return true
		}
	}
	return false
}

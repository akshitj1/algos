package ds

// MergeSort n*log(n) memory
func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	l := MergeSort(a[:mid])
	r := MergeSort(a[mid:])
	return merge(l, r)
}

func merge(l, r []int) []int {
	sorted := make([]int, len(l)+len(r))
	for i, li, ri := 0, 0, 0; i < len(sorted); i++ {
		// we can also append intmax to l,r to make clean
		if ri == len(r) || (li < len(l) && l[li] < r[ri]) {
			sorted[i] = l[li]
			li++
		} else {
			sorted[i] = r[ri]
			ri++
		}
	}
	return sorted
}

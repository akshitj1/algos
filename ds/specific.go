package ds

// LengthOfLongestSubstring soln. for https://leetcode.com/problems/longest-substring-without-repeating-characters
func LengthOfLongestSubstring(s string) int {
	var si []int
	for _, c := range s {
		si = append(si, int(c))
	}
	vis := make([]bool, 256)
	q := &Queue{}
	maxLen := 0
	for _, num := range si {
		for vis[num] {
			tail, _ := q.Dequeue()
			vis[tail] = false
		}
		q.Enqueue(num)
		vis[num] = true
		curLen := q.Len()
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

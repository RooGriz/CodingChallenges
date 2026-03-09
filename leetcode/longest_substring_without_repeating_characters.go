package leetcode

// Time complexity - O(n), space complexity - O(1)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	unique := make(map[byte]int)
	longestLen := 1
	l, r := 0, 1
	unique[s[l]] = 0
	for r < len(s) {
		rc := s[r]
		pos, ok := unique[rc]
		if !ok || pos < l {
			longestLen = max(longestLen, r-l+1)
		} else {
			l = pos + 1
		}
		unique[rc] = r
		r++
	}
	return longestLen
}

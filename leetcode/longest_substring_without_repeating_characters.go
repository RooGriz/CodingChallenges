package leetcode

// Time complexity - O(n), space complexity - O(1)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	unique := make(map[byte]int)
	longestLen := 1
	curLen := 1
	l, r := 0, 1
	unique[s[l]] = 0
	for r < len(s) {
		rc := s[r]
		if _, ok := unique[rc]; !ok {
			unique[rc] = r
			curLen++
			longestLen = max(longestLen, curLen)
		} else {
			for lc := s[l]; l < r; lc = s[l] {
				l++
				if lc != rc {
					delete(unique, lc)
					curLen--
				} else {
					break
				}
			}
		}
		unique[rc] = r
		r++
	}
	return longestLen
}

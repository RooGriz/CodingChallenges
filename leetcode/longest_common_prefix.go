package leetcode

// Time complexity - O(s), space complexity - O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	l := 0
	for ; l < len(strs[0]); l++ {
		c := strs[0][l]
		ok := true
		for _, str := range strs {
			if l == len(str) || str[l] != c {
				ok = false
				break
			}
		}
		if !ok {
			break
		}
	}
	return strs[0][:l]
}

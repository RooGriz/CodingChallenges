package leetcode

// Time complexity - O(len(t)), space complexity - O(1)
func isSubsequence(s string, t string) bool {
	left := len(s)
	for _, c := range t {
		if left == 0 {
			break
		}
		if c == rune(s[len(s)-left]) {
			left--
		}
	}
	return left == 0
}

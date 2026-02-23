package leetcode

import "unicode"

// Time complexity - O(n), space complexity - O(1)
func isPalindrome(s string) bool {
	isAlphaNumeric := func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	}

	l := 0
	r := len(s) - 1
	ok := true
	for l < r {
		lr := rune(s[l])
		rr := rune(s[r])
		if !isAlphaNumeric(lr) {
			l++
			continue
		}
		if !isAlphaNumeric(rr) {
			r--
			continue
		}
		if unicode.ToLower(lr) == unicode.ToLower(rr) {
			l++
			r--
		} else {
			ok = false
			break
		}
	}
	return ok
}

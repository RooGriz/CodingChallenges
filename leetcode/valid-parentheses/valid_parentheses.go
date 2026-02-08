package leetcode

// Time complexity - O(n), space complexity: O(n)
func isValid(s string) bool {
	p := make([]rune, 0)
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			p = append(p, c)
			continue
		}
		if len(p) == 0 {
			return false
		}
		lp := len(p) - 1
		lc := p[lp]
		if (c == ')' && lc != '(') ||
			(c == '}' && lc != '{') ||
			(c == ']' && lc != '[') {
			return false
		}
		p = p[:lp]
	}
	return len(p) == 0
}

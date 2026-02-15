package leetcode

// Time complexity - O(n), space complexity - O(log(n))
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return p == nil && q == nil
}

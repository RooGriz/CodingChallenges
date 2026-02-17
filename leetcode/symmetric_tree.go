package leetcode

// Time complexity - O(n), space complexity - O(n)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricCheck(root.Left, root.Right)
}

func isSymmetricCheck(l, r *TreeNode) bool {
	if l != nil && r != nil && l.Val == r.Val {
		return isSymmetricCheck(l.Left, r.Right) && isSymmetricCheck(l.Right, r.Left)
	}
	return l == nil && r == nil
}

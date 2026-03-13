package leetcode

// Time complexity - O(n), space complexity - O(n)
func inorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	inorderTraversalWalk(root, &values)
	return values
}

func inorderTraversalWalk(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorderTraversalWalk(root.Left, values)
	}
	*values = append(*values, root.Val)
	if root.Right != nil {
		inorderTraversalWalk(root.Right, values)
	}
}

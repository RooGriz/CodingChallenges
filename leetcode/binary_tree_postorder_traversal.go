package leetcode

// Time complexity - O(n), space complexity - O(n)
func postorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	postorderTraversalWalk(root, &values)
	return values
}

func postorderTraversalWalk(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		postorderTraversalWalk(root.Left, values)
	}
	if root.Right != nil {
		postorderTraversalWalk(root.Right, values)
	}
	*values = append(*values, root.Val)
}

package leetcode

// Time complexity - O(n), space complexity - O(n)
func preorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	preorderTraversalWalk(root, &values)
	return values
}

func preorderTraversalWalk(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	*values = append(*values, root.Val)
	if root.Left != nil {
		preorderTraversalWalk(root.Left, values)
	}
	if root.Right != nil {
		preorderTraversalWalk(root.Right, values)
	}
}

package leetcode

// Time complexity - O(n), space complexity - O(n)
func maxDepth(root *TreeNode) int {
	return maxDepthSearch(root, 0)
}

func maxDepthSearch(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	return max(maxDepthSearch(root.Left, depth+1), maxDepthSearch(root.Right, depth+1))
}

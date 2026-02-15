package leetcode

import (
	"math"
)

// Time complexity - O(n), space complexity - O(n)
func isValidBST(root *TreeNode) bool {
	return isValidBSTCheck(root, math.MinInt32-1, math.MaxInt32+1)
}

func isValidBSTCheck(root *TreeNode, minVal int64, maxVal int64) bool {
	if root == nil {
		return true
	}
	val := int64(root.Val)
	if val <= minVal || val >= maxVal {
		return false
	}
	return isValidBSTCheck(root.Left, minVal, val) &&
		isValidBSTCheck(root.Right, val, maxVal)
}

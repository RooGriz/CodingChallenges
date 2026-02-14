package leetcode

// Time complexity - O(n), space complexity - O(1)
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	k := 0
	for i := range n {
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

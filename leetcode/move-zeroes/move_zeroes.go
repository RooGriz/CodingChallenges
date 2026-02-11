package leetcode

// Time complexity - O(n), space complexity - O(1)
func moveZeroes(nums []int) {
	l := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
}

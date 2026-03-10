package leetcode

// Time complexity - O(n), space complexity - O(1)
func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l := 0
	for i := range n {
		if nums[i] != val {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}

package leetcode

// Time complexity - O(n), space complexity - O(1)
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l := 0
	curLen := 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			curLen++
		} else {
			curLen = 1
		}
		if curLen <= 2 {
			l++
			nums[l] = nums[i]
		}
		prev = nums[i]
	}

	return l + 1
}

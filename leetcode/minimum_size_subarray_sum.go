package leetcode

// Time complexity - O(n), space complexity - O(1)
func minSubArrayLen(target int, nums []int) int {
	minLen := len(nums) + 1

	l, r := 0, 0
	curSum := 0
	for r < len(nums) {
		if curSum+nums[r] < target {
			curSum += nums[r]
			r++
			continue
		}
		curLen := r - l + 1
		if curLen == 1 {
			minLen = curLen
			break
		}
		minLen = min(minLen, curLen)
		curSum -= nums[l]
		l++
	}

	if minLen == len(nums)+1 {
		minLen = 0
	}

	return minLen
}

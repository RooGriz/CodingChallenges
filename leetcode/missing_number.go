package leetcode

// Time complexity - O(n), space complexity - O(1)
func missingNumber(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	size := len(nums)
	return size*(size+1)/2 - sum
}

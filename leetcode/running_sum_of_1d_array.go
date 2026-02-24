package leetcode

func runningSum(nums []int) []int {
	sums := make([]int, 0, len(nums))
	var prev int
	for _, num := range nums {
		sum := num + prev
		sums = append(sums, sum)
		prev = sum
	}
	return sums
}

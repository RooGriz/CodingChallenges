package leetcode

// Time complexity - O(n), space complexity - O(n)
func twoSum(nums []int, target int) []int {
	numsDict := make(map[int]int)
	for curPos, num := range nums {
		if foundPos, ok := numsDict[target-num]; ok {
			return []int{curPos, foundPos}
		}
		numsDict[num] = curPos
	}
	return []int{0, 0}
}

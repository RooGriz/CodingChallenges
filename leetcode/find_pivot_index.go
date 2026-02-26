package leetcode

// Time complexity - O(n), space complexity - O(n)
func pivotIndex(nums []int) int {
	psum := make([]int, len(nums)+1)
	psum[0] = 0
	for i, num := range nums {
		psum[i+1] = psum[i] + num
	}

	res := -1
	for i := range nums {
		lsum := psum[i]
		rsum := psum[len(psum)-1] - psum[i+1]
		if lsum == rsum {
			res = i
			break
		}
	}

	return res
}

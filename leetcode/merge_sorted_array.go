package leetcode

// Time complexity - O(m+n), space complexity - O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := range m {
		nums1[m+n-1-i] = nums1[m-1-i]
	}
	i := 0
	i1 := 0
	i2 := 0
	for i1 < m && i2 < n {
		if nums1[n+i1] < nums2[i2] {
			nums1[i] = nums1[n+i1]
			i1++
		} else {
			nums1[i] = nums2[i2]
			i2++
		}
		i++
	}
	for ; i1 < m; i1++ {
		nums1[i] = nums1[n+i1]
		i++
	}
	for ; i2 < n; i2++ {
		nums1[i] = nums2[i2]
		i++
	}
}

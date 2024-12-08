package arrays

func Merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, m+n)
	k := 0
	addedFromTwo := 0
	for i := 0; i < m; i++ {
		for j := addedFromTwo; j < n; j++ {
			if nums2[j] <= nums1[i] {
				result[k] = nums2[j]
				addedFromTwo++
				k++
			} else {
				break
			}
		}
		result[k] = nums1[i]
		k++
	}
	for j := addedFromTwo; j < n; j++ {
		result[k] = nums2[j]
		k++
	}
	copy(nums1, result)
}

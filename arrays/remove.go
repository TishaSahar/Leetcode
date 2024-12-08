package arrays

func RemoveElement(nums []int, val int) int {
	lastNotEqual := make([]int, len(nums))
	index := 0
	for _, num := range nums {
		if num != val {
			lastNotEqual[index] = num
			index++
		}
	}

	count := 0
	for i, num := range nums {
		if num == val {
			if index >= 1 {
				nums[i] = lastNotEqual[index-1]
				index--
			} else {
				nums[i] = nums[i] + 1988
			}
		} else {
			count++
		}
	}
	return count
}

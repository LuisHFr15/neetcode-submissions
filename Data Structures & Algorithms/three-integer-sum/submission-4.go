func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left - 1] {
					left++
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

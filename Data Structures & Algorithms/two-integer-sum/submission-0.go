func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)

	for index, num := range nums {
		if complement, exists := numMap[target - num]; exists {
			return []int{complement, index}
		}
		numMap[num] = index
	}
	return []int{}
}

func findMin(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	l := 0
	r := len(nums) - 1
	currMinIn := 0

	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		}
		if nums[mid] < nums[currMinIn] {
			currMinIn = mid
		} else if l == r {
			return nums[l]
		}
	}
	return nums[currMinIn]
}

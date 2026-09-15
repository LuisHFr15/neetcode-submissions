func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	rows := len(matrix)
	columns := len(matrix[0])
	end := rows * columns - 1
	start := 0
	for start <= end {
		mid := (start + end) / 2
		row := mid / columns
		column := mid % columns
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}

package question

//74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	//行
	row := len(matrix)
	//列
	col := len(matrix[0])
	if col == 0 {
		return false
	}
	//首先找到在哪一行
	up := 0
	down := row - 1
	for up <= down {
		temp := (up + down) / 2
		if target > matrix[temp][0] {
			up = temp + 1
		} else if target < matrix[temp][0] {
			down = temp - 1
		} else {
			return true
		}
	}
	//在down这一行，再次二分
	if down == -1 {
		return false
	}
	left := 1
	right := col - 1
	for left <= right {
		temp := (left + right) / 2
		if target > matrix[down][temp] {
			left = temp + 1
		} else if target < matrix[down][temp] {
			right = temp - 1
		} else {
			return true
		}
	}
	return false
}

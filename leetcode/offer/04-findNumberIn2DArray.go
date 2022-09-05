package offer

//剑指 Offer 04. 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	//从右上角判断
	i, j := 0, m-1
	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			//往左找
			j--
		} else {
			//往下找
			i++
		}
	}
	return false
}

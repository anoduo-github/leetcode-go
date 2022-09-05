package interview

//开始之前

//LC 搜索二维矩阵 II
func searchMatrix(matrix [][]int, target int) bool {
	//左下角到右上角的一条线，左边小，右边大
	row := len(matrix) - 1
	col := 0
	for row >= 0 && col < len(matrix[0]) {
		if target > matrix[row][col] {
			col++
		} else if target < matrix[row][col] {
			row--
		} else {
			return true
		}
	}
	return false
}

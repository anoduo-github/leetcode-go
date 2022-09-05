package practice

/* 任给一个数组，其中只有一个元素是单独出现，其他是成对出现，输出单独的元素。
例如： 输入： {2，2，1，1，4，4，7}
输出：7 */

func Pr2(nums []int) int {
	//map
	m := make(map[int]int)
	//遍历
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	//遍历
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

//74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

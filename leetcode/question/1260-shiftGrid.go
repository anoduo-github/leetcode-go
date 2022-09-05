package question

//1260. 二维网格迁移
func shiftGrid(grid [][]int, k int) [][]int {
	if k == 0 {
		return grid
	}
	//转为一维数组
	temp1 := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		temp1 = append(temp1, grid[i]...)
	}
	//移动
	length := len(temp1)
	temp2 := make([]int, length)
	for i := 0; i < length; i++ {
		temp2[(i+k)%length] = temp1[i]
	}
	//重构二维数组
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = temp2[i*n : (i+1)*n]
	}
	return res
}

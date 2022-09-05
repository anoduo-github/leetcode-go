package question

//200. 岛屿数量
func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				island(grid, i-1, j)
				island(grid, i, j-1)
				island(grid, i, j+1)
				island(grid, i+1, j)
			}
		}
	}
	return res
}

func island(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	if grid[i][j] == '1' {
		//遍历到的置0
		grid[i][j] = '0'
		//遍历上边的
		island(grid, i-1, j)
		//遍历左边的
		island(grid, i, j-1)
		//遍历右边的
		island(grid, i, j+1)
		//遍历下边的
		island(grid, i+1, j)
	}
}

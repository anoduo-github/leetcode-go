package question

//64. 最小路径和
func minPathSum(grid [][]int) int {
	//动态规划

	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 && j != 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			} else if j == 0 && i != 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

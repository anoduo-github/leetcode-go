package offer

/* dp[i][0]=grid[i][0]+grid[i-1][0]
dp[0][j]=grid[0][j]+grid[0][j-1]
dp[0][0]=grid[0][0]
dp[i][j]=max(dp[i-1][j],dp[i][j-1])+grid[i][j] */

//剑指 Offer 47. 礼物的最大价值
func maxValue(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	i, j := 0, 0
	for i < m {
		if j == 0 {
			if i > 0 {
				grid[i][0] = grid[i][0] + grid[i-1][0]
			}
		} else {
			if i == 0 {
				grid[0][j] = grid[0][j] + grid[0][j-1]
			} else {
				grid[i][j] = max(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
		j++
		if j == n {
			j = 0
			i++
		}
	}
	return grid[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package question

//279. 完全平方数
func numSquares(n int) int {
	nums := squares(n)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
	}
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		for j := temp; j <= n; j++ {
			dp[j] = min2(dp[j], dp[j-nums[i]]+1)
		}
	}
	return dp[n]
}

// 获取小于n的完全平方数
func squares(n int) []int {
	res := []int{1}
	temp := n / 2
	for i := 2; i <= temp; i++ {
		num := i * i
		if num <= n {
			res = append(res, num)
		}
	}
	return res
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

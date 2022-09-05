package primary

import "math"

//动态规划

//LC 爬楼梯
//f(n)=f(n-1)+f(n-2)
func climbStairs(n int) int {
	//超时了
	/* if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2) */

	//非递归
	if n == 1 {
		return 1
	}
	res := make([]int, n+1)
	res[1] = 1
	res[2] = 2
	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}

//LC 买卖股票的最佳时机
func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		max = int(math.Max(float64(max), float64(prices[i]-min)))
		min = int(math.Min(float64(min), float64(prices[i])))
	}
	if prices[len(prices)-1]-min > max {
		max = prices[len(prices)-1] - min
	}
	return max
}

//LC 最大子序和
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

//LC 打家劫舍
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	res := make([][]int, 2)
	res[0] = make([]int, length) //表示当前i这个没偷
	res[1] = make([]int, length) //表示当前i这个偷了
	res[0][0] = 0
	res[1][0] = nums[0]
	for i := 1; i < length; i++ {
		//当前这个没偷的
		if res[0][i-1] > res[1][i-1] {
			res[0][i] = res[0][i-1] //要么等于前一个没偷的
		} else {
			res[0][i] = res[1][i-1] //要么等于前一个偷的
		}
		//当前偷的，等于当前值加上一个没偷的
		res[1][i] = res[0][i-1] + nums[i]
	}
	if res[1][length-1] > res[0][length-1] {
		return res[1][length-1]
	} else {
		return res[0][length-1]
	}
}

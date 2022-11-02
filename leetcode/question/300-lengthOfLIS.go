package question

//300. 最长递增子序列
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	res := 1
	for i := 1; i < n; i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				max = getMax(max, dp[j]+1)
			}
		}
		dp[i] = max
		res = getMax(res, max)
	}
	return res
}

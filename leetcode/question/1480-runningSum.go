package question

//1480. 一维数组的动态和
func runningSum(nums []int) []int {
	sum := 0
	res := make([]int, 0)
	for _, v := range nums {
		sum = sum + v
		res = append(res, sum)
	}
	return res
}

/* 输入：nums = [1,2,3,4]
输出：[1,3,6,10]
解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。 */

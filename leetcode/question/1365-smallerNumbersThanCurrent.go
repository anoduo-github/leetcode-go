package question

//1365. 有多少小于当前数字的数字
func smallerNumbersThanCurrent(nums []int) []int {
	//计数
	count := [101]int{}
	//遍历数组，记录每位数出现的次数
	for _, v := range nums {
		count[v]++
	}
	//计算比i小的数之和，即i-1、i-2....0
	for i := 1; i <= 100; i++ {
		count[i] = count[i] + count[i-1]
	}
	res := make([]int, len(nums))
	for k, v := range nums {
		if v > 0 {
			res[k] = count[v-1]
		}
	}
	return res
}

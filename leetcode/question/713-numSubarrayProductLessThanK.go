package question

//713. 乘积小于 K 的子数组
func numSubarrayProductLessThanK(nums []int, k int) int {

	//滑动窗口
	count := 0
	prod, i := 1, 0
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		count += j - i + 1
	}

	return count

	//超时了
	/* count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			//第一个数小于k,表示只有nums[i]的数组符合条件
			count++
			//用这个数当除数，用k去除，得到的数就是数量大于1的目标数组其他几数的乘积不大于的数
			var temp int
			if k%nums[i] == 0 {
				temp = k/nums[i] - 1	//减一是因为不能等于temp,否则结果就是等于k而不是小于k了
			} else {
				temp = k / nums[i]
			}
			//把j当作第二个数开始计算
			j := i + 1
			for j < len(nums) {
				//是可以等于的
				if nums[j] <= temp {
					//表示i，j两个都可以
					count++
					//缩小temp
					temp = temp / nums[j]
					j++
				} else {
					break
				}
			}
		}
	}
	return count */
}

package question

//31. 下一个排列
func nextPermutation(nums []int) {
	//从往后前找
	for last := len(nums) - 1; last >= 1; last-- {
		if nums[last-1] < nums[last] {
			//表示当前的大于前一个，那么从当前往最后找，找到到一个最小的值与last-1所在的值替换
			//又由于下一个排列是最小的大于当前的排列，所以将last往最后的进行递增排序，即符合最小的
			//同时根据判断条件，直到last到最后的都是前面的大于后面的，所以直接反转即可
			index := last
			min := nums[last]
			for i := last + 1; i < len(nums); i++ {
				if nums[i] > nums[last-1] {
					if nums[i] <= min {
						min = nums[i]
						index = i
					}
				}
			}
			//交换
			nums[last-1], nums[index] = nums[index], nums[last-1]
			//反转
			k, j := last, len(nums)-1
			for k < j {
				nums[k], nums[j] = nums[j], nums[k]
				k++
				j--
			}
			return
		}
	}
	//表示是降序的，直接反转
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

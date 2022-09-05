package question

//33. 搜索旋转排序数组
func search31(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return look(nums, 0, len(nums)-1, target)
}

func look(nums []int, index, last, target int) int {
	if index > last {
		return -1
	}
	mid := (index + last) / 2    //获取中间下标
	if nums[mid] > nums[index] { //表示[index,mid]递增,[mid+1,last]非递增
		i, j := index, mid
		for i <= j {
			temp := (i + j) / 2
			if nums[temp] > target {
				j = temp - 1
			} else if nums[temp] < target {
				i = temp + 1
			} else {
				return temp
			}
		}
		return look(nums, mid+1, last, target) //递归处理非递增的情况
	} else if nums[mid] < nums[index] {
		i, j := mid+1, last
		for i <= j {
			temp := (i + j) / 2
			if nums[temp] > target {
				j = temp - 1
			} else if nums[temp] < target {
				i = temp + 1
			} else {
				return temp
			}
		}
		return look(nums, index, mid, target)
	} else {
		//表示mid和index相对，即区间缩小到只有一个或两个的情况，遍历下
		for i := index; i <= last; i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}
}

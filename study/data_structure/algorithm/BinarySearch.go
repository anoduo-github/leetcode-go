package algorithm

//NoRecv 二分查找（非递归）
func NoRecv(arr []int, target int) int {
	left := 0             //左下标
	right := len(arr) - 1 //右下标
	for left <= right {
		mid := (left + right) / 2 //中间下标
		if arr[mid] == target {
			return mid //找到值
		} else if arr[mid] > target { //说明值在左边
			right = mid - 1 //右下标重新赋值
		} else if arr[mid] < target {
			left = mid + 1 //左下标重新赋值
		}
	}
	return -1 //没找到
}

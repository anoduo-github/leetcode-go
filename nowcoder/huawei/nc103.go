package huawei

//solve 反转字符串
func solve(str string) string {
	//双指针
	if len(str) < 2 {
		return str
	}
	arr := []rune(str)
	//最左
	left := 0
	//最右
	right := len(str) - 1
	for left <= right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return string(arr)
}

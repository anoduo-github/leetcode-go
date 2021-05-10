package question

//1221. 分割平衡字符串
func balancedStringSplit(s string) int {
	if len(s) <= 1 {
		return 0
	}
	//双指针
	left := 0
	right := 1
	count := 0
	for left <= right {
		if s[left] != s[right] {
			count++
		}
	}
	return -999
}

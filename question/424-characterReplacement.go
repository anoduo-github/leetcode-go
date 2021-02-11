package question

//424. 替换后的最长重复字符
func characterReplacement(s string, k int) int {
	//定义一个字母数组存放每个字母出现的次数
	var arr [26]int
	//定义一个变量保存滑块长度
	max := 0
	//定义滑块左下标
	left := 0
	//遍历
	for right, v := range s {
		arr[v-'A']++
		max = getMax(max, arr[v-'A'])
		if right-left+1 > max+k {
			arr[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func getMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

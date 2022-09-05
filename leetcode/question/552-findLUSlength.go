package question

//522. 最长特殊序列 II
func findLUSlength(strs []string) int {
	res := -1
LABEL:
	for i := 0; i < len(strs); i++ {
		temp := strs[i]
		for j := 0; j < len(strs); j++ {
			if i != j && isSub(temp, strs[j]) {
				continue LABEL
			}
		}
		if res < len(temp) {
			res = len(temp)
		}
	}
	return res
}

//子序列 t中的字符在s中都有，且顺序一致
//isSub t是否是s的子序列
func isSub(t, s string) bool {
	l1 := len(t)
	l2 := len(s)
	if l1 > l2 {
		return false
	}
	index := 0
	for i := 0; i < l2; i++ {
		//判断t中的字符是否在s中存在
		if t[index] == s[i] {
			index++
			if index == l1 {
				return true
			}
		}
	}
	return false
}

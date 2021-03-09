package question

//1047. 删除字符串中的所有相邻重复项
func removeDuplicates(S string) string {
	if len(S) < 2 {
		return S
	}
	res := make([]rune, 0)
	for _, v := range S {
		if len(res) > 0 && v == res[len(res)-1] {
			res = res[:len(res)-1]
		} else {
			res = append(res, v)
		}
	}
	return string(res)
}

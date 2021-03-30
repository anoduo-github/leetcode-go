package question

//1528. 重新排列字符串
func restoreString(s string, indices []int) string {
	/* m := make(map[int]byte)
	for i := 0; i < len(s); i++ {
		m[indices[i]] = s[i]
	}
	str := ""
	for j := 0; j < len(s); j++ {
		str = str + string(m[j])
	}
	return str */
	l := len(s)
	res := make([]byte, l)
	for i := 0; i < l; i++ {
		res[indices[i]] = s[i]
	}
	return string(res)
}

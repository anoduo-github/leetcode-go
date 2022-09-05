package question

//392. 判断子序列
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if t[j] == s[i] {
			i++
		}
		j++
	}
	return i == len(s)
}

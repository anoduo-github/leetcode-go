package question

//242. 有效的字母异位词
func isAnagram(s string, t string) bool {
	var i, j [26]int
	for _, v := range []byte(s) {
		i[v-'a']++
	}
	for _, v := range []byte(t) {
		j[v-'a']++
	}
	return i == j
}

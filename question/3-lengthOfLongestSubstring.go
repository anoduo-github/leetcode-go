package question

//lengthOfLongestSubstring  3.无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	i := 0
	j := 1
	m := make(map[byte]int)
	m[s[i]] = 1
	for j < len(s) {
		if _, ok := m[s[j]]; ok {
			delete(m, s[i])
			i++
			j++
		} else {
			m[s[j]] = 1
			j++
		}
	}
	return j - i + 1
}

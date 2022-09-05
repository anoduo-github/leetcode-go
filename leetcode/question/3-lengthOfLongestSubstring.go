package question

//lengthOfLongestSubstring  3.无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length < 2 {
		return length
	}
	m := map[byte]int{}
	index := -1
	max := 0
	for i := 0; i < length; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for index+1 < length && m[s[index+1]] == 0 {
			m[s[index+1]]++
			index++
		}
		if max < index-i+1 {
			max = index - i + 1
		}
	}
	return max
}

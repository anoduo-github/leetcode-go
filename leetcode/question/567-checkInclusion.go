package question

// 567. 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	length1, length2 := len(s1), len(s2)
	if length1 > length2 {
		return false
	}
	var c1, c2 [26]int
	//遍历
	for i := 0; i < length1; i++ {
		c1[s1[i]-'a']++
		c2[s2[i]-'a']++
	}
	if c1 == c2 {
		return true
	}
	for i := length1; i < length2; i++ {
		c2[s2[i]-'a']++
		c2[s2[i-length1]-'a']--
		if c1 == c2 {
			return true
		}
	}
	return false
}

package question

//1684. 统计一致字符串的数目
func countConsistentStrings(allowed string, words []string) int {
	//1.将allowed数组字符放入哈希
	m := make(map[rune]int)
	for _, v := range allowed {
		m[v] = 1
	}
	count := 0
	//2.遍历words
	for _, v := range words {
		flag := true
		for _, v2 := range v {
			if _, ok := m[v2]; !ok {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}
	return count
}

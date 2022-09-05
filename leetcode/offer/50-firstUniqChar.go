package offer

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	temp := make([]int, 26)
	for _, v := range s {
		temp[v-'a']++
	}
	for _, v := range s {
		if temp[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}

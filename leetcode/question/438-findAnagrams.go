package question

//438. 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	lenS := len(s)
	lenP := len(p)
	if lenS < lenP {
		return []int{}
	}
	res := make([]int, 0)
	var arrS, arrP [26]int
	//遍历p，同时也算遍历[0,lenP]的s
	for i, v := range p {
		arrS[s[i]-'a']++
		arrP[v-'a']++
	}
	//相同长度的数组是可以比较的
	if arrS == arrP {
		//表示初始遍历就满足了，索引为0
		res = append(res, 0)
	}
	//遍历后面的s
	for i := 1; i < lenS-lenP+1; i++ {
		arrS[s[i-1]-'a']--
		arrS[s[i+lenP-1]-'a']++
		if arrS == arrP {
			res = append(res, i)
		}
	}
	return res
}

package question

//205. 同构字符串
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//找映射关系
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	//遍历
	for i := 0; i < len(s); i++ {
		if (s2t[s[i]] > 0 && s2t[s[i]] != t[i]) || (t2s[t[i]] > 0 && t2s[t[i]] != s[i]) {
			return false
		}
		s2t[s[i]] = t[i]
		t2s[t[i]] = s[i]
	}
	return true
}

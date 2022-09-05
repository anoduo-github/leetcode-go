package question

//804. 唯一摩尔斯密码词
func uniqueMorseRepresentations(words []string) int {
	if len(words) == 1 {
		return 1
	}
	mos := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	m := make(map[string]int, 0)
	for _, v := range words {
		s := ""
		for _, b := range v {
			s += mos[b-'a']
		}
		if _, ok := m[s]; !ok {
			m[s] = 1
		}
	}
	return len(m)
}

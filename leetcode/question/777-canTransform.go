package question

//777. 在LR字符串中交换相邻字符
func canTransform(start string, end string) bool {
	if (len(start) == 1 || len(end) == 1) || len(start) != len(end) {
		return false
	}
	s := []byte(start)
	e := []byte(end)

	for i := 0; i < len(s)-1; i++ {
		if s[i] != e[i] {
			//左右交换
			if (s[i] == 'X' && s[i+1] == 'L') || (s[i] == 'R' && s[i+1] == 'X') {
				s[i], s[i+1] = s[i+1], s[i]
			}
			//重新判断
			if s[i] != e[i] {
				return false
			}
		}
	}

	return s[len(s)-1] == e[len(e)-1]
}

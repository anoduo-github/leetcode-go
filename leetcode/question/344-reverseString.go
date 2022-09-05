package question

//344. 反转字符串
func reverseString(s []byte) {
	//双指针
	index, last := 0, len(s)-1
	for index < last {
		s[index], s[last] = s[last], s[index]
		index++
		last--
	}
}

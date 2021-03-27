package question

//1662. 检查两个字符串数组是否相等
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1 := ""
	str2 := ""
	for _, v := range word1 {
		str1 = str1 + v
	}
	for _, v := range word2 {
		str2 = str2 + v
	}
	return str1 == str2
}

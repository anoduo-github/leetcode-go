package primary

import "strconv"

//LC最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

//LC 外观数列
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	//存放数据
	s := ""
	//数组下标
	index := 0
	//计数器
	count := 0
	//获取前一列字符串
	str := countAndSay(n - 1)
	i := 0
	for i < len(str) {
		for str[i] == str[index] {
			count++
			index++
			if index >= len(str) {
				break
			}
		}
		s = s + strconv.Itoa(count) + string(str[index-1])
		//移动下标
		i = index
		//清零
		count = 0
	}
	return s
}
